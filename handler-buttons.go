package gamepadclient

import "syscall/js"

func (g *gamepadClient) handleButtons(buttons js.Value) {

	for i := 0; i < buttons.Length(); i++ {
		button := buttons.Index(i)
		if button.Get("pressed").Bool() && !g.previousButtonState[i] {

			js.Global().Get("console").Call("log", "BOTÓN PRESIONADO", i)

		}
		g.previousButtonState[i] = button.Get("pressed").Bool()
	}

}
