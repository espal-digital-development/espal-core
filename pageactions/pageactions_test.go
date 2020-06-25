package pageactions_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts/contextsmock"
)

var (
	context *contextsmock.ContextMock
)

func initMocks() {
	context = &contextsmock.ContextMock{
		HasUserRightFunc: func(code string) bool {
			return true
		},
		TranslateFunc: func(name string) string {
			return "stubTranslation"
		},
		AdminURLFunc: func() string {
			return "stubAdminURL"
		},
	}
}

func TestNew(t *testing.T) {
	initMocks()
	pageactions.New(context, "User", true)
}

func TestRenderOverviewActions(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	out := a.RenderOverviewActions()
	if out != "" {
		t.Fatal("Output should be empty when no actions are present")
	}
	a.AddDelete()
	out = a.RenderOverviewActions()
	if out == "" {
		t.Fatal("Output shouldn't be empty when actions are present")
	}
}

func TestIsFilled(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	if a.IsFilled() {
		t.Fatal("There shouldn't be any filling yet")
	}
	a.AddCreate()
	if !a.IsFilled() {
		t.Fatal("There should be filling by now")
	}
}

func TestAddCreate(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddCreate()
}

func TestAddCreateWithPath(t *testing.T) {
	initMocks()
	var seenTranslation bool
	context.TranslateFunc = func(name string) string {
		if name == "user" {
			seenTranslation = true
		}
		return ""
	}
	a := pageactions.New(context, "User", true)
	a.AddCreateWithPath("stubPath")
	if !seenTranslation {
		t.Fatal("Expected subject to be translated")
	}
}

func TestAddCreateWithPathNoRights(t *testing.T) {
	initMocks()
	context.HasUserRightFunc = func(code string) bool {
		return false
	}
	context.TranslateFunc = func(name string) string {
		t.Fatal("Shouldn't have being able to reach anywere further than having no rights for it")
		return ""
	}
	a := pageactions.New(context, "User", true)
	a.AddCreateWithPath("stubPath")
}

func TestAddCreateWithFieldAndPathEmptyFieldEmptyPath(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddCreateWithFieldAndPath("", "")
}

func TestAddToggle(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddToggle()
}

func TestAddToggleWithField(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddToggleWithField("stubField")
}

func TestAddToggleWithPathNoRights(t *testing.T) {
	initMocks()
	context.HasUserRightFunc = func(code string) bool {
		return false
	}
	context.TranslateFunc = func(name string) string {
		t.Fatal("Shouldn't have being able to reach anywere further than having no rights for it")
		return ""
	}
	a := pageactions.New(context, "User", true)
	a.AddToggleWithPath("stubPath")
}

func TestAddToggleWithFieldAndPathEmptyFieldEmptyPath(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddToggleWithFieldAndPath("", "")
}

func TestAddDelete(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddDelete()
}

func TestAddDeleteWithPath(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddDeleteWithPath("stubPath")
}

func TestAddDeleteWithPathNoRights(t *testing.T) {
	initMocks()
	context.HasUserRightFunc = func(code string) bool {
		return false
	}
	context.TranslateFunc = func(name string) string {
		t.Fatal("Shouldn't have being able to reach anywere further than having no rights for it")
		return ""
	}
	a := pageactions.New(context, "User", true)
	a.AddDeleteWithPath("stubPath")
}

func TestAddDeleteWithPathEmptyPath(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddDeleteWithPath("")
}

func TestAddUpdate(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddUpdate()
}

func TestAddUpdateWithField(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddUpdateWithField("stubField")
}

func TestAddUpdateWithPath(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddUpdateWithPath("stubPath")
}

func TestAddUpdateWithPathNoRights(t *testing.T) {
	initMocks()
	context.HasUserRightFunc = func(code string) bool {
		return false
	}
	context.TranslateFunc = func(name string) string {
		t.Fatal("Shouldn't have being able to reach anywere further than having no rights for it")
		return ""
	}
	a := pageactions.New(context, "User", true)
	a.AddUpdateWithPath("stubPath")
}

func TestAddUpdateWithFieldAndPathEmptyFieldEmptyPath(t *testing.T) {
	initMocks()
	a := pageactions.New(context, "User", true)
	a.AddUpdateWithFieldAndPath("", "")
}
