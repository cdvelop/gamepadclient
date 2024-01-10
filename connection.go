package gamepadclient

import "syscall/js"

func (g *gamepadClient) Connected(this js.Value, p []js.Value) interface{} {
	// g.gamepad = p[0].Get("gamepad")
	// g.Log("Gamepad connected", p[0].Get("gamepad"))
	js.Global().Get("console").Call("log", "Gamepad connected")
	g.connected = true

	if g.GamepadConfig != nil && g.GamepadConfig.Connected != nil {
		// g.Log("Gamepad connected")
		g.GamepadConfig.Connected()
	}

	return nil
}

func (g *gamepadClient) Disconnected(this js.Value, p []js.Value) interface{} {
	js.Global().Get("console").Call("log", "Gamepad disconnected")
	g.connected = false
	if g.GamepadConfig != nil && g.GamepadConfig.Disconnected != nil {
		// g.Log("Gamepad disconnected")
		g.GamepadConfig.Disconnected()
	}

	return nil
}
