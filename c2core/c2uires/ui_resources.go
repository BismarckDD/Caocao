package c2uires

import (
	"image/color"

	"github.com/blizzy78/ebitenui/image"
	"github.com/golang/freetype/truetype"
)

const (
	normalLishuFont  = "resources/fonts/lishu.ttf"
	normalKaishuFont = "resources/fonts/lishu.ttf"
)

const (
	backgroundColor = "131a22"

	buttonIdleColor                = "000000" // pure black
	buttonDisabledColor            = "aaaaaa" // gray
	buttonIdleBackground           = "eeeeee" // pure black
	buttonHoverBackground          = "dddddd" // pure black
	buttonPressedBackground        = "dddddd" // pure black
	buttonDisabledBackground       = "cccccc" // gray
	textIdleColor                  = "dff4ff"
	textDisabledColor              = "5a7a91"
	labelIdleColor                 = textIdleColor
	labelDisabledColor             = textDisabledColor
	listSelectedBackground         = "4b687a"
	listDisabledSelectedBackground = "2a3944"
	textInputCaretColor            = "e7c34b"
	textInputDisabledCaretColor    = "766326"
	toolTipColor                   = backgroundColor
	separatorColor                 = listDisabledSelectedBackground
)

// UIManager manages a collection of UI elements (buttons, textboxes, labels)
type UIResources struct {
	FontMap            map[string]*truetype.Font
	WindowMenu         *MenuUI
	GameMenu           *MenuUI
	Header             *MenuUI
	Background         *image.NineSlice
	SeparatorColor     color.Color
	Text               *TextUI
	WindowsStyleButton *ButtonUI
	GameStyleButton    *ButtonUI
	MainMenuButton     *ButtonUI
}

func NewUIResources() (*UIResources, error) {

	uiRes := &UIResources{}

	bg, err := LoadImageNineSlice("resources/normal_res/menu/bg_bird.png", 48, 48)
	if err != nil {
		return nil, err
	}
	uiRes.Background = bg

	// initialize font resources.
	lishuFont, err := LoadFont(normalLishuFont)
	if err != nil {
		return nil, err
	}

	kaishuFont, err := LoadFont(normalKaishuFont)
	if err != nil {
		return nil, err
	}

	fontMap := map[string]*truetype.Font{}
	fontMap["lishu"] = lishuFont
	fontMap["kaishu"] = kaishuFont
	uiRes.FontMap = fontMap

	mainMenuButton, err := CreateMainMenuButton(uiRes)
	if err != nil {
		return nil, err
	}

	uiRes.MainMenuButton = mainMenuButton

	gameStyleButton, err := CreateGameStyleButton(uiRes)
	if err != nil {
		return nil, err
	}

	uiRes.GameStyleButton = gameStyleButton

	windowsStyleButton, err := CreateWindowsStyleButton(uiRes)
	if err != nil {
		return nil, err
	}

	uiRes.WindowsStyleButton = windowsStyleButton

	wsMenuUI, err := CreateWindowsStyleMenuUI(uiRes)
	if err != nil {
		return nil, err
	}

	uiRes.Header = wsMenuUI
	uiRes.WindowMenu = wsMenuUI

	return uiRes, nil
}

func (ui *UIResources) Close() {
}

// func (game *Game) Update() error {
// 	game.frameNum++
// 	return nil
// }

// func (game *Game) Draw(screen *ebiten.Image) {

// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(float64(0), float64(80)) // translate: vertically and horizontally move image.
// 	// let face img shows on the screen.
// 	screen.DrawImage(mainMenuBgImg, op)
// 	// fill a rect as menu on the screen.
// 	ebitenutil.DrawRect(screen, 0, 0, 640, 80, color.RGBA{0, 255, 255, 127})
// 	//
// 	currentFPS := fmt.Sprintf("%0.2f", ebiten.CurrentFPS())
// 	// if !game.flag {
// 	ebitenutil.DebugPrintAt(screen, "Current FPS: "+currentFPS, 0, -2)
// 	img0a, _, err1 := ebitenutil.NewImageFromFile("resources/unit/73_move_down_0.png")
// 	img0b, _, err2 := ebitenutil.NewImageFromFile("resources/unit/73_move_down_1.png")
// 	img1a, _, err3 := ebitenutil.NewImageFromFile("resources/unit/74_move_down_0.png")
// 	img1b, _, err4 := ebitenutil.NewImageFromFile("resources/unit/74_move_down_1.png")
// 	img2a, _, err5 := ebitenutil.NewImageFromFile("resources/unit/75_move_down_0.png")
// 	img2b, _, err6 := ebitenutil.NewImageFromFile("resources/unit/75_move_down_1.png")
// 	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
// 		log.Fatal("Failed to get the move img.")
// 	}

// 	flag := (game.frameNum/c2config.FramesPerAction)%2 == 1
// 	var img0, img1, img2 *ebiten.Image
// 	if flag {
// 		img0 = img0a
// 		img1 = img1a
// 		img2 = img2a
// 	} else {
// 		img0 = img0b
// 		img1 = img1b
// 		img2 = img2b
// 	}
// 	op.GeoM.Reset()
// 	op.GeoM.Translate(float64(64), float64(16))
// 	screen.DrawImage(img0, op)
// 	op.GeoM.Reset()
// 	op.GeoM.Translate(float64(128), float64(16))
// 	screen.DrawImage(img1, op)
// 	op.GeoM.Reset()
// 	op.GeoM.Translate(float64(192), float64(16))
// 	screen.DrawImage(img2, op)

// 	ebitenutil.DebugPrintAt(screen, "Current FrameNo: "+fmt.Sprintf("%d", game.frameNum), 0, 12)
// 	if ebiten.IsWindowBeingClosed() {

// 		ebiten.SetWindowClosingHandled(true)
// 	}
// }
