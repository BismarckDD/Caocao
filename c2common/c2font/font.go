package c2font

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/c2common/c2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/c2common/c2interface"
	"github.com/OpenDiablo2/OpenDiablo2/c2common/c2math"
	"golang.org/x/image/font"
)

const (
	knownSignature = "Woo!\x01"
)

const (
	numHeaderBytes          = 12
	bytesPerGlyph           = 14
	signatureBytesCount     = 5
	unknownHeaderBytesCount = 7
	unknown1BytesCount      = 1
	unknown2BytesCount      = 3
	unknown3BytesCount      = 4
)

// Font represents a displayable font
type Font struct {
	sheet  c2interface.Animation
	table  []byte
	Glyphs map[rune]*font.Face
	color  color.Color
}

// Load loads a new font from byte slice
func Load(data []byte) (*Font, error) {
	sr := c2datautils.CreateStreamReader(data)

	signature, err := sr.ReadBytes(signatureBytesCount)
	if err != nil {
		return nil, err
	}

	if string(signature) != knownSignature {
		return nil, fmt.Errorf("invalid font table format")
	}

	font := &Font{
		table: data,
		color: color.White,
	}

	sr.SkipBytes(unknownHeaderBytesCount)

	err = font.initGlyphs(sr)
	if err != nil {
		return nil, err
	}

	return font, nil
}

// SetBackground sets font's background
func (f *Font) SetBackground(sheet c2interface.Animation) {
	f.sheet = sheet

	// recalculate max height
	_, h := f.sheet.GetFrameBounds()

	for i := range f.Glyphs {
		f.Glyphs[i].SetSize(f.Glyphs[i].Width(), h)
	}
}

// SetColor sets the fonts color
func (f *Font) SetColor(c color.Color) {
	f.color = c
}

// GetTextMetrics returns the dimensions of the Font element in pixels
func (f *Font) GetTextMetrics(text string) (width, height int) {
	var (
		lineWidth  int
		lineHeight int
	)

	for _, c := range text {
		if c == '\n' {
			width = c2math.MaxInt(width, lineWidth)
			height += lineHeight
			lineWidth = 0
			lineHeight = 0
		} else if glyph, ok := f.Glyphs[c]; ok {
			lineWidth += glyph.Width()
			lineHeight = c2math.MaxInt(lineHeight, glyph.Height())
		}
	}

	width = c2math.MaxInt(width, lineWidth)
	height += lineHeight

	return width, height
}

// RenderText prints a text using its configured style on a Surface (multi-lines are left-aligned, use label otherwise)
func (f *Font) RenderText(text string, target c2interface.Surface) error {
	f.sheet.SetColorMod(f.color)

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		var (
			lineHeight int
			lineLength int
		)

		for _, c := range line {
			glyph, ok := f.Glyphs[c]
			if !ok {
				continue
			}

			if err := f.sheet.SetCurrentFrame(glyph.FrameIndex()); err != nil {
				return err
			}

			f.sheet.Render(target)

			lineHeight = c2math.MaxInt(lineHeight, glyph.Height())
			lineLength++

			target.PushTranslation(glyph.Width(), 0)
		}

		target.PopN(lineLength)
		target.PushTranslation(0, lineHeight)
	}

	target.PopN(len(lines))

	return nil
}

func (f *Font) initGlyphs(sr *c2datautils.StreamReader) error {
	glyphs := make(map[rune]*c2fontglyph.FontGlyph)

	// for i := numHeaderBytes; i < len(f.table); i += bytesPerGlyph {
	for i := numHeaderBytes; true; i += bytesPerGlyph {
		code, err := sr.ReadUInt16()
		if err != nil {
			break
		}

		// byte of 0
		sr.SkipBytes(unknown1BytesCount)

		width, err := sr.ReadByte()
		if err != nil {
			return err
		}

		height, err := sr.ReadByte()
		if err != nil {
			return err
		}

		// 1, 0, 0
		sr.SkipBytes(unknown2BytesCount)

		frame, err := sr.ReadUInt16()
		if err != nil {
			return err
		}

		// 1, 0, 0, character code repeated, and further 0.
		sr.SkipBytes(unknown3BytesCount)

		glyph := c2fontglyph.Create(int(frame), int(width), int(height))

		glyphs[rune(code)] = glyph
	}

	f.Glyphs = glyphs

	return nil
}

// Marshal encodes font back into byte slice
func (f *Font) Marshal() []byte {
	sw := c2datautils.CreateStreamWriter()

	sw.PushBytes([]byte("Woo!\x01")...)

	// unknown header bytes - constant
	sw.PushBytes([]byte{1, 0, 0, 0, 0, 1}...)

	// Expected Height of character cell and Expected Width of character cell
	// not used in decoder
	sw.PushBytes([]byte{0, 0}...)

	for c, i := range f.Glyphs {
		sw.PushUint16(uint16(c))
		sw.PushBytes(i.Unknown1()...)
		sw.PushBytes(byte(i.Width()))
		sw.PushBytes(byte(i.Height()))
		sw.PushBytes(i.Unknown2()...)
		sw.PushUint16(uint16(i.FrameIndex()))
		sw.PushBytes(i.Unknown3()...)
	}

	return sw.GetBytes()
}
