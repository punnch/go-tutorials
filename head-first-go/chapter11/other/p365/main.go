package main

import (
	"fmt"

	"github.com/punnch/gadgets"
)

type Player interface {
	Play(string)
	Stop()
}

// Test function
func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadgets.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

// Get songs and call play, stop methods func
func Playlist(device Player, songs []string) {
	fmt.Printf("%v is turned on:\n", device)
	for _, s := range songs {
		device.Play(s)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Отвратительный король", "Овердоз", "Цветы"}
	var device Player

	// Tape Recorder
	device = gadgets.TapeRecorder("Some Recorder")
	Playlist(device, mixtape)

	// Tape Player
	device = gadgets.TapePlayer("Some Player")
	Playlist(device, mixtape)

	// Test of both
	TryOut(gadgets.TapePlayer("test player"))
	TryOut(gadgets.TapeRecorder("test recorder"))
}
