package component

import (
	"time"
)

type ASCIIAnimation struct {
	Frames       []string
	currentFrame int
	lastUpdate   time.Time
	FrameDelay   time.Duration
}

func NewASCIIAnimation() *ASCIIAnimation {
	return &ASCIIAnimation{
		lastUpdate: time.Now(),
		FrameDelay: 150 * time.Millisecond,
	}
}

func (a *ASCIIAnimation) Update() {
	if a.Frames != nil {
		if time.Since(a.lastUpdate) > a.FrameDelay {
			a.currentFrame = (a.currentFrame + 1) % len(a.Frames)
			a.lastUpdate = time.Now()
		}
	}
}

func (a *ASCIIAnimation) StringOutput() string {
	if a.Frames == nil {
		return ""
	}
	return a.Frames[a.currentFrame]
}

func (a *ASCIIAnimation) SetFramesAndSpeed(frames []string, speed time.Duration) {
	a.Frames = frames
	a.FrameDelay = speed
}
