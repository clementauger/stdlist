package stdlist

import "strings"

var data = map[string][]string{}

//Has returns if given import path p exists for given go version v.
func Has(v, p string) bool {
	v = strings.TrimPrefix(v, "go")
	if x, ok := data[p]; ok {
		for _, y := range x {
			if y == v {
				return true
			}
		}
	}
	return false
}
