# WordsFilter-go
A high performance text filter.

## Download & Install
```shell
go get github.com/syyongx/WordsFilter-go
```

## Instructions for use
```
import (
    "fmt"
    "github.com/syyongx/WordsFilter-go"
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

    // Remove
    wf.Remove("shif", root)

    // Contains
    c1 := wf.Contains("アン", root)
    fmt.Println(c1) // false
    c2 := wf.Contains("アンジェラ", root)
    fmt.Println(c1) // true

    // Replace
    r1 := wf.Replace("Game ความรุ่งโรจน์ i like 妲己 heroMiyamotoMusashi", root)
    fmt.Println(r1) // Game*************ilike**hero***************
}
```

## LICENSE
WordsFilter-go source code is licensed under the [MIT](https://github.com/syyongx/WordsFilter-go/blob/master/LICENSE) Licence.