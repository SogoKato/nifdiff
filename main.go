package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/nifcloud/nifcloud-sdk-go/nifcloud"
	nifcloudnrn "github.com/nifcloud/nifcloud-sdk-go/nifcloud/nrn"
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

func getResource(plugins []plugin.Plugin, nrnString string) (resource any, err error) {
	nrn, err := nifcloudnrn.Parse(nrnString)
	if err != nil {
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
	if !slices.Contains(regions, nrn.Region) {
		err = fmt.Errorf("The region '%s' is not available.", nrn.Region)
		return
	}
	resourceSplitted := strings.SplitN(nrn.Resource, ":", 2)
	resourceType := resourceSplitted[0]
	resourceName := resourceSplitted[1]
	pluginResourceMap := map[string]map[string]plugin.Resource{}
	for i := 0; i < len(plugins); i++ {
		pluginName := plugins[i].Name()
		if pluginName == "base" {
			continue
		}
		pluginResourceMap[pluginName] = plugins[i].GetResourceMap()
	}
	for svcName, resourceMap := range pluginResourceMap {
		if svcName != nrn.Service {
			continue
		}
		for rt, r := range resourceMap {
			if rt != resourceType {
				continue
			}
			cfg := newConfigWithRegion(nrn.Region)
			resource, err = r.Fetch(cfg, resourceName)
			return
		}
		err = fmt.Errorf("The resource type '%s' is not a supported resource of computing service.", resourceType)
		return
	}
	err = fmt.Errorf("The service '%s' is not available.", nrn.Service)
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
