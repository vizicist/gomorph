package main

import (
	"log"

	"github.com/vizicist/gomorph/morph"
)

func main() {
	morphs, err := morph.Init()
	if err != nil {
		log.Printf("gomorph: err=%s\n", err)
	}
	morph.Start(morphs, handleMorph, 1.0)
	log.Printf("gomorph: end\n")
}

func handleMorph(e morph.CursorDeviceEvent) {
	log.Printf("cursor %s %s %f %f %f %f\n", e.CID, e.Ddu, e.X, e.Y, e.Z, e.Area)
}
