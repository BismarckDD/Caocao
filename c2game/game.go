package c2game

import (
	"image"
	_ "image/png" // we must import this to decode png file.
	"log"

	"github.com/BismarckDD/Caocao/c2core/c2uires"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
)

var mainMenuBgImg *ebiten.Image
var menuImg *ebiten.Image

type Option struct {
	windowTitle string
}

type Game struct {
	logger *log.Logger
	// c2renderer *c2renderer.Renderer
	option     *Option
	ebitenUI   *ebitenui.UI
	containers []*widget.PreferredSizeLocateableWidget
	uiRes      *c2uires.UIResources
	closeGame  func()
}

// 这里可以增加一个configuration的路径，用于初始化游戏的configuration
func CreateGame() (*Game, bool, error) {

	option := &Option{windowTitle: "三国志曹操传 20220126 made by Michael Ding."}
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowResizable(false)
	// ebiten.SetWindowFloating(true)
	// ebiten.SetWindowIcon() // not work on MacOS.
	ebiten.SetWindowTitle(option.windowTitle)
	ebiten.SetMaxTPS(30) // SLG 游戏 30 is enough.
	ebiten.SetScreenClearedEveryFrame(false)
	// set = true, will inhibit close operation.
	// need to handle on ebiten.
	ebiten.SetWindowClosingHandled(false)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)

	var err error
	// mainMenuBgImg, _, err = ebitenutil.NewImageFromFile(c2util.GetAppPath() + "/resources/map/scenario_map/main_menu.png")
	// if err != nil {
	// 	log.Fatal("Failed to get main menu bg img. ", err.Error())
	// }

	res, err := c2uires.NewUIResources()

	if res == nil || err != nil {
		log.Fatalln("Failed to create ui resource.", err.Error())
	}

	var game *Game = &Game{
		uiRes:    res,
		option:   option,
		ebitenUI: &ebitenui.UI{},
	}
	game.ToMainMenu()
	return game, true, nil
}

func (Game *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (game *Game) Update() error {
	game.ebitenUI.Update()
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	if game.ebitenUI == nil {
		log.Fatalf("ui is nil.")
	}
	if screen == nil {
		log.Fatalf("Image is nil.")
	}

	game.ebitenUI.Draw(screen)
}

func (game *Game) GetUIRes() *c2uires.UIResources {
	return game.uiRes
}

func (game *Game) GetEbitenUI() *ebitenui.UI {
	return game.ebitenUI
}

func (game *Game) ToMainMenu() {

	// the outermost container muse use grid layout
	// otherwise the inner container will collapse (disappear.).
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Padding(widget.Insets{Top: 0, Left: 0, Bottom: 0, Right: 0}),
			widget.GridLayoutOpts.Spacing(0, 0),
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, false, true}),
		)),
		// widget.ContainerOpts.BackgroundImage(game.GetUIRes().Background)
	)

	// container := widget.NewContainer(
	// 	widget.ContainerOpts.Layout(widget.NewRowLayout(
	// 		widget.RowLayoutOpts.Padding(widget.Insets{Top: 0, Left: 0, Bottom: 0, Right: 0}),
	// 		widget.RowLayoutOpts.Spacing(0),
	// 		widget.RowLayoutOpts.Direction(widget.DirectionVertical),
	// 	)),
	// 	widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
	// 		widget.RowLayoutData{
	// 			Stretch:   true,
	// 			MaxWidth:  1280,
	// 			MaxHeight: 960,
	// 			Position:  0,
	// 		},
	// 	)),
	// widget.ContainerOpts.BackgroundImage(game.GetUIRes().Background))

	game.containers = []*widget.PreferredSizeLocateableWidget{}
	c1 := CreateWindowsStyleMenu(game.uiRes)
	game.containers = append(game.containers, &c1)
	container.AddChild(*game.containers[0])
	c2 := CreateGameStyleMenu(game.uiRes)
	game.containers = append(game.containers, &c2)
	container.AddChild(*game.containers[1])
	c3 := CreateMainMenuBackgroundContainer(game)
	game.containers = append(game.containers, &c3)
	container.AddChild(*game.containers[2])
	container.SetLocation(image.Rect(0, 0, 1280, 960))
	game.ebitenUI.Container = container
	game.ebitenUI.ToolTip = nil
}

func (game *Game) ToScenario(scenarioId int) {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 0, Left: 0, Bottom: 0, Right: 0}),
			widget.RowLayoutOpts.Spacing(0),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:  true,
				Position: 0,
			}),
		),
		widget.ContainerOpts.BackgroundImage(game.uiRes.Background))

	container.AddChild(CreateWindowsStyleMenu(game.uiRes))
	container.AddChild(CreateGameStyleMenu(game.uiRes))
	container.AddChild(CreateMainMenuBackgroundContainer(game))
	game.ebitenUI.Container = container
}

func (game *Game) ToBattle(battleId int) {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 0, Left: 0, Bottom: 0, Right: 0}),
			widget.RowLayoutOpts.Spacing(0),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:  true,
				Position: 0,
			}),
		),
		widget.ContainerOpts.BackgroundImage(game.uiRes.Background))

	container.AddChild(CreateWindowsStyleMenu(game.uiRes))
	container.AddChild(CreateGameStyleMenu(game.uiRes))
	container.AddChild(CreateMainMenuBackgroundContainer(game))
	game.ebitenUI.Container = container
}

func (game *Game) ToPreBattle(preBattleId int) {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 0, Left: 0, Bottom: 0, Right: 0}),
			widget.RowLayoutOpts.Spacing(0),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:  true,
				Position: 0,
			}),
		),
		widget.ContainerOpts.BackgroundImage(game.uiRes.Background))

	container.AddChild(CreateWindowsStyleMenu(game.uiRes))
	container.AddChild(CreateGameStyleMenu(game.uiRes))
	container.AddChild(CreateMainMenuBackgroundContainer(game))
	game.ebitenUI.Container = container
}

func (game *Game) ToEndMovie(endId int) {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 0, Left: 0, Bottom: 0, Right: 0}),
			widget.RowLayoutOpts.Spacing(0),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:  true,
				Position: 0,
			}),
		),
		widget.ContainerOpts.BackgroundImage(game.uiRes.Background))

	container.AddChild(CreateWindowsStyleMenu(game.uiRes))
	container.AddChild(CreateGameStyleMenu(game.uiRes))
	container.AddChild(CreateMainMenuBackgroundContainer(game))
	game.ebitenUI.Container = container
}

func (game *Game) ToExitPic() {
}
