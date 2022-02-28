package session

import (
	"testing"

	"github.com/codingeasygo/util/xhttp"
	"github.com/codingeasygo/web"
	"github.com/codingeasygo/web/httptest"

	"github.com/Centny/rediscache"
)

func init() {
	redisURI := "redis.loc:6379?db=1"
	rediscache.InitRedisPool(redisURI)
	xhttp.EnableCookie()
}

func TestSessionBuilder(t *testing.T) {
	ts := httptest.NewMuxServer()
	sb := NewDbSessionBuilder()
	sb.Redis = rediscache.C
	sb.ShowLog = true
	ts.Mux.Builder = sb
	xhttp.ClearCookie()
	//
	ts.Mux.HandleFunc("^/set.*$", func(hs *web.Session) web.Result {
		hs.Clear()
		hs.SetValue("abc", 1)
		err := hs.Flush()
		return hs.Printf("%v", err)
	})
	ts.Mux.HandleFunc("^/err.*$", func(hs *web.Session) web.Result {
		hs.SetValue("abc", TestSessionBuilder)
		err := hs.Flush()
		if err == nil {
			panic(err)
		}
		return hs.Printf("%v", "OK")
	})
	ts.Mux.HandleFunc("^/get.*$", func(hs *web.Session) web.Result {
		return hs.Printf("%v", hs.Value("abc"))
	})

	//
	res, err := ts.GetText("/set")
	if err != nil || res != "<nil>" {
		t.Errorf("err:%v,res:%v", err, res)
		return
	}
	res, err = ts.GetText("/get")
	if err != nil || res != "1" {
		t.Errorf("err:%v,res:%v", err, res)
		return
	}
	//
	for i := 0; i < 1000; i++ {
		sb.sessiones = map[string]*DbSession{}
		res, err = ts.GetText("/get")
		if err != nil || res != "1" {
			t.Errorf("err:%v,res:%v", err, res)
			return
		}
	}
	//
	sb.sessiones = map[string]*DbSession{}
	conn := rediscache.C()
	conn.Do("FLUSHDb")
	conn.Close()
	res, err = ts.GetText("/get")
	if err != nil || res != "<nil>" {
		t.Errorf("err:%v,res:%v", err, res)
		return
	}
	//
	res, err = ts.GetText("/err")
	if err != nil || res != "OK" {
		t.Errorf("err:%v,res:%v", err, res)
		return
	}
	//
	//error
	rediscache.MockerStart()
	defer rediscache.MockerStop()
	//
	sb.Redis = rediscache.C
	sb.sessiones = map[string]*DbSession{}
	//
	rediscache.MockerSet("Conn.Do", 1)
	res, err = ts.GetText("/get")
	if err != nil || res != "<nil>" {
		t.Errorf("err:%v,res:%v", err, res)
		return
	}
	rediscache.MockerClear()
	//
	rediscache.MockerSet("Conn.Do", 1)
	xhttp.ClearCookie()
	res, err = ts.GetText("/set")
	if err != nil || res == "<nil>" {
		t.Errorf("err:%v,res:%v", err, res)
		return
	}
	rediscache.MockerClear()
	//
	sb.SetEventHandler(nil)
}
