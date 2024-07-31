package main

import (
	"encoding/base64"
	"io"
	"strings"
)

var untLogo = []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xE, 0x60, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xB, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0x90, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xB, 0xB0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xC, 0x60, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xD, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x74, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x6D, 0xB7, 0x77, 0x77, 0x77, 0x76, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x60, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0x1, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xC0, 0x3C, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xF8, 0xE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0x1, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x1, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xF8, 0x0, 0x0, 0x1, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xC, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x55, 0x55, 0x55, 0x55, 0x7F, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0x1, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xF0, 0x0, 0x0, 0x3, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0x7, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0xF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xF0, 0x0, 0x7, 0xFF, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xF0, 0x0, 0xF, 0xCF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xF0, 0x0, 0x1F, 0xBF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xF0, 0x0, 0x3F, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xF0, 0x0, 0x7F, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xE0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xE0, 0x1, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xE0, 0x3, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0xC0, 0x7, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFF, 0xC0, 0xF, 0xFF, 0xFC, 0x0, 0xF0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0xF, 0xFF, 0xFC, 0x0, 0xFC, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x3F, 0xFF, 0xFC, 0x1, 0xF4, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x3F, 0xFF, 0xFC, 0x1, 0xE0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x7F, 0x2, 0xD8, 0x1, 0xE0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x0, 0xF0, 0x7F, 0xFC, 0x1, 0xE0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0xE9, 0x87, 0xFF, 0xFC, 0x0, 0xF0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x1F, 0xFF, 0xFE, 0x3, 0xF0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0xFF, 0xFF, 0xFE, 0xF, 0xF8, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE1, 0xFF, 0xFF, 0xFF, 0x1F, 0xFC, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x87, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xEE, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x3F, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE3, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x87, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFF, 0x8F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xFF, 0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x0, 0x0, 0xF, 0xFF, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xFE, 0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFC, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x7, 0xFF, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0xFF, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0xF0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x1F, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0xE1, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x7, 0xFE, 0x7F, 0xC3, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFC, 0xFF, 0x87, 0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xF9, 0xFF, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xF3, 0xFE, 0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xE7, 0xFC, 0x3F, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xCF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0x9F, 0xF0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x3F, 0xF1, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFE, 0x7F, 0xC3, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFC, 0xFF, 0xC7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xF9, 0xFF, 0x8F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xF3, 0xFF, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE7, 0xFE, 0x3F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0x8F, 0xFC, 0x3F, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x9F, 0xF8, 0x7F, 0xFB, 0xFF, 0x7F, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFE, 0x3F, 0xF0, 0xFF, 0xF7, 0xFE, 0x7F, 0xFF, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFE, 0x7F, 0xE1, 0xFF, 0xE7, 0xFC, 0xFF, 0x3F, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xF8, 0xFF, 0xC3, 0xFF, 0xCF, 0xFD, 0xFF, 0x7F, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xF1, 0xFF, 0x87, 0xFF, 0xDF, 0xF9, 0xFE, 0x3F, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xE3, 0xFF, 0xF, 0xFF, 0xBF, 0xF3, 0xFE, 0x7F, 0xF7, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xC7, 0xFE, 0x1F, 0xFF, 0x3F, 0xF7, 0xFC, 0x7F, 0xF3, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0x8F, 0xFC, 0x3F, 0xFE, 0x7F, 0xE7, 0xF8, 0x7F, 0xFB, 0xF0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x1F, 0xF8, 0x7F, 0xFC, 0xFF, 0xCF, 0xF8, 0x7F, 0xF3, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFE, 0x3F, 0xF8, 0xFF, 0xFD, 0xFF, 0xDF, 0xF0, 0x7F, 0xFB, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFC, 0x7F, 0xE1, 0xFF, 0xFB, 0xFF, 0x9F, 0xF0, 0x7F, 0xFB, 0xF8, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xF8, 0xFF, 0xE3, 0xFF, 0xF3, 0xFF, 0x3F, 0xE0, 0x7F, 0xF9, 0xFC, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xF1, 0xFF, 0xC3, 0xFF, 0xE7, 0xFE, 0x7F, 0xC0, 0x7F, 0xF9, 0xF8, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xE3, 0xFF, 0x8F, 0xFF, 0xCF, 0xFE, 0x7F, 0xC0, 0x7F, 0xFD, 0xFC, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xC7, 0xFF, 0xF, 0xFF, 0x9F, 0xFC, 0xFF, 0x80, 0x7F, 0xFD, 0xFC, 0x0, 0x0, 0x0, 0xF, 0xFF, 0x8F, 0xFE, 0x1F, 0xFF, 0x9F, 0xF9, 0xFF, 0x80, 0x7F, 0x7D, 0xFC, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0x1F, 0xFC, 0x3F, 0xFF, 0x3F, 0xF9, 0xFF, 0x0, 0x7F, 0xFC, 0xF8, 0x0, 0x0, 0x0, 0x7F, 0xFE, 0x3F, 0xF8, 0x7F, 0xFE, 0x7F, 0xF3, 0xFE, 0x0, 0x7E, 0x7C, 0xF8, 0x0, 0x0, 0x0, 0x7F, 0xFC, 0x7F, 0xF0, 0xFF, 0xFC, 0xFF, 0xE7, 0xFE, 0x0, 0x7E, 0xFC, 0xF0, 0x0, 0x0, 0x1, 0xFF, 0xF0, 0xFF, 0xE1, 0xFF, 0xF9, 0xFF, 0xE7, 0xFC, 0x0, 0x7E, 0x7E, 0xE0, 0x0, 0x0, 0xF, 0xFF, 0xF3, 0xFF, 0xC3, 0xFF, 0xF9, 0xFF, 0xCF, 0xFC, 0x0, 0x7F, 0x7E, 0xC0, 0x0, 0x0, 0xF, 0xFF, 0xCF, 0xFF, 0x87, 0xFF, 0xF3, 0xFF, 0x9F, 0xF8, 0x0, 0x7E, 0x7E, 0x80, 0x0, 0x0, 0x1, 0xFF, 0x87, 0xFF, 0xF, 0xFF, 0xE7, 0xFF, 0x1F, 0xF0, 0x0, 0x7E, 0x7E, 0x0, 0x0, 0x0, 0x0, 0x3F, 0x0, 0xFE, 0x1F, 0xFF, 0xCF, 0xFF, 0x3F, 0xF0, 0x0, 0x7F, 0x7E, 0x0, 0x0, 0x0, 0x0, 0xC, 0x0, 0x3C, 0x3F, 0xFF, 0x8F, 0xFE, 0x7F, 0xE0, 0x0, 0x7E, 0x7C, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0x9F, 0xFC, 0x7F, 0xE0, 0x0, 0x7F, 0x78, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0x3F, 0xFC, 0xFF, 0xC0, 0x0, 0x7E, 0x70, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFE, 0x7F, 0xF9, 0xFF, 0x80, 0x0, 0x7F, 0x60, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFC, 0xFF, 0xF1, 0xFF, 0x80, 0x0, 0x7E, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xF8, 0xFF, 0xF3, 0xFF, 0x0, 0x0, 0x7F, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xF1, 0xFF, 0xE7, 0xFF, 0x0, 0x0, 0x7E, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xF3, 0xFF, 0xC7, 0xFE, 0x0, 0x0, 0x7E, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xE7, 0xFF, 0x8F, 0xFC, 0x0, 0x0, 0x7C, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xC7, 0xFF, 0x9F, 0xFC, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0x8F, 0xFF, 0x1F, 0xF8, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0x1F, 0xFE, 0x3F, 0xF8, 0x0, 0x0, 0x60, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0x3F, 0xFE, 0x7F, 0xF0, 0x0, 0x0, 0x60, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFE, 0x3F, 0xFC, 0x7F, 0xE0, 0x0, 0x0, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFC, 0x7F, 0xF8, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xF8, 0xFF, 0xF9, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xF1, 0xFF, 0xF1, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xF3, 0xFF, 0xE3, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xE3, 0xFF, 0xE7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0xC7, 0xFF, 0xC7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0x8F, 0xFF, 0x8F, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFF, 0x1F, 0xFF, 0x1F, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFE, 0x1F, 0xFF, 0x1F, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xFF, 0xFE, 0x3F, 0xFE, 0x3F, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFC, 0x7F, 0xFC, 0x7F, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1F, 0xFF, 0xF8, 0xFF, 0xFC, 0x7F, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xFF, 0xF0, 0xFF, 0xF8, 0xFF, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xE1, 0xFF, 0xF1, 0xFF, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFF, 0xE3, 0xFF, 0xF1, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xFF, 0xFF, 0xC7, 0xFF, 0xE3, 0xFF, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2F, 0xFF, 0xFF, 0x8F, 0xFF, 0xC7, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7F, 0xFF, 0xFF, 0xF, 0xFF, 0x87, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF, 0xFF, 0xFE, 0x1F, 0xFF, 0x8F, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xFF, 0xFE, 0x3F, 0xFF, 0x7F, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFF, 0xFD, 0xFF, 0xFE, 0x7F, 0xFE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3F, 0xF9, 0xFF, 0xFE, 0xF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xF0, 0x3F, 0xFC, 0x3, 0xF8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0xF, 0xF8, 0x0, 0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xE0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

const untEagle = `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`

const untEagle2 = `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`

func base64Reader(s string) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))
}
