package client

import "github.com/dontpanic92/wxGo/wx"

func (f *TheFrame) SendEvent(id int) {
	threadEvent := wx.NewThreadEvent(wx.EVT_THREAD, THE_WORKER_ID)
	threadEvent.SetInt(id)
	f.frame.QueueEvent(threadEvent)
}
