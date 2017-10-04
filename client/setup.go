package client

import "github.com/dontpanic92/wxGo/wx"
import "github.com/andrewarrow/simplecoin.life/words"
import "strconv"
import "fmt"

var ui_from wx.TextCtrl
var ui_amount wx.TextCtrl
var ui_to wx.TextCtrl
var ui_add wx.Button
var ui_bar wx.Gauge
var ui_percent wx.StaticText

var balance wx.StaticText
var currentUser string

func setupFeeds(f *TheFrame) {
	f.sizer.Clear(true)
}

func take(f *TheFrame) {
	coins, _ := strconv.ParseFloat(balance.GetLabelText(), 10)
	coins += 0.01
	balance.SetLabelText(fmt.Sprintf("%f", coins))
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

func login(f *TheFrame) {
	currentUser = ui_from.GetLineText(0)
	f.sizer.Clear(true)
	row := wx.NewBoxSizer(wx.HORIZONTAL)
	msg := wx.NewStaticText(f.frame, wx.ID_ANY, "Public Name", wx.DefaultPosition, wx.DefaultSize, 0)
	row.Add(msg, 0, wx.ALL|wx.EXPAND, 5)
	name := wx.NewStaticText(f.frame, wx.ID_ANY, currentUser, wx.DefaultPosition, wx.DefaultSize, 0)
	row.Add(name, 0, wx.ALL|wx.EXPAND, 5)

	row3 := wx.NewBoxSizer(wx.HORIZONTAL)
	ui_add = wx.NewButton(f.frame, wx.ID_ANY, "Logout", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(ui_add, 0, wx.ALL|wx.FIXED_MINSIZE, 5)
	balanceLabel := wx.NewStaticText(f.frame, wx.ID_ANY, "Balance: ☀", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(balanceLabel, 0, wx.ALL|wx.EXPAND, 5)
	balance = wx.NewStaticText(f.frame, wx.ID_ANY, "0.00", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(balance, 0, wx.ALL|wx.EXPAND, 5)
	take := wx.NewButton(f.frame, wx.ID_ANY, "Take Coin", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(take, 0, wx.ALL|wx.FIXED_MINSIZE, 5)

	f.sizer.Add(row, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row3, 0, wx.ALL|wx.EXPAND, 5)

	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtLogout, ui_add.GetId())
	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtTake, take.GetId())
	f.frame.Layout()
}

func promptForLogin(f *TheFrame) {
	f.sizer.Clear(true)
	row := wx.NewBoxSizer(wx.HORIZONTAL)
	msg := wx.NewStaticText(f.frame, wx.ID_ANY, "Public Name", wx.DefaultPosition, wx.DefaultSize, 0)
	row.Add(msg, 0, wx.ALL|wx.EXPAND, 5)
	ui_from = wx.NewTextCtrl(f.frame, wx.ID_ANY, words.BigWords(), wx.DefaultPosition, wx.NewSize(380, 25), 0)
	row.Add(ui_from, 0, wx.ALL|wx.EXPAND, 5)

	row3 := wx.NewBoxSizer(wx.HORIZONTAL)
	ui_add = wx.NewButton(f.frame, wx.ID_ANY, "Login", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(ui_add, 0, wx.ALL|wx.FIXED_MINSIZE, 5)

	f.sizer.Add(row, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row3, 0, wx.ALL|wx.EXPAND, 5)

	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtLogin, ui_add.GetId())
	f.frame.Layout()
}
