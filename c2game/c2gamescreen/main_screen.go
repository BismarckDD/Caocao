package c2gamescreen

import (
	"os"

	"github.com/BismarckDD/Caocao/c2core/c2hero"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2resource"
	"github.com/BismarckDD/Caocao/c2common/c2util"
	"github.com/BismarckDD/Caocao/c2core/c2asset"
	"github.com/BismarckDD/Caocao/c2core/c2screen"
	"github.com/BismarckDD/Caocao/c2core/c2ui"
	"github.com/BismarckDD/Caocao/c2script"
)

type mainMenuScreenMode int

// mainMenuScreenMode types
const (
	ScreenModeUnknown mainMenuScreenMode = iota
	ScreenModeTrademark
	ScreenModeMainMenu
	ScreenModeMultiplayer
	ScreenModeTCPIP
	ScreenModeServerIP
)

const (
	joinGameDialogX, joinGameDialogY         = 318, 245
	serverIPbackgroundX, serverIPbackgroundY = 270, 175
	backgroundX, backgroundY                 = 0, 0
	versionLabelX, versionLabelY             = 795, -10
	commitLabelX, commitLabelY               = 2, 2
	copyrightX, copyrightY                   = 400, 500
	copyright2X, copyright2Y                 = 400, 525
	oc2LabelX, oc2LabelY                     = 400, 580
	tcpOptionsX, tcpOptionsY                 = 400, 23
	joinGameX, joinGameY                     = 400, 190
	diabloLogoX, diabloLogoY                 = 400, 120
	exitDiabloBtnX, exitDiabloBtnY           = 264, 535
	creditBtnX, creditBtnY                   = 264, 505
	cineBtnX, cineBtnY                       = 401, 505
	singlePlayerBtnX, singlePlayerBtnY       = 264, 290
	githubBtnX, githubBtnY                   = 264, 400
	mapTestBtnX, mapTestBtnY                 = 264, 440
	tcpBtnX, tcpBtnY                         = 33, 543
	srvCancelBtnX, srvCancelBtnY             = 285, 305
	srvOkBtnX, srvOkBtnY                     = 420, 305
	multiplayerBtnX, multiplayerBtnY         = 264, 330
	tcpNetBtnX, tcpNetBtnY                   = 264, 280
	networkCancelBtnX, networkCancelBtnY     = 264, 540
	tcpHostBtnX, tcpHostBtnY                 = 264, 200
	tcpJoinBtnX, tcpJoinBtnY                 = 264, 240
	errorLabelX, errorLabelY                 = 400, 250
	machineIPX, machineIPY                   = 400, 90
	tipX, tipY                               = 400, 300
)

const (
	white       = 0xffffffff
	lightYellow = 0xffff8cff
	gold        = 0xd8c480ff
	red         = 0xff0000ff
)

const (
	joinGameCharacterFilter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890._:"
)

const (
	logPrefix = "Game Screen"
)

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(
	navigator c2interface.Navigator,
	asset *c2asset.AssetManager,
	renderer c2interface.Renderer,
	inputManager c2interface.InputManager,
	audioProvider c2interface.AudioProvider,
	uiManager *c2ui.UIManager,
	l c2util.LogLevel,
	errorMessageOptional ...string,
) (*MainMenu, error) {
	heroStateFactory, err := c2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	mainMenu := &MainMenu{
		asset:          asset,
		screenMode:     ScreenModeUnknown,
		leftButtonHeld: true,
		renderer:       renderer,
		inputManager:   inputManager,
		audioProvider:  audioProvider,
		navigator:      navigator,
		uiManager:      uiManager,
		heroState:      heroStateFactory,
	}

	mainMenu.Logger = c2util.NewLogger()
	mainMenu.Logger.SetPrefix(logPrefix)
	mainMenu.Logger.SetLevel(l)

	if len(errorMessageOptional) != 0 {
		mainMenu.errorLabel = uiManager.NewLabel(c2resource.FontFormal12, c2resource.PaletteUnits)
		mainMenu.errorLabel.SetText(errorMessageOptional[0])
	}

	return mainMenu, nil
}

// MainMenu represents the main menu
type MainMenu struct {
	background          *c2ui.Sprite
	diabloLogoLeft      *c2ui.Sprite
	diabloLogoRight     *c2ui.Sprite
	diabloLogoLeftBack  *c2ui.Sprite
	diabloLogoRightBack *c2ui.Sprite
	serverIPBackground  *c2ui.Sprite
	// main button panel.
	newGameButton  *c2ui.Button
	loadGameButton *c2ui.Button
	confGameButton *c2ui.Button
	exitGameButton *c2ui.Button
	// img button bar.
	imgExitGameButton *c2ui.Button
	imgLoadGameButton *c2ui.Button
	imgSaveGameButton *c2ui.Button
	imgConfGameButton *c2ui.Button
	imgUnitInfoButton *c2ui.Button
	imgTreaInfoButton *c2ui.Button
	imgCharInfoButton *c2ui.Button
	imgLandInfoButton *c2ui.Button
	// Battle
	btnServerIPCancel *c2ui.Button
	btnServerIPOk     *c2ui.Button
	copyrightLabel    *c2ui.Label
	copyrightLabel2   *c2ui.Label
	openDiabloLabel   *c2ui.Label
	versionLabel      *c2ui.Label
	commitLabel       *c2ui.Label
	errorLabel        *c2ui.Label
	joinTipLabel      *c2ui.Label
	hostTipLabel      *c2ui.Label
	tcpJoinGameEntry  *c2ui.TextBox
	screenMode        mainMenuScreenMode
	leftButtonHeld    bool

	asset         *c2asset.AssetManager
	inputManager  c2interface.InputManager
	renderer      c2interface.Renderer
	audioProvider c2interface.AudioProvider
	scriptEngine  *c2script.ScriptEngine
	navigator     c2interface.Navigator
	uiManager     *c2ui.UIManager
	heroState     *c2hero.HeroStateFactory
	*c2util.Logger
}

// OnLoad is called to load the resources for the main menu
func (v *MainMenu) OnLoad(loading c2screen.LoadingState) {

	v.audioProvider.PlayBGM(c2resource.BGMTitle)

	loading.Progress(twentyPercent)

	v.createMainMenuLabels(loading)
	v.createMultiplayerLabels()
	v.loadBackgroundSprites()
	v.createLogos(loading)
	v.createMainMenuButtons(loading)
	v.createMultiplayerMenuButtons()

	v.tcpJoinGameEntry = v.uiManager.NewTextbox()
	v.tcpJoinGameEntry.SetPosition(joinGameDialogX, joinGameDialogY)
	v.tcpJoinGameEntry.SetFilter(joinGameCharacterFilter)
	loading.Progress(ninetyPercent)

	if v.screenMode == ScreenModeUnknown {
		v.SetScreenMode(ScreenModeTrademark)
	} else {
		v.SetScreenMode(ScreenModeMainMenu)
	}

	if err := v.inputManager.BindHandler(v); err != nil {
		v.Error("failed to add main menu as event handler")
	}
}

func (v *MainMenu) loadBackgroundSprites() {
	var err error

	v.background, err = v.uiManager.NewSprite(c2resource.GameSelectScreen, c2resource.PaletteSky)
	if err != nil {
		v.Error(err.Error())
	}

	v.background.SetPosition(backgroundX, backgroundY)

	v.trademarkBackground, err = v.uiManager.NewSprite(c2resource.TrademarkScreen, c2resource.PaletteSky)
	if err != nil {
		v.Error(err.Error())
	}

	v.trademarkBackground.SetPosition(backgroundX, backgroundY)

	v.tcpIPBackground, err = v.uiManager.NewSprite(c2resource.TCPIPBackground, c2resource.PaletteSky)
	if err != nil {
		v.Error(err.Error())
	}

	v.tcpIPBackground.SetPosition(backgroundX, backgroundY)

	v.serverIPBackground, err = v.uiManager.NewSprite(c2resource.PopUpOkCancel, c2resource.PaletteFechar)
	if err != nil {
		v.Error(err.Error())
	}

	v.serverIPBackground.SetPosition(serverIPbackgroundX, serverIPbackgroundY)
}

func (v *MainMenu) createMainMenuLabels(loading c2screen.LoadingState) {

	v.versionLabel = v.uiManager.NewLabel(c2resource.FontFormal12, c2resource.PaletteStatic)
	v.versionLabel.Alignment = c2ui.HorizontalAlignRight
	v.versionLabel.SetText("OpenDiablo2 - " + v.buildInfo.Branch)
	v.versionLabel.Color[0] = c2util.Color(white)
	v.versionLabel.SetPosition(versionLabelX, versionLabelY)

	v.commitLabel = v.uiManager.NewLabel(c2resource.FontFormal10, c2resource.PaletteStatic)
	v.commitLabel.Alignment = c2ui.HorizontalAlignLeft
	v.commitLabel.SetText(v.buildInfo.Commit)
	v.commitLabel.Color[0] = c2util.Color(white)
	v.commitLabel.SetPosition(commitLabelX, commitLabelY)

	v.copyrightLabel = v.uiManager.NewLabel(c2resource.FontFormal12, c2resource.PaletteStatic)
	v.copyrightLabel.Alignment = c2ui.HorizontalAlignCenter
	v.copyrightLabel.SetText("@BismarckDD 2022")
	v.copyrightLabel.Color[0] = c2util.Color(lightBrown)
	v.copyrightLabel.SetPosition(copyrightX, copyrightY)
	loading.Progress(thirtyPercent)

	v.copyrightLabel2 = v.uiManager.NewLabel(c2resource.FontFormal12, c2resource.PaletteStatic)
	v.copyrightLabel2.Alignment = c2ui.HorizontalAlignCenter
	v.copyrightLabel2.SetText(v.asset.TranslateString(c2enum.AllRightsReservedLabel))
	v.copyrightLabel2.Color[0] = c2util.Color(lightBrown)
	v.copyrightLabel2.SetPosition(copyright2X, copyright2Y)

	v.openDiabloLabel = v.uiManager.NewLabel(c2resource.FontFormal10, c2resource.PaletteStatic)
	v.openDiabloLabel.Alignment = c2ui.HorizontalAlignCenter
	v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
	v.openDiabloLabel.Color[0] = c2util.Color(lightYellow)
	v.openDiabloLabel.SetPosition(oc2LabelX, oc2LabelY)
	loading.Progress(fiftyPercent)

	if v.errorLabel != nil {
		v.errorLabel.SetPosition(errorLabelX, errorLabelY)
		v.errorLabel.Alignment = c2ui.HorizontalAlignCenter
		v.errorLabel.Color[0] = c2util.Color(red)
	}
}

func (v *MainMenu) createLogos(loading c2screen.LoadingState) {
	var err error

	v.diabloLogoLeft, err = v.uiManager.NewSprite(c2resource.Diablo2LogoFireLeft, c2resource.PaletteUnits)
	if err != nil {
		v.Error(err.Error())
	}

	v.diabloLogoLeft.SetEffect(c2enum.DrawEffectModulate)
	v.diabloLogoLeft.PlayForward()
	v.diabloLogoLeft.SetPosition(diabloLogoX, diabloLogoY)
	loading.Progress(sixtyPercent)

	v.diabloLogoRight, err = v.uiManager.NewSprite(c2resource.Diablo2LogoFireRight, c2resource.PaletteUnits)
	if err != nil {
		v.Error(err.Error())
	}

	v.diabloLogoRight.SetEffect(c2enum.DrawEffectModulate)
	v.diabloLogoRight.PlayForward()
	v.diabloLogoRight.SetPosition(diabloLogoX, diabloLogoY)

	v.diabloLogoLeftBack, err = v.uiManager.NewSprite(c2resource.Diablo2LogoBlackLeft, c2resource.PaletteUnits)
	if err != nil {
		v.Error(err.Error())
	}

	v.diabloLogoLeftBack.SetPosition(diabloLogoX, diabloLogoY)

	v.diabloLogoRightBack, err = v.uiManager.NewSprite(c2resource.Diablo2LogoBlackRight, c2resource.PaletteUnits)
	if err != nil {
		v.Error(err.Error())
	}

	v.diabloLogoRightBack.SetPosition(diabloLogoX, diabloLogoY)
}

func (v *MainMenu) createMainMenuButtons(loading c2screen.LoadingState) {
	v.exitDiabloButton = v.uiManager.NewButton(c2ui.ButtonTypeWide, v.asset.TranslateString(c2enum.ExitGameLabel))
	v.exitDiabloButton.SetPosition(exitDiabloBtnX, exitDiabloBtnY)
	v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })

	v.creditsButton = v.uiManager.NewButton(c2ui.ButtonTypeShort, v.asset.TranslateString(c2enum.CreditsLabel))
	v.creditsButton.SetPosition(creditBtnX, creditBtnY)
	v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })

	v.cinematicsButton = v.uiManager.NewButton(c2ui.ButtonTypeShort, v.asset.TranslateString(c2enum.CinematicsLabel))
	v.cinematicsButton.SetPosition(cineBtnX, cineBtnY)
	v.cinematicsButton.OnActivated(func() { v.onCinematicsButtonClicked() })
	loading.Progress(seventyPercent)

	v.singlePlayerButton = v.uiManager.NewButton(c2ui.ButtonTypeWide, v.asset.TranslateString(c2enum.SinglePlayerLabel))
	v.singlePlayerButton.SetPosition(singlePlayerBtnX, singlePlayerBtnY)
	v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })

	v.githubButton = v.uiManager.NewButton(c2ui.ButtonTypeWide, "开始新游戏")
	v.githubButton.SetPosition(newGameBtnX, newGameBtnY)
	v.githubButton.OnActivated(func() { v.onNewGameButtonClicked() })

	v.mapTestButton = v.uiManager.NewButton(c2ui.ButtonTypeWide, "作战测试")
	v.mapTestButton.SetPosition(battleTestBtnX, battleTestBtnY)
	v.mapTestButton.OnActivated(func() { v.onBattleTestClicked() })
}

func (v *MainMenu) onBattleTestButtonClicked() {
	v.navigator.ToBattleTest() // 一个作战测试场景
}

func (v *MainMenu) onNewGameButtonClicked() {
	v.SetScreenMode(ScreenModeUnknown)
	v.navigator.ToScenario(1) // 去第一个场景
}

func (v *MainMenu) onLoadGameButtonClicked() {
}

func (v *MainMenu) onGameConfigurationButtonClicked() {
}

func (v *MainMenu) onExitButtonClicked() {
	os.Exit(0)
}

// Render renders the main menu
func (v *MainMenu) Render(screen c2interface.Surface) {
	v.renderBackgrounds(screen)
	v.renderLogos(screen)
	v.renderLabels(screen)
}

func (v *MainMenu) renderBackgrounds(screen c2interface.Surface) {
	switch v.screenMode {
	case ScreenModeTrademark:
		v.trademarkBackground.RenderSegmented(screen, 4, 3, 0)
	default:
		v.background.RenderSegmented(screen, 4, 3, 0)
	}
}

func (v *MainMenu) renderLogos(screen c2interface.Surface) {
	switch v.screenMode {
	case ScreenModeTrademark, ScreenModeMainMenu, ScreenModeMultiplayer:
		v.diabloLogoLeftBack.Render(screen)
		v.diabloLogoRightBack.Render(screen)
		v.diabloLogoLeft.Render(screen)
		v.diabloLogoRight.Render(screen)
	}
}

func (v *MainMenu) renderLabels(screen c2interface.Surface) {
	switch v.screenMode {
	case ScreenModeServerIP:
		v.tcpIPOptionsLabel.Render(screen)
		v.tcpJoinGameLabel.Render(screen)
		v.machineIP.Render(screen)
	case ScreenModeTCPIP:
		v.tcpIPOptionsLabel.Render(screen)
		v.machineIP.Render(screen)
	case ScreenModeTrademark:
		v.copyrightLabel.Render(screen)
		v.copyrightLabel2.Render(screen)

		if v.errorLabel != nil {
			v.errorLabel.Render(screen)
		}
	case ScreenModeMainMenu:
		v.openDiabloLabel.Render(screen)
		v.versionLabel.Render(screen)
		v.commitLabel.Render(screen)
	}
}

// Advance runs the update logic on the main menu
func (v *MainMenu) Advance(tickTime float64) error {
	switch v.screenMode {
	case ScreenModeMainMenu, ScreenModeTrademark, ScreenModeMultiplayer:
		if err := v.diabloLogoLeftBack.Advance(tickTime); err != nil {
			return err
		}

		if err := v.diabloLogoRightBack.Advance(tickTime); err != nil {
			return err
		}

		if err := v.diabloLogoLeft.Advance(tickTime); err != nil {
			return err
		}

		if err := v.diabloLogoRight.Advance(tickTime); err != nil {
			return err
		}
	}

	return nil
}

// OnMouseButtonDown is called when a mouse button is clicked
func (v *MainMenu) OnMouseButtonDown(event c2interface.MouseEvent) bool {
	if v.screenMode == ScreenModeTrademark && event.Button() == c2enum.MouseButtonLeft {
		v.SetScreenMode(ScreenModeMainMenu)
		return true
	}

	return false
}

func (v *MainMenu) onEscapePressed(event c2interface.KeyEvent, mode mainMenuScreenMode) {
	if event.Key() == c2enum.KeyEscape {
		v.SetScreenMode(mode)
	}
}

// OnKeyUp is called when a key is released
func (v *MainMenu) OnKeyUp(event c2interface.KeyEvent) bool {
	preventKeyEventPropagation := false

	switch v.screenMode {
	case ScreenModeTrademark: // On retail version of D2, some specific key events (Escape, Space and Enter) puts you onto the main menu.
		switch event.Key() {
		case c2enum.KeyEscape, c2enum.KeyEnter, c2enum.KeySpace:
			v.SetScreenMode(ScreenModeMainMenu)
		}

		preventKeyEventPropagation = true
	case ScreenModeMainMenu: // pressing escape in Main Menu close the game
		if event.Key() == c2enum.KeyEscape {
			v.onExitButtonClicked()
		}
	case ScreenModeMultiplayer: // back to previous menu
		v.onEscapePressed(event, ScreenModeMainMenu)

		preventKeyEventPropagation = true
	case ScreenModeTCPIP: // back to previous menu
		v.onEscapePressed(event, ScreenModeMultiplayer)

		preventKeyEventPropagation = true
	case ScreenModeServerIP: // back to previous menu
		v.onEscapePressed(event, ScreenModeTCPIP)

		preventKeyEventPropagation = true
	}

	return preventKeyEventPropagation
}

// SetScreenMode sets the screen mode (which sub-menu the screen is on)
func (v *MainMenu) SetScreenMode(screenMode mainMenuScreenMode) {
	v.screenMode = screenMode
	isMainMenu := screenMode == ScreenModeMainMenu
	isMultiplayer := screenMode == ScreenModeMultiplayer
	isTCPIP := screenMode == ScreenModeTCPIP
	isServerIP := screenMode == ScreenModeServerIP

	v.exitDiabloButton.SetVisible(isMainMenu)
	v.creditsButton.SetVisible(isMainMenu)
	v.cinematicsButton.SetVisible(isMainMenu)
	v.singlePlayerButton.SetVisible(isMainMenu)
	v.githubButton.SetVisible(isMainMenu)
	v.mapTestButton.SetVisible(isMainMenu)
	v.multiplayerButton.SetVisible(isMainMenu)
	v.networkTCPIPButton.SetVisible(isMultiplayer)
	v.networkCancelButton.SetVisible(isMultiplayer)
	v.btnTCPIPCancel.SetVisible(isTCPIP)
	v.btnTCPIPHostGame.SetVisible(isTCPIP)
	v.btnTCPIPJoinGame.SetVisible(isTCPIP)
	v.tcpJoinGameEntry.SetVisible(isServerIP)

	if isServerIP {
		v.tcpJoinGameEntry.Activate()
	}

	v.btnServerIPOk.SetVisible(isServerIP)
	v.btnServerIPCancel.SetVisible(isServerIP)
}
