package nginxconf

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	f, err := OpenFile("example.conf")
	if err != nil {
		panic(err)
	}

	listen := NewDirective("listen", "80")
	err = f.WriteDirective(listen)
	if err != nil {
		panic(err)
	}

	sname := NewDirective("server_name", "example.com", "www.example.com")
	err = f.WriteDirective(sname)
	if err != nil {
		panic(err)
	}

	f.Close()

	rf, err := os.OpenFile("example.conf", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(rf)
	if err != nil {
		panic(err)
	}

	expected := fmt.Sprintf("%s\n%s\n", listen, sname)

	if string(b) != expected {
		t.Fail()
	}

	err = os.Remove("example.conf")
	if err != nil {
		panic(err)
	}
}
