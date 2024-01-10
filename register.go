package gamepadclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (g *gamepadClient) GamepadCallFunRegisterButton(config *model.GamepadConfig) {
	// g.Log("registrando función botón")

	if g.GamepadConfig == nil { // solo registro una ves
		g.GamepadConfig = config

		js.Global().Set("gamepadButtonHandler", js.FuncOf(gamepad.gamepadButtonHandler))
		js.Global().Get("window").Call("addEventListener", "gamepadconnected", js.FuncOf(gamepad.connected))
		js.Global().Get("window").Call("addEventListener", "gamepaddisconnected", js.FuncOf(gamepad.disconnected))

	}
}
