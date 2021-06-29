package experiments

import (
    "go.k6.io/k6/js/modules"
	"github.com/simskij/xk6-chaos/internal/experiments/podkillers"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/chaos/experiments".
func init() {
	modules.Register("k6/x/chaos/experiments", &Experiments{
		Podkillers: &podkillers.Podkillers{Ready: true},
	})
}

// This exposes experiment metadata for use in displaying results.
type Experiments struct {
    Podkillers *podkillers.Podkillers
}