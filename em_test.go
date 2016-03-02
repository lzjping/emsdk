package emsdk_test

import (
	"os"
	"testing"

	"github.com/wpajqz/emsdk"
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
	err := client.CreateUser("217", "123456")
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
	err := client.ResetPassword("217", "456789")
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
	err := client.DeleteUser("217")
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
