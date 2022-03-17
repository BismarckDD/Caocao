package c2screen

import (
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2util"
)

const (
	logPrefix = "Screen Manager"
)

// ScreenManager manages game screens (main menu, scenario, camp, battle, etc)
type ScreenManager struct {
	currentScreen Screen
	nextScreen    Screen
	loadingScreen Screen
	loadingState  LoadingState
	uiManager     *c2ui.UIManager
	guiManager    *c2gui.GuiManager
	*c2util.Logger
}

// NewScreenManager creates a screen manager
func NewScreenManager(uiManager *c2ui.UIManager, guiManager *c2gui.GuiManager, l c2util.LogLevel) *ScreenManager {
	sm := &ScreenManager{
		uiManager:  uiManager,
		guiManager: guiManager,
	}

	sm.Logger = c2util.NewLogger()
	sm.Logger.SetPrefix(logPrefix)
	sm.Logger.SetLevel(l)

	return sm
}

// SetNextScreen is about to set a given screen as next
func (sm *ScreenManager) SetNextScreen(screen Screen) {
	sm.nextScreen = screen
}

// Advance updates the UI on every frame
func (sm *ScreenManager) Advance(elapsed float64) error {
	switch {
	case sm.loadingScreen != nil: // 适用于当前处于loading的状态
		// this call blocks execution and could lead to deadlock if a screen implements OnLoad incorreclty
		load, ok := <-sm.loadingState.updates // updates is a chan, so will block if no data.
		if !ok {
			sm.Warning("loadingState chan should not be closed while in a loading screen")
		}

		if load.err != nil {
			sm.Errorf("PROBLEM LOADING THE SCREEN: %v", load.err)
			return load.err
		}

		sm.guiManager.ShowLoadScreen(load.progress)

		if load.done {
			sm.currentScreen = sm.loadingScreen
			sm.loadingScreen = nil

			sm.guiManager.ShowCursor()
			sm.guiManager.HideLoadScreen()
		}
	case sm.nextScreen != nil: // 适用于即将进入切屏的状态
		if handler, ok := sm.currentScreen.(ScreenUnloadHandler); ok {
			if err := handler.OnUnload(); err != nil {
				return err
			}
		}

		sm.uiManager.Reset()
		sm.guiManager.SetLayout(nil)

		// 这里其实就是判断 NextScreen 是否是 LoadingScreen
		if handler, ok := sm.nextScreen.(ScreenLoadHandler); ok {
			sm.guiManager.ShowLoadScreen(0)
			sm.guiManager.HideCursor()

			sm.loadingState = LoadingState{updates: make(chan loadingUpdate)}

			go func() {
				handler.OnLoad(sm.loadingState)
				sm.loadingState.Done()
			}()

			sm.currentScreen = nil
			sm.loadingScreen = sm.nextScreen
		} else {
			sm.currentScreen = sm.nextScreen
			sm.loadingScreen = nil
		}

		sm.nextScreen = nil
	case sm.currentScreen != nil: // 正常游戏
		if handler, ok := sm.currentScreen.(ScreenAdvanceHandler); ok {
			if err := handler.Advance(elapsed); err != nil {
				return err
			}
		}
	}

	return nil
}

// Render renders the UI by a given surface
func (sm *ScreenManager) Render(surface c2interface.Surface) {
	if handler, ok := sm.currentScreen.(ScreenRenderHandler); ok {
		handler.Render(surface)
	}
}
