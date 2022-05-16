package experiments

import (
	"github.com/dop251/goja"
	"github.com/grafana/xk6-chaos/internal/experiments/podkillers"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/chaos/experiments".
func init() {
	modules.Register("k6/x/chaos/experiments", &ExperimentsRoot{
		// Podkiller: podkillers.New(),
	})
	// i++
}

var _ modules.Module = ExperimentsRoot{}

type ExperimentsRoot struct{}

func (k ExperimentsRoot) NewModuleInstance(vu modules.VU) modules.Instance {
	return &Experiments{
		vu: vu,
	}
}

// This exposes experiment metadata for use in displaying results.
type Experiments struct {
	Podkiller *podkillers.Podkillers
	Summary   string
	vu        modules.VU
}

func (e *Experiments) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Podkillers": e.XPodkillers,
		},
	}
}

// XPodkillers serves as a constructor of the Podkillers js class
func (e *Experiments) XPodkillers(goja.ConstructorCall) *goja.Object {
	rt := e.vu.Runtime()
	p := podkillers.New()
	return rt.ToValue(p).ToObject(rt)
}
