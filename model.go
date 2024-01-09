package gamepadclient

import (
	"github.com/cdvelop/model"
)

type gamepadClient struct {
	model.Logger

	*model.GamepadConfig
}
