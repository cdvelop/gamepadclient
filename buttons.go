package gamepadclient

import "syscall/js"

func (g *gamepadClient) gamepadButtonHandler(this js.Value, p []js.Value) interface{} {
	// g.Log("BOTON:", p[0])
	// g.Log("gamepad:", p[1])

	if g.pressAnyButton != nil {
		g.GamepadPressAnyButton()

	} else if g.pressSpecificButtons != nil {

		if funCall, ok := g.GamepadPressSpecificButton()[p[0].Int()]; ok {
			funCall()
		}

	}

	return nil

}
