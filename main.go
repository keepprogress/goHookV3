package main

import (
	"fmt"
	"time"

	hook "github.com/robotn/gohook"
)

var recording bool
var actions []string

func main() {
	fmt.Println("Press F2 to start recording, F10 to stop, and F4 to play.")

	// Start listening for global keyboard events
	go listenForGlobalKeys()

	// Keep the program running
	select {}
}

func listenForGlobalKeys() {
	// Start the hook
	chanHook := hook.Start()
	defer hook.End()

	for ev := range chanHook {
		switch ev.Kind {
		case hook.KeyHold:
			fmt.Printf("Key pressed: %d\n", ev.Keycode)
			handleKeyHold(int(ev.Keycode))
		}
	}
}

func handleKeyHold(keycode int) {

	if recording {
		actions = append(actions, fmt.Sprintf("Key: %d", keycode))
		fmt.Printf("Recorded key: %d\n", keycode)
	}

	// Check for specific keys to control recording and playback
	switch keycode {
	case 60: // F2 key
		startRecording()
	case 68: // F10 key
		stopRecording()
	case 62: // F4 key
		playRecording()
	}
}

func startRecording() {
	recording = true
	actions = []string{}
	fmt.Println("Recording started...")
}

func stopRecording() {
	recording = false
	fmt.Println("Recording stopped.")
}

func playRecording() {
	fmt.Println("Playing back recorded actions...")
	for _, action := range actions {
		fmt.Println(action)                // Display the action
		time.Sleep(200 * time.Millisecond) // Delay between actions
	}
	fmt.Println("Playback finished.")
}
