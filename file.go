package nginxconf

import "os"

// File represents a nginx config file.
// It wraps os.File
type File struct {
	*os.File
}

// OpenFile opens a nginx config file for writing only.
// If the file does not exist at that path, it is created
// with the permissions 0666.
func OpenFile(path string) (*File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &File{f}, nil
}

// WriteDirective writes one directive to the file.
// Call this function for each directive to write.
func (f *File) WriteDirective(d *Directive) error {
	b := []byte(d.String() + "\n")
	_, err := f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
