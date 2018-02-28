// Code generated by go-bindata.
// sources:
// resources/capturedio.gtl
// resources/contents.gtl
// resources/htmltemplate.gtl
// resources/index-content.gtl
// resources/index.gtl
// resources/interestinggivens.gtl
// resources/javascript.gtl
// resources/style.gtl
// resources/test-body.gtl
// DO NOT EDIT!

package htmlspec

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resourcesCapturedioGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xcb\x0a\xc2\x30\x10\x45\xf7\x85\xfe\xc3\x10\xba\xf4\xb1\xe8\x4e\xda\xfe\x80\x3b\x17\xba\x8e\xcd\x35\x09\x1d\x1b\xe9\x0b\xcb\x90\x7f\x97\x14\x05\x77\x17\x86\x73\xce\x88\x18\x3c\x7c\x0f\x52\xad\x7e\x4d\xf3\x00\xb3\xf7\x41\xc5\x98\x67\x79\x26\x32\xe8\xde\x82\x8a\x0e\xeb\x8e\x8a\x45\xf3\x0c\x3a\xd5\x74\xa0\x74\xaf\x5c\x49\x2d\xeb\x71\xac\x15\x07\x7b\xc6\xaa\x88\x83\xed\xb0\xd6\x4a\x24\x21\x31\xaa\xe6\xb7\xaa\xa3\x2b\x9b\x3c\xab\x8c\x5f\xfe\xa0\xeb\x66\x74\xde\x3a\xf6\xd6\x4d\xf4\x7e\x32\xdd\x70\xbf\x60\x0c\xf3\xd0\x62\xc3\xb7\x6a\x12\x18\xbf\x34\xe9\x27\xf4\x26\xe5\xbf\xe3\x13\x00\x00\xff\xff\x3c\x1c\x63\x49\xc0\x00\x00\x00")

func resourcesCapturedioGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesCapturedioGtl,
		"resources/capturedio.gtl",
	)
}

func resourcesCapturedioGtl() (*asset, error) {
	bytes, err := resourcesCapturedioGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/capturedio.gtl", size: 192, mode: os.FileMode(438), modTime: time.Unix(1516043421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesContentsGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\x41\x6b\x03\x21\x10\x85\xef\x82\xff\x61\xd8\x16\x72\x4a\x84\x1c\x8b\xeb\xa5\xa1\xd0\x6b\x9b\x3f\x20\x71\xb6\x3b\x74\x30\xcb\xaa\xb9\x88\xff\xbd\xec\xd6\x76\x4b\x49\x72\xdb\x9b\xe3\xfb\x9e\xef\xe1\xe4\xec\xb0\x23\x8f\xd0\x9c\xce\x3e\xa2\x8f\xa1\x29\x45\x0a\xed\xe8\x02\x27\xb6\x21\xb4\x8b\x60\xa4\x00\x00\xd0\xfd\xde\x3c\xd7\x3b\xad\xfa\xbd\x91\xa2\x0a\x89\x6f\x79\x72\x1e\xad\xff\x40\x78\x8c\x18\xe2\xc1\x46\x0b\x4f\x2d\xec\x8e\x75\x98\x12\xbf\x29\xea\x16\x64\x96\xdf\x30\x24\x8e\xbb\x17\x4b\x8c\xee\x87\x9b\xc3\x98\x6a\xd8\x66\x32\x6c\xbb\x99\xd8\x18\x6d\xa1\x1f\xb1\x6b\x9b\x87\x9c\xaf\x3e\x35\x1d\x5f\x0f\xa5\x34\xe6\x3f\x70\xa4\xc8\x58\x8a\x56\xd6\x68\xc5\xf4\x5b\x1d\x39\x20\xdc\x6a\xf6\xfe\x49\xc3\x70\xbf\x9a\x3f\xc7\xed\x98\xfc\x6a\xdd\xee\x65\x0f\x36\x84\xb5\xbe\xc5\xbb\x65\x71\x7f\x06\xad\x12\x1b\x29\xb4\x72\x74\x31\x52\x54\xed\x2b\x00\x00\xff\xff\xaa\x8f\x02\x51\x69\x02\x00\x00")

func resourcesContentsGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesContentsGtl,
		"resources/contents.gtl",
	)
}

func resourcesContentsGtl() (*asset, error) {
	bytes, err := resourcesContentsGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/contents.gtl", size: 617, mode: os.FileMode(438), modTime: time.Unix(1517867335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesHtmltemplateGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x6b\xc2\x30\x14\xc7\xef\x85\x7e\x87\x2c\x77\xf3\x56\x3c\x29\xb1\x87\xa9\xb0\x81\xdb\x64\x64\x6c\x3b\x46\xf3\x66\x3b\x62\x2a\xe6\xb1\x2a\x21\xdf\x7d\xc4\x16\x51\x66\x4e\x7f\xf2\xfe\xf9\xfd\x92\x84\x60\xf0\xbb\x76\xc8\xf8\x4a\x7b\xe4\x31\xe6\x99\xbc\x9b\xbd\x4e\xd5\xd7\x72\xce\x2a\xda\x5a\xb6\x7c\x7f\x58\x3c\x4d\x19\x1f\x00\x7c\x0c\xa7\x00\x33\x35\x63\x9f\x8f\xea\x79\xc1\x0a\x71\xcf\xd4\x5e\x3b\x5f\x53\xdd\x38\x6d\x01\xe6\x2f\x3c\xcf\x58\xbf\x78\x45\xb4\x1b\x03\xb4\x6d\x2b\xda\xa1\x68\xf6\x1b\x50\x6f\x70\x48\xd0\x22\x51\xfa\x38\xa0\x0b\x84\x30\x64\x78\x99\x67\xf2\xa4\x3e\x6c\xad\xf3\x93\x1b\x9c\x62\x34\x1a\x75\xc7\x79\x2a\x8d\xad\x76\x9b\x09\x47\xc7\xd9\x39\x9d\x20\xa8\x4d\xd9\x5d\x48\x52\x4d\x16\xcb\x10\x84\x4a\x21\x46\x09\xdd\x4e\x9e\xe5\x59\x08\x84\xdb\x9d\xd5\x84\x8c\xff\xe8\x5f\xed\xd7\xfb\x7a\x47\x9c\x89\xf4\x1d\xd7\x63\x4f\x47\x8b\xe7\x89\x84\x5e\x21\x57\x8d\x39\x9e\x58\xb2\x2a\xae\x2c\x55\xf1\x4f\xb1\x6e\x1c\xa1\x23\x7f\x53\x40\xe8\x69\x90\x68\x17\x92\x1e\x2e\x21\xbd\xb8\x4c\x7d\x74\x26\xc6\xbf\x00\x00\x00\xff\xff\x01\x0b\xc6\xde\xbc\x01\x00\x00")

func resourcesHtmltemplateGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesHtmltemplateGtl,
		"resources/htmltemplate.gtl",
	)
}

func resourcesHtmltemplateGtl() (*asset, error) {
	bytes, err := resourcesHtmltemplateGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/htmltemplate.gtl", size: 444, mode: os.FileMode(438), modTime: time.Unix(1515789047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesIndexContentGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xb1\x6a\xc3\x30\x10\x86\x77\x3f\xc5\xe1\x16\x32\xd9\x86\x8c\x45\xd6\xd2\x50\xe8\x9a\xe6\x05\x44\x74\xaa\x45\x85\x62\xac\x4b\x28\x08\xbd\x7b\x89\xac\x44\xb6\x69\x42\xa6\x4c\xb6\xee\xfe\xff\xf4\xfd\x87\xbc\x97\xa8\xb4\x45\x28\xf7\x07\x4b\x68\xc9\x95\x21\x14\x4c\xea\x13\xec\x8d\x70\xae\xcd\x75\x5e\x00\x00\xb0\x6e\xcd\xdf\x53\x89\x35\xdd\x9a\x17\x63\xf9\xb6\xc3\xfb\x41\xd8\x6f\x84\x57\x6d\x25\xfe\x6e\x04\x09\x78\x6b\xa1\x0e\x21\x76\x47\xb3\xb9\x78\x09\x1d\x55\x4a\x68\x83\x12\xa2\xbe\x1a\xd0\x1d\x0d\xa5\x59\xd9\x41\x9c\x09\xe8\x06\x54\x6d\xe9\x7d\x1e\x5d\x6f\x51\x85\x50\xf2\x59\x6d\xa7\xc9\x60\x08\xac\x11\x9c\x35\x92\xf2\xa8\x2b\xda\xf9\xda\x2b\xd9\x2e\x1d\x26\x84\xa3\x58\xab\xac\x8c\xaa\x6d\x64\xab\x3f\x22\xef\x42\x3e\x62\xca\x14\x6c\x35\x09\xb6\xba\x87\xfe\xe2\xfd\xbf\x77\x9c\x7f\x3f\x37\x29\xda\x4c\x30\x4f\x27\xf9\x02\x1a\x8d\x43\xb8\x45\xfe\xf5\xa3\xfb\xfe\x21\x74\x7b\xa0\x6a\x38\xda\xe7\xb3\x3f\xc0\xd6\x0b\xe7\x9e\xbe\x56\x3b\xdd\xda\xf2\xcc\x1a\x69\x2e\xaf\x3f\x77\x58\x23\xf5\x89\x17\xe9\x93\x3a\x7f\x01\x00\x00\xff\xff\xd0\x21\x3f\xb8\x82\x03\x00\x00")

func resourcesIndexContentGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesIndexContentGtl,
		"resources/index-content.gtl",
	)
}

func resourcesIndexContentGtl() (*asset, error) {
	bytes, err := resourcesIndexContentGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/index-content.gtl", size: 898, mode: os.FileMode(438), modTime: time.Unix(1519838158, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesIndexGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd0\x51\x4f\x83\x30\x10\x00\xe0\x77\x12\xfe\x43\xed\xfb\x38\xc9\x9e\x58\x3a\x1e\x84\x25\x2e\x99\xba\x18\x8c\xfa\x58\xd7\x73\xd4\x94\x42\xe0\x22\x2c\x84\xff\x6e\x0a\xb8\xcc\xb8\x3e\x5d\x73\x77\xdf\xf5\xda\xf7\x0a\x3f\xb5\x45\xc6\xb5\x55\xd8\xf1\x61\xf0\x3d\x71\x93\x3e\x25\xd9\xfb\x7e\xc3\x72\x2a\x0c\xdb\xbf\xdc\xed\xb6\x09\xe3\x0b\x80\xd7\x65\x02\x90\x66\x29\x7b\xbb\xcf\x1e\x76\x2c\x0c\x6e\x59\x56\x4b\xdb\x68\xd2\xa5\x95\x06\x60\xf3\xc8\x7d\x8f\xcd\x87\xe7\x44\xd5\x0a\xa0\x6d\xdb\xa0\x5d\x06\x65\x7d\x84\xec\x19\x3a\x87\x86\x4e\x99\xc3\x05\x5d\x10\x81\x22\xc5\x63\xdf\x13\xe3\xe8\xae\x30\xb6\x59\x5f\x71\xc2\x28\x8a\xa6\x76\xee\x8a\x56\x46\xda\xe3\x9a\xa3\xe5\xec\x1c\x8d\x08\x4a\x15\x4f\x0f\x12\xa4\xc9\x60\xbc\x75\x5b\x0a\x98\x2e\xbe\xe7\x7b\x7d\x4f\x58\x54\x46\x12\x32\xfe\x25\xbf\x65\x73\xa8\x75\x45\x9c\x05\xee\x27\xfe\xa6\x1b\x3a\x19\x3c\x67\x04\xcc\xba\xf8\x28\xd5\x69\xb4\x44\x1e\xfe\x0e\xc8\xc3\x7f\xfa\xa1\xb4\x84\x96\x9a\x0b\x61\xee\x14\xe0\x36\x89\x5d\x39\x5a\x35\x0c\x3f\x01\x00\x00\xff\xff\x6e\x91\x52\x86\x95\x01\x00\x00")

func resourcesIndexGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesIndexGtl,
		"resources/index.gtl",
	)
}

func resourcesIndexGtl() (*asset, error) {
	bytes, err := resourcesIndexGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/index.gtl", size: 405, mode: os.FileMode(438), modTime: time.Unix(1517951949, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesInterestinggivensGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x4d\x8b\xc3\x20\x10\x86\xef\x81\xfc\x87\x41\x72\xdc\x5d\x0f\xb9\x2d\x26\xd7\x52\x7a\xef\xdd\xd6\xa9\x4a\xc4\x40\xb4\x81\x30\xf8\xdf\x8b\x26\xe9\xc7\xa1\x47\x9d\xe7\x99\x77\x5e\x22\x85\x37\xeb\x11\x98\xf5\x11\x27\x0c\xd1\x7a\xfd\xab\xed\x8c\x3e\xb0\x94\xea\xaa\xae\x84\x69\xe1\xea\x64\x08\x1d\x73\xa3\x3e\xe1\xc2\xfa\xe3\x8b\x85\x43\x61\x05\x37\x6d\x5f\xe8\x28\x2f\x0e\x77\xe1\x6d\xe9\xca\x81\x1b\xf5\x59\xba\x3b\xb2\xbe\xae\x00\x00\x88\x26\xe9\x35\x42\x33\xe0\xf2\x03\xcd\x9c\x67\xf0\xdf\xc1\x1f\xe4\xf4\x4c\x88\x38\x6d\xec\xfa\x32\xfb\xf2\x21\x9f\x42\x94\xcd\x94\x04\x8f\xe6\x03\x53\xdf\x6e\x28\x4e\xc9\x29\x96\xda\x2c\xc1\x9f\x39\x44\xe8\x55\x8e\x17\xbc\xb4\x29\xc5\xb6\xcf\x47\x00\x00\x00\xff\xff\x45\xb5\xf8\x13\x31\x01\x00\x00")

func resourcesInterestinggivensGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesInterestinggivensGtl,
		"resources/interestinggivens.gtl",
	)
}

func resourcesInterestinggivensGtl() (*asset, error) {
	bytes, err := resourcesInterestinggivensGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/interestinggivens.gtl", size: 305, mode: os.FileMode(438), modTime: time.Unix(1516043421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesJavascriptGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\xdf\x6f\xdb\x36\x10\x7e\x0f\x90\xff\xe1\x4a\x04\x15\xb5\xa8\x52\x5b\xec\x29\x8e\x52\x04\x5d\x51\x0c\xfb\xf1\x50\x14\x1b\x30\x5b\x1d\x58\xe9\x2c\x31\xa5\x49\x95\xa4\x13\x1b\xae\xfe\xf7\x81\x92\x65\xeb\x87\x9d\xa4\xc1\x1e\x36\x3d\xc9\x77\xbc\xef\x8e\x77\x1f\x3f\xd1\x9b\x4d\x86\x73\x2e\x11\xc8\x0d\xbb\x65\x26\xd5\xbc\xb4\xa4\xaa\x4e\x4f\x2e\x9b\x77\x30\x3a\x8d\x49\x61\x6d\x79\x11\x45\xec\x86\xad\xc2\x5c\xa9\x5c\x20\x2b\xb9\x09\x53\xb5\xa8\x6d\x91\xe0\x9f\x4d\x74\xf3\x75\x89\x7a\x1d\xbd\x0a\x7f\x0c\x5f\x6f\x7f\x84\x0b\x2e\xc3\x1b\x43\xc0\xae\x4b\x8c\x89\xc5\x95\x8d\x3a\x79\xae\x2e\xa3\xe6\xed\xea\xf4\x64\x9f\xf1\xc8\xda\xd3\x13\x00\x80\x28\x82\xcb\x67\xd3\xb7\x3f\x5d\x7f\xbc\x9e\x36\x96\x0f\x98\xbf\x5b\x95\x61\xa9\x95\x55\x2e\x34\xd4\x58\x0a\x96\x22\xc4\x30\x5f\xca\xd4\x72\x25\x81\x1a\xab\x03\xd8\x3a\x74\x00\x52\xc9\xdf\x98\x4d\x0b\xcc\x3e\x6c\x6d\x3e\x6c\x1a\x38\xf7\x8c\xdd\x10\x1f\x32\x7e\xfb\xd6\x49\x71\xcb\xc4\x12\x7b\x30\xee\xd1\x68\x97\x5a\x42\xed\xdc\x7b\xaa\xc9\xfe\xfd\x96\x69\xd0\x68\x96\xc2\x42\x0c\xd3\x64\xe2\x5a\xd1\xf5\x95\xca\xf0\x3a\x43\x0c\x2f\x07\x61\x0b\x57\x4f\xc7\x76\x57\x70\x81\x40\x69\x6d\x87\x18\x6c\xc1\x4d\x88\x2b\x4c\xdd\xfe\x7d\x1f\x9e\xc5\x20\x97\x42\x1c\x28\xd2\xa5\x0f\xcb\xa5\x29\xe8\x78\x9b\x2e\x38\x34\xcb\xcf\xc6\x6a\x2e\x73\xda\xd6\x13\x34\xe9\x43\x2e\x33\x5c\xf9\xbe\x3f\x39\x8e\xd9\x36\xbe\x29\x6c\xb4\xb4\xb3\xc3\xba\x62\xc1\x8c\xfd\xd9\xa1\x76\xd6\x55\xfb\xd7\x27\x55\xdb\x14\xd8\x05\xa9\xe7\xb2\xc5\xba\x51\x5c\x52\x42\xda\xc2\xaa\x76\xe5\x6e\xba\x6b\x66\x4d\x89\x29\xdd\xb5\x6e\xb7\x64\xeb\x71\xfc\x4b\xd1\x18\xcc\x20\x06\xaf\xe0\x79\x21\x78\x5e\x58\xcc\xbc\xc9\x70\xe5\xce\xd9\x63\x28\x0a\x5c\xa0\xb4\x01\x94\x8c\x6b\xd3\x1b\x11\x9f\x03\xad\xad\xa1\x40\x99\xdb\x02\xe2\x18\x5e\x1e\x61\x5a\xbf\x65\x7d\x8c\xb3\x36\x89\x1f\x16\xcc\xbc\x15\xcc\x18\x3a\x2a\xdf\xff\x4e\x60\xc7\xc3\xd4\x41\xa1\x69\xf9\x3b\xa2\xe8\x7b\xad\x96\xe5\xc8\x7d\x16\x22\x4b\x8b\x66\x67\x41\xa7\x13\xa3\x02\x3a\x18\xcd\xd4\x09\x25\x41\x43\x95\x92\x59\x8b\x5a\x06\x40\x7c\x12\x00\xf9\x46\x86\xdc\xda\x96\xd6\xc4\xd5\x21\xa9\x69\xb6\xde\x5d\x59\x75\x7f\xf4\xd2\xa9\x92\xfa\xa3\xc3\x9a\xe3\xca\x09\x02\xde\x6d\xc5\x87\x76\x43\x5a\x2e\x05\x40\x72\xd2\x27\x5d\x77\x00\x76\x21\x68\x8d\xd4\xea\x15\x1d\x7a\xfd\x6e\x53\x9a\x93\x33\xec\xcc\xae\xc3\x75\xf3\x9b\xf3\x68\x04\x4f\x91\xbe\x1a\x36\x62\xae\xb4\xd3\x28\x0d\xbc\x16\x12\xe0\x70\xd9\x86\x6e\x69\x35\x01\x7e\x7e\x3e\x4a\xd1\x92\x67\xbb\x76\xca\x93\x83\x4b\xf6\x44\x01\xef\xd2\x94\x4c\x36\x9d\x8f\x89\x07\xe7\xed\x10\xa6\x3c\x81\x73\xf0\xc8\x95\xb3\xed\xf1\x9c\xed\x32\x72\x31\x57\xde\x18\xb8\xea\x9b\xaa\xee\xd0\x8e\x76\x97\x65\xd9\x31\x7a\x0f\x8f\xf8\x19\xcd\x54\xba\x6c\xc2\x34\xb2\x6c\x4d\x8f\x30\xf1\x8c\x7a\xfb\xa3\x1b\x3a\x54\x3e\xe7\x29\x73\x2b\x3d\xbf\xa1\xf2\x3d\x1c\x1e\x1d\xfe\x9a\x8c\x01\x4c\xc7\x3b\xde\x6c\x49\x7d\x01\x1e\x99\x7e\x22\xc9\x0f\xc4\x0b\xa0\x65\xed\x05\x90\xaf\x4b\x65\x91\x54\xc1\x7d\x91\xe4\x3d\xbf\x45\x49\x7a\x71\x5f\x70\x7d\xa7\x74\xf6\x50\xe4\xb5\xcc\x9e\x14\xf7\x67\xf1\xc4\x84\x1f\x9f\x1a\x38\xbd\x7e\xf1\xd7\xdf\xc9\xe6\x75\x50\xf5\xc3\x53\x25\x8d\x65\xd2\x3e\x18\x3f\x9b\x65\xc9\x79\x3f\x56\x70\x8b\x9a\x09\x32\xa0\x5d\xd2\x57\x8b\x3e\xed\xba\xbc\x68\x3e\x28\xe6\x3f\xc8\x88\x77\xab\x12\x53\x8b\x4f\x1b\x6e\xae\xec\xff\x7c\x44\x26\x45\xc9\x34\x57\x8f\x98\x4c\xad\x92\xd2\xa2\x46\x63\xb9\xcc\xeb\xa3\xe4\xf4\xd5\xc1\x0c\xed\x5e\xf3\x21\xf2\xc3\x39\x17\x16\x35\xf5\x2e\xa4\xb2\xf4\x02\x17\xa5\x5d\xfb\x9e\x1f\x2e\x58\x79\x5f\x2a\xd8\xcb\xe6\x81\xb9\x0f\x67\xef\x74\xf3\x8c\x36\xf9\xdc\xfd\x98\xfa\xb5\x9e\xf6\xc9\x30\xac\xf0\x60\x8b\x07\xc8\xb3\xd9\xe7\xc3\xd8\xce\xf1\xfd\xe8\xc9\xe0\xdb\x53\xf9\x61\x8e\x96\xf6\xe5\xba\x1d\x8c\x50\xf9\x2f\xb8\xde\xf5\x31\x15\x3c\xfd\xf2\x50\xcb\xda\x42\xa5\x2b\x94\x38\x88\x3f\xdc\x05\x9b\xf8\xa1\x55\x79\x2e\xb0\x11\x7f\x52\xf0\x0c\x47\x37\x82\xea\x51\x65\x3c\xc4\x91\x51\x15\xde\xae\x8a\x8e\x1a\xac\x16\xe2\x11\x7c\x6b\x1f\xf7\x99\x6d\xae\x28\x05\x17\x99\x46\xd9\xbd\xef\xb9\x3f\x1b\x3d\xe7\xf4\x65\x12\x5a\x96\xff\xce\x16\x18\x5a\xf5\xab\xba\x43\xfd\x96\x19\xa4\xf5\x35\x9f\x98\xdb\x9c\x1c\x4d\x04\xf7\xc8\xcf\x88\xf8\x61\xaa\x64\xca\x2c\x3d\xc2\xcf\xf6\x79\xba\x46\x1d\x46\x21\xcf\x85\x9d\x4c\x3f\xcd\x66\xe6\xf9\x50\x00\xee\xd3\x9e\x23\x60\xb3\xd9\x9b\x37\xcf\x73\x3b\xf9\x17\x80\xcc\x74\x36\xbb\xbb\x78\x91\x9c\xc7\xc7\x24\xed\x38\x58\x32\xfa\xef\xd3\x3e\x07\x82\x9c\x86\xf5\x7f\x3f\x86\xb8\xdf\x7b\x24\xf6\x4a\x59\x05\x30\x67\xc2\xe0\x3e\x4d\x14\x41\x92\x5c\x9d\x9e\x74\xfe\xa6\x6f\x36\x28\xb3\xaa\xfa\x27\x00\x00\xff\xff\x30\xf1\xd0\xb1\x37\x10\x00\x00")

func resourcesJavascriptGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesJavascriptGtl,
		"resources/javascript.gtl",
	)
}

func resourcesJavascriptGtl() (*asset, error) {
	bytes, err := resourcesJavascriptGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/javascript.gtl", size: 4151, mode: os.FileMode(438), modTime: time.Unix(1515789047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesStyleGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xc9\x8e\xe3\x36\x10\xbd\x37\xd0\xff\x50\xe9\x41\x2e\x0d\xc9\x96\xd7\xe9\xd1\x74\x1a\xf0\xd8\x56\x02\xe4\x1a\xe4\x32\x98\x03\x2d\x96\x2c\x62\x28\x92\x21\xe9\x6d\x0c\xff\x7b\x20\xcb\x8b\x6c\x51\x12\x3a\x8b\x4f\x22\xfd\x5e\x2d\xaa\xc7\x62\x69\xbf\xa7\x98\x30\x81\xf0\x64\xec\x8e\xe3\xd3\xe1\xf0\xf8\xf0\x7a\x7c\x04\xbb\x53\xf8\xcb\x93\xc5\xad\xed\xc6\xc6\x3c\xbd\x3d\x3e\x00\x00\x74\x9f\xe1\xf5\xa7\xaf\xd3\xd9\xe4\x8f\xc9\x57\x78\xee\x16\x9b\xa9\xcd\xb8\x07\x0b\x49\x77\xb0\x2f\x76\xf2\x5f\x46\xf4\x92\x89\x10\x7a\x98\x41\x1f\xb7\xd0\xc7\xec\xf3\xf5\x5f\x45\x28\x65\x62\x19\x42\x50\xda\x5c\x90\xf8\xfb\x52\xcb\x95\xa0\x7e\x2c\xb9\xd4\x21\x7c\x48\x92\xa4\x04\x38\xef\x06\x41\x99\x96\x48\x61\xfd\x84\x64\x8c\xef\x42\x98\x68\x46\xb8\x07\xbf\x21\x5f\xa3\x65\x31\xf1\xe0\x4f\xd4\x94\x08\xe2\x81\x21\xc2\xf8\x06\x35\x4b\xee\xc9\x86\xfd\xc0\x10\x7a\x81\xb2\xa7\x7f\x0e\x8f\x0f\xa7\xdc\x7a\x1e\xa4\x7d\x0f\xd2\x81\x07\xe9\xd0\x83\x74\xe4\x41\x3a\xf6\xc0\xa6\xe5\x5c\xf3\xd7\xe4\x5b\x4d\x84\x49\xa4\xce\x42\x88\x89\x62\x96\x70\xf6\x03\xab\xf6\xca\xbc\xb2\xf3\x8f\xc1\xcf\x15\x70\xbf\x0e\x3c\x72\x80\x07\x75\xe0\xa1\x03\x3c\xac\x03\x0f\x1c\xe0\x51\x1d\xb8\xef\x00\x8f\xeb\xc0\xbd\x2a\xb8\x93\xb2\x65\xca\xd9\x32\xb5\x65\x92\x43\x06\xf3\x97\xf9\x3c\xfa\x58\x96\x8a\xd4\x14\x75\x08\x3d\xb5\x05\x23\x39\xa3\xf0\x61\x3a\x98\x7d\x8a\x22\x97\xc6\x46\x6a\x7b\x57\xf1\xa2\xd8\x67\xb1\x54\x84\xe1\x0a\xb0\xc3\x99\x45\x4d\x78\x39\xd2\x8b\x1e\xc7\xe3\x26\xe6\x77\xdc\x6d\xa4\xa6\x4e\x66\xf0\xd2\xc4\x8c\xa5\x30\x96\x08\xeb\xa2\x8e\x83\x46\xa7\x7f\xad\xa4\x45\xa7\xcb\x97\xa0\x9e\xd7\x31\x0a\x63\x96\xb0\x98\x58\x26\xc5\x25\x69\xaf\x0d\x77\x4a\xb1\x15\x77\x4e\xa8\x15\x58\x09\xff\x28\xa5\x0d\xe6\xd8\x10\x16\x92\xd3\xfb\x24\x94\xc6\x8e\x45\x63\x7d\x21\xad\xaf\x57\xc2\x83\xfb\x25\x13\x16\x35\x1a\xcb\xc4\xf2\x57\xb6\x46\xd1\x22\xba\x28\x8a\x86\xd3\xa9\x43\x74\xfd\xab\xe8\xa2\xe8\xcb\x17\x47\xf1\x9b\x3c\x5d\x5b\x5f\x6e\xa8\x36\x8b\x84\x30\x8e\xf4\x9c\x44\xb1\x6a\x0f\x78\x7c\x73\x02\x5c\x01\xcf\x26\x2f\x93\x8a\xe6\x2e\x5e\x15\x31\xe6\xea\xb5\x58\xb5\x79\xed\x47\xd1\x7c\xde\xec\x75\x32\x9a\xcd\x26\xb3\x7b\xaf\x9c\xdd\x55\xe8\xbc\xe1\xf0\x7b\xb2\x29\xa4\xa8\xf4\xd4\x15\xcf\x65\x65\x51\x58\x03\xe4\x9f\xb5\x65\xe2\x81\x25\x0b\x8e\xb7\xfc\xcb\x49\x1b\x07\xd7\x13\x7e\x31\x4c\x31\x96\xfa\xa8\x56\x77\x58\x24\x4c\xe5\x1a\xf5\xc5\x72\xb1\xac\xc4\x57\x36\xb3\x12\x14\x35\x67\x55\x5b\x85\x89\x96\x3a\x8c\xa2\x4f\xd1\xec\x9d\x3d\xb2\xc0\xe4\x36\x38\x51\x06\x43\x38\x3f\x95\x30\x98\x29\xbb\xf3\x63\xe4\xdc\x84\x60\x52\xb9\xa9\x44\x97\xfe\xd7\xed\xfb\x6a\xda\xa9\x82\x56\x1a\x65\xeb\x8b\x26\xbc\xe3\x2a\xd7\x55\x86\x36\x95\xb4\x3a\xa4\xf8\x0b\x69\xad\xcc\x42\x18\x06\xd5\x03\xd9\x31\x31\x0a\xa2\x99\x34\x77\xc1\x94\xcf\x71\x70\x73\xc3\x1c\xeb\x4a\x38\x5b\x8a\x10\x62\xcc\x5b\x41\xbd\x51\xef\xfa\xdc\x10\x58\xbf\x29\xb0\xff\x43\x17\xa5\xe4\x7a\x0e\xdf\xb7\x93\xc9\x29\x58\x2b\x55\xe8\x44\x2b\x8d\x0d\xb9\x0d\x1c\x5d\x30\xaf\x18\x6e\x49\xa6\x78\x13\x73\x3c\x74\xbc\x95\x94\xd1\x1b\x0e\x65\x46\x71\xb2\x73\x9f\xd1\x0e\x97\xcb\xdf\xf1\x66\x6e\x8d\x57\xda\xe4\xef\x4d\x49\xe6\xac\x5c\xc1\xf8\x37\x87\xb9\x72\x3b\x18\xd7\x38\x79\xd2\x0f\xc7\xa4\x32\x94\xba\x0c\xd0\x3a\x03\x3a\xbf\x31\xdd\x95\xbd\x9d\x8b\xd6\xa8\xf3\x81\x99\x9f\x89\x19\xa3\x94\x57\x83\xdf\x66\x37\x23\xd0\x26\x65\x16\x7d\xa3\x48\x8c\x61\x5e\x68\x7f\xa3\x89\xfa\xfc\xae\x62\x77\x9f\xe1\xdb\xb7\xb7\xe3\xe7\xc4\x6b\xf7\xf8\xf5\xf1\xf6\xf8\xb0\xdf\xa3\xa0\x87\xc3\xdf\x01\x00\x00\xff\xff\x0e\xf0\xb5\x6a\xa2\x0c\x00\x00")

func resourcesStyleGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesStyleGtl,
		"resources/style.gtl",
	)
}

func resourcesStyleGtl() (*asset, error) {
	bytes, err := resourcesStyleGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/style.gtl", size: 3234, mode: os.FileMode(438), modTime: time.Unix(1515789047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesTestBodyGtl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\xdf\x6b\xdb\x30\x18\x7c\x0f\xe4\x7f\x10\x62\xd0\xa7\xd4\x90\xc7\xe2\x18\x4a\xd2\x0e\xb3\x41\xc3\x6c\xd8\xe3\xd0\xac\x2f\x96\x98\x23\x0b\xe9\x73\x60\x08\xfd\xef\x43\x8a\x7f\x28\xa3\xa5\x94\x3e\x45\x9f\x72\x77\xba\x3b\x4b\xce\x71\x38\x49\x05\x84\x22\x58\xdc\xfc\xee\xf9\x5f\xea\xfd\x7a\xe5\x9c\x61\xaa\x05\xf2\x25\x6c\x1f\x18\x32\xf2\xb0\x23\xf7\xf5\x38\x04\x44\xce\xe5\x85\x34\x1d\xb3\x76\x17\xb9\x67\x40\xd1\x73\x5a\xac\x57\x84\x10\x92\x33\x22\xf9\x8e\x3a\x37\x0b\x44\xf2\x0f\xb0\x43\x87\x71\x59\x1e\xbc\xa7\x45\x9e\xb1\x89\x21\xb6\xc5\xff\xf0\x5a\x62\x07\xde\xe7\x99\xd8\x8e\xa8\xd9\x57\xd3\x9f\xcf\xa0\x30\xd8\x5a\x38\x47\x66\x2c\xf0\xc0\xdc\xf7\x0a\x41\xe1\xfd\xfe\x0a\xf3\x3e\xd7\x41\xbd\x99\xc7\x2c\xcc\xa0\x78\x88\x12\x8f\x4f\xe2\xd8\x06\x14\x33\xb2\xa7\x31\xc3\x34\xfd\x7a\x37\xcc\x55\x29\x89\x9f\xe6\x9b\x32\x56\x1a\x1a\x79\x92\x0d\x43\xd9\xab\x31\x59\x82\xd0\x06\x26\x1b\x42\xb6\xa2\x93\xad\x40\x62\x53\x0e\x2d\xe6\x12\x5a\x79\x01\xf5\x53\x80\xaa\x05\xa8\xf7\xaa\xf8\x9a\x82\xa7\xd8\xd7\x4e\x6f\x85\xbc\x1f\x9b\xc9\x33\x6d\xe0\xd6\x9e\xd8\x16\x41\x93\x98\x98\xdd\x3e\xdc\x04\x70\x4e\x9e\xc8\xab\x1d\x3d\x33\xd9\x01\x4f\x0f\x4d\x82\xde\x2d\x41\x47\x59\x12\x2f\xe3\x29\x92\xee\x8a\xfa\xa9\xaa\xc9\xf3\x63\xf9\xfd\xe9\x70\x63\xfa\xad\x8f\xf1\x32\xa0\x1e\x70\xb1\x7f\x85\x43\x67\x81\xbc\xe5\xaf\xfa\x23\xb5\xfe\xb0\x41\xd5\xe3\xc6\x0c\x6a\x74\x58\x7d\x2b\x8f\xc7\x4f\x5b\xfc\xa0\x05\xcd\xac\x9d\x3b\x3a\x3e\x56\xd5\xe7\x0c\x2c\xef\xc1\x39\x84\xb3\xee\x18\x02\xa1\x52\x21\x18\xb0\x28\x55\xbb\x89\x37\xc5\xd2\xa4\xc6\x72\xf9\x37\x5e\x31\xfb\x9a\x44\xc3\x34\x0e\x06\xf8\x26\xbc\xaa\x85\xbb\x1f\xb7\xcb\x97\xf9\x1d\x66\x5c\x5e\x8a\xf5\x6a\xfa\x9d\x4d\x8d\x8b\x7f\x01\x00\x00\xff\xff\xd1\x08\x31\x21\xb1\x04\x00\x00")

func resourcesTestBodyGtlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesTestBodyGtl,
		"resources/test-body.gtl",
	)
}

func resourcesTestBodyGtl() (*asset, error) {
	bytes, err := resourcesTestBodyGtlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/test-body.gtl", size: 1201, mode: os.FileMode(438), modTime: time.Unix(1517867372, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"resources/capturedio.gtl": resourcesCapturedioGtl,
	"resources/contents.gtl": resourcesContentsGtl,
	"resources/htmltemplate.gtl": resourcesHtmltemplateGtl,
	"resources/index-content.gtl": resourcesIndexContentGtl,
	"resources/index.gtl": resourcesIndexGtl,
	"resources/interestinggivens.gtl": resourcesInterestinggivensGtl,
	"resources/javascript.gtl": resourcesJavascriptGtl,
	"resources/style.gtl": resourcesStyleGtl,
	"resources/test-body.gtl": resourcesTestBodyGtl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"resources": &bintree{nil, map[string]*bintree{
		"capturedio.gtl": &bintree{resourcesCapturedioGtl, map[string]*bintree{}},
		"contents.gtl": &bintree{resourcesContentsGtl, map[string]*bintree{}},
		"htmltemplate.gtl": &bintree{resourcesHtmltemplateGtl, map[string]*bintree{}},
		"index-content.gtl": &bintree{resourcesIndexContentGtl, map[string]*bintree{}},
		"index.gtl": &bintree{resourcesIndexGtl, map[string]*bintree{}},
		"interestinggivens.gtl": &bintree{resourcesInterestinggivensGtl, map[string]*bintree{}},
		"javascript.gtl": &bintree{resourcesJavascriptGtl, map[string]*bintree{}},
		"style.gtl": &bintree{resourcesStyleGtl, map[string]*bintree{}},
		"test-body.gtl": &bintree{resourcesTestBodyGtl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

