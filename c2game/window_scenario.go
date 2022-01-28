package c2game

import (
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"

	"github.com/BismarckDD/Caocao/c2core/c2uires"
)

func CreateGameMenuBattleField(uiRes *c2uires.UIResources, ui func() *ebitenui.UI) *c2uires.Page {
	return &c2uires.Page{
		Title:   "",
		Content: nil,
	}
}

func CreateGameMenuNormal(uiRes *c2uires.UIResources, ui func() *ebitenui.UI) *c2uires.Page {

	menuContainer := widget.NewContainer()

	exitButton := widget.NewButton()
	loadButton := widget.NewButton()
	saveButton := widget.NewButton()
	configButton := widget.NewButton()
	whitespace := widget.NewButton()
	charButton := widget.NewButton()
	itemButton := widget.NewButton()
	item1Button := widget.NewButton()
	item2Button := widget.NewButton()

	menuContainer.AddChild(exitButton)
	menuContainer.AddChild(loadButton)
	menuContainer.AddChild(saveButton)
	menuContainer.AddChild(configButton)
	menuContainer.AddChild(whitespace)
	menuContainer.AddChild(charButton)
	menuContainer.AddChild(itemButton)
	menuContainer.AddChild(loadButton)
	menuContainer.AddChild(saveButton)
	menuContainer.AddChild(configButton)
	menuContainer.AddChild(item1Button)
	menuContainer.AddChild(item2Button)

	// chapter info -- part info
	// red-blue progress
	infoContainer := widget.NewContainer()

	chapterInfoContainer := widget.NewContainer()
	redBlueProgressBarContainer := widget.NewContainer()

	infoContainer.AddChild((chapterInfoContainer))
	infoContainer.AddChild((redBlueProgressBarContainer))
	menuContainer.AddChild(infoContainer)

	return &c2uires.Page{
		Title:   "",
		Content: menuContainer,
	}
}
