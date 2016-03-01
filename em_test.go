package emsdk_test

import (
	"testing"

	"github.com/wpajqz/emsdk"
)

func TestNew(t *testing.T) {
	client, err := emsdk.New("links123", "linker", "YXA617WqYLg2EeWieSNUhFCrvg", "YXA6WCpzHwPQ8S2oMhozRy4PXFu2YGo")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Message)
			t.Log(info.Description)
		}
		t.Fatal(err)
	}

	if client == nil {
		t.Fail()
	}
}
