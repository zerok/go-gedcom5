package gedcom5

import "strings"

type fieldTag struct {
	Tag       string
	OmitEmpty bool
}

func parseFieldTag(s string) fieldTag {
	elems := strings.Split(s, ",")
	res := fieldTag{}
	for idx, elem := range elems {
		if idx == 0 {
			res.Tag = elem
		} else {
			if elem == "omitempty" {
				res.OmitEmpty = true
			}
		}
	}
	return res
}
