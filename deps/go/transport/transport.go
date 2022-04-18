package transport

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/centny/demoservice/deps/go/xlog"
	"github.com/codingeasygo/util/xnet"
	"github.com/codingeasygo/web"
	"golang.org/x/net/websocket"
)

type ForwardH struct {
	Remote      string
	server      *websocket.Server
	transporter xnet.Transporter
}

func NewTransportH(remote string) (forward *ForwardH, err error) {
	var transporter xnet.Transporter
	if strings.HasPrefix(remote, "tcp://") {
		transporter = xnet.RawDialerF(net.Dial)
	} else if strings.HasPrefix(remote, "ws://") || strings.HasPrefix(remote, "wss://") {
		transporter = xnet.NewWebsocketDialer()
	} else {
		err = fmt.Errorf("not supported remote %v", remote)
		return
	}
	forward = &ForwardH{
		Remote:      remote,
		transporter: transporter,
	}
	forward.server = &websocket.Server{Handler: websocket.Handler(forward.wsHandler)}
	return
}

func (f *ForwardH) SrvHTTP(w *web.Session) web.Result {
	remote, err := url.Parse(f.Remote)
	if err != nil {
		xlog.Errorf("TransportH parse remove %v fail with %v", f.Remote, err)
		w.W.WriteHeader(500)
		return w.SendPlainText(err.Error())
	}
	if expectUsername := remote.User.Username(); len(expectUsername) > 0 {
		expectPassword, _ := remote.User.Password()
		havingUsername, havingPassword, ok := w.R.BasicAuth()
		if !ok || expectUsername != havingUsername || expectPassword != havingPassword {
			xlog.Warnf("TransportH check basic auth fail with expect(%v:%v),having(%v:%v)", expectUsername, expectPassword, havingUsername, havingPassword)
			w.W.WriteHeader(401)
			return w.SendPlainText("not acccess")
		}
	}
	f.server.ServeHTTP(w.W, w.R)
	return web.Return
}

func (f *ForwardH) wsHandler(ws *websocket.Conn) {
	xlog.Infof("ForwardH start forward %v to %v", ws.Request().RemoteAddr, f.Remote)
	err := f.transporter.Transport(ws, f.Remote)
	xlog.Infof("ForwardH forward %v to %v is stopped by %v", ws.Request().RemoteAddr, f.Remote, err)
}
