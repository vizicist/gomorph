package morph

import (
	"log"
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
	idx              uint8
	opened           bool
	serialNum        string
	width            float32
	height           float32
	fwVersionMajor   uint8
	fwVersionMinor   uint8
	fwVersionBuild   uint8
	fwVersionRelease uint8
	deviceID         int
}

// MaxForce Might need to be adjusted at some point
var MaxForce float32 = 1500.0

// CursorDeviceCallbackFunc xxx
type CursorDeviceCallbackFunc func(e CursorDeviceEvent)

func Init() ([]OneMorph, error) {
	// The initialize func is platform-specific,
	// See windowsmorph.go
	morphs, err := initialize()
	return morphs, err
}

// Start xxx
func Start(morphs []OneMorph, callback CursorDeviceCallbackFunc, forceFactor float32) {

	if len(morphs) == 0 {
		log.Printf("No Morphs were found\n")
		return
	}
	for {
		for _, m := range morphs {
			m.readFrames(callback, forceFactor)
		}
		time.Sleep(time.Millisecond)
	}
}
