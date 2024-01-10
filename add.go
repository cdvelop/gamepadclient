package gamepadclient

import (
	"github.com/cdvelop/model"
)

var gamepad *gamepadClient

func AddGamePadHandler(l model.Logger) (g *gamepadClient) {

	if gamepad == nil {

		gamepad = &gamepadClient{
			Logger: l,
		}

	}

	return gamepad
}
