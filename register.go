package gamepadclient

func (g *gamepadClient) NotifyRegister(c notifyConnection, a pressAnyButton, s pressSpecificButtons) {
	// g.Log("registrando función botón")

	g.notifyConnection = c

	g.pressAnyButton = a

	g.pressSpecificButtons = s

	if g.connected && g.notifyConnection != nil {
		g.GamepadConnected()
	}

}
