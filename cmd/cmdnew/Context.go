package cmdnew

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

const (
	// MatchKey .
	MatchKey = `[+-\?\*\/\w\s\.!@#\$%\^&\(\);:',]+`
	// MatchString .
	MatchString = `.*?[\s\(]?msg\s+?[\.\$\w]+?\s+?"(` + MatchKey + `)".*?`
)

// Context .
type Context struct {
	// revel views directory
	Src string
	// revel messages directory
	Dist string

	// zh zh-TW zh-HK ...
	Locale string

	// 如果存在 翻頁檔案 覆蓋
	Touch bool

	// 不需要 寫入 翻頁來源 檔案
	NotLine bool

	// go template delimiters
	Delimiters string
}

var matchLocale = regexp.MustCompile(`^[a-zA-Z]+(-[a-zA-Z]+)?$`)

func (c *Context) String() string {
	b, _ := json.MarshalIndent(c, "", "\t")
	return string(b)
}

// GetLocaleRegion 返回 語言-區域
func (c *Context) GetLocaleRegion() (l string, r string, e error) {
	str := strings.TrimSpace(c.Locale)
	if !matchLocale.MatchString(str) {
		e = fmt.Errorf("%s is not the locale value", str)
		return
	}

	strs := strings.Split(str, "-")
	switch len(strs) {
	case 1:
		l = strs[0]
	case 2:
		l = strs[0]
		r = strs[1]
	default:
		e = fmt.Errorf("%s is not the locale value", str)
	}
	return
}

// GetDist .
func (c *Context) GetDist() string {
	str := strings.TrimSpace(c.Dist)
	if !strings.HasSuffix(str, `/`) {
		if runtime.GOOS == "windows" && !strings.HasSuffix(str, `\`) {
			str += "/"
		}
	}
	return str
}

// GetSrc .
func (c *Context) GetSrc() string {
	str := strings.TrimSpace(c.Src)
	if !strings.HasSuffix(str, `/`) {
		if runtime.GOOS == "windows" && !strings.HasSuffix(str, `\`) {
			str += "/"
		}
	}
	return str
}

// NewMatch .
func (c *Context) NewMatch() (r *regexp.Regexp, e error) {
	str := strings.TrimSpace(c.Delimiters)
	strs := strings.Split(str, " ")
	if len(strs) != 2 {
		e = fmt.Errorf("%s is not the delimiters", str)
		return
	}
	strs[0] = strings.TrimSpace(strs[0])
	if strs[0] == "" {
		e = fmt.Errorf("%s is not the delimiters", str)
		return
	}
	strs[1] = strings.TrimSpace(strs[1])
	if strs[1] == "" {
		e = fmt.Errorf("%s is not the delimiters", str)
		return
	}

	s0 := []byte(strs[0])
	s1 := []byte(strs[1])

	b0 := make([]byte, 0, len(s0))
	b1 := make([]byte, 0, len(s1))

	for _, b := range s0 {
		if c.match(b) {
			b0 = append(b0, '\\')
		}
		b0 = append(b0, b)
	}
	for _, b := range s1 {
		if c.match(b) {
			b1 = append(b1, '\\')
		}
		b1 = append(b1, b)
	}
	str = string(b0) + MatchString + string(b1)

	r, e = regexp.Compile(str)
	return
}
func (c *Context) match(b byte) bool {
	switch b {
	case '\\':
		return true
	case '\'':
		return true
	case '"':
		return true
	case '[':
		return true
	case ']':
		return true
	case '(':
		return true
	case ')':
		return true
	case '{':
		return true
	case '}':
		return true
	case '?':
		return true
	}
	return false
}
