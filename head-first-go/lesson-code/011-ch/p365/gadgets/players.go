package gadgets

import "fmt"

// Tape Recorder
type TapeRecorder string

func (t TapeRecorder) Play(song string) {
	fmt.Printf("Playing '%s'\n", song)
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stop!")
}
func (t TapeRecorder) Record() {
	fmt.Println("Track is recorded!")
}

// Tape Player
type TapePlayer string

func (t TapePlayer) Play(song string) {
	fmt.Printf("Playing '%s'\n", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stop!")
}
