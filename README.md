# go-wordsfilter
A high performance text filter.

## Download & Install
```shell
go get github.com/syyongx/go-wordsfilter
```

## Quick Start
```go
import (
    "github.com/syyongx/go-wordsfilter"
)

func main() {
    texts := []string{
        "Miyamoto Musashi",
        "妲己",
        "アンジェラ",
        "ความรุ่งโรจน์",
    }
    wf := wordsfilter.New()

    // Generate
    root := wf.Generate(texts)
    // Generate with file
    // root := wf.GenerateWithFile(path)

    // Contains
    c1 := wf.Contains("アン", root)
    // c1: false
    c2 := wf.Contains("アンジェラ", root)
    // c2: true

    // Remove
    wf.Remove("アンジェラ", root)
    c3 := wf.Contains("アンジェラ", root)
    // c3: false

    // Replace
    r1 := wf.Replace("Game ความรุ่งโรจน์ i like 妲己 heroMiyamotoMusashi", root)
    // r1: Game*************ilike**hero***************
}
```

## Apis
```go
New() *WordsFilter
Generate(texts []string) map[string]*Node
GenerateWithFile(path string) (map[string]*Node, error)
Add(text string, root map[string]*Node)
Replace(text string, root map[string]*Node) string
Contains(text string, root map[string]*Node) bool
Remove(text string, root map[string]*Node)
```

## LICENSE
go-wordsfilter source code is licensed under the [MIT](https://github.com/syyongx/go-wordsfilter/blob/master/LICENSE) Licence.
