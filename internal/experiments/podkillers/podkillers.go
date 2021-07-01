package podkillers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/simskij/xk6-chaos/internal/k8s/pods"
	"github.com/simskij/xk6-chaos/pkg/k8s/client"
)

// This exposes podkiller metadata for use in displaying results.
type Podkillers struct {
	Ready           bool
	Id              int    `json:"id"`
	ExperimentType  string `json:"expType"`
	NumOfPodsBefore int    `json:"numPodsBefore"`
	NumOfPodsAfter  int    `json:"numPodsAfter"`
	Victims         string `json:"victims"`
	pod             *pods.Pods
}

var victims []string = []string{}

// var numOfPodsBefore int

// New creates a new podkiller
func New(Ready bool, ExperimentNum int) *Podkillers {
	c, err := client.New()
	p := pods.New(c)
	if err != nil {
		return nil
	}
	// var expType = "Podkiller"
	// var desc = ""
	podkiller := Podkillers{
		Ready:           true,
		Id:              ExperimentNum,
		ExperimentType:  "Pod termination",
		NumOfPodsBefore: 0,
		NumOfPodsAfter:  0,
		Victims:         "",
		pod:             p,
	}
	// return &Podkillers{Ready, ExperimentNum, expType, desc, p}
	return &podkiller
}

// AddVictim saves the names of pods selected to be terminated.
func (p *Podkillers) AddVictim(victim string) {
	victims = append(victims, victim)
	p.Victims = p.Victims + victim
}

// GetVictims returns the string array of all pods selected to be terminated.
func (p *Podkillers) GetVictims() []string {
	return victims
}

// SetStartingPods saves the number of pods at the beginning of a test as the "before" state.
func (p *Podkillers) SetStartingPods(number int) {
	p.NumOfPodsBefore = number
}

// GetStartingPods returns and saves the number of pods at the beginning of the test for reporting purposes.
func (p *Podkillers) GetStartingPods(namespace string) int {
	var podsAlive, _ = p.pod.List(context.Background(), namespace)
	p.NumOfPodsBefore = len(podsAlive)
	// p.SetStartingPods(p.NumOfPodsBefore)
	return p.NumOfPodsBefore
}

// GetNumOfPods returns and saves the current number of pods.
func (p *Podkillers) GetNumOfPods(namespace string) int {
	time.Sleep(5 * time.Second)
	var podsAlive, _ = p.pod.List(context.Background(), namespace)
	p.NumOfPodsAfter = len(podsAlive)
	fmt.Println("Num of pods after: " + strconv.Itoa(p.NumOfPodsAfter))
	return p.NumOfPodsAfter
}

// KillPod terminates a k8s pod identified by name.
func (p *Podkillers) KillPod(namespace string, podName string) error {
	p.GetStartingPods(namespace)
	err := p.pod.KillByName(context.Background(), namespace, podName)
	p.AddVictim(podName)
	p.GetNumOfPods(namespace)
	return err
}

// KillPodLike terminates a k8s pod whose name contains string.
func (p *Podkillers) KillPodLike(namespace string, keyword string) error {
	p.GetStartingPods(namespace)
	podName, err := p.pod.KillByKeyword(context.Background(), namespace, keyword)
	p.AddVictim(podName)
	p.GetNumOfPods(namespace)
	return err
}

// KillRandomPod terminates a pod at random.
func (p *Podkillers) KillRandomPod(namespace string) error {
	p.GetStartingPods(namespace)
	podName, err := p.pod.KillRandom(context.Background(), namespace)
	p.AddVictim(podName)
	p.GetNumOfPods(namespace)
	return err
}
