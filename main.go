package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/uc8151"
)

var (
	display = uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	btnA    = machine.BUTTON_A
	btnB    = machine.BUTTON_B
	btnC    = machine.BUTTON_C
	btnUp   = machine.BUTTON_UP
	btnDown = machine.BUTTON_DOWN
	led     = machine.LED
)

var (
	black = color.RGBA{R: 1, G: 1, B: 1, A: 255}
	white = color.RGBA{A: 255}
)

const (
	WIDTH  = 296
	HEIGHT = 128
)

func main() {
	configGlobals()
	splashScreen()
	showBadge()
}

func configGlobals() {
	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display.Configure(uc8151.Config{
		Rotation:    uc8151.ROTATION_270,
		Speed:       uc8151.MEDIUM,
		Blocking:    true,
		FlickerFree: true,
	})

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
}

func splashScreen() {
	display.ClearDisplay()
	display.DrawBuffer(0, 0, 128, 296, untLogo)
	display.Display()
	display.WaitUntilIdle()
	time.Sleep(3 * time.Second)
}
