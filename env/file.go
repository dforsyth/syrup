package syrup

import (
	"io"
	"os"
	"path"
	// "launchpad.net/goamz/s3"	
)

type Env interface {
	OpenFile(string) (File, error)
}

type File interface {
        io.Closer
        io.Reader
        io.ReaderAt
        io.Writer
        Stat() (*os.FileInfo, error)
}

type LocalEnv struct {
	basepath string
}

func NewLocalEnv(basepath string) *LocalEnv {
	return &LocalEnv{pathname: pathname}
}

func (e *LocalEnv) OpenFile(pathname string) (File, error) {
	return os.OpenFile(path.Join(e.basepath, pathname), os.O_CREATE | os.O_RDWR, 0666)
}

/*
// XXX goamz looks like it doesn't track tip, so it will need to be gofix'd before use
type S3Env struct {
	s3 *s3.S3
}

func NewS3Env(s3 *s3.S3) *S3Env {
	return &S3Env{
		s3: s3,
	}
}

type S3File struct {
	info os.FileInfo
}

func (e *S3Env) OpenFile(pathname string) (*S3File, error) {
	
}

func (f *S3File) Close() error {
	
}

func (f *S3File) Read(p []byte) (int, error) {
	return 0, nil
}

func (f *S3File) ReadAt(p []byte, off int64) (int, error) {
	return 0, nil
}

func (f *S3File) Write(p []byte) (int, error) {
	return 0, nil
}

func (f *S3File) Stat() (*os.FileInfo, error) {
	return &f.info, nil
}
*/