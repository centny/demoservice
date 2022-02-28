package define

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := NewError(NotAccess, "not access", nil)
	fmt.Printf("err->%v\n", err)
	fmt.Printf("err->%v\n", err.String())
	err.Inner = fmt.Errorf("inner")
	fmt.Printf("err->%v\n", err)
	fmt.Printf("code->%v\n", err.Code())
	if !IsCodeError(err, NotAccess) {
		t.Error(err)
		return
	}
}
