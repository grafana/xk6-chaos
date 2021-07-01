package experiments

import (
	"github.com/simskij/xk6-chaos/internal/experiments/podkillers"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/chaos/experiments".
func init() {
	modules.Register("k6/x/chaos/experiments", &Experiments{
		Podkiller: podkillers.New(true),
	})
}

// This exposes experiment metadata for use in displaying results.
type Experiments struct {
	Podkiller *podkillers.Podkillers
}
