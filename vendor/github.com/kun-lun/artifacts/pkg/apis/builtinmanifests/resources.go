// Code generated by "esc -pkg builtinmanifests -prefix  -ignore  -include  -o resources.go manifests"; DO NOT EDIT.

package builtinmanifests

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/manifests/large_php.yml": {
		name:    "large_php.yml",
		local:   "manifests/large_php.yml",
		size:    3586,
		modtime: 1542074042,
		compressed: `
H4sIAAAAAAAC/+xWUYsbNxB+318x3L3ExXbORwhl39rS0kCgkFzoQymLVhp71dNKqma0PufXF2nXu+uz
nSalDy2Eg8OabzQz+mY+rVarVUGywVaU0N2tN0WBtqusaLGEx2hNtEVAcjFIrHbBRT9gL15cMC8WhXFS
sHY2eRx/LxZFoYWgEsTHGBBugRtNIIWFGgebC0CuRXDcYABvBG9daIuia/v4VBYAK+iT/xFbX7unAgCg
RRZl/gXQV8IH/8xHumi5hE1e0GMs4T0Lq0RQ1fcbytZ+U9c+K65rU2VdS7QEGUNAy+YAzppDgjQBRe9d
YFQ5iqPKB7fVBo8lCdVqO3I2rRaLwcFoG58q6exW72LouRsgAKJmWgD4WBstq0c80NyciDnGJmrWk9uQ
hdgFsRtr0u1sAeC2WwwlfKij5fgeQ4ehOMlITcJ/ENZZLYWZqktMbl6v716t3j68H80dBsoTYAQj8WB3
VClNj1PaVlixQ5WtQ8/Grrx9N4WTQjba7kp4h0L9GjTjBAUUjJXz/cT9FFz7Jp1tcFCCRQ4/Y2v12Xk/
mfks94+t58MMztFJf8RqV5ewubsbsdthjFPPO7Q9VTFqNRu9GrPlZR7oJehtXi7B2R6AvYtGpQHdWQyC
UYGgHDWnsch7F6Zjr4BibZGHOSSLvNqMBQ3eFaGMQfPhROdETWVpVzwbQe1LIBasZTHT5R7rFeUBoi/V
3Vf5fJXP/1M+xglV1cIIKzEMLqe2a761kI9oVSWUCkhUeefMEOAi9JmS3WPdi3AUbnAGqTwSpAkioQJ2
sMP+/MANgkJv3KFFy0AyaM80nn/+Hlj7xhdFZ5FPvsndxMqxaPJCYureOv+93Lzu5ZTJnLF7kdYg7G6+
+f7VCO0E414cRnBTFCfEntR1qRUnV1G2XOL7rMJPNKVBYbhJt1eNZ/saZt9D08UQHDvpTAk/M3u4hV4I
wkAnTEzN+u1B+mVG+//0+7TZBS7h20kTAf+MSFx5wU0JNy9v4DbbdECVpv+YLT9YkFPnc1oX+tBr+CU9
u/aacAmak5t1DMIYt0e1ft4oaq6e5kH6Z2Xe3/cjGM0VYhJyOdJ1Wj4ov4TvjJk42QZnOTXnjJxj164C
Xy6/ebersx4Xl9V5MpXnGj3bdJGx3JJVSjajTLu0oYT72TWpdEDZ361vbO2iVSMkpESiMvHn9n/Tw+GJ
n7irBkXefHMzZUFibfNX98RnxvIQYWQy4FY/XY9y2W8+ev+EMKLmAl+b/wxfg0j+Jb6KTgQt6oGNkYnj
A2f27kqrvwIAAP//W4XumgIOAAA=
`,
	},

	"/manifests/maximum_php.yml": {
		name:    "maximum_php.yml",
		local:   "manifests/maximum_php.yml",
		size:    3586,
		modtime: 1542074017,
		compressed: `
H4sIAAAAAAAC/+xWUYsbNxB+318x3L3ExXbORwhl39rS0kCgkFzoQymLVhp71dNKqma0PufXF2nXu+uz
nSalDy2Eg8OabzQz+mY+rVarVUGywVaU0N2tN0WBtqusaLGEx2hNtEVAcjFIrHbBRT9gL15cMC8WhXFS
sHY2eRx/LxZFoYWgEsTHGBBugRtNIIWFGgebC0CuRXDcYABvBG9daIuia/v4VBYAK+iT/xFbX7unAgCg
RRZl/gXQV8IH/8xHumi5hE1e0GMs4T0Lq0RQ1fcbytZ+U9c+K65rU2VdS7QEGUNAy+YAzppDgjQBRe9d
YFQ5iqPKB7fVBo8lCdVqO3I2rRaLwcFoG58q6exW72LouRsgAKJmWgD4WBstq0c80NyciDnGJmrWk9uQ
hdgFsRtr0u1sAeC2WwwlfKij5fgeQ4ehOMlITcJ/ENZZLYWZqktMbl6v716t3j68H80dBsoTYAQj8WB3
VClNj1PaVlixQ5WtQ8/Grrx9N4WTQjba7kp4h0L9GjTjBAUUjJXz/cT9FFz7Jp1tcFCCRQ4/Y2v12Xk/
mfks94+t58MMztFJf8RqV5ewubsbsdthjFPPO7Q9VTFqNRu9GrPlZR7oJehtXi7B2R6AvYtGpQHdWQyC
UYGgHDWnsch7F6Zjr4BibZGHOSSLvNqMBQ3eFaGMQfPhROdETWVpVzwbQe1LIBasZTHT5R7rFeUBoi/V
3Vf5fJXP/1M+xglV1cIIKzEMLqe2a761kI9oVSWUCkhUeefMEOAi9JmS3WPdi3AUbnAGqTwSpAkioQJ2
sMP+/MANgkJv3KFFy0AyaM80nn/+Hlj7xhdFZ5FPvsndxMqxaPJCYureOv+93Lzu5ZTJnLF7kdYg7G6+
+f7VCO0E414cRnBTFCfEntR1qRUnV1G2XOL7rMJPNKVBYbhJt1eNZ/saZt9D08UQHDvpTAk/M3u4hV4I
wkAnTEzN+u1B+mVG+//0+7TZBS7h20kTAf+MSFx5wU0JNy9v4DbbdECVpv+YLT9YkFPnc1oX+tBr+CU9
u/aacAmak5t1DMIYt0e1ft4oaq6e5kH6Z2Xe3/cjGM0VYhJyOdJ1Wj4ov4TvjJk42QZnOTXnjJxj164C
Xy6/ebersx4Xl9V5MpXnGj3bdJGx3JJVSjajTLu0oYT72TWpdEDZ361vbO2iVSMkpESiMvHn9n/Tw+GJ
n7irBkXefHMzZUFibfNX98RnxvIQYWQy4FY/XY9y2W8+ev+EMKLmAl+b/wxfg0j+Jb6KTgQt6oGNkYnj
A2f27kqrvwIAAP//W4XumgIOAAA=
`,
	},

	"/manifests/medium_php.yml": {
		name:    "medium_php.yml",
		local:   "manifests/medium_php.yml",
		size:    3586,
		modtime: 1542073989,
		compressed: `
H4sIAAAAAAAC/+xWUYsbNxB+318x3L3ExXbORwhl39rS0kCgkFzoQymLVhp71dNKqma0PufXF2nXu+uz
nSalDy2Eg8OabzQz+mY+rVarVUGywVaU0N2tN0WBtqusaLGEx2hNtEVAcjFIrHbBRT9gL15cMC8WhXFS
sHY2eRx/LxZFoYWgEsTHGBBugRtNIIWFGgebC0CuRXDcYABvBG9daIuia/v4VBYAK+iT/xFbX7unAgCg
RRZl/gXQV8IH/8xHumi5hE1e0GMs4T0Lq0RQ1fcbytZ+U9c+K65rU2VdS7QEGUNAy+YAzppDgjQBRe9d
YFQ5iqPKB7fVBo8lCdVqO3I2rRaLwcFoG58q6exW72LouRsgAKJmWgD4WBstq0c80NyciDnGJmrWk9uQ
hdgFsRtr0u1sAeC2WwwlfKij5fgeQ4ehOMlITcJ/ENZZLYWZqktMbl6v716t3j68H80dBsoTYAQj8WB3
VClNj1PaVlixQ5WtQ8/Grrx9N4WTQjba7kp4h0L9GjTjBAUUjJXz/cT9FFz7Jp1tcFCCRQ4/Y2v12Xk/
mfks94+t58MMztFJf8RqV5ewubsbsdthjFPPO7Q9VTFqNRu9GrPlZR7oJehtXi7B2R6AvYtGpQHdWQyC
UYGgHDWnsch7F6Zjr4BibZGHOSSLvNqMBQ3eFaGMQfPhROdETWVpVzwbQe1LIBasZTHT5R7rFeUBoi/V
3Vf5fJXP/1M+xglV1cIIKzEMLqe2a761kI9oVSWUCkhUeefMEOAi9JmS3WPdi3AUbnAGqTwSpAkioQJ2
sMP+/MANgkJv3KFFy0AyaM80nn/+Hlj7xhdFZ5FPvsndxMqxaPJCYureOv+93Lzu5ZTJnLF7kdYg7G6+
+f7VCO0E414cRnBTFCfEntR1qRUnV1G2XOL7rMJPNKVBYbhJt1eNZ/saZt9D08UQHDvpTAk/M3u4hV4I
wkAnTEzN+u1B+mVG+//0+7TZBS7h20kTAf+MSFx5wU0JNy9v4DbbdECVpv+YLT9YkFPnc1oX+tBr+CU9
u/aacAmak5t1DMIYt0e1ft4oaq6e5kH6Z2Xe3/cjGM0VYhJyOdJ1Wj4ov4TvjJk42QZnOTXnjJxj164C
Xy6/ebersx4Xl9V5MpXnGj3bdJGx3JJVSjajTLu0oYT72TWpdEDZ361vbO2iVSMkpESiMvHn9n/Tw+GJ
n7irBkXefHMzZUFibfNX98RnxvIQYWQy4FY/XY9y2W8+ev+EMKLmAl+b/wxfg0j+Jb6KTgQt6oGNkYnj
A2f27kqrvwIAAP//W4XumgIOAAA=
`,
	},

	"/manifests/small_php.yml": {
		name:    "small_php.yml",
		local:   "manifests/small_php.yml",
		size:    4409,
		modtime: 1542073948,
		compressed: `
H4sIAAAAAAAC/+xXTY/bNhO+61fM671kX9jOerObtro1bZoGSJEgH+ihKASKHFvsUiTDIe04v74gRUuy
rU3SoocWCAwY0sxwZvjMpxaLRUG8wZaVsL1arooC9bbSrMUS7oJWQRcOyQTHsdo4E2zmPXgwQb68LJTh
zEujo8Th+fKyKCRjVAL7GBzCBfhGEnCmocZMMw7ItAjGN+jAKubXxrVFsW07/VQWAAvojP8RWlubDwUA
QIuelekJoPPE7+2JDDdB+xJWwwFIj3QXSnjjmRbMierJihK1O79tT/zcttHJbUs0Bx6cQ+3VHoxW+8iS
BBSsNc6jSFoMVdaZtVR48I6JVuoevuHt8jILKKnDh4obvZab4DoYMwuAqBleAGyoleTVHe5pTI4YHXQT
NctBLFshbxzb9D7JdvQCYNZrdCW8q4P24Q26LbriyCI1kf8D00ZLztTgXURy9Xh5dbN48fZNT96io5QM
inkkn+mGKiHpbjDbMs02KBI1h6+PyovXgzrOeCP1poTXyMSvTnocWA6Zx8rYLvl+cqZ9Hu+WBQTzLKkf
obX4YruftHxm+2lr/X7ETtpJfsRqU5ewurrqeRc5o2PMt6g7qEKQYpR6NSbKw5Tbc5Dr9DoHozsG7ExQ
IiboRqNjHgUwSlqTGY1+Z9xw7QVQqDX6nIek0S9WvUNZuiLkwUm/Pyp5oqbStClOUlDaEsgzL3kxKtEd
1gtKCUTnJfiZuvtaP1/r579ZP8owUdVMMc3RZZFj2n2yNeN3qEXFhHBIVFljVFYwyfrCmt1h3VVhX7nO
KKTyAJAkCIQCvIENdvcH3yAItMrsW9QeiDtpPfX3H+8GS9vYothq9EfzeTugcnCaLOMYo7dMv4erx105
JTBH6E7C6pjejA9f3/SsDfO4Y/ueuSqKI2CP/JoKxVEvSpQpvM88/ERQGmTKN7F71Xh2rvHedqyhMTjj
DTeqhJ+9t3ABXSEwBVumQgzWb2+5nSdu90+/D4eN8yV8O9SEw/cByVeW+aaE2cMZXCSadChi9h+spY0F
fYx8Mmtcp3oJL+MKtpOEc5A+imnjgSlldiiWp4Gi5t7bvOX2xM3r6y4Fg7oHmMiZ1nQ/LO+EncP3Sg2Y
rJ3RPgbnDJxD1O5l/PXyG0e7OotxMV2dR1l5XqNnhyYRSyFZRGMjyKSJB0q4HrVJIR3yrrc+17UJWvQs
xjkSlRE/s/tMDPO6H7GrckXO/j8brCB5qdPUPZIZoZw19Eg6XMsP92uZlhun3t8BjKiZwGv1r8ErF8k/
hFeRF5aK8bSCHWXeCS9n9T2fcH2zpEOzLIYhkj6D1sal4RGzWXKEmsXREleH+JRm794E2DGd2k43Y9KJ
+HV4WHaWxQVYhYwwjqbETjMLpCaPTCyLdk/vVXXQm2/Uyk2cXtXamTaON6mF5HGc7Rp0CDsEatKkz4Kx
S7SHdmac3EhdNYZ8CYzEMflgaYIVCF1OrlOWZUQ748TA6gRn3ehcpFt0gezXvNnt8ptZ3pddbHFdLngZ
t8ZnaUCrV8FZQ/Er+pUhkrXCoSHOnjCSfDaH2bFwpPyCrXH7l9bLVn5EMev65Zq1Uu2T8ttplc9Q32SN
t/nQYQmGeCS/gNTw7Mn8tEsnmOE2xvvm6rvH/YANtnLoUafEFWxPJTy67ZSTqlCvjeMY948SnmpWKxTT
3mVmdPBHSd1z5+QQmox4JHTgDqGZsZqL1fWjm//NimLLnIwajmqk/xYYfaPEtz8DAAD//8w4kFA5EQAA
`,
	},

	"/manifests": {
		name:  "manifests",
		local: `manifests`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"manifests": {
		_escData["/manifests/large_php.yml"],
		_escData["/manifests/maximum_php.yml"],
		_escData["/manifests/medium_php.yml"],
		_escData["/manifests/small_php.yml"],
	},
}
