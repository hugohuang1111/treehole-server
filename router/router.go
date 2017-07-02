package router

import "github.com/hugohuang1111/treehole/module"
import "github.com/golang/glog"

var (
	mailChan chan *module.Mail
	runFlag  bool
)

// Start router start
func Start() {
	mailChan = make(chan *module.Mail, 8)
	runFlag = true
	go run()
}

// End router end
func End() {
	runFlag = false
}

//Route route message
func Route(msg *module.Mail) {
	glog.Infof("Router %s->%s %s", msg.Sender, msg.Recver, msg.Type)
	mailChan <- msg
}

func run() {
	for runFlag {
		mail := <-mailChan
		m := module.Find(mail.Recver)
		if nil == m {
			glog.Warning("Router can't find module: ", mail.Recver)
			continue
		}
		m.OnMail(mail)
	}
}
