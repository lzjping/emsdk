package emsdk_test

import "testing"

func TestFriend(t *testing.T) {
	client.CreateUser("test_217", "test")
	client.CreateUser("test_45", "test")
	result, err := client.AddFriend("test_217", "test_45")
	if err != nil {
		t.Fatal(err)
	}

	if result {
		_, err := client.DeleteFriend("test_217", "test_45")
		if err != nil {
			t.Fatal(err)
		}
	}

	client.DeleteUser("test_45")
	client.DeleteUser("test_217")
}

func TestBlacklist(t *testing.T) {
	client.CreateUser("test_217", "test")
	client.CreateUser("test_45", "test")
	client.CreateUser("test_2", "test")

	result, err := client.AddBlacklist("test_217", []string{"test_45", "test_2"})
	if err != nil {
		t.Fatal(err)
	}

	if result {
		a, _ := client.DeleteBlacklist("test_217", "test_45")
		b, _ := client.DeleteBlacklist("test_217", "test_2")

		if a && b {
			client.DeleteUser("test_45")
			client.DeleteUser("test_217")
			client.DeleteUser("test_2")
		}
	}
}
