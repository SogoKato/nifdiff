package computing

import "github.com/SogoKato/nifdiff/pkg/plugin"

type ComputingPlugin struct{}

func (c *ComputingPlugin) Name() string {
	return "computing"
}

func (c *ComputingPlugin) GetResourceMap() map[string]plugin.Resource {
	return map[string]plugin.Resource{
		"security_group": &SecurityGroup{},
	}
}
