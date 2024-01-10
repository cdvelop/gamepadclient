package gamepadclient

import "syscall/js"

func (g gamepadClient) connected(this js.Value, p []js.Value) interface{} {
	// g.gamepad = p[0].Get("gamepad")
	// g.Log("Gamepad connected", p[0].Get("gamepad"))

	if g.GamepadConfig != nil && g.GamepadConfig.Connected != nil {
		// g.Log("Gamepad connected")
		g.GamepadConfig.Connected()
	}

	return nil
}

func (g gamepadClient) disconnected(this js.Value, p []js.Value) interface{} {
	if g.GamepadConfig != nil && g.GamepadConfig.Disconnected != nil {
		// g.Log("Gamepad disconnected")
		g.GamepadConfig.Disconnected()
	}

	return nil
}
