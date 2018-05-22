package animals

import (
	"testing"
)

func TestElefantFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
