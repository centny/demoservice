package upgrade

import (
	"sort"
	"testing"

	"github.com/codingeasygo/util/converter"
)

func TestVersionSorter(t *testing.T) {
	list := VersionSorter{
		"v1.1.0",
		"v1.0.1",
		"v1.0",
		"v1.1.10",
		"1.1.2",
	}
	sort.Sort(list)
	t.Log(converter.JSON(list))
}
