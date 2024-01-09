package gamepadclient

import "github.com/cdvelop/model"

func (g *gamepadClient) GamepadCallFunRegisterButton(config *model.GamepadConfig) {
	// g.Log("registrando función botón")
	g.GamepadConfig = config
}
