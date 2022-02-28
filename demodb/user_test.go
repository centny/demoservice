package demodb

import "testing"

func TestFindUserByUsrPass(t *testing.T) {
	user, err := FindUserByUsrPass("admin", "123")
	if err != nil {
		t.Error(err)
		return
	}
	if user.Username != "admin" {
		t.Errorf("username is not correct by %v, expect %v", user.Username, "admin")
		return
	}
}
