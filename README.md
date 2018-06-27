# go-wordsfilter
A high performance text filter.

## Download & Install
```shell
go get github.com/syyongx/go-wordsfilter
```

## Usage Instructions
```
import (
    "fmt"
    "github.com/syyongx/go-wordsfilter"
)

func main() {
    texts := []string{
        "Miyamoto Musashi",
        "妲己",
        "アンジェラ",
        "ความรุ่งโรจน์",
    }
    wf := NewWordsFilter("*", true)

    // Generate
    root := wf.Generate(texts)
    // Generate with file
    // root := wf.GenerateWithFile(path)

    // Contains
    c1 := wf.Contains("アン", root)
    fmt.Println(c1) // false
    c2 := wf.Contains("アンジェラ", root)
    fmt.Println(c2) // true

    // Remove
    wf.Remove("アンジェラ", root)
    c3 := wf.Contains("アンジェラ", root)
    fmt.Println(c3) // false

    // Replace
    r1 := wf.Replace("Game ความรุ่งโรจน์ i like 妲己 heroMiyamotoMusashi", root)
    fmt.Println(r1) // Game*************ilike**hero***************
}
```

## LICENSE
go-wordsfilter source code is licensed under the [MIT](https://github.com/syyongx/go-wordsfilter/blob/master/LICENSE) Licence.