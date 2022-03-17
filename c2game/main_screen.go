package c2game

import (
	"image"
	"log"
	"os"
	"time"

	"github.com/BismarckDD/Caocao/c2core/c2uires"
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame(res *c2uires.UIResources, ui func() *ebitenui.UI) {
	return
}

func LoadGame(res *c2uires.UIResources, ui func() *ebitenui.UI) {
	return
}

func ConfigGame(res *c2uires.UIResources, ui func() *ebitenui.UI) {
	return
}

func ExitGame(game *Game) {
	game.ToExitPic()
	time.Sleep(time.Duration(2) * time.Second)
	log.Println("Exit the game.")
	os.Exit(0)
}

// MainMenuContainer
// ---- Background Image
// ---- ButtonGroup
func CreateMainMenuBackgroundContainer(game *Game) widget.PreferredSizeLocateableWidget {

	// 9-slice will only fully streach the center part
	// partly or not streach the boarder part to keep the image vision normal.
	// so we'd better input the actual size of the pic at the last-2 params.
	backgroundImage, _ := c2uires.LoadImageNineSlice("resources/high_res/cover/main_menu.png", 2560, 1600)
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(backgroundImage),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Padding(widget.Insets{Top: 0, Bottom: 0, Left: 0, Right: 0}),
			widget.GridLayoutOpts.Spacing(10, 10),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}), // 需要strecth一下
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.GridLayoutData{
				MaxWidth:           1280,
				MaxHeight:          800,
				HorizontalPosition: 400,
				VerticalPosition:   0,
			},
		)),
	)
	// container.AddChild(CreateMainPageMenuWindow(game))
	// SetLocation可以控制大小和位置
	container.SetLocation(image.Rect(0, 0, 1280, 800))
	// CreateMainPageMenuWindow(game)
	return container
}

func CreateMainPageMenuWindow(game *Game) {

	var removeWindow ebitenui.RemoveWindowFunc

	buttonGroupContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(game.GetUIRes().Background),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(15),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 1, Left: 40}),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:  true,
				Position: 20,
			},
		)),
	)

	newGameButton := widget.NewButton(
		widget.ButtonOpts.Image(game.uiRes.MainMenuButton.Image),
		widget.ButtonOpts.TextPadding(game.uiRes.MainMenuButton.Padding),
		widget.ButtonOpts.Text("开始游戏", game.uiRes.MainMenuButton.Face, game.uiRes.MainMenuButton.TextColor),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			removeWindow()
			ExitGame(game)
		}),
	)

	loadGameButton := widget.NewButton(
		widget.ButtonOpts.Image(game.uiRes.MainMenuButton.Image),
		widget.ButtonOpts.TextPadding(game.uiRes.MainMenuButton.Padding),
		widget.ButtonOpts.Text("读取进度", game.uiRes.MainMenuButton.Face, game.uiRes.MainMenuButton.TextColor),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			// removeWindow()
			ExitGame(game)
		}),
	)

	configGameButton := widget.NewButton(
		widget.ButtonOpts.Image(game.uiRes.MainMenuButton.Image),
		widget.ButtonOpts.TextPadding(game.uiRes.MainMenuButton.Padding),
		widget.ButtonOpts.Text("环境设定", game.uiRes.MainMenuButton.Face, game.uiRes.MainMenuButton.TextColor),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			// removeWindow()
			ExitGame(game)
		}),
	)

	exitGameButton := widget.NewButton(
		widget.ButtonOpts.Image(game.uiRes.MainMenuButton.Image),
		widget.ButtonOpts.TextPadding(game.uiRes.MainMenuButton.Padding),
		widget.ButtonOpts.Text("结束游戏", game.uiRes.MainMenuButton.Face, game.uiRes.MainMenuButton.TextColor),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			removeWindow()
			ExitGame(game)
		}),
	)

	buttonGroupContainer.AddChild(newGameButton)
	buttonGroupContainer.AddChild(loadGameButton)
	buttonGroupContainer.AddChild(configGameButton)
	buttonGroupContainer.AddChild(exitGameButton)

	buttonGroupWindow := widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(buttonGroupContainer),
	)

	// 这里的location: window和rect的位置完全相同
	ww, wh := ebiten.WindowSize()
	location := image.Rect(0, 0, ww/3, wh/2)
	location = location.Add(image.Point{ww / 3, wh / 4})
	buttonGroupWindow.SetLocation(location)
	removeWindow = game.ebitenUI.AddWindow(buttonGroupWindow)
}

// This menu contains "file", "function", "info" button and never changes.
func CreateWindowsStyleMenu(res *c2uires.UIResources) widget.PreferredSizeLocateableWidget {

	menuContainer := widget.NewContainer(

		// widget.ContainerOpts.BackgroundImage(res.Header.Background),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(res.MainMenuButton.Padding),
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(0),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:   true,
				MaxWidth:  1280,
				MaxHeight: 60,
				Position:  0,
			})),
	)

	menuContainer.AddChild(widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: 40,
		})),
		widget.TextOpts.Text("文件", res.WindowMenu.Face, res.WindowMenu.Color),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	))

	menuContainer.AddChild(widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: 40,
		})),
		widget.TextOpts.Text("    ", res.WindowMenu.Face, res.WindowMenu.Color),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	))

	menuContainer.AddChild(widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: 40,
		})),
		widget.TextOpts.Text("功能", res.WindowMenu.Face, res.WindowMenu.Color),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	))

	menuContainer.AddChild(widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: 40,
		})),
		widget.TextOpts.Text("    ", res.WindowMenu.Face, res.WindowMenu.Color),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	))

	menuContainer.AddChild(widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: 40,
		})),
		widget.TextOpts.Text("情报", res.WindowMenu.Face, res.WindowMenu.Color),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	))

	menuContainer.SetLocation(image.Rect(0, 0, 1280, 20)) // size: 1280 * 40
	return menuContainer
}

// create non-battle game-style menu
// exit/save/load/info, xxxx/xxxx/xxxx/xxx,
// scenairo - round(current/total)
func CreateGameStyleMenu(res *c2uires.UIResources) widget.PreferredSizeLocateableWidget {

	menuContainer := widget.NewContainer(
		// widget.ContainerOpts.BackgroundImage(),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(res.MainMenuButton.Padding),
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(0),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:   true,
				MaxWidth:  1280,
				MaxHeight: 120,
				Position:  0,
			})),
	)

	for i := 0; i < 4; i++ {
		button := widget.NewButton(
			widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch:   true,
				MaxWidth:  120,
				MaxHeight: 120,
				Position:  0,
			})),
			widget.ButtonOpts.Image(res.MainMenuButton.Image),
			// widget.ButtonOpts.Text(fmt.Sprintf("Button %d", i+1), res.MainMenuButton.Face, res.MainMenuButton.TextColor),
			widget.ButtonOpts.TextPadding(widget.Insets{Top: 16, Left: 16}),
			// widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			// 	log.Printf("Click %d button.", i)
			// }),
		)
		menuContainer.AddChild(button)
	}

	menuContainer.SetLocation(image.Rect(0, 0, 1280, 140)) // size: 1280 * 120
	return menuContainer
}
