package youdaofanyi

import "testing"

func Test_fanyi_1(t *testing.T) {
	ex1 := "I am the god of war"
	result, err := Fanyi(ex1, Plain)

	if err != nil {
		t.Fatal(err)
	}

	if result != "我是战争之神" {
		t.Log(result)
	}

}

func Test_fanyi_2(t *testing.T) {
	ex1 := "I am the god of war"
	result, err := Fanyi(ex1, HTML)

	t.Log(result)
	if err != nil {
		t.Fatal(err)
	}

}
