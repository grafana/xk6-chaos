package podkillers

import (
	"context"

	"github.com/simskij/xk6-chaos/internal/k8s/pods"
	"github.com/simskij/xk6-chaos/pkg/k8s/client"
)

// This exposes podkiller metadata for use in displaying results.
type Podkillers struct {
	Ready bool
	pod   *pods.Pods
}

var victims []string = []string{}
var numOfPodsBefore int
var numOfPodsAfter int

// New creates a new podkiller
func New(Ready bool) *Podkillers {
	c, err := client.New()
	p := pods.New(c)
	if err != nil {
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
	numOfPodsBefore = number
}

// GetStartingPods returns and saves the number of pods at the beginning of the test for reporting purposes.
func (p *Podkillers) GetStartingPods(namespace string) int {
	var podsAlive, _ = p.pod.List(context.Background(), namespace)
	numOfPodsBefore = len(podsAlive)
	p.SetStartingPods(numOfPodsBefore)
	return numOfPodsBefore
}

// GetNumOfPods returns and saves the current number of pods.
func (p *Podkillers) GetNumOfPods(namespace string) int {
	var podsAlive, _ = p.pod.List(context.Background(), namespace)
	numOfPods := len(podsAlive)
	return numOfPods
}

// KillPod terminates a k8s pod identified by name.
func (p *Podkillers) KillPod(namespace string, podName string) error {
	p.GetStartingPods(namespace)
	err := p.pod.KillByName(context.Background(), namespace, podName)
	p.AddVictim(podName)
	return err
}

// KillPodLike terminates a k8s pod whose name contains string.
func (p *Podkillers) KillPodLike(namespace string, keyword string) error {
	p.GetStartingPods(namespace)
	podName, err := p.pod.KillByKeyword(context.Background(), namespace, keyword)
	p.AddVictim(podName)
	return err
}

// KillRandomPod terminates a pod at random.
func (p *Podkillers) KillRandomPod(namespace string) error {
	p.GetStartingPods(namespace)
	podName, err := p.pod.KillRandom(context.Background(), namespace)
	p.AddVictim(podName)
	return err
}
