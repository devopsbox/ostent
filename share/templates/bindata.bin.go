// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

// +build bin

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\x6d\x73\xdb\x36\xf2\x7f\x9f\x4f\x81\xf2\x9f\xfe\x5f\x74\x4a\xa9\x7e\xc8\x5d\x2e\x95\x7c\xe3\x58\x6e\xca\x69\xec\x68\x6c\xd9\x77\xbd\x4e\xa7\x03\x91\x90\x84\x9a\x22\x59\x12\x92\xed\x6a\xf4\xdd\x6f\xf1\x40\x8a\xa4\x48\x49\x24\x61\xd7\xba\x24\x33\xb1\x48\x00\xbb\xd8\x05\x76\x7f\xbb\x20\x40\x76\xbe\xea\x7d\x3a\x1b\xfc\xdc\x3f\x47\x13\x36\x75\x5f\x9d\x74\xc4\x0f\x42\x70\x41\xb0\x03\x17\xfc\x72\x4a\x18\x46\xf6\x04\x87\x11\x61\x5d\x63\xc6\x46\xe6\x5b\x43\x55\x31\xca\x5c\x22\xae\xe1\x6e\xb1\x68\xf5\x30\xc3\xad\x1f\xfd\x88\x79\x78\x4a\x96\x4b\x04\x57\xc4\x63\x9d\xf6\xaa\x9d\x62\x37\x61\x2c\x30\xc9\x1f\x33\x3a\xef\x1a\xff\x36\x6f\x4e\xcd\x33\x7f\x1a\x60\x46\x87\x2e\x31\x90\xed\x7b\x9c\xac\x6b\x58\xe7\x5d\xe2\x8c\x89\x91\xa6\xe4\x9c\xbb\xc6\x9c\x92\xfb\xc0\x0f\x59\xaa\xf1\x3d\x75\xd8\xa4\xeb\x90\x39\xb5\x89\x29\x6e\xbe\x45\xd4\xa3\x8c\x62\xd7\x8c\x6c\xec\x92\xee\x81\x62\xb4\x58\xfc\x38\xb8\xf8\x88\x8c\xce\x57\xc6\x72\x69\x9a\xbf\xd0\x11\xb2\xce\x7f\x3d\xe9\xb8\xd4\xbb\x43\x21\x71\xbb\x46\x34\x01\xde\xf6\x8c\x21\x0a\xec\x0d\x34\x09\xc9\xa8\x6b\xb4\x47\x78\xce\xef\x5b\xf0\xa7\x88\xd3\x2f\xc4\x73\xe8\xe8\x57\xd3\x4c\x73\x92\x0c\x22\xfa\x27\x89\xba\xc6\xd1\xe1\xc3\xd1\x61\x9e\x5d\x14\x5f\x98\x47\x87\xad\xc0\x1b\x1b\x88\x3d\x06\xa0\x22\x9d\xe2\x31\x69\xf3\x02\xa5\xfe\x8a\x29\x0e\x02\x97\x98\xcc\x9f\xd9\x13\x33\xd3\xc1\xc1\xe1\x77\x0f\xf0\xbf\xbc\x0b\xa8\x6c\x55\x64\x79\x7c\xfc\x00\xff\x37\xb0\x3c\x3e\xae\xca\xf2\xcd\xe1\x03\xfc\xdf\xc0\xf2\xcd\x61\x55\x96\x6f\x41\xf1\xb7\x9b\x14\x7f\x5b\xa2\x78\xc4\x1e\x5d\x12\x4d\x08\x61\xf1\xc0\x33\xf2\xc0\xda\x76\x14\x25\xcc\xe0\xba\x4d\x3d\x87\x3c\xb4\x78\xa9\xe2\x10\xd9\x21\x0d\x58\x9a\xe4\x77\x3c\xc7\xb2\xd4\xc8\xbb\x0b\x8a\x42\x1b\x18\xfd\x1e\xb5\x43\x6e\xf4\x21\x81\xab\xc3\xd6\x41\xeb\xe0\x6d\x5c\xd0\x9a\x52\xaf\xf5\x3b\xf4\xe9\x80\x0b\x99\x53\x4c\x3d\xd9\x7e\xb1\x00\xeb\x6c\x0d\x4e\x3f\x7c\x38\xef\x0d\xa9\xb7\x5c\x42\x3b\x25\x8c\xa4\x58\x2c\x88\x1b\x81\xaf\x41\x0f\xed\x29\x75\xef\x54\xa5\xa8\xf0\x9c\xe5\xd2\x88\xdd\xb3\xd3\x96\xc2\x29\xf9\xdb\xca\xc3\x4f\x3a\x43\xdf\x79\x54\x85\x1e\x9e\x23\xdb\xc5\x11\x8c\x28\x5c\x0e\x71\x88\xe4\x8f\xe9\x90\x11\x9e\xb9\x2c\xbe\x8d\x18\xf8\xaa\x0d\x13\x11\x18\x28\xf4\xc1\xb5\x78\x73\x3a\x86\x42\x98\x93\xa4\x43\x87\x26\xdc\xb8\x97\x82\x4e\x24\x34\x47\xee\x8c\x3a\x71\x9b\x5c\x2b\xc5\x9d\x4b\x46\x42\xf0\x6c\xd7\x9c\x3a\xe6\x01\x0a\xb0\xe3\x50\x6f\x6c\xba\x64\x04\x30\x12\x4f\x41\x4c\x3f\x9c\x31\xe6\x7b\x39\x16\xcc\x1f\x8f\x5d\xc2\x59\xb8\x38\x88\x88\x13\xcf\xad\x6c\xac\x46\x59\x36\xe2\xc2\xc9\x56\x71\x31\x0e\xc7\x7c\xea\xfe\x4f\xf1\x4a\xaa\x53\xdd\x0a\x0b\x08\x70\xd2\x6d\x14\x9a\xbe\xe7\x3e\x66\x9b\x40\xa3\x81\x94\x63\x35\x3a\x30\x0b\x40\xb6\x81\x93\xb0\x56\xe8\x76\x8d\xd5\x8b\xa1\x6c\xcb\x41\xcc\x4c\x02\x46\xd4\xe9\x1a\x13\x18\xd9\xec\x3c\x0c\x43\xec\x39\x50\x46\xc1\x50\x84\xa7\x75\x8d\x29\x7e\x90\xf0\xfc\x0e\x1d\x1c\xda\x93\xc4\xcb\x60\x8a\x78\x94\x00\x2e\x2a\x7e\xa0\xf5\x80\x92\x9b\x82\xf5\x06\x9d\x36\xce\xc8\xd5\x06\xeb\xca\xd9\x1a\x17\x34\x3f\xb3\x2b\x33\x95\x05\x68\xc3\xd4\x9f\x74\x66\x6e\x4a\xcb\xb8\x29\xfc\xe4\x0d\xc4\xa5\x71\x3b\x6c\x33\x3a\x27\xeb\xc3\x8b\x95\xf2\x3c\x1c\x46\xef\xda\xed\xfb\xfb\xfb\x16\xe8\x12\xc2\xff\x96\xed\x4f\xdb\x32\x7e\x02\x46\xb8\x04\x47\x24\x6a\xbb\x98\x91\x88\xfd\xd3\x9e\x06\xdd\x58\xf7\xdb\xf3\xab\x6b\xeb\xd3\x65\x7e\x6c\x04\xff\x38\xfc\xe2\xf5\x79\x75\x69\xa9\xb0\xb1\x0b\x01\xa6\xa1\xd4\x35\x80\x52\x38\xa6\x1e\x77\x40\x34\xa2\x61\xc4\x44\xe9\xba\x4e\xb6\xef\x90\x1c\x2b\x5e\x04\x88\xbb\x36\x00\x42\xc4\x64\x9a\x7b\xd6\xf5\xe0\xca\x7a\xcf\x27\x91\x13\x68\x96\xb9\xa9\xa0\x34\xf8\x7f\x6f\x18\x05\xdf\x4b\xc7\xe1\x46\x04\x36\xbd\xd6\x2c\x65\x94\x56\x9f\x6b\xb2\xe6\x3f\x4a\x97\x17\xa8\xe1\x6c\x4d\xc3\xd9\x66\x0d\x6f\x02\x46\xa5\xd3\xed\x91\x96\x2e\xce\x6b\xe9\xe2\x8d\x5a\x7e\x3c\xd5\xa0\x61\x7b\xe6\x6e\xc2\xa5\xd4\x2d\xdc\x80\x12\x2a\x1e\x6f\x8f\xa0\x99\x36\xa1\x7f\x9f\x8f\xac\x5c\xc1\x29\x99\x42\x50\x93\xa9\x75\x82\x73\x3c\xb4\x1e\xc7\x31\xd6\x1f\x8d\x20\x5b\x31\x0f\x32\x38\xb7\x58\x30\x32\x0d\x38\xe4\x20\x03\xe2\x3f\xf4\x1d\xbd\x7b\x27\x2f\x7e\x83\xd1\x20\x2e\xf0\x35\x50\x8b\x8f\x4e\x31\xca\x3a\xa3\xe2\x6e\xdf\x34\xeb\xd6\x19\x15\xf6\x9a\x19\xc3\xed\x83\x62\x07\xb3\x27\x19\x14\xe0\xbb\x71\x50\xe8\xd3\x0c\x0a\xd5\x32\x28\x41\x54\x2c\xdc\xc1\x77\xcd\xa4\x0b\x22\x1d\xd2\xcd\xc7\x4f\x23\xdd\x7c\xbc\x5d\xba\xe4\x66\x4b\xe6\x1f\x13\xcc\x21\x77\xe6\xf0\x81\xba\x71\x7c\x5b\x2e\xbf\xcf\xa7\xe0\x2a\xf3\xe6\xc9\xb8\x58\x7e\x2f\x16\xed\x6f\x5e\x7d\xd3\x5e\x2e\x17\x0b\x29\x5f\x4a\xe0\x21\x86\x95\xb0\xc3\xd7\x2d\x3c\x0d\x86\xe5\x66\x9c\x3c\x00\xfb\x1f\xe1\x97\xaf\xb6\xbd\x33\x97\xda\x77\xdd\x05\x9b\xd0\xa8\x35\x81\xcc\xcb\x25\xa2\x64\x09\x32\x88\x21\x5b\x2e\x79\xf3\x33\x79\x09\x3d\xbf\x4a\x32\xc1\x55\x03\xe3\xc1\xc5\x43\xe2\x22\xf9\x13\x67\xfe\xc9\x82\xb7\xf5\x9e\x0b\x92\x82\xc4\x13\xae\xe0\x00\x86\x41\x66\x5f\x27\x6a\xdd\xb1\x9b\x2e\x21\x1d\x4f\x74\x2a\x93\x88\x82\x1a\xab\xb5\xab\x32\xca\x92\xc0\xef\xc1\x6c\x23\x83\xb7\xb9\xa7\x6c\x82\xf8\x1d\x18\x23\xb7\x2c\x60\xcf\xff\xc8\x85\xda\x62\x01\x39\xf1\x98\xa0\xd7\xf4\x5b\xf4\xda\xf6\x43\x82\xde\x75\x91\x0c\x35\x67\xfd\x9b\xd6\x47\x1a\xf1\x91\x64\x21\xc8\x7e\xca\x58\xf8\x13\x79\x44\x02\xac\x80\xdf\xf0\xd1\xbc\x34\x24\x51\xeb\x52\xe8\x7c\xd2\x61\x4e\xec\x0d\x22\x50\x8a\x01\x45\x9e\x7f\x1f\xe2\x20\x51\x2e\xa1\xe8\xb4\x99\x53\x4a\x95\xac\x6c\x73\x03\x37\x8b\x48\x40\x42\x1b\x32\x49\x99\xf1\xc9\x55\xd2\xaa\xb4\x1b\x77\x70\x13\x91\x50\x48\x95\xe9\x56\x96\xa6\x22\xe8\xc9\xd3\x4a\x71\xfd\x18\xad\x0b\x21\x0a\x9f\x4f\x86\x7f\x61\xca\xd6\x85\x90\xa5\xba\xa5\x30\xa9\x37\x27\x61\xb2\x7a\x2d\x90\xc6\x02\xd7\x59\x97\x46\x96\x16\x4a\xd3\x66\x61\xca\xf4\xb3\x3f\xbb\x38\x82\x44\x69\xde\xe8\x75\xb0\xb2\xee\x3e\x0e\xf1\x94\x4f\x83\xa8\x06\x8d\x84\xeb\x5e\xfa\xde\x9f\x24\xf4\xd1\xeb\x80\x9b\xbf\x87\x0c\x59\x2b\xfe\x9a\x13\x1f\x54\x33\x50\xb2\x2e\x33\x27\xd4\x71\x88\x67\xc4\xb6\x9f\x7a\x4a\x09\x22\xc7\x58\xcd\x26\x69\x14\x4f\xee\x78\xc5\x16\x5f\x11\xcd\xb8\xc1\x66\x96\xc7\x33\x8f\x66\xd6\x39\x27\x5f\xe7\xf2\xc1\x3a\xbd\x80\x45\x3e\x7d\x27\xdc\xe2\x9e\xbe\x17\x6e\x49\x4d\x7a\x69\xab\xb9\x13\xa5\xf1\x53\x29\xb6\x7a\x2c\xb5\x29\x80\xc7\xb0\x2b\x03\x38\x5b\x45\x54\x61\x40\x55\x00\xdc\x19\x0d\x1f\x61\x75\xfd\x5b\x45\x14\x77\x68\x74\xb7\xb2\xf1\xde\x0f\x82\x49\x31\x8a\xab\x1e\x14\x92\x3b\x34\xe4\xcf\x2c\x0c\xc9\xa2\xd5\xa3\xe1\xa5\x78\x84\xa1\xf0\x20\xc1\x6e\x59\x4b\xe6\x97\xea\x01\x47\x82\x18\xb9\x16\x31\xfd\x2e\x98\x12\x13\x9d\xce\x31\x75\x77\x25\x29\x9f\x84\x15\xe6\xc4\xda\x80\x0f\xf5\x65\xc9\x72\x29\xd7\x5b\x71\x97\x50\xe3\x54\x13\x72\xe0\x33\x9c\x16\xb2\x11\x3a\xc5\xb3\x5c\x05\xa2\x06\x78\xc8\xe1\xa9\x37\xf2\xe4\x0f\x43\x87\x1a\x60\xaa\x70\x40\xe5\xa3\x4f\xe2\xcd\xa6\x3c\x3b\x32\xd0\xad\x78\x50\x0d\x22\x8a\x8e\xef\xd0\x01\x32\x7a\x62\x87\x03\x7a\x13\xd2\xd7\xe0\x01\xc2\x5f\xf8\x33\x8f\xf1\xa7\xa2\xb5\x99\x1c\x21\x43\x18\x0f\xb0\x48\x4d\x5c\x4d\x66\xc7\xc8\xb8\x11\x0f\x69\x35\xf0\x7a\x83\x0c\x61\x30\x79\x66\x8d\x50\x26\x83\x0d\x3a\xa0\x86\x7a\xbe\xd3\x1c\x6b\x24\x97\x32\xb0\x91\xb5\x2f\x08\x6d\xac\x51\x48\x76\x26\xa9\x82\x36\xd6\xac\x14\x6e\x78\x55\x45\xbc\xb1\xc4\xb8\xe9\x03\x1c\x35\xd7\xcd\x10\xe7\x40\x4f\x62\x14\x2b\x2f\xbd\x28\x15\xc3\x25\xaa\x94\x04\xff\xb5\xd6\x0a\x3f\x36\x36\x47\xf9\x01\x16\xcd\x04\x68\xd4\xa0\xe3\xf8\x50\x83\x4c\x40\x81\xb6\x64\x23\xeb\xb7\xcd\x71\x80\xd6\x4b\x39\xe8\x68\x65\x3f\xd6\xa6\x84\x83\x66\x12\x0e\xe5\xff\x74\xd4\x2a\x76\xfd\xa4\x62\x47\x67\x81\xf6\x3d\xe2\x32\x6c\x79\x95\x49\x3e\xcd\x58\x15\x9a\x6a\x3d\x64\x99\x37\x72\x5e\x5a\x3f\x5b\xb0\xa4\xef\x5a\xe0\xbb\x47\xba\x17\x35\x16\x38\x5f\x38\xc2\xa5\xfe\xba\x9e\xac\xc7\x7b\x6b\xef\xad\xc1\x35\x02\xf4\x44\x11\xb1\x7d\x2f\xbd\x0f\x6b\x79\x9b\x53\xf8\x4e\xee\x91\xf6\xb0\xd3\x4e\x97\x9c\x04\x51\xed\x95\xc4\x76\xe1\x60\x4a\xff\x72\xe9\x18\x87\x12\xf4\xfe\xe7\xc1\xf9\x35\x9a\xfa\xce\xcc\xf5\xd1\xf1\x87\x06\x03\xf8\x3e\x27\xe2\xd7\xc7\x1f\x9e\x5c\xc6\xca\xe3\x58\x55\xc8\x26\xe8\x4a\xf5\x26\x59\x74\x44\xc2\xd0\x0f\x1b\xa2\xab\xe4\x51\x06\xaf\xb2\xf6\x0b\xbe\xd6\xc7\x57\x35\x47\xcd\x00\x56\xc7\x72\x4c\x13\xc0\x96\x60\x6b\x7a\xcf\xb1\xcc\xf7\x34\x60\x54\x19\x78\x3e\x53\xf7\xc5\xb0\xb8\x4b\xe7\x3a\xc0\xaf\x04\xf0\x74\x74\xdf\x0c\xd6\xd2\x38\xa4\x03\xd7\x02\x6c\xdf\x11\xd6\x10\xd8\x14\x93\x32\x64\x53\xd5\x5f\xa0\xad\x3e\xb4\xc5\xd3\xd4\x0c\xdb\x34\x2d\xfc\xbe\x60\x9b\x96\xd4\xea\xaf\x45\xb8\x72\x19\x5e\x04\xce\x65\x70\xa9\x39\xd0\x4d\xc9\xb4\xc6\x8e\x2a\x50\xad\x9c\xec\xe2\xfc\xa2\x18\xe0\xf8\x91\x18\x09\x6d\x77\x14\x0c\x4a\x90\xb5\x7e\xa2\x5c\xb2\x3c\xb6\xad\x6a\x76\xc4\x11\x4e\xf0\x83\xfe\x27\x5d\x9c\x6d\xe1\x63\x75\x55\x51\x4d\x40\xad\x0f\xd5\x81\x61\xcd\x2d\xbf\x0b\x32\x7d\xfe\x2d\xbf\x94\xdf\xf0\x79\xaa\xd2\x7e\xc3\x63\xa8\x67\x78\xfe\x14\x7b\x44\x73\xdf\x4a\x8e\x2a\x95\xce\x18\x3f\x74\xc3\xff\xa9\x39\xeb\xfb\x11\xe5\x67\x58\x57\xfb\xb4\x30\x47\x82\x0b\x12\x7f\x93\x03\x15\x72\x9e\x10\x52\x0c\xb6\x92\x8b\xd3\xde\xd4\x1b\x1b\x6a\x87\xfa\x84\xff\xe9\x60\x41\xca\x8f\x82\x0c\xd4\xb3\x7d\x4e\xb5\x4c\x9f\xfd\x50\xe4\x02\x1a\xd1\x90\x79\xe6\xd0\xf5\xed\xbb\xe4\x00\x8e\xfc\xe9\x0c\xcb\x25\x98\x1c\x1b\xfc\x0f\x1a\x8e\x4d\xea\x8d\x7c\x23\xd9\x22\x07\x6a\x68\xd1\x69\x0f\xd3\xbc\xe2\x53\xb4\x5c\x35\x75\x18\x88\x5f\x66\x0c\x7b\xbd\x0f\x59\x9d\xb7\x65\x55\x9e\xa8\xcc\xc1\x49\x59\x91\x54\x0a\xa2\xdc\x88\x8e\x57\x67\xfc\x56\xfe\xac\x0e\x3a\x1d\x96\x9d\xa0\x17\x01\xa3\xf0\x59\x28\xe2\x0f\x79\x5d\xfc\x58\x10\x27\xc4\x79\x9c\xb5\xb3\xd5\x52\x09\xa7\xe0\xf4\xe3\xfa\xf9\x45\x85\x97\x29\xe8\x5c\x13\x8d\x4f\xd1\x38\xf4\x67\x81\x8a\x6a\xf2\x26\x7e\x7f\x40\xde\xbc\xca\x08\x50\xe0\x08\xb9\x53\x50\xe8\x23\x89\xa2\xde\xca\x42\x1c\xc4\xbb\x11\xd6\xb0\xb2\xc7\x2d\x8c\xe4\x30\xa1\x0b\x3f\x24\xdb\x38\x6d\x56\x7b\xed\x54\x9a\xc1\xad\xd5\x77\xf9\xe0\x71\x7b\xcd\xcc\x59\x9a\x74\x05\x09\xcf\x64\x07\x57\x00\x22\x15\xcd\xc0\x6b\x9d\x0e\x23\xdf\x9d\xb1\xa2\xf3\xbe\x2f\xcb\x1e\x2e\x53\x88\xd1\xcc\x1e\xb6\x71\x7a\x01\x6a\x2b\x69\xff\x03\xf1\xb4\x99\xb4\xeb\xd6\xbb\xcd\x5a\x39\x10\xca\xc0\x23\x6e\xb6\x9c\xd2\x50\x27\x37\x93\x73\x96\x55\x83\x95\x33\xaa\x15\xab\xf8\x16\x5a\xfd\x50\x95\xa1\xde\x35\x52\x01\x91\xae\x40\x25\xfa\xdf\x14\xa7\x7a\x7c\x47\x78\x16\xe1\x31\x69\x18\xae\x44\x4f\x7b\x1e\xad\x7a\xa3\xfd\x09\x56\x20\xab\x9e\x58\x55\xc2\xe8\x73\x0d\x55\x60\xc7\xfb\x17\xa9\x84\xf3\xe9\x08\x54\x25\x8c\x5e\x80\xd2\x05\x71\xaa\x96\xb0\xd5\xc3\x94\x3e\x4b\x2d\x31\xd1\x6b\xe2\x12\x9b\x6d\x11\x3b\xe3\x5e\x47\xe9\x0e\xd6\xde\x26\x34\x03\xea\xba\xb9\x97\x5c\x5d\xba\xf1\x68\x88\x7a\xbf\x50\x1c\xe7\xca\x3b\x45\xc1\xb4\x94\x9d\x9e\x12\xbc\xe4\xf1\x17\xc9\x2b\xf7\x26\xd3\x66\x39\x0e\x35\xca\x01\xbc\xde\xf3\xcd\xc0\x12\x31\xd2\xaf\x53\x35\xcf\x4e\x72\x47\x74\x78\x8e\xf2\x6a\x1b\x45\xfa\x5c\x40\xfd\xa4\x86\xd6\x4b\x6a\xac\x46\x49\x8d\x55\x27\xa9\xb1\xf4\x25\x35\xd6\xb6\xa4\x26\x79\x42\x1d\x35\x4c\x6a\xac\xff\x81\xa4\xc6\xda\xa3\xa4\xc6\xd2\x95\xd4\x94\x30\xfa\x5c\x93\x1a\x6b\x1f\x93\x1a\x4b\x57\x52\x53\xc2\xe8\x05\x28\x5d\x90\xd4\xd4\x12\xf6\x4b\x52\x53\xb4\xed\xa9\x25\x99\x50\xbc\xfa\x72\x6b\x6a\xd7\xac\x26\x7b\xb6\x44\x9b\x20\xc0\xeb\x5c\x9c\x05\xa8\x21\xc7\x91\x46\x39\x8e\x9e\x33\xbb\xca\x6f\x84\xef\x90\x5e\xe5\x8e\x05\xed\x44\xa1\x27\x21\xe3\xaf\xb4\xd7\xc9\xc8\xe4\x3e\x56\xed\x94\x2c\x4b\xbe\x6b\x4e\xc6\xa9\x74\x25\x65\x52\x82\x4d\x59\x19\xb4\xf0\xc3\xc7\x86\x19\x59\x7a\xbf\x6f\x6f\x53\x32\x50\x62\x7f\x72\x32\x2e\xac\x9e\xa4\xac\x8c\xd3\xe7\x9a\x95\x71\x5b\xde\xbf\xb4\x4c\x7a\xa0\x8e\xbc\xac\x8c\xd3\x0b\x50\xbb\x20\x31\xab\x27\xed\x93\xee\x8a\x24\xc7\x26\xea\xc7\xab\x20\xaa\x15\xae\xfa\x51\x93\x68\x95\xa1\xde\x35\x58\x01\x91\xae\x58\x25\xfa\xdf\x14\xaa\xfa\xa1\x6f\x83\xc9\x37\x7e\x7e\x20\x3a\xda\xf3\x60\xd5\x8f\xf6\x27\x56\x81\xac\x7a\x42\x55\x09\xa3\xcf\x35\x52\x81\x1d\xef\x5f\xa0\x12\xce\xa7\x23\x4e\x95\x30\x7a\x01\x4a\x17\x84\xa9\x5a\xc2\x3e\x69\x94\x0a\xa2\xa6\x41\x6a\x3e\xae\x15\xa4\x6e\xc7\x4d\x82\x54\x86\x7a\xd7\x20\x05\x44\xba\x82\x94\xe8\x7f\x53\x90\xba\xc5\xe3\x10\xf3\xaf\x33\x36\x0a\x51\xa2\x9b\x3d\x0f\x51\xb7\xe3\xfd\x09\x51\x20\xab\x9e\x10\x55\xc2\xe8\x73\x0d\x51\x60\xc7\xfb\x17\xa2\x84\xf3\xe9\x08\x51\x25\x8c\x5e\x80\xd2\x05\x21\xaa\x96\xb0\x4f\x1a\xa2\xe6\xe3\x26\x21\x2a\xaa\xf8\x96\x41\x00\xab\x8b\x54\x24\xbb\x16\x7d\x17\xbf\x66\x10\xc4\x2f\x50\x05\x94\xbf\x64\xc0\x29\x5b\x7d\xab\xb7\xe1\xd3\x6d\xea\x94\x3e\xe2\x4e\x91\xb4\xde\xe5\x5c\x7f\x42\x71\x93\xa3\x78\x95\x38\x99\xa8\x55\x5f\x63\xab\xc0\xb0\x1f\x52\x3f\xa4\xec\xb1\x22\xd9\x25\xb5\x77\x7c\x05\x22\x21\xb9\xa6\x7f\x56\x25\xb9\x22\x11\x75\xc4\xfb\x10\x65\x64\xfc\x6d\x89\xf8\xe3\x15\xc9\x30\x0c\x68\xd9\xe7\x4d\xa4\xec\xd9\x97\xd9\x1a\xbd\x1c\x11\xe7\x4f\x95\xdf\x8d\x48\xad\x3e\x9f\xf3\x33\x43\xfd\x48\x7c\x66\x08\x0c\xaf\xf1\xf7\x78\x38\xab\xb7\xc8\xb8\xd1\xc4\xea\x1f\xc0\xea\xfa\xfc\xaa\xe6\x57\x8b\x38\x87\x43\xd0\xeb\x4a\x8b\x2c\x47\xc8\xb8\xb4\xb4\x70\x3a\x46\xc6\xad\x75\x35\xd0\xc2\xeb\x0d\x32\xae\xce\xaf\xb5\xb0\xfa\x1b\x32\x06\xd6\xc5\x79\xcc\x4b\x79\x51\x4d\x66\x7f\x47\xc6\xd9\xa7\x8b\x8b\xd3\xcb\x5e\xbc\xe5\xd4\xe0\x2d\x18\x05\xd8\xcd\x5f\x82\x49\xbd\x5e\x05\xac\xf2\x9f\x59\x14\xdf\x09\x2d\xfa\xa8\xa2\xdc\x8d\x3a\x11\x17\xf1\x87\xed\x2a\xf4\x3a\x1f\x8b\xcd\x2d\x43\x06\x8b\x3b\xf2\xc8\x3f\x6c\x2b\x8b\x62\xf8\xca\xe4\x74\x6f\x12\xe8\x52\x1f\x7c\x97\xeb\x06\xb1\x91\x58\x8a\x51\x3b\x48\x91\x44\x3d\x15\xd6\xa6\xd8\x9e\xac\xd0\x49\xf5\x72\x01\x85\x9c\xae\x38\xbc\xcd\x65\x23\x15\xe3\x66\x33\x11\xe4\x38\x9f\xd6\xcd\x4d\x2a\xca\x25\xd8\x9a\xaa\x2a\x82\x5e\x51\x5d\xfe\xe1\x29\x51\xdd\x0f\xfd\x39\x00\x7e\xb8\xa1\x09\x04\x64\xb6\x89\x85\x52\x6d\x44\x5d\xbe\x3e\x65\x93\x6d\x40\xbf\xc3\x58\x46\x90\x39\xd8\x93\x5d\x72\x08\x3a\x42\xe4\x0f\xb4\x3e\x93\x0e\x62\xe1\x8c\x94\xb8\x57\xce\x72\x50\x86\xe3\xa6\xf6\x89\xa3\xd4\x0d\x5e\x71\x66\x55\x39\x78\xa5\xd6\xa5\xda\x5e\x5c\xee\xe5\xde\xbc\x5b\x55\x71\xa3\x29\xad\x8c\x4d\xa6\xb4\x81\x30\x98\xd2\xda\x1e\x0d\x89\xcd\xc4\xd6\xa7\xa6\xd7\xf8\x12\x7b\x69\x00\x61\x09\xd6\x2e\xd3\xdf\x4e\xfe\x08\x25\x5b\xbf\x9f\xac\xc0\x8b\xb7\xad\xf6\xd9\xe6\x1c\xca\xf3\xd9\x9f\xa4\x11\x53\x7d\x88\x0b\x78\x9f\xba\x74\xec\x9d\xc9\x72\x89\x6a\x85\x32\x66\x3e\xe4\xcc\xcb\x15\x49\xbc\x00\xc8\x4a\x99\x03\x68\xfe\xed\x67\x1c\x12\x96\xa6\x11\x83\x98\xfe\xa0\x2d\x56\xe3\x3a\x49\x14\xfc\x6f\x00\x00\x00\xff\xff\x18\xf3\xd2\x7d\xf3\x6a\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 27379, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"index.html": indexHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{
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

