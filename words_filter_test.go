package WordsFilter

import (
	"testing"
	"fmt"
)

func TestWordsFilter(t *testing.T) {
	//root := make(map[string]*Word)
	texts := []string{
		"Miyamoto Musashi",
		"妲己",
		"アンジェラ",
		"ความรุ่งโรจน์",
	}
	wf := NewWordsFilter("*", true)
	root := wf.Generate(texts)
	wf.Remove("shif", root)
	fmt.Println(wf.Contains("shift", root))
	fmt.Println(wf.Replace("Game ความรุ่งโรจน์ i like 妲己 heroMiyamotoMusashi", root))
}
