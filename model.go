package gamepadclient

type gamepadClient struct {
	connected bool

	notifyConnection

	pressAnyButton

	pressSpecificButtons
}

type notifyConnection interface {
	GamepadConnected()    //  al conectar gamepad
	GamepadDisconnected() //  al desconectar gamepad
}

type pressAnyButton interface {
	GamepadPressAnyButton() //  al presionar cualquier botón
}

type pressSpecificButtons interface {
	GamepadPressSpecificButton() map[int]func() // función a ejecutar según numero de botón especifico
}
