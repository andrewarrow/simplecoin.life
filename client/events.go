package client

import "github.com/dontpanic92/wxGo/wx"
import "os"

func (f *TheFrame) evtThread(e wx.Event) {
	te := wx.ToThreadEvent(e)
	switch {
	case te.GetInt() == 1:
		login(f)
	case te.GetInt() == 2:
		promptForLogin(f)
	case te.GetInt() == 3:
		take(f)
	case te.GetInt() == 4:
		send(f)
	}
}

func (f *TheFrame) evtQuit(wx.Event) {
	os.Exit(0)
}

func (f *TheFrame) evtLogout(wx.Event) {
	go func() {
		f.SendEvent(2)
	}()
}

func (f *TheFrame) evtTake(wx.Event) {
	go func() {
		f.SendEvent(3)
	}()
}
func (f *TheFrame) evtSend(wx.Event) {
	go func() {
		f.SendEvent(4)
	}()
}

func (f *TheFrame) evtLogin(wx.Event) {
	go func() {
		f.SendEvent(1)
	}()
}
