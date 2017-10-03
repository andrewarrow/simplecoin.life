package client

import "github.com/dontpanic92/wxGo/wx"

//import "fmt"
//import "time"

const THE_WORKER_ID = wx.ID_HIGHEST + 1

var username, password, team string

type TheFrame struct {
	frame wx.Frame
	sizer wx.BoxSizer
}

func NewTheFrame() TheFrame {
	f := TheFrame{}
	f.frame = wx.NewFrame(wx.NullWindow, -1, "simplecoin.life", wx.DefaultPosition, wx.NewSize(650, 300), wx.DEFAULT_FRAME_STYLE|wx.TAB_TRAVERSAL)

	menubar := wx.NewMenuBar()
	menuFile := wx.NewMenu()
	menuHelp := wx.NewMenu()

	logout := wx.NewMenuItem(menuFile, wx.ID_ANY, "&Logout", "Logout", wx.ITEM_NORMAL)
	menuFile.Append(logout)
	wx.Bind(f.frame, wx.EVT_MENU, f.evtLogout, logout.GetId())

	quit := wx.NewMenuItem(menuFile, wx.ID_ANY, "&Quit", "Quit", wx.ITEM_NORMAL)
	menuFile.Append(quit)
	wx.Bind(f.frame, wx.EVT_MENU, f.evtQuit, quit.GetId())

	menuItemAbout := wx.NewMenuItem(menuFile, wx.ID_ANY, "&About", "About", wx.ITEM_NORMAL)
	menuHelp.Append(menuItemAbout)
	menubar.Append(menuFile, "&File")
	menubar.Append(menuHelp, "&Help")
	f.frame.SetMenuBar(menubar)

	f.sizer = wx.NewBoxSizer(wx.VERTICAL)

	f.frame.SetSizer(f.sizer)

	promptForLogin(&f)

	wx.Bind(f.frame, wx.EVT_THREAD, f.evtThread, THE_WORKER_ID)
	return f
}

func Setup() {
	/*
		i := 0
		for {
			fmt.Printf("%s 10000 %d %s\n", xWords(), time.Now().Unix(), "genesis")
			i++
			if i == 1000 {
				break
			}
		}*/
	wx1 := wx.NewApp()
	f := NewTheFrame()
	f.frame.Show()
	wx1.MainLoop()
}
