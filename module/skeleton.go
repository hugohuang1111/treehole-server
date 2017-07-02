package module

import (
	"github.com/golang/glog"
)

// SkelectionProcess skelection interface
type SkelectionProcess interface {
	OnProcess(evt *Mail)
}

// Skelecton module skelecton
type Skelecton struct {
	mailList chan *Mail
	running  bool
	process  SkelectionProcess
}

// OnInit module init
func (s *Skelecton) OnInit(proc SkelectionProcess) {
	s.running = true
	s.mailList = make(chan *Mail)
	s.process = proc
	go s.work()
}

// OnDestroy module destroy
func (s *Skelecton) OnDestroy() {
	s.running = false
}

// OnMail module event
func (s *Skelecton) OnMail(mail *Mail) {
	s.mailList <- mail
}

func (s *Skelecton) work() {
	for s.running {
		mail := <-s.mailList
		if nil != s.process {
			s.process.OnProcess(mail)
		} else {
			glog.Warning("skelecton process is nil, skip one event")
		}
	}
}
