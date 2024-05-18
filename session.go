package main

import (
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

type Session struct {
	id           string
	currUrl      string
	sessRepeat   int
	sessDispatch *SessDispatch
}

func NewSession(id string, sessDispatch *SessDispatch, sessRepeat int) *Session {
	return &Session{
		id:           id,
		currUrl:      "",
		sessRepeat:   sessRepeat,
		sessDispatch: sessDispatch,
	}
}

type TsSegment struct {
	url    string
	extinf int64 //播放时长毫秒
}

func (sess *Session) DoNgod() (err error) {
	logrus.Info("id:[", sess.id, "]--url: ", sess.currUrl)

	return errors.New("i am error")
}

func (sess *Session) Do() {
	var err error

	for {
		if len(sess.currUrl) == 0 {
			sess.currUrl, err = sess.sessDispatch.getUrlFromPlayList()
			if err != nil {
				logrus.Error(err)
				break
			}
		}

		err = sess.DoNgod()

		if err != nil {
			time.Sleep(time.Second * 1)
		}

	}
}
