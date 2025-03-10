package component

import (
	"math/rand"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type LogViewer struct {
	Title        string
	Log          []string
	displayedLog string
	lastUpdate   time.Time
	charIndex    int
	speed        time.Duration
	currentSpeed time.Duration
}

func NewLogViewer() *LogViewer {
	return &LogViewer{
		Log:        []string{},
		speed:      50 * time.Millisecond,
		lastUpdate: time.Now(),
	}
}

func (l *LogViewer) SetLogs(logs []string) {
	l.Log = logs
	l.displayedLog = ""
	l.charIndex = 0
	l.lastUpdate = time.Now()
	l.randomizeSpeed()
}

func (l *LogViewer) SetLog(log string) {
	l.Log = []string{log}
	l.displayedLog = ""
	l.charIndex = 0
	l.lastUpdate = time.Now()
	l.randomizeSpeed()
}

func (l *LogViewer) AddLog(log string) {
	l.Log = append(l.Log, log)
}

func (l *LogViewer) ClearLog() {
	l.Log = []string{}
	l.displayedLog = ""
	l.charIndex = 0
}

func (l *LogViewer) Update() error {
	if len(l.Log) == 0 {
		return nil
	}

	fullText := strings.Join(l.Log, "\n")
	if l.charIndex < len(fullText) {
		if time.Since(l.lastUpdate) > l.currentSpeed {
			step := rand.Intn(3) + 1
			if l.charIndex+step > len(fullText) {
				step = len(fullText) - l.charIndex
			}
			l.charIndex += step
			l.displayedLog = fullText[:l.charIndex]

			l.randomizeSpeed()
			l.lastUpdate = time.Now()
		}
	}

	return nil
}

func (l *LogViewer) StringOutput(screen *ebiten.Image) string {
	charWidth := 7
	screenWidth := screen.Bounds().Dx()
	lineLenth := screenWidth / charWidth
	line := strings.Repeat("-", lineLenth)

	var titleLine string
	if l.Title != "" {
		title := " " + l.Title + " "
		titleLen := len(title)

		if titleLen > lineLenth {
			title = title[:lineLenth]
			titleLen = len(title)
		}

		// Calculate left and right padding for centering
		leftPad := (lineLenth - titleLen) / 2
		rightPad := lineLenth - titleLen - leftPad

		titleLine = strings.Repeat("-", leftPad) + title + strings.Repeat("-", rightPad)
	} else {
		titleLine = line
	}

	output := strings.Builder{}
	output.WriteString(titleLine + "\n")
	output.WriteString(l.displayedLog + "\n")
	output.WriteString(line)

	return output.String()
}

func (l *LogViewer) Draw(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.LineSpacing = 20
	op.GeoM.Translate(float64(screen.Bounds().Min.X)+10, float64(screen.Bounds().Min.Y+200))
	text.Draw(screen, l.StringOutput(screen), text.NewGoXFace(basicfont.Face7x13), op)
}

func (l *LogViewer) randomizeSpeed() {
	l.currentSpeed = l.speed + time.Duration(rand.Intn(50)-25)*time.Millisecond
}
