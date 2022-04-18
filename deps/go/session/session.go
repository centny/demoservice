package session

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	"sync"

	"github.com/centny/demoservice/deps/go/xlog"
	"github.com/codingeasygo/util/uuid"

	"github.com/codingeasygo/util/xmap"
	"github.com/codingeasygo/util/xsql"
	"github.com/codingeasygo/web"
	"github.com/gomodule/redigo/redis"
)

//DbSessionBuilder is the session builder by db
type DbSessionBuilder struct {
	Redis      func() redis.Conn
	Path       string
	Domain     string
	MaxAge     int
	Key        string
	ShowLog    bool
	sessionLck sync.RWMutex
	sessiones  map[string]*DbSession
}

//NewDbSessionBuilder will return new session builder.
func NewDbSessionBuilder() *DbSessionBuilder {
	return &DbSessionBuilder{
		Path:       "/",
		Domain:     "",
		Key:        "session_id",
		MaxAge:     0,
		ShowLog:    false,
		sessionLck: sync.RWMutex{},
		sessiones:  map[string]*DbSession{},
	}
}

func (d *DbSessionBuilder) log(format string, args ...interface{}) {
	if d.ShowLog {
		xlog.Debugf(format, args...)
	}
}

//create the http.Cookie with value and http.ResponseWriter/http.Request.
func (d *DbSessionBuilder) writeCookie(value string, w http.ResponseWriter, r *http.Request) *http.Cookie {
	c := d.newCookie()
	c.Value = value
	http.SetCookie(w, c)
	return c
}

//FindSession is impl func to web.SessionableBuilder.
func (d *DbSessionBuilder) FindSession(w http.ResponseWriter, r *http.Request) web.Sessionable {
	sessionID := r.URL.Query().Get("session_id")
	if len(sessionID) < 1 {
		cookie, err := r.Cookie(d.Key)
		if err == nil && len(cookie.Value) > 0 {
			sessionID = cookie.Value
		}
	}
	if len(sessionID) > 0 {
		//find session by cookie id
		session, err := d.FindSessionByKey(sessionID)
		if err == nil && len(session.SID) > 0 {
			//session found by cookie id
			if w != nil {
				d.writeCookie(session.SID, w, r)
			}
			return session
		}
		if err != redis.ErrNil {
			xlog.Errorf("DbSessionBuilder find session fail with %v", err)
		}
	}
	//not cookie and seesion found
	session := NewDbSession(d)
	session.SID = uuid.New()
	d.sessionLck.Lock()
	d.sessiones[session.SID] = session
	d.sessionLck.Unlock()
	if w != nil {
		d.writeCookie(session.SID, w, r)
	}
	return session
}

//SetEventHandler is impl func to web.SessionableBuilder.
func (d *DbSessionBuilder) SetEventHandler(h web.SessionEventHandler) {
}

//Find find by session id
func (d *DbSessionBuilder) Find(id string) web.Sessionable {
	sesion, _ := d.FindSessionByKey(id)
	if sesion == nil {
		return nil
	}
	return sesion
}

//FindSessionByKey will find the sesssion from memory,redist cache
func (d *DbSessionBuilder) FindSessionByKey(key string) (session *DbSession, err error) {
	d.sessionLck.RLock()
	session = d.sessiones[key]
	d.sessionLck.RUnlock()
	if session != nil {
		return
	}
	conn := d.Redis()
	defer conn.Close()
	var data []byte
	data, err = redis.Bytes(conn.Do("GET", d.Key+"_"+key))
	if err == nil && len(data) > 0 {
		session = NewDbSession(d)
		buf := bytes.NewBuffer(data)
		decoder := gob.NewDecoder(buf)
		err = decoder.Decode(session)
		if err == nil {
			d.sessionLck.Lock()
			d.sessiones[session.SID] = session
			d.sessionLck.Unlock()
		}
	}
	return
}

//new http.Cookie with Key/Path/Domain/MaxAge.
func (d *DbSessionBuilder) newCookie() *http.Cookie {
	c := &http.Cookie{}
	c.Name = d.Key
	c.Path = d.Path
	c.Domain = d.Domain
	c.MaxAge = d.MaxAge
	// d.log("new cookie by name(%v),path(%v),domain(%v),maxage(%v)", c.Name, c.Path, c.Domain, c.MaxAge)
	return c
}

//DbSession is the http session by db
type DbSession struct {
	xmap.M  `json:"values,omitempty"` //the session values.
	Builder *DbSessionBuilder         `json:"-"`
	SID     string                    `json:"id,omitempty"`   //the session id
	Last    int64                     `json:"last,omitempty"` //last update time
	Time    int64                     `json:"time,omitempty"` //create time
}

//NewDbSession will return new session.
func NewDbSession(b *DbSessionBuilder) *DbSession {
	now := xsql.TimeNow().Timestamp()
	return &DbSession{
		Builder: b,
		M:       xmap.M{},
		Last:    now,
		Time:    now,
	}

}

func (s *DbSession) ID() string {
	return s.SID
}

//Flush all updated session value
func (s *DbSession) Flush() (err error) {
	buf := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buf)
	err = encoder.Encode(s)
	if err != nil {
		xlog.Errorf("DbSession flush session to redist fail with %v", err)
		return
	}
	conn := s.Builder.Redis()
	defer conn.Close()
	_, err = conn.Do("MSET",
		s.Builder.Key+"_"+s.SID, buf.Bytes(),
		fmt.Sprintf(s.Builder.Key+"_%v_"+s.SID, s.Value("uid")), xsql.TimeNow().Timestamp(),
	)
	if err != nil {
		xlog.Errorf("DbSession flush session to redist fail with %v", err)
	} else {
		s.Builder.log("DbSession flush session to redist success with key:%v", s.Builder.Key+"_"+s.SID)
	}
	return
}
