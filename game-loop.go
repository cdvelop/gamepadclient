package gamepadclient

import "syscall/js"

func (g *gamepadClient) gameLoop(this js.Value, p []js.Value) interface{} {
	g.gamepads = js.Global().Get("navigator").Call("getGamepads")

	for i := 0; i < g.gamepads.Length(); i++ {
		g.gamepad = g.gamepads.Index(i)
		if g.gamepad.Truthy() && g.gamepad.Get("connected").Bool() {
			g.handleButtons(g.gamepad.Get("buttons"))
		}
	}

	js.Global().Call("requestAnimationFrame", js.Global().Get("gameLoop"))

	return nil
}
