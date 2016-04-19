// Copyright 2016 Sean Callahan. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package nginxconf

// Directive represents a basic Nginx configuration file directive.
// The directive is considered a block if it contains children and con
// be checked by calling the Block() function.
type Directive struct {
	Name     string
	Args     []string
	Parent   *Directive
	Children []*Directive
}

// NewDirective creates a new directive with a specified name and optional
// arguments.
func NewDirective(name string, args ...string) *Directive {
	return &Directive{
		Name: name,
		Args: args,
	}
}

// Block checkes if the Directive is considered a block (it has children).
func (d Directive) Block() bool {
	return len(d.Children) > 0
}

// AddChild adds a child to the directive, effectively making a block.
// The parent directive is returned, so future calls can be chained.
func (d *Directive) AddChild(c *Directive) *Directive {
	d.Children = append(d.Children, c)
	c.Parent = d
	return d
}

// Level returns the number of level down the block tree the directive
// is located. If the directive has not parents, this value is 0.
func (d *Directive) Level() int {
	cur := d
	level := 0
	for cur.Parent != nil {
		level++
		cur = cur.Parent
	}
	return level
}

func (d Directive) String() string {
	if d.Block() {
		return d.blockString()
	}
	out := d.Name
	for _, arg := range d.Args {
		out += " " + arg
	}
	return out + ";"
}

func (d Directive) blockString() string {
	out := d.Name
	for _, arg := range d.Args {
		out += " " + arg
	}
	out += " {\n"
	for _, child := range d.Children {
		for i := 0; i < (d.Level() + 1); i++ {
			out += "    "
		}
		out += child.String() + "\n"
	}
	for i := 0; i < d.Level(); i++ {
		out += "    "
	}
	return out + "}"
}
