package main

import (
	_ "embed"
	"encoding/base64"
	"image/color"
	"io"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
	"tinygo.org/x/drivers/image/png"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/freeserif"
)

func profile() {
	display.ClearDisplay()
	display.WaitUntilIdle()

	badgeProfile()
}

type font []*tinyfont.Font

var (
	sansRegularFont = font{
		&freesans.Regular24pt7b,
		&freesans.Regular18pt7b,
		&freesans.Regular12pt7b,
		&freesans.Regular9pt7b,
	}
	sansBoldFont = font{
		&freesans.Bold24pt7b,
		&freesans.Bold18pt7b,
		&freesans.Bold12pt7b,
		&freesans.Bold9pt7b,
	}
	monoRegularFont = font{
		&freemono.Regular24pt7b,
		&freemono.Regular18pt7b,
		&freemono.Regular12pt7b,
		&freemono.Regular9pt7b,
	}
	monoBoldObliqueFont = font{
		&freemono.BoldOblique24pt7b,
		&freemono.BoldOblique18pt7b,
		&freemono.BoldOblique12pt7b,
		&freemono.BoldOblique9pt7b,
	}
	monoBoldFont = font{
		&freemono.Bold24pt7b,
		&freemono.Bold18pt7b,
		&freemono.Bold12pt7b,
		&freemono.Bold9pt7b,
	}
	serifRegularFont = font{
		&freeserif.Regular24pt7b,
		&freeserif.Regular18pt7b,
		&freeserif.Regular12pt7b,
		&freeserif.Regular9pt7b,
	}
)

const untEagle = `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`

func base64Reader(s string) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))
}

type (
	Item struct {
		label string
		img   pixel.Image[pixel.Monochrome]
	}
	SidebarItems struct {
		currIdx int
		items   []*Item
	}
)

func (q *SidebarItems) Curr() *Item {
	return q.items[q.currIdx]
}

func (q *SidebarItems) Inc() {
	q.currIdx = (q.currIdx + 1) % len(q.items)
}

func (q *SidebarItems) Dec() {
	q.currIdx = (q.currIdx - 1 + len(q.items)) % len(q.items)
}

var sidebar = SidebarItems{
	items: []*Item{
		{
			label: "Email",
			img:   qrToImage("mailto:jacob.hochstetler@unt.edu"),
		},
		{
			label: "CSE WWW",
			img:   qrToImage("https://computerscience.engineering.unt.edu/"),
		},
		{
			label: "LinkedIn",
			img:   qrToImage("https://www.linkedin.com/in/jacob-hochstetler/"),
		},
		//{
		//	img: ConvertPNGToMonochrome2(base64Reader(untEagle)),
		//},
	},
}

const (
	headerBottom = 28
	rMargin      = 208
)

func badgeProfile() {
	fillRect(0, 0, WIDTH, headerBottom, black)
	_ = fitTextToWidth("University of North Texas", 0, WIDTH, 22, white, sansRegularFont...)
	_ = fitTextToWidth("   Jacob   ", 0, rMargin, 66, black, &freesans.Bold24pt7b)
	_ = fitTextToWidth("Hochstetler", 0, rMargin, 98, black, &freesans.Bold18pt7b)
	_ = fitTextToWidth("Professor | he/him", 1, rMargin, 122, black, monoBoldFont...)

	//tinydraw.Line(&display, rMargin, headerBottom, rMargin, HEIGHT-1, black) // divider

	//if err := display.DrawBitmap(WIDTH-qrSize-10, 26, ConvertPNGToMonochrome(base64Reader(untEagle))); err != nil {
	//	display.ClearDisplay()
	//	_ = fitTextToWidth(err.Error(), 0, 0, 30, black, monoBoldFont...)
	//}

	const qrX, qrY = WIDTH - qrSize - 8, HEIGHT - qrSize
	if err := display.Display(); err != nil {
		panic(err)
	}
	var prevQR *Item
	for {
		if qr := sidebar.Curr(); prevQR != qr {
			drawSidebar(qr.label, qr.img)
			prevQR = qr
			if err := display.Display(); err != nil {
				//panic(err)
			}
			led.Low()
		}
		switch {
		case btnUp.Get():
			led.High()
			sidebar.Dec()
		case btnDown.Get():
			led.High()
			sidebar.Inc()
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}

const ledDuration = 100 * time.Millisecond

func blinkLED(d time.Duration) {
	led.High()
	time.Sleep(d)
	led.Low()
}

func drawSidebar(label string, img pixel.Image[pixel.Monochrome]) {
	fillRect(rMargin, headerBottom+2, WIDTH-rMargin, 32, white)
	if label != "" {
		_ = fitTextToWidth(label, rMargin, qrSize, 44, black, monoBoldObliqueFont...)
		if err := display.DrawBitmap(rMargin+2, HEIGHT-qrSize, img); err != nil {
			panic(err)
		}
		return
	}

	if err := display.DrawBitmap(rMargin+2, HEIGHT-qrSize-30, img); err != nil {
		panic(err)
	}
}

const qrSize = 80

func qrToImage(content string) pixel.Image[pixel.Monochrome] {
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		panic(err)
	}
	q.DisableBorder = true
	img := pixel.NewImage[pixel.Monochrome](qrSize, qrSize)
	qr := q.Image(qrSize)
	for y := range qr.Bounds().Dy() {
		for x := range qr.Bounds().Dx() {
			img.Set(x, y, qr.At(x, y) == color.Black)
		}
	}

	return img
}

func fitTextToWidth(text string, x0, x1, y int16, c color.RGBA, fonts ...*tinyfont.Font) *tinyfont.Font {
	for _, f := range fonts {
		if w32, _ := tinyfont.LineWidth(f, text); w32 < uint32(x1) {
			tinyfont.WriteLine(&display, f, x0+(x1-int16(w32))/2, y, text, c)
			return f
		}
	}
	// fallthrough to the last font
	f := fonts[len(fonts)-1]
	tinyfont.WriteLine(&display, f, x0, y, text, c)

	return f
}

func fillRect(x int16, y int16, w int16, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			display.SetPixel(i, j, c)
		}
	}
}

// ConvertPNGToMonochrome converts a png to pixel.Image[pixel.Monochrome].
func ConvertPNGToMonochrome(r io.Reader) pixel.Image[pixel.Monochrome] {
	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	p := pixel.NewImage[pixel.Monochrome](width, height)
	for y := range height {
		for x := range width {
			c := color.GrayModel.Convert(img.At(x-bounds.Min.X, y-bounds.Min.Y)).(color.Gray)
			p.Set(x, y, c.Y < 128)
		}
	}

	return p
}

func ConvertPNGToMonochrome2(r io.Reader) pixel.Image[pixel.Monochrome] {
	image, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	bounds := image.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	img := pixel.NewImage[pixel.Monochrome](width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := color.GrayModel.Convert(image.At(x+bounds.Min.X, y+bounds.Min.Y)).(color.Gray)
			img.Set(x, y, c.Y < 128)
		}
	}

	return img
}
