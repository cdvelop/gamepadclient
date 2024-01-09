package gamepadclient

import (
	"syscall/js"
)

var gamepad *gamepadClient

func AddGamePadHandler() (g *gamepadClient) {

	if gamepad == nil {

		gamepad = &gamepadClient{}

		js.Global().Set("gameLoop", js.FuncOf(g.gameLoop))
	}

	return gamepad
}
