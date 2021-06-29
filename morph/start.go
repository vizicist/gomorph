package morph

import (
	"fmt"
	"time"
)

var DebugMorph bool = false

// CursorDeviceEvent is a single CursorDevice event
type CursorDeviceEvent struct {
	CID       string
	Timestamp int64  // milliseconds
	Ddu       string // "down", "drag", "up"
	X         float32
	Y         float32
	Z         float32
	Area      float32
}

// CursorDown etc match values in sensel.h
const (
	CursorDown = 1
	CursorDrag = 2
	CursorUp   = 3
)

// OneMorph is a single Morph
type OneMorph struct {
	Idx              uint8
	Opened           bool
	SerialNum        string
	Width            float32
	Height           float32
	FwVersionMajor   uint8
	FwVersionMinor   uint8
	FwVersionBuild   uint8
	FwVersionRelease uint8
	DeviceID         int
}

// MaxForce Might need to be adjusted at some point
var MaxForce float32 = 1500.0

// CursorDeviceCallbackFunc xxx
type CursorDeviceCallbackFunc func(e CursorDeviceEvent)

func Init(serial string) ([]OneMorph, error) {
	// The initialize func is platform-specific,
	// See windowsmorph.go
	return initialize(serial)
}

// Start xxx
func Start(morphs []OneMorph, callback CursorDeviceCallbackFunc, forceFactor float32) error {
	if len(morphs) == 0 {
		return fmt.Errorf("morph.Start: morphs array is empty!?")
	}
	for {
		for _, m := range morphs {
			m.readFrames(callback, forceFactor)
		}
		time.Sleep(time.Millisecond)
	}
}
