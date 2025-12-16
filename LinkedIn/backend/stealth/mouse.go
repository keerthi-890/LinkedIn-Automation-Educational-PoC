package stealth

import (
	"math"
	"math/rand"
	"time"
)

// MoveMouseHuman simulates human-like mouse movement using Bezier math (simulation only)
func MoveMouseHuman() {
	rand.Seed(time.Now().UnixNano())

	// Simulated Bezier curve calculation (no real mouse interaction)
	points := rand.Intn(5) + 3
	for i := 0; i < points; i++ {
		_ = bezier(float64(i) / float64(points))
		time.Sleep(time.Duration(rand.Intn(40)+20) * time.Millisecond)
	}
}

// bezier curve helper (simulation)
func bezier(t float64) float64 {
	p0 := rand.Float64()
	p1 := rand.Float64()
	p2 := rand.Float64()

	return math.Pow(1-t, 2)*p0 + 2*(1-t)*t*p1 + math.Pow(t, 2)*p2
}
