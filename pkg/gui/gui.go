package gui

import (
	"errors"
	"fmt"
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/config"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/service"
)

const (
	ServiceView     = "Services"
	InfoView        = "Info"
	KeybindingsView = "Keybindings"
)

var (
	viewArr  = []string{ServiceView, InfoView}
	active   = 0
	services []*service.Service
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

// Selects the line below
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

// Selects the line above
func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func printServiceInfo(v *gocui.View, s *service.Service) error {
	fmt.Fprintf(v, "Selected service: \033[32;7m%s\033[0m\n", s.Name)
	fmt.Fprintf(v, "IDL Path:\n\033[32;4m%s\033[0m\n", s.Path)
	fmt.Fprintln(v, "\nRoutes:")
	for method, routes := range s.Routes {
		for _, route := range routes {
			fmt.Fprintf(v, "\033[33;7m %s \033[0m: \033[35;1m%s\033[0m\n", method, route)

		}
	}
	return nil
}

func selectService(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	infoView, err := g.View(InfoView)
	if err != nil {
		return err
	}
	infoView.Clear()

	for _, s := range services {
		if s.Name == l {
			printServiceInfo(infoView, s)
		}
	}

	return nil
}

// Takes a variable number of functions and calls them in sequence
func appendHandlers(handlers ...func(*gocui.Gui, *gocui.View) error) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		for _, h := range handlers {
			if err := h(g, v); err != nil {
				return err
			}
		}
		return nil
	}
}

// Sets the colors for the selected lines
func setHighlightColors(v *gocui.View) {
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack
}

func refreshServiceView(g *gocui.Gui, v *gocui.View) error {
	serviceView, err := g.View(ServiceView)
	if err != nil {
		return err
	}
	serviceView.Clear()

	cfg := config.ReadConfig()
	ss, err := service.GetServicesFromIDLDir(cfg.IDLDir)
	if err != nil {
		return err
	}

	services = ss
	for _, service := range services {
		fmt.Fprintln(serviceView, service.Name)
	}

	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(ServiceView, 0, 0, maxX*1/3-1, maxY-4, gocui.RIGHT); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Title = "Services"
		v.Autoscroll = true
		v.Highlight = true

		fmt.Fprintln(v, "View with default frame color")
		fmt.Fprintln(v, "It's connected to v2 with overlay RIGHT.")

		setHighlightColors(v)

		if _, err = setCurrentViewOnTop(g, "Services"); err != nil {
			return err
		}

		if err := refreshServiceView(g, v); err != nil {
			return err
		}
	}

	if v, err := g.SetView(InfoView, maxX*1/3-1, 0, maxX-1, maxY-4, gocui.LEFT); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Title = "Info"
		v.Wrap = true
		v.FrameColor = gocui.ColorMagenta
		v.FrameRunes = []rune{'═', '│'}
	}

	if v, err := g.SetView(KeybindingsView, 0, maxY-4, maxX-1, maxY-1, gocui.TOP); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Title = "Keybindings"
		v.Editable = true
		v.TitleColor = gocui.ColorYellow
		v.FrameColor = gocui.ColorRed
		v.FrameRunes = []rune{'─', '│', '┌', '┐', '└', '┘'}
		fmt.Fprint(v, "\033[31;1mCtrl+C\033[0m: Quit\t\t")
		fmt.Fprint(v, "\033[31;1mTab\033[0m: Change current view\t\t")
		fmt.Fprint(v, "\033[31;1mCtrl+R\033[0m: Refresh API Gateway\t\t")
		fmt.Fprintln(v)
		fmt.Fprintln(v, "\033[32;2mSelected frame is highlighted with green color\033[0m")
	}
	return nil
}

func toggleOverlap(g *gocui.Gui, v *gocui.View) error {
	g.SupportOverlaps = !g.SupportOverlaps
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Run() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SelFrameColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlO, gocui.ModNone, toggleOverlap); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlR, gocui.ModNone, refreshServiceView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding(
		ServiceView,
		gocui.KeyArrowDown,
		gocui.ModNone,
		appendHandlers(cursorDown, selectService),
	); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(
		ServiceView,
		gocui.KeyArrowUp,
		gocui.ModNone,
		appendHandlers(cursorUp, selectService),
	); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(ServiceView, gocui.KeyEnter, gocui.ModNone, selectService); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}

}
