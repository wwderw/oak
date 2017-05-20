package particle

import (
	"bitbucket.org/oakmoundstudio/oak/physics"
	"github.com/200sc/go-dist/floatrange"
	"github.com/200sc/go-dist/intrange"
)

var (
	// Inf represents Infinite duration
	Inf = intrange.NewInfinite()
)

type Generator interface {
	GetBaseGenerator() *BaseGenerator
	GenerateParticle(*BaseParticle) Particle
	Generate(int) *Source
	GetParticleSize() (float64, float64, bool)
	ShiftX(float64)
	ShiftY(float64)
	SetPos(float64, float64)
	GetPos() (float64, float64)
}

// A BaseGenerator fulfills all basic requirements to generate
// particles
// Modeled after Parcycle
type BaseGenerator struct {
	physics.Vector
	// This float is currently forced to an integer
	// at new particle rotation. This should be changed
	// to something along the lines of 'new per 30 frames',
	// or allow low fractional values to be meaningful,
	// so that more fine-tuned particle generation speeds are possible.
	NewPerFrame floatrange.Range
	// The number of frames each particle should persist
	// before being removed.
	LifeSpan floatrange.Range
	// 0 - between quadrant 1 and 4
	// 90 - between quadrant 2 and 1
	Angle  floatrange.Range
	Speed  floatrange.Range
	Spread physics.Vector
	// Duration in milliseconds for the particle source.
	// After this many milliseconds have passed, it will
	// stop sending out new particles. Old particles will
	// not be removed until their individual lifespans run
	// out.
	// A duration of -1 represents never stopping.
	Duration intrange.Range
	// Rotational acceleration, to change angle over time
	Rotation floatrange.Range
	// Gravity X and Gravity Y represent particle acceleration per frame.
	Gravity    physics.Vector
	SpeedDecay physics.Vector
	EndFunc    func(Particle)
	LayerFunc  func(physics.Vector) int
}

func (bg *BaseGenerator) SetDefaults() {
	*bg = BaseGenerator{
		Vector:      physics.NewVector(0, 0),
		NewPerFrame: floatrange.Constant(1),
		LifeSpan:    floatrange.Constant(60),
		Angle:       floatrange.Constant(0),
		Speed:       floatrange.Constant(1),
		Spread:      physics.NewVector(0, 0),
		Duration:    Inf,
		Rotation:    nil,
		Gravity:     physics.NewVector(0, 0),
		SpeedDecay:  physics.NewVector(0, 0),
		EndFunc:     nil,
		LayerFunc:   func(physics.Vector) int { return 1 },
	}
}

func (bg *BaseGenerator) ShiftX(x float64) {
	bg.Vector = bg.Vector.ShiftX(x)
}
func (bg *BaseGenerator) ShiftY(y float64) {
	bg.Vector = bg.Vector.ShiftY(y)
}
func (bg *BaseGenerator) SetPos(x, y float64) {
	bg.Vector.X = x
	bg.Vector.Y = y
}
