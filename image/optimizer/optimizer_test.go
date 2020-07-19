package optimizer_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/config/configmock"
	"github.com/espal-digital-development/espal-core/image/optimizer"
)

var (
	configService *configmock.ConfigMock
)

func initMocks() {
	configService = &configmock.ConfigMock{}
}

func TestNew(t *testing.T) {
	initMocks()
	optimizer, err := optimizer.New(configService)
	if err != nil {
		t.Fatal(err)
	}
	if optimizer == nil {
		t.Fatal("expected optimizer to not be nil")
	}
}
