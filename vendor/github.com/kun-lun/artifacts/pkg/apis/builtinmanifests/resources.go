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
		size:    3695,
		modtime: 1542091302,
		compressed: `
H4sIAAAAAAAC/+xXW4vkNhN9968oZl62P7p7p4dl+fBbEhKysBDYC3kIwchSdVsZWVJUJff2/vog2W27
b3sJeUhgGRisqlJV6Zw6Gs1qtSpINtiKErqH9aYo0HaVFS2W8BStibYISC4GidUuuOgH37NnV8yLRWGc
FKydTRHH78WiKLQQVIL4GAPCPXCjCaSwUONgcwHItQiOGwzgjeCtC21RdG2fn8oCYAV98T9i62v3oQAA
aJFFmb8A+k744M9ipIuWS9jkBT3FEt6ysEoEVX2/oWztN3XtWXNdmzrrWqIlyBgCWjYHcNYckksTUPTe
BUaVsziqfHBbbfDYklCttiNm02qxGAKMtvFDJZ3d6l0MPXaDC4ComRYAPtZGy+oJDzQ3J2COuYma9RQ2
VCF2QezGnnQ7WwC47RZDCe/raDm+xdBhKE4qUpP8PwjrrJbCTN0lJDcv1w8vVq/fvR3NHQbKE2AEI/Fg
d1QpTU9T2VZYsUOVrQNnIyuv30zppJCNtrsS3qBQvwbNOLkCCsbK+X7ifgqufZXONgQowSKnn6G1+uK6
n6x8UfvH1vNh5s7ZSX/EaleXsHl4GH33wxgnzju0PVQxajUbvRqz5Xke6CXobV4uwdneAXsXjUoDurMY
BKMCQTlrLmOR9y5Mx14BxdoiD3NIFnm1GRsaoitCGYPmw4nOiZrK0q44G0HtSyAWrGX2BGdwVm1+fawT
CwmNYibgPdYrypNGXyvQbzr7prP/ps6ME6qqhRFWYhhCTm23Ymshn9CqSigVkKjyzpkhwVXXF2p7j3Uv
wlHhvY6PAGmCSKiAHeywPz9wg6DQG3do0TKQDNozfUb51705S/+GkE7h9Sjf+KLoLPLJA6CbkD0enLyQ
mCZgnX+eb172ksyEXNxNZ9QEYXfzzY8vRtdOMO7FYXRuiuKEnJO+rtF5cp1lyzXOLjr8BLENCsNNugHr
y1u3Yfa9a7pcgmMnnSnhZ2YP99CLSRjohImJ8N/eSb/M3v43/T5tdoFL+P+kq4B/RiSuvOCmhLvnd3Cf
bTqgSgo6VsuvI+Q0PbmsC33qNfyS3nh7TbgEzSnMOgZhjNujWp8TRc3N07yT/qzNx8d+jKO5AUzyXM90
G5b3yi/hO2MmTLbBWU7kXIBzZO2m4+slPGe7uuC4uK7wk6m81PnFpquIZUpWqdgMMu3ShhIeZ1et0gFl
fz+/srWLVo0uISUSlQk/t/8Mh8P/Ewm7alDk3f/upipIrG3+y30SM0N5yDAiGXCrP9zOcj1uPnp/BzCi
5gpem38NXoNI/iG8ik4ELeoBjRGJ4yNp9nZLq78CAAD//znUjaRvDgAA
`,
	},

	"/manifests/maximum_php.yml": {
		name:    "maximum_php.yml",
		local:   "manifests/maximum_php.yml",
		size:    3695,
		modtime: 1542091297,
		compressed: `
H4sIAAAAAAAC/+xXW4vkNhN9968oZl62P7p7p4dl+fBbEhKysBDYC3kIwchSdVsZWVJUJff2/vog2W27
b3sJeUhgGRisqlJV6Zw6Gs1qtSpINtiKErqH9aYo0HaVFS2W8BStibYISC4GidUuuOgH37NnV8yLRWGc
FKydTRHH78WiKLQQVIL4GAPCPXCjCaSwUONgcwHItQiOGwzgjeCtC21RdG2fn8oCYAV98T9i62v3oQAA
aJFFmb8A+k744M9ipIuWS9jkBT3FEt6ysEoEVX2/oWztN3XtWXNdmzrrWqIlyBgCWjYHcNYckksTUPTe
BUaVsziqfHBbbfDYklCttiNm02qxGAKMtvFDJZ3d6l0MPXaDC4ComRYAPtZGy+oJDzQ3J2COuYma9RQ2
VCF2QezGnnQ7WwC47RZDCe/raDm+xdBhKE4qUpP8PwjrrJbCTN0lJDcv1w8vVq/fvR3NHQbKE2AEI/Fg
d1QpTU9T2VZYsUOVrQNnIyuv30zppJCNtrsS3qBQvwbNOLkCCsbK+X7ifgqufZXONgQowSKnn6G1+uK6
n6x8UfvH1vNh5s7ZSX/EaleXsHl4GH33wxgnzju0PVQxajUbvRqz5Xke6CXobV4uwdneAXsXjUoDurMY
BKMCQTlrLmOR9y5Mx14BxdoiD3NIFnm1GRsaoitCGYPmw4nOiZrK0q44G0HtSyAWrGX2BGdwVm1+fawT
CwmNYibgPdYrypNGXyvQbzr7prP/ps6ME6qqhRFWYhhCTm23Ymshn9CqSigVkKjyzpkhwVXXF2p7j3Uv
wlHhvY6PAGmCSKiAHeywPz9wg6DQG3do0TKQDNozfUb51705S/+GkE7h9Sjf+KLoLPLJA6CbkD0enLyQ
mCZgnX+eb172ksyEXNxNZ9QEYXfzzY8vRtdOMO7FYXRuiuKEnJO+rtF5cp1lyzXOLjr8BLENCsNNugHr
y1u3Yfa9a7pcgmMnnSnhZ2YP99CLSRjohImJ8N/eSb/M3v43/T5tdoFL+P+kq4B/RiSuvOCmhLvnd3Cf
bTqgSgo6VsuvI+Q0PbmsC33qNfyS3nh7TbgEzSnMOgZhjNujWp8TRc3N07yT/qzNx8d+jKO5AUzyXM90
G5b3yi/hO2MmTLbBWU7kXIBzZO2m4+slPGe7uuC4uK7wk6m81PnFpquIZUpWqdgMMu3ShhIeZ1et0gFl
fz+/srWLVo0uISUSlQk/t/8Mh8P/Ewm7alDk3f/upipIrG3+y30SM0N5yDAiGXCrP9zOcj1uPnp/BzCi
5gpem38NXoNI/iG8ik4ELeoBjRGJ4yNp9nZLq78CAAD//znUjaRvDgAA
`,
	},

	"/manifests/medium_php.yml": {
		name:    "medium_php.yml",
		local:   "manifests/medium_php.yml",
		size:    3695,
		modtime: 1542091287,
		compressed: `
H4sIAAAAAAAC/+xXW4vkNhN9968oZl62P7p7p4dl+fBbEhKysBDYC3kIwchSdVsZWVJUJff2/vog2W27
b3sJeUhgGRisqlJV6Zw6Gs1qtSpINtiKErqH9aYo0HaVFS2W8BStibYISC4GidUuuOgH37NnV8yLRWGc
FKydTRHH78WiKLQQVIL4GAPCPXCjCaSwUONgcwHItQiOGwzgjeCtC21RdG2fn8oCYAV98T9i62v3oQAA
aJFFmb8A+k744M9ipIuWS9jkBT3FEt6ysEoEVX2/oWztN3XtWXNdmzrrWqIlyBgCWjYHcNYckksTUPTe
BUaVsziqfHBbbfDYklCttiNm02qxGAKMtvFDJZ3d6l0MPXaDC4ComRYAPtZGy+oJDzQ3J2COuYma9RQ2
VCF2QezGnnQ7WwC47RZDCe/raDm+xdBhKE4qUpP8PwjrrJbCTN0lJDcv1w8vVq/fvR3NHQbKE2AEI/Fg
d1QpTU9T2VZYsUOVrQNnIyuv30zppJCNtrsS3qBQvwbNOLkCCsbK+X7ifgqufZXONgQowSKnn6G1+uK6
n6x8UfvH1vNh5s7ZSX/EaleXsHl4GH33wxgnzju0PVQxajUbvRqz5Xke6CXobV4uwdneAXsXjUoDurMY
BKMCQTlrLmOR9y5Mx14BxdoiD3NIFnm1GRsaoitCGYPmw4nOiZrK0q44G0HtSyAWrGX2BGdwVm1+fawT
CwmNYibgPdYrypNGXyvQbzr7prP/ps6ME6qqhRFWYhhCTm23Ymshn9CqSigVkKjyzpkhwVXXF2p7j3Uv
wlHhvY6PAGmCSKiAHeywPz9wg6DQG3do0TKQDNozfUb51705S/+GkE7h9Sjf+KLoLPLJA6CbkD0enLyQ
mCZgnX+eb172ksyEXNxNZ9QEYXfzzY8vRtdOMO7FYXRuiuKEnJO+rtF5cp1lyzXOLjr8BLENCsNNugHr
y1u3Yfa9a7pcgmMnnSnhZ2YP99CLSRjohImJ8N/eSb/M3v43/T5tdoFL+P+kq4B/RiSuvOCmhLvnd3Cf
bTqgSgo6VsuvI+Q0PbmsC33qNfyS3nh7TbgEzSnMOgZhjNujWp8TRc3N07yT/qzNx8d+jKO5AUzyXM90
G5b3yi/hO2MmTLbBWU7kXIBzZO2m4+slPGe7uuC4uK7wk6m81PnFpquIZUpWqdgMMu3ShhIeZ1et0gFl
fz+/srWLVo0uISUSlQk/t/8Mh8P/Ewm7alDk3f/upipIrG3+y30SM0N5yDAiGXCrP9zOcj1uPnp/BzCi
5gpem38NXoNI/iG8ik4ELeoBjRGJ4yNp9nZLq78CAAD//znUjaRvDgAA
`,
	},

	"/manifests/small_php.yml": {
		name:    "small_php.yml",
		local:   "manifests/small_php.yml",
		size:    4518,
		modtime: 1542091284,
		compressed: `
H4sIAAAAAAAC/+xXS4/bthO/61PM33vJ/mE7681u2urWtGkaIEWCPNBDUQgUObbYpUiFQ9pxPn3BhyXZ
1iZp0UMLBAssrJnhzPA3Ty4Wi4J4gy0rYXu1XBUF6m2lWYsl3HmtvC4skvGWY7WxxneZ9+DBBPnyslCG
MyeNDhKH35eXRSEZoxLYR28RLsA1koAzDTVmmrFApkUwrkELnWJubWxbFNs26aeyAFhAMv6Hb7vafCgA
AFp0rIy/AJInbt+dyHDjtSthNRyA+JPufAlvHNOCWVE9WVGkpvPb9sTPbRuc3LZEc+DeWtRO7cFotQ8s
SUC+64x1KKIWQ1VnzVoqPHjHRCt1D9/wdXmZBZTU/kPFjV7LjbcJxswCIGqGD4DO10ry6g73NCYHjA66
iZrlIJatkDOWbXqfZDv6ADDrNdoS3tVeO/8G7RZtcWSRmsD/gWmjJWdq8C4guXq8vLpZvHj7pidv0VJM
BsUckst0Q5WQdDeYbZlmGxSRmsPXR+XF60EdZ7yRelPCa2TiVysdDiyLzGFlupR8P1nTPg93ywKCORbV
j9BafLHdT1o+s/207dx+xI7aSX7EalOXsLq66nkXOaNDzLeoE1TeSzFKvRoj5WHM7TnIdfycg9GJATvj
lQgJutFomUMBjKLWaEaj2xk7XHsB5GuNLuchaXSLVe9Qlq4IubfS7Y9KnqipNG2KkxSUXQnkmJM8cqxR
OLI27iTLEIWARjGq5R3WC4qZRue1+pkC/VpoXwvtv1loyjBR1UwxzdFmkWPafbI143eoRcWEsEhUdcao
rGCS9YXFvcM6VWFf4qmQDwBJAk8owBnYYLo/uAZBYKfMvkXtgLiVnaPPlP40N2pJ+wQ3AqeluqYriq1G
d7QMbAdkDxenjnEMGbCMfw9Xj1NJxoCcNaeT0FimN+PD1zc9a8Mc7ti+Z66K4ig4R35NhfOon0XKVMzO
PPxEYBtkyjWhA9bnbbdxrkusoblY4ww3qoSfnevgAlIxMQVbpnwI+G9veTeP3PSffh8OG+tK+HaoK4vv
PZKrOuaaEmYPZ3ARadKiCBV0sBbXI3Qhe6JZY5PqJbwM+95OEs5BuiCmjQOmlNmhWJ4Gipp7b/OWdydu
Xl+nNPbqHmACZ1rT/bC8E90cvldqwGRtjXYhOGfgHKJ2L+Ovl/A42tVZjIvpCj/KyvM6Pzs0iVgMySIY
G0EmTThQwvWo1Qppkaf+/FzXxmvRsxjnSFQG/MzuMzHMb4uAXZUrcvb/2WAFyUkdJ/eRzAjlrKFH0uJa
frhfy7TcOPX+DmBEzQReq38NXrlI/iG8irz0VIzHNe4o8054OavveS/2zZIOzbIYBlF8c62NjQMoZLPk
CDUL4ymMmfArzu+98bBjOradNGHiifAUPSxMy+ICOoWMMIy3yI5zD6Qmh0wsi3ZP71V10Jtv1MpNmIDV
2po2jEipheRhJO4atAg7BGritpAFQ5doD+3MWLmRumoMuRIYiWPywdIEyxPanFynrI4R7YwVAysJztLo
XMRbpED2q+LsdvnNLO/cNrS4lAtOhs3zWRzy6pW3naHwZH9liGStcGiIsyeMJJ/NYXYsHCi/YGvs/mXn
ZCs/opilfrlmrVT7qPx2WuUz1DdZ420+dFikIRzJHyA1PHsyP+3SEWa4DfG+ufrucT9gfVdZdKhj4gq2
pxIe3SblpCrUa2M5hh2mhKea1QrFtHeZGRz8UVL6nZwcQpMRD4QE7hCaGau5WF0/uvnfrCi2zMqg4ahG
+vfE6J0Tvv4MAAD//3/YPKemEQAA
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
