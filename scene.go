package oak

import (
	"time"

	"github.com/oakmound/oak/v2/scene"
	"github.com/oakmound/oak/v2/timing"
)

var (
	// SceneMap is a global map of scenes referred to when scenes advance to
	// determine what the next scene should be.
	// It can be replaced or modified so long as these modifications happen
	// during a scene or before oak has started.
	SceneMap = scene.NewMap()
)

func init() {
	// The scene "loading" is reserved
	err := SceneMap.AddScene("loading", loadingScene)
	if err != nil {
		panic(err)
	}
}

// AddScene is shorthand for oak.SceneMap.AddScene
func AddScene(name string, s scene.Scene) error {
	return SceneMap.AddScene(name, s)
}

// Add is shorthand for oak.SceneMap.Add
func Add(name string,
	start func(context *scene.Context),
	loop func() (cont bool),
	end func() (nextScene string, result *scene.Result)) error {

	return AddScene(name, scene.Scene{Start: start, Loop: loop, End: end})
}

func sceneTransition(result *scene.Result) {
	if result.Transition != nil {
		i := 0
		tx, _ := screenControl.NewTexture(winBuffer.Bounds().Max)
		cont := true
		for cont {
			cont = result.Transition(winBuffer.RGBA(), i)
			drawLoopPublish(tx)
			i++
			time.Sleep(timing.FPSToDuration(FrameRate))
		}
	}
}
