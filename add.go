package gamepadclient

import "syscall/js"

var gamepad *gamepadClient

func AddGamePadHandler() *gamepadClient {

	if gamepad == nil {

		gamepad = &gamepadClient{
			connected:            false,
			notifyConnection:     nil,
			pressAnyButton:       nil,
			pressSpecificButtons: nil,
		}

		js.Global().Set("gamepadButtonHandler", js.FuncOf(gamepad.gamepadButtonHandler))
		js.Global().Get("window").Call("addEventListener", "gamepadconnected", js.FuncOf(gamepad.Connected))
		js.Global().Get("window").Call("addEventListener", "gamepaddisconnected", js.FuncOf(gamepad.Disconnected))

	}

	return gamepad
}
