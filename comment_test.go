package nginxconf

import "testing"

var testComments = []string{
	"Foo Bar",
	"### Bar #",
	"こんにちは",
	" #",
}

var expectedText = []string{
	"#Foo Bar",
	"#### Bar #",
	"#こんにちは",
	"# #",
}

func TestComment(t *testing.T) {
	for i, c := range testComments {
		cmt := Comment(c)
		if cmt.String() != expectedText[i] {
			t.Fail()
		}
	}
}
