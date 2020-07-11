package meta_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/modules/meta"
)

const incorrectValueErrTemplate = "set value is not equal for UniqueIdentifier: gave `%s` but expected `%s`"

func TestNew(t *testing.T) {
	config := &meta.Config{
		UniqueIdentifier:             "testA",
		Version:                      "testB",
		MinimumCompatibleCoreVersion: "testC",
		MaximumCompatibleCoreVersion: "testD",
		Name:                         "testE",
		Author:                       "testF",
		Contact:                      "testG",
	}
	meta, err := meta.New(config)
	if err != nil {
		t.Fatal(err)
	}
	if meta == nil {
		t.Fatal("expected meta to not be nil")
	}
	gave := meta.UniqueIdentifier()
	if gave != config.UniqueIdentifier {
		t.Fatalf(incorrectValueErrTemplate, gave, config.UniqueIdentifier)
	}
	gave = meta.Version()
	if gave != config.Version {
		t.Fatalf(incorrectValueErrTemplate, gave, config.Version)
	}
	gave = meta.MinimumCompatibleCoreVersion()
	if gave != config.MinimumCompatibleCoreVersion {
		t.Fatalf(incorrectValueErrTemplate, gave, config.MinimumCompatibleCoreVersion)
	}
	gave = meta.MaximumCompatibleCoreVersion()
	if gave != config.MaximumCompatibleCoreVersion {
		t.Fatalf(incorrectValueErrTemplate, gave, config.MaximumCompatibleCoreVersion)
	}
	gave = meta.Name()
	if gave != config.Name {
		t.Fatalf(incorrectValueErrTemplate, gave, config.Name)
	}
	gave = meta.Author()
	if gave != config.Author {
		t.Fatalf(incorrectValueErrTemplate, gave, config.Author)
	}
	gave = meta.Contact()
	if gave != config.Contact {
		t.Fatalf(incorrectValueErrTemplate, gave, config.Contact)
	}
}
