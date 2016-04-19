// Copyright 2016 Sean Callahan. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

/*
Package nginxconf provides simple Nginx configuration generation.

Directives:

        index := nginxconf.NewDirective("index", "index.html", "index.htm", "index.php")

Blocks:

        server := nginxconf.NewDirective("server")
        ...
        server.AddChild(listen)

String representation:

        fmt.Println(server)
*/
package nginxconf
