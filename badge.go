package main

import (
	"fmt"
	"image/color"
	"image/png"
	"io"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/freeserif"
)

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

type Item struct {
	label string
	img   []uint8
}

var sidebar = []*Item{
	{
		label: Pronouns,
		img:   untEagle,
	},
	{
		label: "Email",
		img:   qrToBuffer(fmt.Sprintf("mailto:%v.%v@unt.edu", FName, LName)),
	},
	{
		label: "W W W",
		img:   qrToBuffer("https://computerscience.engineering.unt.edu/"),
	},
}

const (
	headerBottom = 25
	rMargin      = 208
)

func showBadge() {
	display.ClearDisplay()
	display.WaitUntilIdle()

	fillRect(0, 0, WIDTH, headerBottom, black)
	_ = fitTextToWidth("University of North Texas", 0, WIDTH, 19, white, sansRegularFont...)
	// center first and last name
	fname, lname := CenterStrings(FName, LName)
	_ = fitTextToWidth(fname, 0, rMargin, 66, black, &freesans.Bold24pt7b)
	_ = fitTextToWidth(lname, 0, rMargin, 98, black, &freesans.Bold18pt7b)
	_ = fitTextToWidth(Title, 1, rMargin, 122, black, monoBoldFont...)

	//display.Display()
	//display.WaitUntilIdle()

	if LinkedIn != "" {
		sidebar = append(sidebar, &Item{
			label: "LinkedIn",
			img:   qrToBuffer("https://www.linkedin.com/in/" + LinkedIn + "/"),
		})
	}

	sidebarMenu()
}

func CenterStrings(s1, s2 string) (string, string) {
	maxLen := max(len(s1), len(s2)) // Use built-in max function
	return centerString(s1, maxLen), centerString(s2, maxLen)
}

// centerString centers the string with spaces based on the given length.
func centerString(s string, length int) string {
	padTotal := length - len(s)
	leftPad := padTotal / 2
	rightPad := padTotal - leftPad

	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

func sidebarMenu() {
	selected := int16(0)
	indicatorHeight := int16(qrSize / len(sidebar))
	drawSidebar(sidebar[selected].label, sidebar[selected].img)
	fillRect(WIDTH-2, 48+selected*indicatorHeight, 2, indicatorHeight, black)
	display.Display()

	for {
		switch {
		case btnUp.Get() && selected > 0:
			selected--
			drawSidebar(sidebar[selected].label, sidebar[selected].img)
			fillRect(WIDTH-2, 48+selected*indicatorHeight, 2, indicatorHeight, black)
			fillRect(WIDTH-2, 48+(selected+1)*indicatorHeight, 2, indicatorHeight, white)
			display.Display()
		case btnDown.Get() && selected < int16(len(sidebar)-1):
			selected++
			drawSidebar(sidebar[selected].label, sidebar[selected].img)
			fillRect(WIDTH-2, 48+selected*indicatorHeight, 2, indicatorHeight, black)
			fillRect(WIDTH-2, 48+(selected-1)*indicatorHeight, 2, indicatorHeight, white)
			display.Display()
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func drawSidebar(label string, img []uint8) {
	const (
		headerStart = rMargin - 4
		headerWidth = WIDTH - headerStart
	)
	fillRect(headerStart, headerBottom+1, headerWidth, HEIGHT-headerBottom, white)
	if label != "" {
		// draw header from top of screen
		_, f := lineWidth(label, uint32(headerWidth), monoBoldObliqueFont...)
		lineH := f.GetGlyph(rune(label[0])).Info().Height
		fitTextToWidth(label, headerStart, headerWidth, int16(headerBottom+lineH+1), black, f)
	}

	// DrawBuffer draws from the top right corner.
	display.DrawBuffer(HEIGHT-80, 6, 80, 80, img)
}

const qrSize = 80

func qrToBuffer(context string) []uint8 {
	q, err := qrcode.New(context, qrcode.Medium)
	if err != nil {
		panic(err)
	}
	q.DisableBorder = true

	return ditherImage(q.Image(qrSize))
}

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
	w32, f := lineWidth(text, uint32(x1-x0), fonts...)
	tinyfont.WriteLine(&display, f, x0+(x1-int16(w32))/2, y, text, c)

	return f
}

func lineWidth(text string, maxWidth uint32, fonts ...*tinyfont.Font) (uint32, *tinyfont.Font) {
	for _, f := range fonts {
		if w32, _ := tinyfont.LineWidth(f, text); w32 < maxWidth {
			return w32, f
		}
	}

	// fallthrough to the last font
	return maxWidth, fonts[len(fonts)-1]
}

func fillRect(x, y, w, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			display.SetPixel(i, j, c)
		}
	}
}

// ConvertPNGToMonochrome converts a png to pixel.Image[pixel.Monochrome]
func ConvertPNGToMonochrome(r io.Reader) pixel.Image[pixel.Monochrome] {
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
