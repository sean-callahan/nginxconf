// Copyright 2016 Sean Callahan. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package nginxconf

import (
	"fmt"
	"testing"
)

var testDirectives = []*Directive{
	NewDirective("listen", "80"),
	NewDirective("server_name", "example.com", "www.example.com"),
	NewDirective("server", "127.0.0.1:3000", "weight=5"),
}

var testDirectivesText = []string{
	"listen 80;",
	"server_name example.com www.example.com;",
	"server 127.0.0.1:3000 weight=5;",
}

var testBlocks = []*Directive{
	NewDirective("server").AddChild(NewDirective("listen", "80")).AddChild(NewDirective("root", "/var/www/htdocs")),
	NewDirective("server").AddChild(NewDirective("listen", "80")).AddChild(NewDirective("location", "/").AddChild(NewDirective("proxy_pass", "http://127.0.0.1:3000"))),
}

var testBlocksText = []string{
	`server {
    listen 80;
    root /var/www/htdocs;
}`,
	`server {
    listen 80;
    location / {
        proxy_pass http://127.0.0.1:3000;
    }
}`,
}

func TestNewDirective(t *testing.T) {
	for i, d := range testDirectives {
		if d.String() != testDirectivesText[i] {
			t.Fail()
		}
	}

	for i, b := range testBlocks {
		if b.String() != testBlocksText[i] {
			t.Fail()
		}
	}
}

func ExampleDirective() {
	errorLog := NewDirective("error_log", "logs/error.log")
	fmt.Println(errorLog)
	// Output: error_log logs/error.log;
}

func ExampleDirective_block() {
	server := NewDirective("server")
	listen := NewDirective("listen", "443", "ssl")
	server.AddChild(listen)
	fmt.Println(server)
	// Output: server {
	//     listen 443 ssl;
	// }
}
