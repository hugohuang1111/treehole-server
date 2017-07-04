package internal

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang/glog"
)

//NewConnectHandler new connection handler
type NewConnectHandler func() Connect

//DisConnectHandler disconnect handler
type DisConnectHandler func(id uint64)

var (
	httpServer *http.Server
	newHandler NewConnectHandler
	disHandler DisConnectHandler
)

//StartHTTPServer start http server
func StartHTTPServer(hNew NewConnectHandler, hDis DisConnectHandler) {
	newHandler = hNew
	disHandler = hDis
	httpServer := &http.Server{
		Addr:           ":8000",
		Handler:        http.HandlerFunc(httpHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	glog.Info("Gate HTTP server listen on 8000")
	httpServer.ListenAndServe()
}

//StopHTTPServer stop http server
func StopHTTPServer() {
	// httpServer.Close()
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if nil == newHandler {
		w.WriteHeader(500)
		w.Write([]byte("handler not set"))
		return
	}
	c := newHandler()
	exitCh := make(chan bool, 1)

	go func() {
		//send routine
		runFlag := true
		for runFlag {
			select {
			case data := <-c.SendChan():
				w.Write(data)
				runFlag = false
			case <-c.CloseSig():
				runFlag = false
			}
		}

		exitCh <- true
	}()

	cmd := make(map[string]interface{})
	query := r.URL.Query()

	glog.Info("method ", r.Method)
	glog.Info("url ", r.URL.Path)
	glog.Info("params ", query)

	switch r.URL.Path {
	case "/api/express":
		cmd["cmd"] = "express"
		cmd["word"] = query.Get("word")
		cmd["nickName"] = query.Get("nickName")
		cmd["datetime"] = strconv.FormatInt(time.Now().Unix(), 10)
	case "/api/topexpress":
		cmd["cmd"] = "topexpress"
		cmd["start"], _ = strconv.Atoi(query.Get("start"))
		cmd["length"], _ = strconv.Atoi(query.Get("length"))
	default:
		w.WriteHeader(404)
		glog.Warning("HTTPServer invalid path")
		return
	}

	recv := c.RecvChan()
	recv <- cmd
	<-exitCh
	disHandler(c.ID())
}
