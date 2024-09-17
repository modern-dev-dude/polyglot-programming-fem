package projector_test

import (
	"reflect"
	"testing"

	"github.com/modern-dev-dude/polyglot-programming/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	return &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
	opts := getOpts(args)
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("expected no error %v", err)
	}

	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args to be %v but got %v", args, expectedArgs)
	}

	if config.Operation != operation {
		t.Errorf("operation expected was %v but got %v", operation, config.Operation)

	}
}

func TestConfigPrint(t *testing.T) {
	args := []string{""}
	testConfig(t, args, args, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	args := []string{"foo"}
	testConfig(t, args, args, projector.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	args := []string{"add", "foo", "bar"}
	testConfig(t, args, args, projector.Add)
}

func TestConfigRmoveKey(t *testing.T) {
	args := []string{"rm", "foo"}
	testConfig(t, args, args, projector.Remove)
}
