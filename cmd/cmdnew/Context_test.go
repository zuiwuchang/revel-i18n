package cmdnew

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	r, e := regexp.Compile(`\{\{` + MatchString + `\}\}`)
	if e != nil {
		t.Fatal(e)
	}

	e = _tTestMatch(t, r,
		`"first page":{{msg . "first page"}},`,
		"first page",
	)
	if e != nil {
		t.Fatal(e)
	}

	e = _tTestMatch(t, r,
		`{{set . "title" (msg . "123 About" )}}{{msg . "ok"}} {{yes}}`,
		"123 About",
		"ok",
	)
	if e != nil {
		t.Fatal(e)
	}

	e = _tTestMatch(t, r,
		`{{set . "title" (msg $ctx "User.ViewMenu" )}}{{msg . "ok"}} {{yes}}`,
		"User.ViewMenu",
		"ok",
	)
	if e != nil {
		t.Fatal(e)
	}

}
func _tTestMatch(t *testing.T, r *regexp.Regexp, str string, val ...string) (e error) {
	strs := r.FindAllString(str, -1)
	if len(strs) != len(val) {
		e = fmt.Errorf("len(strs)-%v != len(val)-%v", len(strs), len(val))
		return
	}
	for i, str := range strs {
		str = r.ReplaceAllString(str, "$1")
		if str != val[i] {
			e = fmt.Errorf("val[%v] not match %s != %s", i, str, val)
			return
		}
	}
	return
}
