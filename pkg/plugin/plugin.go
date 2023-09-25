package plugin

import (
	"github.com/google/go-cmp/cmp"
	"github.com/nifcloud/nifcloud-sdk-go/nifcloud"
)

type Plugin interface {
	Name() string
	GetResourceMap() map[string]Resource
}

type Resource interface {
	GetCmpOpts() []cmp.Option
	Fetch(cfg nifcloud.Config, resourceName string) (resource any, err error)
}
