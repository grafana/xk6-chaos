package k8s

import (
	"github.com/dop251/goja"
	"github.com/grafana/xk6-chaos/internal/k8s/pods"
	"github.com/grafana/xk6-chaos/pkg/k8s/client"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/chaos/k8s", K8sRoot{})
}

var _ modules.Module = K8sRoot{}

type K8sRoot struct{}

func (k K8sRoot) NewModuleInstance(vu modules.VU) modules.Instance {
	return &K8s{
		vu: vu,
	}
}

// K8s exports all kubernetes related APIs
type K8s struct {
	vu   modules.VU
	Pods *pods.Pods
}

func (k *K8s) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Pods": k.XPods,
		},
	}
}

// XPods serves as a constructor of the Pods js class
func (k *K8s) XPods(goja.ConstructorCall) *goja.Object {
	rt := k.vu.Runtime()
	c, err := client.New()
	if err != nil {
		common.Throw(rt, err)
	}
	p := pods.New(c)
	return rt.ToValue(p).ToObject(rt)
}
