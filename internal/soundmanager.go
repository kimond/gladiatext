package internal

import (
	"bytes"
	"kimond/gladiatext/assets/sfx"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const sampleRate = 44100

type SoundManager struct {
	audioContext *audio.Context
	swingPlayer  *audio.Player
}

func (s *SoundManager) PlaySwing() {
	if s.swingPlayer == nil {
		wavData, err := wav.DecodeF32(bytes.NewReader(sfx.Swing_wav))
		if err != nil {
			log.Fatal(err)
		}
		player, err := s.audioContext.NewPlayerF32(wavData)
		if err != nil {
			log.Fatal(err)
		}
		s.swingPlayer = player
	}
	s.swingPlayer.Rewind()
	s.swingPlayer.Play()
}

var SFXManager *SoundManager

func init() {
	audioContext := audio.NewContext(sampleRate)

	SFXManager = &SoundManager{audioContext: audioContext}
}
