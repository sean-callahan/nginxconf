package nginxconf

// Comment represents a nginx config file comment.
// This should not include a `#` prefix because it
// will be added when accessed via the String() method.
type Comment string

// String returns comment with a hash `#` prefix.
func (c Comment) String() string {
	return "#" + string(c)
}
