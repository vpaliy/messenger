package media

import (
	"io"
)

type Filer interface {
	io.Reader
	io.Seeker
	io.Closer
}

type MediaHandler interface {
	Init() error

	Upload(filer *Filer) (string, error)

	Download(url string) (*Filer, error)

	Delete(locations []string) error
}
