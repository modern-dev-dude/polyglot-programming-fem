package projector_test

import (
	"testing"

	"github.com/modern-dev-dude/polyglot-programming/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
			"/foo/bar/baz": {
				"foo": "bar4",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Ags:       []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "Hello, mdd",
		},
		data,
	)
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}

	if v != value {
		t.Errorf("expected to find %v but recieved %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)

	test(t, projector, "foo", "bar3")
	test(t, projector, "fem", "is_great")
}

func TestSetValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)
	test(t, projector, "foo", "bar3")

	projector.SetValue("foo", "bar4")
	test(t, projector, "foo", "bar4")

	projector.SetValue("fem", "is_super_great")
	test(t, projector, "fem", "is_super_great")

	projector = getProjector("/", data)
	test(t, projector, "fem", "is_great")

}

func TestRemoveValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)
	test(t, projector, "foo", "bar3")

	projector.RemoveValue("foo")
	test(t, projector, "foo", "bar2")

	projector.RemoveValue("fem")
	test(t, projector, "fem", "is_great")

}
