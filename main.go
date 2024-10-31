package main

import (
	"fmt"
	"time"

	robotgo "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var recording bool
var actions []string
var lastKeyTime time.Time

func main() {
	fmt.Println("Press F2 to start recording, F10 to stop, and F4 to play.")

	// Start listening for global keyboard events
	go listenForGlobalKeys()

	// Keep the program running indefinitely
	select {}
}

// listenForGlobalKeys listens for global keyboard events and processes them.
func listenForGlobalKeys() {
	// Start the hook
	chanHook := hook.Start()
	defer hook.End()

	for ev := range chanHook {
		if ev.Kind == hook.KeyHold {
			fmt.Printf("Key pressed: %d\n", ev.Keycode)
			handleKeyHold(int(ev.Keycode))
		}
	}
}

// handleKeyHold processes the key hold events and manages recording/playback.
func handleKeyHold(keycode int) {
	currentTime := time.Now()

	if recording {
		// Calculate delay since the last key press
		delay := 0
		if !lastKeyTime.IsZero() {
			delay = int(currentTime.Sub(lastKeyTime).Milliseconds())
		}
		// Store the keycode and delay in a simpler format
		actions = append(actions, fmt.Sprintf("%d,%d", keycode, delay))
		fmt.Printf("Recorded key: %d, Delay: %d ms\n", keycode, delay)
	}

	// Update the last key time
	lastKeyTime = currentTime

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

// startRecording initializes the recording process.
func startRecording() {
	recording = true
	actions = []string{}      // Clear previous actions
	lastKeyTime = time.Time{} // Reset last key time
	fmt.Println("Recording started...")
}

// stopRecording ends the recording process.
func stopRecording() {
	recording = false
	fmt.Println("Recording stopped.")
}

// playRecording plays back the recorded actions with the recorded delays.
func playRecording() {
	recording = false
	fmt.Println("Playing back recorded actions...")

	for _, action := range actions {
		var keycode, delay int
		fmt.Sscanf(action, "%d,%d", &keycode, &delay) // Extract keycode and delay

		// Simulate the delay for playback
		time.Sleep(time.Duration(delay) * time.Millisecond)

		// Simulate the key press using robotgo
		robotgo.KeyTap(getKeyString(keycode)) // Convert keycode to string

		// Display the action (simulating pressing the key)
		fmt.Printf("Simulated Key Press: %d\n", keycode)
	}

	fmt.Println("Playback finished.")
}

// getKeyString converts a keycode to its corresponding string representation.
func getKeyString(keycode int) string {
	for k, v := range hook.Keycode {
		if int(v) == keycode {
			return k
		}
	}
	return ""
}
