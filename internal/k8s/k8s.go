package k8s

import (
	"context"

	"github.com/grafana/xk6-chaos/internal/k8s/pods"
	"github.com/grafana/xk6-chaos/pkg/k8s/client"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/chaos/k8s", &K8s{
		Pods: &pods.Pods{},
	})
}

// K8s exports all kubernetes related APIs
type K8s struct {
	Pods *pods.Pods
}

// XPods serves as a constructor of the Pods js class
func (*K8s) XPods(ctx *context.Context) (interface{}, error) {
	rt := common.GetRuntime(*ctx)
	c, err := client.New()
	if err != nil {
		return nil, err
	}
	p := pods.New(c)
	return common.Bind(rt, p, ctx), nil
}
