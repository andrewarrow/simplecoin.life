package client

import "github.com/dontpanic92/wxGo/wx"

var ui_token wx.TextCtrl
var ui_add wx.Button
var ui_bar wx.Gauge
var ui_percent wx.StaticText

func setupFeeds(f *TheFrame) {
	f.sizer.Clear(true)
}

func setupError(f *TheFrame, text string) {
	row3 := wx.NewBoxSizer(wx.VERTICAL)
	msg := wx.NewStaticText(f.frame, wx.ID_ANY, text, wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(msg, 0, wx.ALL|wx.EXPAND, 5)
	msg2 := wx.NewStaticText(f.frame, wx.ID_ANY, "Please quit app and try again.", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(msg2, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row3, 0, wx.ALL|wx.EXPAND, 5)
	f.frame.Layout()
}

func setupAddToken(f *TheFrame) {
	row := wx.NewBoxSizer(wx.HORIZONTAL)
	msg := wx.NewStaticText(f.frame, wx.ID_ANY, "Name", wx.DefaultPosition, wx.DefaultSize, 0)
	row.Add(msg, 0, wx.ALL|wx.EXPAND, 5)
	ui_token = wx.NewTextCtrl(f.frame, wx.ID_ANY, "", wx.DefaultPosition, wx.NewSize(380, 25), 0)
	row.Add(ui_token, 0, wx.ALL|wx.EXPAND, 5)

	row3 := wx.NewBoxSizer(wx.HORIZONTAL)
	ui_add = wx.NewButton(f.frame, wx.ID_ANY, "Generate Name", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(ui_add, 0, wx.ALL|wx.FIXED_MINSIZE, 5)

	f.sizer.Add(row, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row3, 0, wx.ALL|wx.EXPAND, 5)

	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtAddToken, ui_add.GetId())
	f.frame.Layout()
}
