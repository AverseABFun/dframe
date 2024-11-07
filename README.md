# dframe: drawing, input, and audio framework

dframe is a framework that provides abstractions and implementations of drawing to the screen(via a pixel buffer), getting input, and outputting audio for go programs. The current implementations provided include OpenGL for drawing and input(via `github.com/go-gl/gl/v4.6-core/gl` and `github.com/go-gl/glfw/v3.3/glfw`), `github.com/gopxl/beep/v2` for audio, a non-functional implementation of `github.com/gordonklaus/portaudio`, also for audio, and a whole host of interfaces and types for all of things mentioned earlier.

Please note that all features(audio in particular) are very limited as they are meant for programs MEANT to be retro. For example, the audio interface only supports playing pure sine tones with a static amplitude(however, you can play many at once), and nothing else, and the screen renderer only supports drawing pure pixels(excluding stack renderers).
