package client

import "github.com/dontpanic92/wxGo/wx"

import "fmt"

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

func NewLogFrame() TheFrame {
	f := TheFrame{}
	f.frame = wx.NewFrame(wx.NullWindow, -1, "simplecoin.life", wx.NewPoint(200, 70), wx.NewSize(950, 700), wx.DEFAULT_FRAME_STYLE|wx.TAB_TRAVERSAL)

	f.sizer = wx.NewBoxSizer(wx.VERTICAL)

	f.frame.SetSizer(f.sizer)

	grid := wx.NewGrid(f.frame, wx.ID_ANY, wx.DefaultPosition, wx.DefaultSize, 0)
	grid.CreateGrid(20, 5)
	grid.EnableEditing(false)
	grid.EnableGridLines(true)
	grid.SetColLabelValue(0, "created")
	grid.SetColLabelValue(1, "id")
	grid.SetColLabelValue(2, "owner")
	grid.SetColLabelValue(3, "previous")
	grid.SetColLabelValue(4, "transfered")

	db := SqlInit()
	list := TransactionsFrom(db)
	for i, t := range list {
		grid.SetCellValue(i, 0, fmt.Sprintf("%d", t.Created))
		grid.SetCellValue(i, 1, t.Id)
		grid.SetCellValue(i, 2, t.Owner)
		grid.SetCellValue(i, 3, t.Previous)
		grid.SetCellValue(i, 4, fmt.Sprintf("%d", t.Transfered))
	}
	f.sizer.Add(grid, 0, wx.ALL|wx.EXPAND, 5)
	f.frame.Layout()

	wx.Bind(f.frame, wx.EVT_THREAD, f.evtThread, THE_WORKER_ID)
	return f
}

func Setup() {
	wx1 := wx.NewApp()
	f := NewTheFrame()
	f.frame.Show()
	f2 := NewLogFrame()
	f2.frame.Show()
	wx1.MainLoop()
}
