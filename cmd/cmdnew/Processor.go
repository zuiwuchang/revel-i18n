package cmdnew

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var errShortWrite = errors.New("short write")
var matchNotKey = regexp.MustCompile(`(^\s*#)|(^\s*$)|(^\s*\[.*\]\s*$)`)

// Processor .
type Processor struct {
	locale      string
	region      string
	file        string
	root        string
	merge       map[string]string
	regexp      *regexp.Regexp
	description bool
}

// NewProcessor .
func NewProcessor(context *Context) (processor *Processor, e error) {
	p := &Processor{}
	p.locale, p.region, e = context.GetLocaleRegion()
	if e != nil {
		return
	}

	//創建 正則匹配
	p.regexp, e = context.NewMatch()
	if e != nil {
		return
	}

	// 合併 檔案
	if p.region == "" {
		p.file = context.GetDist() + "locale." + p.locale
	} else {
		p.file = context.GetDist() + "locale_" + p.region + "." + p.locale
	}
	if !context.Touch {
		if e = p.initMerge(); e != nil {
			return
		}
	}
	p.root = context.GetSrc()
	p.description = !context.NotLine

	processor = p
	return
}

// Run .
func (p *Processor) Run() (e error) {
	keys := make(map[string]*Item)
	e = filepath.Walk(p.root, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if info.IsDir() {
			return nil
		}
		name := path
		if strings.HasPrefix(name, p.root) {
			name = name[len(p.root):]
		}
		return p.readFile(keys, path, name)
	})
	if e != nil {
		return
	}
	if len(keys) == 0 {
		fmt.Println("not found any translation entries")
		return
	}
	p.writeFile(keys)
	return
}
func (p *Processor) writeFile(keys map[string]*Item) (e error) {
	var f *os.File
	f, e = os.Create(p.file)
	if e != nil {
		return
	}

	// session
	var n int
	var str string
	if p.region != "" {
		str = "[" + p.region + "]\n"
		n, e = f.WriteString(str)
		if e != nil {
			return
		} else if n != len(str) {
			e = errShortWrite
			return
		}
	}
	// key
	merge := 0

	for _, item := range keys {
		if item.Value == "" { // 先寫入 已經 完成 的翻譯
			continue
		}
		merge++
		// description
		if p.description && len(item.Description) != 0 {
			for description := range item.Description {
				str = "#" + description + "\n"
				n, e = f.WriteString(str)
				if e != nil {
					return
				} else if n != len(str) {
					e = errShortWrite
					return
				}
			}
		}
		// key
		str = item.Key + "=" + item.Value + "\n"

		n, e = f.WriteString(str)
		if e != nil {
			return
		} else if n != len(str) {
			e = errShortWrite
			return
		}
	}

	for _, item := range keys {
		if item.Value != "" { //寫入 未翻譯 項
			continue
		}
		// description
		if p.description && len(item.Description) != 0 {
			for description := range item.Description {
				str = "#" + description + "\n"
				n, e = f.WriteString(str)
				if e != nil {
					return
				} else if n != len(str) {
					e = errShortWrite
					return
				}
			}
		}
		// key
		str = item.Key + "=\n"

		n, e = f.WriteString(str)
		if e != nil {
			return
		} else if n != len(str) {
			e = errShortWrite
			return
		}
	}

	f.Close()

	fmt.Printf("%v items %v merge %v need translation\n", len(keys), merge, len(keys)-merge)
	return
}
func (p *Processor) readFile(keys map[string]*Item, path, name string) (e error) {
	var f *os.File
	f, e = os.Open(path)
	if e != nil {
		return
	}
	r := bufio.NewReader(f)
	var b []byte
	line := 1
	var description string
	for {
		b, _, e = r.ReadLine()
		if e != nil {
			if e == io.EOF {
				e = nil
			}
			break
		}

		bb := p.regexp.FindAll(b, -1)
		for _, b := range bb {
			key := string(p.regexp.ReplaceAll(b, []byte("$1")))
			if p.description {
				description = fmt.Sprintf("%s : %v", name, line)
			}
			p.push(keys, key, description)
		}
		line++
	}
	f.Close()
	return
}
func (p *Processor) push(keys map[string]*Item, key, description string) {
	item, _ := keys[key]
	if item == nil {
		item = &Item{Key: key}
		if p.description {
			item.Description = make(map[string]bool)
			item.Description[description] = true
		}
		if val, ok := p.merge[key]; ok {
			item.Value = val
		}
		keys[key] = item
	} else {
		if p.description {
			item.Description[description] = true
		}
	}
}
func (p *Processor) initMerge() (e error) {
	_, e = os.Stat(p.file)
	if e != nil {
		if os.IsNotExist(e) {
			e = nil
			return
		}
		return
	}

	var f *os.File
	f, e = os.Open(p.file)
	if e != nil {
		return
	}
	r := bufio.NewReader(f)
	var b []byte
	merge := make(map[string]string)
	for {
		b, _, e = r.ReadLine()
		if e != nil {
			if e == io.EOF {
				e = nil
			}
			break
		}
		if matchNotKey.Match(b) {
			continue
		}

		str := strings.TrimSpace(string(b))
		strs := strings.SplitN(str, "=", 2)
		strs[0] = strings.TrimSpace(strs[0])
		strs[1] = strings.TrimSpace(strs[1])
		if len(strs) != 2 || strs[0] == "" || strs[1] == "" {
			continue
		}

		merge[strs[0]] = strs[1]
	}
	f.Close()
	if e == nil && len(merge) != 0 {
		p.merge = merge
	}
	return
}
