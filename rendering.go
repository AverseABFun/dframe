package dframe

type RawRenderer interface {
	InitRenderer(windowName string, virtSize IntPoint, realSize IntPoint) error // Initalizes the renderer. The virtSize should be the size of the buffer, and realSize the size of the buffer shown to the user. It should scale linearly between them. No methods should function before this method is called, or after DeinitRenderer() is called.
	GetVirtSize() IntPoint                                                      // Returns the VIRTUAL size of the renderer.
	GetRealSize() IntPoint                                                      // Returns the REAL size of the renderer.
	TickRenderer()                                                              // Ticks the renderer. It should NOT update the display until this is called.
	ShouldQuit() bool                                                           // Returns whether the program should quit. This generally changes to true when the user closes the window.
	DeinitRenderer() error                                                      // Deinitalizes the renderer. This should stop all rendering and after calling this, if the renderer has not been reinitalized, no other method should function.
	DrawPixel(x uint32, y uint32, color Color) error                            // Draws a pixel to the screen. This is in VIRTUAL screen space, so after calling TickRenderer() it should scale all pixels that have been drawn with this or SetBack() to the proper real size.
	SetBack(color Color) error                                                  // Sets the background to the provided color. This should be the default color if a color has not been set with DrawPixel(), and it should overwrite all pixels set with DrawPixel() before calling this method.
}

type AudioRenderer interface {
	InitAudio() error   // Initalizes audio. No method should function before this is called.
	DeinitAudio() error // Deinitalizes audio. No method should function after this is called.

	PlayTone(freqHz float64) (AudioID, error) // Starts playing a pure sine tone with the provide frequency.
	Playing(tone AudioID) bool                // Returns if a certain tone is currently playing.
	PauseTone(tone AudioID) error             // Pauses a tone. It can be resumed with ResumeTone() later.
	ResumeTone(tone AudioID) error            // Resumes a paused tone.
	StopTone(tone AudioID) error              // Stops a tone. The tone should NOT be able to be resumed with ResumeTone().
	StopAll() error                           // Stops ALL tones. None of them can be resumed with ResumeTone().
}

type AudioID uint64 // The identifier of a tone. It will only work for the [AudioRenderer] it was created with.

type StackRenderer interface {
	RawParent() RawRenderer      // Returns the current parent [RawRenderer].
	SetRawParent(rr RawRenderer) // Sets the current parent [RawRenderer].

	StackParent() StackRenderer      // Returns the current parent [StackRenderer].
	SetStackParent(sp StackRenderer) // Sets the current parent [StackRenderer].

	AudioParent() AudioRenderer      // Returns the current parent [AudioRenderer].
	SetAudioParent(ap AudioRenderer) // Sets the current parent [AudioRenderer].

	CanUseCurrentStackRenderer() bool // Returns if it can use the current [StackRenderer]. The reason this method exists but methods for the parent [RawRenderer] and [AudioRenderer] don't is because StackRenderers can(and should!) have other functions that provide extended functionality, but RawRenderers and AudioRenderers generally shouldn't.
}
