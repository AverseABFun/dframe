package dframe

import "github.com/go-gl/glfw/v3.3/glfw"

type KeyGrabber interface {
	GrabKey(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) (continueSearching bool) // This method is called when a [KeyProvider] is searching for [KeyGrabber]s in its grabber stack. It should return if the KeyProvider should continue searching.
}

type MouseGrabber interface {
	GrabMouse(button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey, posX float64, posY float64) (continueSearching bool) // This method is called when a [MouseProvider] is searching for [MouseGrabber]s in its grabber stack. It should return if the MouseProvider should continue searching.
}

type InputProvider interface {
	ProcessInputs() error // Processes all inputs. However, applications should NOT expect input events to only happen after this is called.
}

type KeyProvider interface {
	PushGrabber(grabber KeyGrabber) (index uint32)  // Pushes a [KeyGrabber] onto the stack.
	PopGrabber() (KeyGrabber, error)                // Pops a [KeyGrabber] off the stack.
	PushGrabberAt(grabber KeyGrabber, index uint32) // Pushes a [KeyGrabber] onto the stack at a certain location.
	PopGrabberAt(index uint32) (KeyGrabber, error)  // Pops a [KeyGrabber] off the stack from a certain location.
	InputProvider
}

type MouseProvider interface {
	PushMouseGrabber(grabber MouseGrabber) (index uint32)  // Pushes a [MouseGrabber] onto the stack.
	PopMouseGrabber() (MouseGrabber, error)                // Pops a [MouseGrabber] off the stack.
	PushMouseGrabberAt(grabber MouseGrabber, index uint32) // Pushes a [MouseGrabber] onto the stack at a certain location.
	PopMouseGrabberAt(index uint32) (MouseGrabber, error)  // Pops a [MouseGrabber] off the stack from a certain location.
	InputProvider
}
