package domain

type TransformationType string
type Transformation func()

const (
	Swirl     TransformationType = "swirl"
	Horseshoe TransformationType = "horseshoe"
)

var AvailableTransformations = map[TransformationType]Transformation{
	Swirl:     swirl,
	Horseshoe: horseshoe,
}

func GetTransformation(tt TransformationType) (Transformation, bool) {
	f, ok := AvailableTransformations[tt]
	return f, ok
}

func swirl() {
}

func horseshoe() {
}
