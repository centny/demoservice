package demoapi

import (
	"fmt"
	"testing"

	"github.com/codingeasygo/util/converter"
	"sxbastudio.com/base/go/define"
)

func TestLogin(t *testing.T) {
	login, err := ts.GetMap("/pub/login?username=%v&password=%v", "admin", "123")
	if err != nil || login.Int64("code") != define.Success {
		t.Errorf("err:%v,code:%v", err, login)
		return
	}
	if login.Int64("user/tid") == 0 {
		t.Error("last_login_time err")
		return
	}
	fmt.Printf("login--->%v\n", converter.JSON(login))
	//
	loginErr, err := ts.GetMap("/pub/login?username=%v&password=%v", "abc0", "1x23")
	if err != nil || loginErr.Int64("code") != define.NotFound {
		t.Errorf("err:%v,loginErr:%v", err, loginErr)
		return
	}
	//
	argErr, err := ts.GetMap("/pub/login?username=%v&password=%v", "abcxx", "")
	if err != nil || argErr.Int64("code") != define.ArgsInvalid {
		t.Errorf("err:%v,loginErr:%v", err, loginErr)
		return
	}
}
