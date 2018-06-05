# Discard
因為 microsoft的原因 此項目 已經被轉移到 [gitlab.com/king011/revel-i18n](https://gitlab.com/king011/revel-i18n)

github上的 版本將不在進行 任何維護

# revel-i18n
revel i18n tools

revel-i18n 是孤為 go web 框架 [revel](http://revel.github.io/) 開發的一個 i18n 工具 revel-i18n 可以自動 從 views 中使用正則將 所有 待翻譯的 條目 整理出來 並和 已經存在的翻譯檔案 進行合併

# Rules
因為是 是使用 正則匹配 且 為了實現上的簡單 revel-i18n 對 views 中的 模板 和 翻譯進行了一些 限制 項目必需符合這些規則 revel-i18n 才能正常 工作
1. 傳入 msg 的 待翻譯條目 key 不可以含有 一些特殊符號 比如 \[ " = ... 但 允許 使用 . - _ ... (具體請查看 cmd/cmdnew/Context.go 中的 MatchKey) 你應該儘量使用 英文字符作為key 而使用 . 或 - 來區分作用域 (比如 /app/about 頁面的 title 你可以將其key定義為 App.About.title 或 App-About-title)
2. 如果 你的 go template 使用 {{}} 則 key 也不能包含 { } 同理 如果你自定義了 模板 使用 '' '' 則 ' 也不能作為 key
3. 翻譯的 key value 開始和結尾的空格 會被忽略
4. 對於 翻譯 檔案 messages 檔案 全部使用 locale.xx locale-XX.xx 來命名 比如 zh 的翻譯檔案 為 locale.zh zh-TW 的翻譯檔案 為 locale-TW.zh

# Install
1. go get -u -d github.com/zuiwuchang/revel-i18n
2. cd $GOPATH/src/github.com/zuiwuchang/revel-i18n
3. ./build-go.sh
4. go install

# How To Use
1. cd ${yourProject}
2. revel-i18n new -l zh-TW --no-line
3. revel-i18n new -l zh --no-line

```bash
$ revel-i18n new -h
new message file
	revel-i18n new -v app/views -m messages -l zh-TW

Usage:
  revel-i18n new [flags]

Flags:
      --delimiters string   go template delimiters (default "{{ }}")
  -h, --help                help for new
  -l, --locale string       locale zh zh-TW zh-HK de ... (default "zh")
  -m, --messages string     revel messages directory (default "messages")
      --no-line             if true not write key file:line
  -t, --touch               true (Coverage file) false (Merge file)
  -v, --views string        revel views directory (default "app/views")
```
