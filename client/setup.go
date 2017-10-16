package client

import "github.com/dontpanic92/wxGo/wx"
import "github.com/andrewarrow/simplecoin.life/words"
import "github.com/andrewarrow/simplecoin.life/sql"
import "strconv"
import "fmt"
import "strings"

//import "math"

var ui_from wx.TextCtrl
var ui_amount wx.TextCtrl
var ui_to wx.TextCtrl
var ui_add wx.Button
var ui_bar wx.Gauge
var ui_percent wx.StaticText

var balance wx.StaticText
var currentUser string

func send(f *TheFrame) {
	to := ui_to.GetLineText(0)
	db := sql.SqlInit()
	id := sql.FindAvailableCoin(currentUser, db)
	if id != "" {
		fmt.Println(id)
		sql.TransferCoin(to, id, db)
		bump(-0.01)
	}
}

func bump(amount float64) {
	coins, _ := strconv.ParseFloat(balance.GetLabelText(), 10)
	coins += amount
	c := fmt.Sprintf("%f", coins)
	index := strings.Index(c, ".")
	balance.SetLabelText(c[0 : index+3])
}

func take(f *TheFrame) {
	db := sql.SqlInit()
	sql.TransferCoinFromGenesis(currentUser, db)
	bump(0.01)
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

	db := sql.SqlInit()
	coins := (float64(sql.CountByOwner(currentUser, db))) / 100.0
	fmt.Println(coins)
	c := fmt.Sprintf("%f", coins)
	fmt.Println(c)
	index := strings.Index(c, ".")
	cc := c[0 : index+3]

	balance = wx.NewStaticText(f.frame, wx.ID_ANY, cc, wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(balance, 0, wx.ALL|wx.EXPAND, 5)
	take := wx.NewButton(f.frame, wx.ID_ANY, "Take Coin", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(take, 0, wx.ALL|wx.FIXED_MINSIZE, 5)

	row4 := wx.NewBoxSizer(wx.HORIZONTAL)
	ttg := wx.NewStaticText(f.frame, wx.ID_ANY, "Send To", wx.DefaultPosition, wx.DefaultSize, 0)
	row4.Add(ttg, 0, wx.ALL|wx.EXPAND, 5)
	ui_to = wx.NewTextCtrl(f.frame, wx.ID_ANY, words.BigWords(), wx.DefaultPosition, wx.NewSize(380, 25), 0)
	row4.Add(ui_to, 0, wx.ALL|wx.EXPAND, 5)
	send := wx.NewButton(f.frame, wx.ID_ANY, "Send", wx.DefaultPosition, wx.DefaultSize, 0)
	row4.Add(send, 0, wx.ALL|wx.FIXED_MINSIZE, 5)

	f.sizer.Add(row, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row3, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row4, 0, wx.ALL|wx.EXPAND, 5)

	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtLogout, ui_add.GetId())
	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtTake, take.GetId())
	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtSend, send.GetId())
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
