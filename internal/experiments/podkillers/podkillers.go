package podkillers

// This exposes podkill metadata for use in displaying results.
type Podkillers struct {
    Ready bool
}

// New creates a new podkillers struct
// func New(Victim string) *Podkillers {
//     return &podkiller{Victim}
// }

var text []string = []string{}

func (p *Podkillers) AddVictim(victim string) {
    text = append(text, victim)
}

func (p *Podkillers) GetVictims() []string {
    return text
}