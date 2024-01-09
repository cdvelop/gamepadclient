package gamepadclient

import "syscall/js"

func (g *gamepadClient) gamepadButtonHandler(this js.Value, p []js.Value) interface{} {
	// g.Log("BOTON:", p[0])
	// g.Log("gamepad:", p[1])
	if g.GamepadConfig != nil {

		if g.GamepadConfig.ButtonAny != nil {
			g.GamepadConfig.ButtonAny()

		} else if g.GamepadConfig.ButtonSpecific != nil {

			if funCall, ok := g.GamepadConfig.ButtonSpecific[p[0].Int()]; ok {
				funCall()
			}

		}

	}

	return nil

}
