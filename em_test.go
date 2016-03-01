package emsdk_test

import (
	"testing"

	"github.com/wpajqz/emsdk"
)

var (
	client *emsdk.Client
)

func TestNew(t *testing.T) {
	var err error
	client, err = emsdk.New("links123", "linker", "YXA617WqYLg2EeWieSNUhFCrvg", "YXA6WCpzHwPQ8S2oMhozRy4PXFu2YGo")
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
		t.Fail()
	}
}

func TestCreateUser(t *testing.T) {
	result, err := client.CreateUser("217", "123456")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}

	t.Log("create user:", result)
}

func TestResetPassword(t *testing.T) {
	result, err := client.ResetPassword("217", "456789")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}

	t.Log("reset password:", result)
}

func TestDeleteUser(t *testing.T) {
	result, err := client.DeleteUser("217")
	if err != nil {
		if err == emsdk.ErrEM {
			info := err.(emsdk.EMError)
			t.Log(info.Code)
			t.Log(info.Description)
			t.FailNow()
		}
		t.Fatal(err)
	}

	t.Log("delete user:", result)
}
