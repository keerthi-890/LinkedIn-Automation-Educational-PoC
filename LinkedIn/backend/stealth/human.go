package stealth

import (
	"math/rand"
	"time"
)

// Humanizer struct simulates human behavior
type Humanizer struct {
	Enabled bool
}

// NewHumanizer returns a new Humanizer
func NewHumanizer() *Humanizer {
	return &Humanizer{Enabled: true}
}

// SimulateTypingDelay simulates the time between keystrokes
func (h *Humanizer) SimulateTypingDelay() {
	if !h.Enabled {
		return
	}
	// WPM ~ 40-60 => 100ms - 250ms per key roughly
	base := 100 // ms
	variance := rand.Intn(150)
	time.Sleep(time.Duration(base+variance) * time.Millisecond)
}

// SimulateReadingDelay simulates reading content
func (h *Humanizer) SimulateReadingDelay() {
	if !h.Enabled {
		return
	}
	// 2-5 seconds
	SleepRandom(2000, 5000) // milliseconds
}

// SimulateThinkTime simulates deciding on an action
func (h *Humanizer) SimulateThinkTime() {
	if !h.Enabled {
		return
	}
	// 1-3 seconds
	SleepRandom(1000, 3000) // milliseconds
}

// SleepRandom simulates a human-like random delay in milliseconds
func SleepRandom(minMs, maxMs int) {
	rand.Seed(time.Now().UnixNano())
	ms := rand.Intn(maxMs-minMs+1) + minMs
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
