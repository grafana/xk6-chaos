package chaos

import (
	"github.com/loadimpact/k6/js/modules"
	_ "github.com/simskij/k6-extension-chaos/internal/k8s" // Register the k8s module as well
)

const version = "v0.0.1"

func init() {
	modules.Register("k6/x/chaos", &Chaos{
		Version: version,
	})
}

// Chaos is the main export of the chaos engineering extension
type Chaos struct {
	Version string
}