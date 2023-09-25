package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/nifcloud/nifcloud-sdk-go/nifcloud"
	"golang.org/x/exp/slices"

	"github.com/SogoKato/nifdiff/pkg/plugin"
	baseplugin "github.com/SogoKato/nifdiff/pkg/plugin/base"
	computingplugin "github.com/SogoKato/nifdiff/pkg/plugin/computing"
)

func main() {
	plugins := []plugin.Plugin{
		&baseplugin.BasePlugin{},
		&computingplugin.ComputingPlugin{},
	}
	if len(os.Args) != 3 {
		panic(fmt.Errorf("usage: nifdiff NRN_A NRN_B"))
	}
	nrnResourceA := os.Args[1]
	nrnResourceB := os.Args[2]
	resourceA, err := getResource(plugins, nrnResourceA)
	if err != nil {
		panic(err)
	}
	resourceB, err := getResource(plugins, nrnResourceB)
	if err != nil {
		panic(err)
	}
	cmpOpts := getCmpOpts(plugins)
	if diff := cmp.Diff(resourceA, resourceB, cmpOpts...); diff != "" {
		fmt.Printf("Mismatch:\n%s", diff)
	}
}

func getResource(plugins []plugin.Plugin, nrn string) (resource any, err error) {
	nrnArr := strings.Split(nrn, ":")
	if len(nrnArr) != 7 || nrnArr[0] != "nrn" || nrnArr[1] != "nifcloud" {
		err = fmt.Errorf("The argument '%s' is not a valid NRN.", nrn)
		return
	}
	regions := []string{
		"jp-east-1",
		"jp-east-2",
		"jp-east-3",
		"jp-east-4",
		"jp-west-1",
		"jp-west-2",
		"us-east-1",
	}
	if !slices.Contains(regions, nrnArr[3]) {
		err = fmt.Errorf("The region '%s' is not available.", nrnArr[3])
		return
	}
	pluginResourceMap := map[string]map[string]plugin.Resource{}
	for i := 0; i < len(plugins); i++ {
		pluginName := plugins[i].Name()
		if pluginName == "base" {
			continue
		}
		pluginResourceMap[pluginName] = plugins[i].GetResourceMap()
	}
	for svcName, resourceMap := range pluginResourceMap {
		if svcName != nrnArr[2] {
			continue
		}
		for resourceType, r := range resourceMap {
			if resourceType != nrnArr[5] {
				continue
			}
			cfg := newConfigWithRegion(nrnArr[3])
			resource, err = r.Fetch(cfg, nrnArr[6])
			return
		}
		err = fmt.Errorf("The resource type '%s' is not a supported resource of computing service.", nrnArr[5])
		return
	}
	err = fmt.Errorf("The service '%s' is not available.", nrnArr[2])
	return
}

func newConfigWithRegion(region string) nifcloud.Config {
	return nifcloud.NewConfig(
		os.Getenv("NIFCLOUD_ACCESS_KEY_ID"),
		os.Getenv("NIFCLOUD_SECRET_ACCESS_KEY"),
		region,
	)
}

func getCmpOpts(plugins []plugin.Plugin) []cmp.Option {
	cmpOptsList := [][]cmp.Option{}
	for i := 0; i < len(plugins); i++ {
		for _, v := range plugins[i].GetResourceMap() {
			cmpOptsList = append(cmpOptsList, v.GetCmpOpts())
		}
	}
	optsLength := 0
	for i := 0; i < len(cmpOptsList); i++ {
		optsLength += len(cmpOptsList[i])
	}
	cmpOpts := make([]cmp.Option, optsLength)
	optIndex := 0
	for i := 0; i < len(cmpOptsList); i++ {
		for j := 0; j < len(cmpOptsList[i]); j++ {
			cmpOpts[optIndex] = cmpOptsList[i][j]
			optIndex++
		}
	}
	return cmpOpts
}
