package emsdk_test

import (
	"os"
	"testing"

	"github.com/links123com/emsdk"
)

var (
	client *emsdk.Client
	err    error
)

func TestMain(m *testing.M) {
	println("setup")
	client, err = emsdk.New("links123", "linker", "YXA617WqYLg2EeWieSNUhFCrvg", "YXA6WCpzHwPQ8S2oMhozRy4PXFu2YGo")
	code := m.Run()
	println("teardown")
	os.Exit(code)
}

func TestNew(t *testing.T) {
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}

	if client == nil {
		t.Fatal("instance emsdk fail.")
	}
}

func TestCreateUser(t *testing.T) {
	err := client.CreateAccount("217", "123456")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}
}

func TestResetPassword(t *testing.T) {
	err := client.ChangePassword("217", "456789")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}
}

func TestResetNickname(t *testing.T) {
	err := client.ChangeNickname("217", "linker_景")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := client.DeleteAccount("217")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}
}

func TestIsOnline(t *testing.T) {
	ok := client.IsOnline("217")
	t.Log(ok)
}
