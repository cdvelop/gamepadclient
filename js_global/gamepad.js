let previousButtonState = [];

function gameLoop() {
    const gamepads = navigator.getGamepads();

    for (let i = 0; i < gamepads.length; i++) {
        const gamepad = gamepads[i];
        
        if (gamepad && gamepad.connected) {
            handleButtons(gamepad,  gamepad.buttons);
        }
    }

    requestAnimationFrame(gameLoop);
}

function handleButtons(gamepad, buttons) {
    buttons.forEach((button, index) => {
        if (button.pressed && !previousButtonState[index]) {
            gamepadButtonHandler(index,gamepad) 
        }
        previousButtonState[index] = button.pressed;
    });
}

gameLoop();