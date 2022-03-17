package c2ui

import (
	"image/color"
	"regexp"
	"strings"

	"github.com/BismarckDD/Caocao/c2common/c2fileformats/c2font"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2util"
)

// static check Label implements Widget interface.
var _ Widget = &Label{}

// Label represents a user interface label
type Label struct {
	*BaseWidget
	text            string
	Alignment       HorizontalAlign
	font            *c2font.Font
	Color           map[int]color.Color
	backgroundColor color.Color

	*c2util.Logger
}

// NewLabel creates a new instance of a UI label
func (ui *UIManager) NewLabel(fontPath, palettePath string) *Label {
	font, err := c2utils.LoadFont(fontPath)
	if err != nil {
		ui.Error(err.Error())
		return nil
	}

	base := NewBaseWidget(ui)

	result := &Label{
		BaseWidget: base,
		Alignment:  HorizontalAlignLeft,
		Color:      map[int]color.Color{0: color.White},
		font:       font,
		Logger:     ui.Logger,
	}

	result.bindManager(ui)

	result.SetVisible(false)

	ui.addWidget(result)

	return result
}

// Render draws the label on the screen, respliting the lines to allow for other alignments.
func (v *Label) Render(target c2interface.Surface) {
	target.PushTranslation(v.GetPosition())

	lines := strings.Split(v.text, "\n")
	yOffset := 0

	lastColor := v.Color[0]
	v.font.SetColor(lastColor)

	for _, line := range lines {
		lw, lh := v.GetTextMetrics(line)
		characters := []rune(line)

		target.PushTranslation(v.getAlignOffset(lw), yOffset)

		for idx := range characters {
			character := string(characters[idx])
			charWidth, charHeight := v.GetTextMetrics(character)

			if v.Color[idx] != nil {
				lastColor = v.Color[idx]
				v.font.SetColor(lastColor)
			}

			if v.backgroundColor != nil {
				target.DrawRect(charWidth, charHeight, v.backgroundColor)
			}

			err := v.font.RenderText(character, target)
			if err != nil {
				v.Error(err.Error())
			}

			target.PushTranslation(charWidth, 0)
		}

		target.PopN(len(characters))

		yOffset += lh

		target.Pop()
	}

	target.Pop()
}

// GetTextMetrics returns the width and height of the enclosing rectangle in Pixels.
func (v *Label) GetTextMetrics(text string) (width, height int) {
	return v.font.GetTextMetrics(text)
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	v.text = v.processColorTokens(newText)
	v.BaseWidget.width, v.BaseWidget.height = v.font.GetTextMetrics(v.text)
}

// GetText returns label text
func (v *Label) GetText() string {
	return v.text
}

// SetBackgroundColor sets the background highlight color
func (v *Label) SetBackgroundColor(c color.Color) {
	v.backgroundColor = c
}

func (v *Label) processColorTokens(str string) string {
	tokenMatch := regexp.MustCompile(colorTokenMatch)
	tokenStrMatch := regexp.MustCompile(colorStrMatch)
	empty := []byte("")

	tokenPosition := 0

	withoutTokens := string(tokenMatch.ReplaceAll([]byte(str), empty)) // remove tokens from string

	matches := tokenStrMatch.FindAll([]byte(str), -1)

	if len(matches) == 0 {
		v.Color[0] = getColor(ColorTokenWhite)
	}

	// we find the index of each token and update the color map.
	// the key in the map is the starting index of each color token, the value is the color
	for idx := range matches {
		match := matches[idx]
		matchToken := tokenMatch.Find(match)
		matchStr := string(tokenMatch.ReplaceAll(match, empty))
		token := ColorToken(matchToken)

		theColor := getColor(token)
		if theColor == nil {
			continue
		}

		if v.Color == nil {
			v.Color = make(map[int]color.Color)
		}

		v.Color[tokenPosition] = theColor

		tokenPosition += len(matchStr)
	}

	return withoutTokens
}

func (v *Label) getAlignOffset(textWidth int) int {
	switch v.Alignment {
	case HorizontalAlignLeft:
		return 0
	case HorizontalAlignCenter:
		// nolint:gomnd // center of label = 1/2 of it
		return -textWidth / 2
	case HorizontalAlignRight:
		return -textWidth
	default:
		v.Fatal("Invalid Alignment")
		return 0
	}
}

// Advance is a no-op
func (v *Label) Advance(elapsed float64) error {
	return nil
}

func getColor(colorToken ColorToken) color.Color {

	colors := map[ColorToken]color.Color{
		ColorTokenGrey:   c2util.Color(colorGrey100Alpha),
		ColorTokenWhite:  c2util.Color(colorWhite100Alpha),
		ColorTokenBlue:   c2util.Color(colorBlue100Alpha),
		ColorTokenYellow: c2util.Color(colorYellow100Alpha),
		ColorTokenGreen:  c2util.Color(colorGreen100Alpha),
		ColorTokenGold:   c2util.Color(colorGold100Alpha),
		ColorTokenOrange: c2util.Color(colorOrange100Alpha),
		ColorTokenRed:    c2util.Color(colorRed100Alpha),
		ColorTokenBlack:  c2util.Color(colorBlack100Alpha),
	}

	chosen := colors[colorToken]

	if chosen == nil {
		return nil
	}

	return chosen
}
