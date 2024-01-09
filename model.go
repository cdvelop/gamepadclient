package gamepadclient

import "syscall/js"

type gamepadClient struct {
	gamepads js.Value
	gamepad  js.Value

	previousButtonState []bool
}
