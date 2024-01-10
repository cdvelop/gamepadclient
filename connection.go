package gamepadclient

import "syscall/js"

func (g *gamepadClient) Connected(this js.Value, p []js.Value) interface{} {
	// g.gamepad = p[0].Get("gamepad")
	js.Global().Get("console").Call("log", "Gamepad Connected")
	g.connected = true

	if g.notifyConnection != nil {
		g.GamepadConnected()
	}

	return nil
}

func (g *gamepadClient) Disconnected(this js.Value, p []js.Value) interface{} {
	js.Global().Get("console").Call("log", "Gamepad Disconnected")
	g.connected = false
	if g.notifyConnection != nil {
		g.GamepadDisconnected()
	}

	return nil
}
