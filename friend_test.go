package emsdk_test

import "testing"

func TestFriend(t *testing.T) {
	client.CreateUser("test_217", "test")
	client.CreateUser("test_45", "test")
	err := client.AddFriend("test_217", "test_45")
	if err != nil {
		t.Fatal(err)
	}

	err = client.DeleteFriend("test_217", "test_45")
	if err != nil {
		t.Fatal(err)
	}

	client.DeleteUser("test_45")
	client.DeleteUser("test_217")
}

func TestBlacklist(t *testing.T) {
	client.CreateUser("test_217", "test")
	client.CreateUser("test_45", "test")
	client.CreateUser("test_2", "test")

	err := client.AddBlacklist("test_217", []string{"test_45", "test_2"})
	if err != nil {
		t.Fatal(err)
	}

	a := client.DeleteBlacklist("test_217", "test_45")
	b := client.DeleteBlacklist("test_217", "test_2")

	if a == nil && b == nil {
		client.DeleteUser("test_45")
		client.DeleteUser("test_217")
		client.DeleteUser("test_2")
	}
}
