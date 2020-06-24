package userrights_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/repositories/userrights"
)

func TestGetName(t *testing.T) {
	userRights := userrights.New()
	name, err := userRights.GetName(30000)
	if err != nil {
		t.Fatal(name)
	}
	if name == "" {
		t.Error("name returned empty")
	}
	if name != "AccessAuth" {
		t.Error("name didn't return the expected `AccessAuth`")
	}
}

func TestGetNameUnexisting(t *testing.T) {
	userRights := userrights.New()
	name, err := userRights.GetName(0)
	if err == nil {
		t.Fatal("getting an unexisting name should return an error")
	}
	if name != "" {
		t.Fatal("getting an unexisting name shouldn't return a filled value")
	}
}

func TestGetCode(t *testing.T) {
	userRights := userrights.New()
	code, err := userRights.GetCode("AccessAuth")
	if err != nil {
		t.Fatal(err)
	}
	if code == 0 {
		t.Error("code shouldn't be 0")
	}
	if code != 30000 {
		t.Error("code didn't return the expected `30000`")
	}
}

func TestGetCodeUnexisting(t *testing.T) {
	userRights := userrights.New()
	code, err := userRights.GetCode("")
	if err == nil {
		t.Fatal("getting an unexisting name should return an error")
	}
	if code != 0 {
		t.Fatal("getting an unexisting code shouldn't return a filled value")
	}
}

func TestAllByCode(t *testing.T) {
	userRights := userrights.New()
	all := userRights.AllByCode()

	if len(all) == 0 {
		t.Fatal("AllByCode shouldn't be empty")
	}

	for code := range all {
		name, err := userRights.GetName(code)
		if err != nil {
			t.Fatal(err)
		}
		if name == "" {
			t.Errorf("expected a name for code `%d`, but got an empty string", code)
		}
	}
}

func TestAllByName(t *testing.T) {
	userRights := userrights.New()
	all := userRights.AllByName()

	if len(all) == 0 {
		t.Fatal("shouldn't be empty")
	}

	for name := range all {
		code, err := userRights.GetCode(name)
		if err != nil {
			t.Fatal(err)
		}
		if code == 0 {
			t.Errorf("expected a positive number for name `%s`, but got 0", name)
		}
	}
}

func TestUserRightCodes(t *testing.T) {
	userRights := userrights.New()
	all := userRights.UserRightCodes()

	if len(all) == 0 {
		t.Fatal("shouldn't be empty")
	}

	for k := range all {
		name, err := userRights.GetName(all[k])
		if err != nil {
			t.Fatal(err)
		}
		if name == "" {
			t.Errorf("expected a name for code `%d`, but got an empty string", all[k])
		}
	}
}
