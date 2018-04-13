package WordsFilter

import (
	"testing"
)

func TestWordsFilter(t *testing.T) {
	texts := []string{
		"Miyamoto Musashi",
		"妲己",
		"アンジェラ",
		"ความรุ่งโรจน์",
	}
	wf := NewWordsFilter("*", true)
	root := wf.Generate(texts)
	wf.Remove("shif", root)
	c1 := wf.Contains("アン", root)
	if c1 != false {
		t.Errorf("TestContains expect false,get %T,%v", c1, c1)
	}
	c2 := wf.Contains("アンジェラ", root)
	if c2 != true {
		t.Errorf("TestContains expect true,get %T,%v", c2, c2)
	}
	r1 := wf.Replace("Game ความรุ่งโรจน์ i like 妲己 heroMiyamotoMusashi", root)
	if r1 != "Game*************ilike**hero***************" {
		t.Errorf("TestContains expect Game*************ilike**hero***************,get %T,%v", r1, r1)
	}
}
