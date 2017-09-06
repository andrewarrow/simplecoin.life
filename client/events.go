package client

import "github.com/dontpanic92/wxGo/wx"
import "os"

func (f *TheFrame) evtThread(e wx.Event) {
	te := wx.ToThreadEvent(e)
	switch {
	case te.GetInt() == 1:
		setupFeeds(f)
	case te.GetInt() == 2:
		setupFeeds(f)
	}
}

func (f *TheFrame) evtQuit(wx.Event) {
	os.Exit(0)
}

func (f *TheFrame) evtLogout(wx.Event) {
}

func (f *TheFrame) evtAddToken(wx.Event) {
	go func() {
		f.SendEvent(1)
	}()
}
