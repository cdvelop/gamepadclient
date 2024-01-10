package gamepadclient

func (g *gamepadClient) GamepadRegister(NotifyConnection, PressAnyButton, PressSpecificButtons any) {
	// g.Log("registrando función botón")

	if value, ok := NotifyConnection.(notifyConnection); ok {
		g.notifyConnection = value
	}

	if value, ok := PressAnyButton.(pressAnyButton); ok {
		g.pressAnyButton = value
	}

	if value, ok := PressSpecificButtons.(pressSpecificButtons); ok {
		g.pressSpecificButtons = value
	}

	if g.connected && g.notifyConnection != nil {
		g.GamepadConnected()
	}

}
