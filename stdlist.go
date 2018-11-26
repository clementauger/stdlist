package stdlist

var data = map[string][]string{}

//Has returns if given import path p exists for given go version v.
func Has(v, p string) bool {
	if x, ok := data[v]; ok {
		for _, y := range x {
			if y == p {
				return true
			}
		}
	}
	return false
}
