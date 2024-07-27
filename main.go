package main

import (
	"bytes"
	//_ "embed"
	"image"
	"image/color"
	"machine"

	"tinygo.org/x/drivers/uc8151"
)

var display uc8151.Device
var btnA, btnB, btnC, btnUp, btnDown machine.Pin
var led = machine.LED

var black = color.RGBA{1, 1, 1, 255}
var white = color.RGBA{0, 0, 0, 255}

const WIDTH = 296
const HEIGHT = 128

//go:embed"assets/unt_logo.png"
//var untLogo []byte

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display = uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation:    uc8151.ROTATION_270,
		Speed:       uc8151.MEDIUM,
		Blocking:    true,
		FlickerFree: true,
	})

	btnA = machine.BUTTON_A
	btnB = machine.BUTTON_B
	btnC = machine.BUTTON_C
	btnUp = machine.BUTTON_UP
	btnDown = machine.BUTTON_DOWN
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	//display.ClearDisplay()
	//img, err := pngToImage(untLogo)
	//if err != nil {
	//	print(err)
	//	panic(err)
	//}
	//print(img)
	//time.Sleep(3 * time.Second)

	profile()
}

func pngToImage(pngData []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(pngData))
	if err != nil {
		return nil, err
	}
	return img, nil
}
