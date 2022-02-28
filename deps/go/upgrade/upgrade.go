package upgrade

import (
	"fmt"
	"sort"
	"strings"

	"github.com/codingeasygo/util/converter"
	"sxbastudio.com/base/go/xlog"
)

// VersionSorter -
type VersionSorter []string

// GetIdx -
func (s VersionSorter) GetIdx(version string) int {
	for i, v := range s {
		if v == version {
			return i
		}
	}
	return -1
}

// Len -
func (s VersionSorter) Len() int {
	return len(s)
}

// Less -
func (s VersionSorter) Less(x, y int) bool {
	vx := strings.Split(strings.TrimPrefix(s[x], "v"), ".")
	vy := strings.Split(strings.TrimPrefix(s[y], "v"), ".")
	length := len(vx)
	if len(vy) < len(vx) {
		length = len(vy)
	}

	for i := 0; i < length; i++ {
		vxi, vyi := converter.Int(vx[i]), converter.Int(vy[i])
		if vxi == vyi {
			continue
		}
		return vxi < vyi
	}

	return len(vx) < len(vy)
}

// Swap -
func (s VersionSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ProcUpgrade -
func ProcUpgrade(versions map[string]string, startVersion string, upgrade func(script string) error) error {
	xlog.Infof("Upgrade is starting base %v", startVersion)
	var list VersionSorter
	for k := range versions {
		list = append(list, k)
	}
	if len(list) == 0 {
		return fmt.Errorf("version list is null")
	}
	sort.Sort(list)
	start := list.GetIdx(startVersion)
	if start == -1 {
		return fmt.Errorf("invalid startVersion")
	}
	for i := start; i < list.Len(); i++ {
		xlog.Infof("start upgrade to %v", list[i])
		err := upgrade(versions[list[i]])
		if err != nil {
			xlog.Errorf("upgrade to %v is fail with %v", list[i], err)
			return err
		}
		xlog.Infof("upgrade to %v is success", list[i])
	}
	return nil
}
