package computing

import (
	"context"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nifcloud/nifcloud-sdk-go/nifcloud"
	"github.com/nifcloud/nifcloud-sdk-go/service/computing"
	"github.com/nifcloud/nifcloud-sdk-go/service/computing/types"
)

type SecurityGroup struct{}

func (sg *SecurityGroup) GetCmpOpts() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(types.SecurityGroupInfo{}, "InstancesSet", "InstanceUniqueIdsSet"),
		cmpopts.IgnoreFields(types.IpPermissions{}, "AddDatetime"),
		cmpopts.SortSlices(func(i, j types.IpPermissions) bool {
			return sg.createRuleId(i) < sg.createRuleId(j)
		}),
	}
}

func (sg *SecurityGroup) createRuleId(v types.IpPermissions) string {
	source := ""
	for i := 0; i < len(v.IpRanges); i++ {
		source += *v.IpRanges[i].CidrIp
	}
	for i := 0; i < len(v.Groups); i++ {
		source += *v.Groups[i].GroupName
	}
	return fmt.Sprintf("%s_%s_%d_%d_%s", *v.InOut, *v.IpProtocol, *v.FromPort, *v.ToPort, source)
}

func (sg *SecurityGroup) Fetch(cfg nifcloud.Config, resourceName string) (resource any, err error) {
	svc := computing.NewFromConfig(cfg)
	groupName := []string{resourceName}
	resp, err := svc.DescribeSecurityGroups(context.TODO(), &computing.DescribeSecurityGroupsInput{GroupName: groupName})
	if err != nil {
		return
	}
	if len(resp.SecurityGroupInfo) == 0 {
		err = fmt.Errorf("Security group '%s' does not exist.", resourceName)
		return
	}
	resource = resp.SecurityGroupInfo[0]
	return
}
