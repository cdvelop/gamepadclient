package gamepadclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

var gamepad *gamepadClient

func AddGamePadHandler(l model.Logger) (g *gamepadClient) {

	if gamepad == nil {

		gamepad = &gamepadClient{
			Logger: l,
		}

		js.Global().Set("gamepadButtonHandler", js.FuncOf(gamepad.gamepadButtonHandler))
		js.Global().Get("window").Call("addEventListener", "gamepadconnected", js.FuncOf(gamepad.connected))
		js.Global().Get("window").Call("addEventListener", "gamepaddisconnected", js.FuncOf(gamepad.disconnected))

	}

	return gamepad
}

func (g gamepadClient) connected(this js.Value, p []js.Value) interface{} {
	// g.gamepad = p[0].Get("gamepad")
	// g.Log("Gamepad connected", p[0].Get("gamepad"))
	// g.Log("Gamepad connected")

	if g.GamepadConfig != nil && g.GamepadConfig.Connected != nil {
		g.GamepadConfig.Connected()
	}

	return nil
}

func (g gamepadClient) disconnected(this js.Value, p []js.Value) interface{} {
	// g.Log("Gamepad disconnected")
	if g.GamepadConfig != nil && g.GamepadConfig.Disconnected != nil {
		g.GamepadConfig.Disconnected()
	}

	return nil
}
