package gamepadclient

func (g *gamepadClient) GamepadClientNotifyRegister(gamepadConfig any) {
	// g.Log("registrando función botón")

	if config, ok := gamepadConfig.(*GamepadConfig); ok {

		g.GamepadConfig = config

		if g.connected && config.Connected != nil {
			config.Connected()
		}

	}

}
