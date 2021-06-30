package podkillers

import (
    "context"
    
	"github.com/simskij/xk6-chaos/internal/k8s/pods"
    "github.com/simskij/xk6-chaos/pkg/k8s/client"
)



// This exposes podkiller metadata for use in displaying results.
type Podkillers struct {
    Ready bool
    pod *pods.Pods
}

var victims []string = []string{}
var numOfPods int

// New creates a new podkiller
func New(Ready bool) *Podkillers {
    c, err := client.New()
    p := pods.New(c)
    if err !=nil {
        return nil
    }
	return &Podkillers{Ready, p}
}

// AddVictim saves the names of pods selected to be terminated.
func (p *Podkillers) AddVictim(victim string) {
    victims = append(victims, victim)
}

// GetVictims returns the string array of all pods selected to be terminated.
func (p *Podkillers) GetVictims() []string {
    return victims
}

// SetStartingPods saves the number of pods at the beginning of a test as the "before" state.
func (p *Podkillers) SetStartingPods(number int) {
    numOfPods = number
}

// GetStartingPods returns the number of pods at the beginning of the test for reporting purposes.
func (p *Podkillers) GetStartingPods() int {
    return numOfPods
}

// KillPod terminates a k8s pod identified by name.
func (p *Podkillers) KillPod(namespace string, podName string) error {
    ctx := context.Background()
    err := p.pod.KillByName(ctx, namespace, podName)
    return err
} 