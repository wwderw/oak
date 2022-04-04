// Package window provides a common interface for oak-created windows.
package window

import (
	"image"

	"github.com/oakmound/oak/v3/alg/intgeom"
	"github.com/oakmound/oak/v3/event"
)

// Window is an interface of methods on an oak.Window
type Window interface {
	SetFullScreen(bool) error
	SetBorderless(bool) error
	SetTopMost(bool) error
	SetTitle(string) error
	SetIcon(image.Image) error
	MoveWindow(x, y, w, h int) error
	HideCursor() error

	Width() int
	Height() int
	Viewport() intgeom.Point2
	SetViewportBounds(intgeom.Rect2)

	NextScene()
	GoToScene(string)

	InFocus() bool
	ShiftScreen(int, int)
	SetScreen(int, int)
	Quit()

	EventHandler() event.Handler
	GetCallerMap() *event.CallerMap
}
