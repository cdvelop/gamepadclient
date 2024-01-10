package gamepadclient

type gamepadClient struct {
	connected bool

	*GamepadConfig
}

type GamepadConfig struct {
	Connected      func()         //  al conectar gamepad
	Disconnected   func()         //  al desconectar gamepad
	ButtonAny      func()         //  al presionar cualquier botón
	ButtonSpecific map[int]func() // función a ejecutar según numero de botón especifico
}
