package experiments

import (
	"context"

	"github.com/simskij/xk6-chaos/internal/experiments/podkillers"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

var ChaosSummary string

// Register the extension on module initialization, available to
// import from JS as "k6/x/chaos/experiments".
func init() {
	modules.Register("k6/x/chaos/experiments", &Experiments{
		// Podkiller: podkillers.New(),
	})
	// i++
}

// This exposes experiment metadata for use in displaying results.
type Experiments struct {
	Podkiller *podkillers.Podkillers
	Summary   string
}

// XPodkillers serves as a constructor of the Podkillers js class
func (*Experiments) XPodkillers(ctx *context.Context) (interface{}, error) {
	rt := common.GetRuntime(*ctx)
	p := podkillers.New()
	return common.Bind(rt, p, ctx), nil
}

func (*Experiments) GenerateChaosSummary() string {
	return ChaosSummary
}
