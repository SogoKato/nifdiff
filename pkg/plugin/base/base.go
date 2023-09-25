package base

import (
	"github.com/SogoKato/nifdiff/pkg/plugin"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nifcloud/nifcloud-sdk-go/nifcloud"
)

type BasePlugin struct{}

func (b *BasePlugin) Name() string {
	return "base"
}

func (b *BasePlugin) GetResourceMap() map[string]plugin.Resource {
	return map[string]plugin.Resource{
		"base": &BaseResouce{},
	}
}

type BaseResouce struct{}

func (b *BaseResouce) GetCmpOpts() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
	}
}

func (b *BaseResouce) Fetch(cfg nifcloud.Config, resourceName string) (resource any, err error) {
	return
}
