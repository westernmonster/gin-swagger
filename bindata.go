// Code generated by go-bindata.
// sources:
// templates/api.gotmpl
// templates/config.gotmpl
// templates/parameter.gotmpl
// DO NOT EDIT!

package main

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

var _templatesApiGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x39\x4b\x73\xdb\x36\xb7\x6b\xf1\x57\x9c\x72\x92\x94\xf4\xc8\x54\xdb\xa5\xe6\x7a\x91\x3a\x2f\xdf\x26\x69\xc6\x52\x6e\x97\x2d\x4c\x1e\x51\xb8\xa6\x00\x16\x00\x2d\x2b\x1a\xfe\xf7\x6f\xce\x01\x40\x51\xb2\x92\xb4\x9f\x16\x12\x09\x9c\xf7\x1b\x50\x2b\xca\x7b\x51\x23\xec\xf7\x50\xbc\xfc\x74\xf3\x29\xbc\xf6\x7d\x92\xc8\x4d\xab\x8d\x83\x2c\x99\xa4\xa5\x56\x0e\x1f\x5d\x9a\x4c\x52\x85\x6e\xb6\x76\xae\xa5\x67\x6d\xfd\xf7\xcc\xca\x5a\x89\x86\x5e\xec\xce\x96\xa2\xe1\x47\x27\x37\xc8\x4b\xce\x48\x55\xdb\x34\x49\x26\x69\xad\x1b\xa1\xea\x42\x9b\x7a\xf6\x38\xd3\xa2\x73\xeb\x5f\xfc\xba\x74\xeb\xee\xae\x28\xf5\x66\x56\x4b\x75\x49\xfc\x8c\xbc\x9b\xb5\xad\xd1\xab\xf4\xe9\x7e\xad\x95\x2c\xe9\x29\x4d\x26\xb5\x54\x9e\x10\x8c\xa1\xbe\x88\x46\xa8\x4a\x33\x74\xe4\x73\x44\x66\x23\xef\xef\xb1\xd1\xb6\x14\x86\x81\xec\x56\xd4\x35\x9a\xd9\x46\x56\x55\x83\x5b\x61\xf0\x1f\x22\x88\x56\xa6\xc9\xa4\xd1\xf5\x11\x7f\x2b\x4d\xd7\x5a\x54\xb3\x46\xd7\xa6\xb3\xff\x90\x96\x33\xa2\x94\xaa\x4e\x93\x89\x6e\x51\x85\xb7\x23\xba\xa3\xf5\xf1\xf3\x65\xad\xd3\x64\xb2\xdf\x1b\xa1\x6a\x84\xe2\x15\xae\x44\xd7\xb8\x1b\xf6\xa0\xed\xfb\xfd\xbe\x35\x52\xb9\x15\xa4\xcf\xff\x4e\xa1\xe8\x7b\x82\x45\x55\xf9\x07\x8f\xf4\xec\x1e\x77\x53\x78\xf6\x20\x9a\x0e\x61\x7e\x05\xc5\x08\x9b\xf6\xfa\x9e\x82\x64\x4c\xc7\xc3\x8e\x89\xe5\x49\x32\x9b\xc1\xad\xee\x1c\x5a\xa8\x70\x25\x15\x5a\x10\x4d\x03\x6e\x8d\x60\xfc\xb2\x5e\xf1\xdb\x02\xcd\x03\x1a\xb0\x68\x1e\x64\x89\x45\xe2\x76\x2d\x46\x4c\xeb\x4c\x57\x3a\xd8\x27\x93\x8b\x5a\xaa\xe2\xb5\xaa\xa5\xc2\x91\x7a\xbf\xb7\x68\x84\x93\x5a\xb1\x70\xd0\x0a\x8a\x39\xf9\x05\xa1\xf8\x28\x36\x14\xbc\x23\x0a\x9e\x04\x13\x36\x6f\x8d\xee\xda\x64\x32\xd9\xef\x41\xae\xa0\x78\xd9\xb9\xb5\x36\xf2\x0b\x56\xd0\xf7\xf4\x02\x04\xfa\x4e\xa8\xaa\x41\xf3\xa6\x53\x25\x0c\x46\xea\x93\xf8\xd8\xb3\x8a\xa5\x56\x2b\x59\x77\x06\xff\xc0\xa6\xf9\x4d\xe9\xad\x02\x54\xe2\xae\x21\x75\x55\x75\xd8\xb6\x30\x2b\xb6\xd8\x34\x97\xf7\x01\xa6\x6a\xb5\x54\xce\x16\xc9\x8a\xe8\x67\x06\x2e\xbc\xce\xf9\x19\x92\xd9\x1a\x45\xe3\xd6\x2c\x09\x81\x67\x39\xdc\x69\xdd\xe4\xa4\xd6\x76\xe0\x3b\xbf\x02\x53\xb0\x66\x59\x3a\x66\x96\xe6\xc9\x84\xf4\x1f\x20\x8b\xb7\xaf\x97\x59\x3a\xb3\xe5\x1a\x37\xe2\xb2\x92\xb6\xd4\x0f\x68\x76\xe9\xd4\x13\x2f\xdd\x23\xb0\xad\xae\x7d\xb2\x33\x9b\xc9\x64\x80\x23\x46\x23\xb3\x4e\x26\x0b\x26\xf4\xf9\xf6\x3d\x80\x4f\x71\xf8\xeb\xff\xad\x56\xf3\xd4\x73\xf8\xb3\x33\x4d\xfa\xd7\x08\x72\x49\x1e\x3e\x0b\x49\xbe\x0f\xa0\x9f\x6f\x98\x20\x7d\x8e\x41\x3b\x79\x20\xd8\x9f\x08\x30\x07\x48\x67\x21\x87\x0a\x02\x4f\xa7\x27\x7c\xe7\x90\x86\xfd\xcb\x5f\x8a\x9f\xfc\x76\x4f\x5f\xa5\x7b\x2c\xfe\x77\xf1\xfb\xc7\x8c\x6a\x5a\xb1\x70\xc2\x75\xf6\xf7\xdf\xa6\xf0\x62\xd0\x3b\x4f\x26\x93\x3e\x3f\x63\x48\xef\x9d\x74\x0a\xfe\x21\xc4\xcd\xc8\x69\x79\x4e\x81\x93\x4c\x4c\xb4\xfc\x91\x84\xdf\xb2\x3a\x49\xb5\x60\xf5\x4f\xe5\xf2\x46\xc9\x16\x9e\x14\x49\xce\x5c\xf2\x10\x97\x47\xa2\x80\xb4\x9c\x6a\x7e\x11\xde\x2d\x97\x9f\x60\x1d\xb6\x3a\x8b\x15\xac\xb4\x61\x80\x71\xd8\x04\xb5\x88\x18\x27\x2c\x65\x8a\x0f\xed\x2a\x04\xed\xd7\xd4\x3d\x8e\xd1\x27\xa9\x94\x4c\x3c\x28\x07\xd2\x01\x2b\x99\x18\x74\x9d\x51\xdf\x0c\xc2\xa0\xc1\x69\x04\xbe\xf3\xcb\xc4\x30\x86\x49\x70\xca\x38\x4c\x3c\xd4\x3c\xf0\xdc\x65\x79\xf0\x3e\x7d\xcb\x15\xfc\xe0\xd7\x8b\x40\xcc\xe3\x9c\x0b\x8b\x85\x2f\x55\x9f\x95\x78\x10\xb2\x21\x8b\x4c\xe1\x85\x47\xce\x99\x22\x60\x63\xf1\x1b\x04\x38\xae\xc6\x08\x14\x1e\xde\x6f\x81\x76\xf4\x98\x54\x0e\xcd\x4a\x94\x08\x6e\x2d\x1c\x6c\x3a\xeb\xe0\x0e\x41\x6e\xda\x06\x37\xa8\x1c\x56\x20\x15\x68\x53\xa1\x01\xa7\xa1\x35\xfa\x41\x56\x48\x84\xee\x3a\x4b\x45\xd7\x42\xa3\x6b\x59\x0e\x1e\x3e\x5b\x6d\x07\xa6\x03\xb7\x7d\x12\xac\xb5\x0b\x7e\xfc\x37\x45\xf7\x89\xf3\x42\x99\xfd\x24\x8c\xd8\x58\xe8\xfb\x29\xb4\xfe\xf1\x62\xbf\x2f\xc2\xa4\xd1\xf7\xc5\x59\x62\x1e\x69\xbf\xa7\x92\x09\x7d\x9f\xc3\x85\x68\x65\x71\x8b\xb6\xd5\xca\xe2\xb8\x1a\x73\x4c\xd6\x52\xc9\x2f\xf8\x49\xb8\x75\xd6\x0a\xb7\x0e\x69\x92\xc7\x1a\xb2\x1f\xa2\x2c\x0c\x23\xc5\x2d\xb6\x8d\x28\x31\x3b\x7d\x27\xec\x29\xa4\xfb\x74\x0a\xe9\x3c\x9d\xc2\xe5\xcf\xf9\x14\xd2\x9e\x5e\xfd\x5b\x70\x98\x54\xd2\x49\x16\x39\xf4\xac\xc3\x82\x3d\x74\xbb\x10\xae\x9d\xc1\xaf\x7a\x82\xc5\x3f\xa5\x96\xf9\x8c\xe3\x9e\x44\x6e\x98\x82\xd3\xf7\xa8\xa8\x34\x7a\x81\xa7\x40\x8d\x1f\x0d\x8c\x66\x80\x62\xc9\x4b\x79\x6c\x29\xa4\x35\x72\xeb\xa4\xc4\x21\xbf\x7c\xc4\x6d\x96\xc7\xc5\xe2\xb3\xc5\x8c\x7b\x23\xfa\x42\x97\x51\x25\x09\x4d\x7a\x7e\x05\x2f\x3c\x95\xbd\xef\xbe\x73\xf0\x58\x94\x36\x67\x43\xc2\x23\x9e\x77\xe6\xb8\xfd\xc2\x55\x18\x04\x62\xcf\x4a\x07\xae\xdf\x47\x66\x99\x0f\x13\x5a\xf1\x9e\x87\xab\xf7\x9a\x6a\x21\x8b\x2f\x57\xd1\x30\x3f\x5c\x81\x92\x0d\x67\xe3\xbf\xa2\x1e\x8d\x79\xa3\xa4\x5b\xb4\x42\x65\x9e\x1e\xc5\xc4\x1e\xac\x12\xf7\x63\xec\xd4\x97\xf8\x30\x50\x50\xe7\xcf\x14\x42\xf1\x01\xdd\x5a\x57\x90\xbe\x7d\xbd\x4c\x73\x28\xde\x09\xfb\xab\xae\x76\x43\x1e\xfc\xb7\xda\x72\x5a\x29\x47\xfd\xcc\x66\xc3\xd8\x26\x55\x85\x8f\x53\x78\xe6\x76\x2d\x8f\x6d\xd7\x5a\xd9\x6e\x83\xf6\x03\x56\x92\x7b\x1f\xe7\xab\x5c\x05\xc8\xbe\x9f\x86\xe4\x49\xf7\x7b\x42\xe2\x07\x5e\xc8\xf3\x64\x48\xb8\xb3\x33\xd2\xd0\x0a\x38\x30\xf7\x7e\x92\x7a\xc6\xea\xb0\xf8\xf3\x2b\xc8\x4e\x54\xca\xe9\x28\xe1\xe1\x2c\x96\x9d\x91\x6e\xf7\x8a\xe6\x42\xc9\x61\xc3\x02\x2f\xce\xac\x47\xa4\x10\x68\x11\x04\xb8\xf4\x84\x45\xff\xf2\xac\xc2\x15\x91\x61\xe5\xce\x33\x89\xb6\x1d\x26\x3f\xfc\x9b\xd1\x06\xb2\x3c\x27\xe0\x27\xa3\x5b\x5b\xf0\x94\x92\x86\x23\x83\x47\x62\x05\x97\x31\x01\xe7\x57\x43\x32\x26\xdc\x40\x8e\xb7\xaf\xae\x20\x4d\x7d\x13\x38\xd9\x80\xe3\x29\xfc\xeb\x12\x04\x0c\xe6\xdd\x1f\xc5\xef\xc8\xd6\x7d\xcf\xbe\x01\xce\x6c\x2f\x2e\x2f\x64\xc4\x79\x14\x34\x8b\x52\xb7\x68\x79\xe7\x49\xc8\x58\xda\xf3\x3e\x60\xa8\x6f\x04\x0a\x83\x8e\x42\x85\x7b\x68\x60\xfb\x3a\x4c\xb5\xbe\xf3\x45\xf1\xe7\xc7\x86\xf1\x4d\x97\xbe\x73\xef\x07\x1f\x67\xe3\xe7\xd1\xe2\x61\xde\x1e\x46\xf0\x50\xc0\xbd\x31\x46\x6d\x13\xcd\x70\xd2\xf8\x66\xa7\xa3\xc5\x61\x78\x08\x15\x32\x54\xca\x64\xe2\x87\x6f\xb8\xb8\xe6\xdf\x64\x62\x3d\xc2\x85\xef\xde\xfc\xe2\x17\xa9\x5f\x86\xbe\x79\x18\x68\x7c\xa7\x0c\xdb\xa1\x81\xbe\x51\xe3\x79\x28\x99\x90\xa9\x5e\x49\xcb\x83\x54\x58\x5a\x4a\xd7\xc4\x89\x38\x99\xfc\x1f\x1a\x2b\x75\xec\x50\x41\xc1\x8f\xb8\x0d\xb2\x8f\xfb\x8b\x00\x85\xdb\xf3\x9d\x64\x40\xc8\xec\x43\x19\x25\x9d\xc2\xb1\x7e\x39\x5c\x04\xe4\x3d\x97\xcc\x1f\xfc\x76\xf1\x0a\xef\x3a\x6e\x96\x74\xa8\x2e\x16\xe8\x3e\xe8\x2a\xf6\x88\x06\x85\x45\x7a\x0f\x43\x6d\xb0\x10\x35\x0a\x4f\x8b\xd0\xbc\x39\xe7\x4f\xfb\x19\x39\x3f\x72\x79\x39\xb2\x04\x47\x45\x58\x3f\x8a\x94\xb8\xe6\x8b\x2f\x45\x0d\x7d\x05\x5d\xe7\x60\x1f\x4a\x7a\xf7\x50\xf3\xa0\x1e\xad\xb0\x4d\xe7\x5c\xac\x8b\x1b\xb5\xd2\x85\x37\x72\xdf\xf3\xcc\x1f\x6c\x3c\xde\x8f\x66\x0f\x10\x63\x37\x45\xba\xa7\x12\x93\xfa\xb3\x59\x28\x85\xc0\x77\x14\x40\x81\x72\x38\xde\xd1\x74\x56\xb1\x31\x37\xba\x42\xb6\xf1\x13\x13\x33\x5e\x71\x8b\xb5\xb4\x8e\xdc\xc5\x46\xf4\xe5\xdf\x86\x23\xef\x94\x5a\x58\x3e\x30\xb4\xe8\xc0\x5f\x27\xd0\x4f\xed\x67\xbf\x25\x3e\xba\x37\xda\x6c\x84\x73\x68\x60\x2b\xdd\x1a\x94\x86\x52\x37\xda\x58\xbe\x96\x20\x3f\x0e\x00\xd9\x0b\x5a\x39\xc2\xd9\x07\xcd\xae\x19\x65\x0e\xce\x74\xd8\xe7\x83\x87\x8b\xe0\xe8\x2b\x78\x31\x4a\x06\x52\xe0\x65\x55\x99\x39\xc4\x4f\x34\x55\x55\x19\xb4\x96\x4c\x19\xa6\xff\x00\x72\x56\x3f\x0a\x19\x14\xd5\x52\x6e\x50\x77\x6e\x0e\xf0\xf3\x4f\x70\x01\x4e\x6e\x90\x2a\xa3\x56\x1c\x1f\x7f\x18\xe9\x70\x00\x39\x03\xd1\x1f\x0b\x7b\x94\x83\x57\x14\x2a\x61\xb6\xdf\x25\x47\xd1\x3e\x9c\xb3\x87\xb4\xdc\x87\x10\x3b\xc8\x79\xe6\x50\x1e\x00\xa4\x0d\x44\x0f\xfe\x19\x80\xc3\x51\xe3\xcb\x10\x10\xc9\x09\xd9\xf1\x01\xf2\xcb\x93\x13\xe4\x13\x0e\xe4\x8d\x38\xbf\xfa\x5a\x14\x86\xd0\x08\x01\x7e\xd7\xb2\xf3\xa8\xc3\xdd\x69\xb7\xe6\x72\x18\xbc\x47\xa3\x49\x7c\xa5\x12\x66\x90\xef\x76\xe2\x99\x28\x5e\x47\xd8\x58\x18\xf2\x03\xed\x50\xc1\xc6\x33\x74\x11\x0b\xdf\x8b\x17\x60\x9f\xd8\x3c\x8b\x33\xf2\x75\xb4\xc7\x70\xad\x23\x88\xe7\x70\xc0\x51\xa2\x39\xd8\x2c\xc0\x6c\x78\x74\xc2\x33\x02\x9d\x50\xcb\xf8\x68\x38\xf2\x50\x5c\x3e\xbd\xa1\x09\x84\x47\x57\x32\xa3\xcb\xa8\xaf\x4d\xe5\x23\xe9\xe9\xfc\xa0\x57\x11\x41\xaa\xb2\xe9\x2a\x52\x06\x9d\xa3\xad\xae\x05\xee\xc3\x72\x05\xd2\xd1\xf9\xed\xf8\xb0\x3c\xd6\xe0\x89\xa8\xb1\xfc\xda\xe2\xa8\x3b\xec\xbf\x7e\xea\x7a\x32\x92\xd9\x18\x52\xe7\xa7\x49\x9a\x20\xbf\x03\x42\xe4\xf2\xe1\x0a\x2f\xfc\x1c\x2a\x4e\xd7\xf2\x05\xde\x10\x38\xde\x0c\x62\x45\xd5\x86\x2c\x47\x92\xa3\x72\xb2\x14\x0e\xe1\x30\x77\xc0\x5a\x58\xb8\x43\x54\x4c\xe7\xd0\x0f\xaa\xe2\x2b\xca\x7d\x47\xca\xfd\x3e\x4c\xd5\x7d\x9f\x8d\x0e\x7b\x27\x17\x9b\xb4\x46\xc3\x09\x7c\xff\x80\x19\x87\x96\x6c\x08\xe0\xf3\x70\x3c\x14\x47\x9b\xcc\x66\x70\xdb\x29\x30\x9d\x1a\x0f\x1b\x05\xdc\x38\xd8\xca\xa6\x81\x86\x6a\xb9\x02\xad\x00\xa5\x5b\xa3\xf1\x77\x2e\xda\xff\x2e\xa0\xc2\x16\x55\xc5\x01\xa5\x08\xff\x10\xa7\xc4\xd8\x62\x45\x05\x7d\xe8\xe0\x67\x22\xe8\xb6\x53\x59\x0e\x68\x8c\xe6\xd6\x7d\x54\x73\x8e\x3d\x74\x3e\x35\x7c\x43\xa0\xbe\xb7\xca\x52\x1e\x0e\x54\x0d\x3f\x3e\xb7\x70\x09\xcf\xed\x8f\x24\x95\xf0\xf5\x1b\x9e\xdb\x74\x0a\xd6\x77\x4f\x7a\x08\x6d\x72\x1a\xf2\x1d\x0d\x57\xfa\x3c\x44\x89\x9f\x50\x38\x29\x48\x85\x58\x1e\xb6\x6b\x54\x3e\xeb\xd9\xef\x87\xba\x71\xc5\x85\x8a\x83\x3f\x8a\x59\xdc\x28\x9e\xdd\x91\x4d\xc6\xc7\xb6\x58\x6d\x02\xc3\xf7\x6c\xdc\x97\xaa\x62\x6b\x64\xfe\xe4\xf5\x6d\xa0\xe5\xfb\x45\x36\x30\x58\xbe\x5f\x5c\xa3\x71\x6f\xa4\xd7\xe8\xb0\xfa\x1b\xee\x68\x31\x56\x8e\xc5\xba\x73\x95\xde\x2a\xef\xd2\x9a\x26\x90\x55\xd7\x34\x3b\xb0\x71\xe3\xa4\x62\x9c\x75\x55\xa4\x72\xe2\xaf\x27\xb6\xea\xd4\xb1\xb5\xd6\x9d\xaf\x2b\x84\x7b\x6c\xb2\x95\x68\x2c\x3e\xd5\x78\x60\x14\xfe\xa4\x29\x7e\x15\xe5\x7d\x6d\x74\xa7\x2a\x3a\x12\x0f\x51\xfb\x87\x74\xeb\x85\xac\xe3\xdd\xe0\x49\x0c\x47\xc1\x78\x80\x58\xdc\xbc\x5d\xbe\xbe\xfd\xe0\x2f\x0b\x49\x18\xd1\x39\xbd\x11\x94\xe4\x4d\xb3\x4b\x86\xd9\xa7\x2a\x60\x79\xe8\x31\xe3\x0c\xa0\xba\x2a\x06\x3a\xfe\xef\x21\x6e\x42\x67\xcc\x49\xf4\x48\x90\x2d\xde\x0d\xe6\x9c\xcd\xe0\xc6\xfd\x68\xa1\xd5\xd6\x4a\x9a\xb2\x9c\x06\xdd\x52\xa9\x20\x09\x38\x5b\x40\xa8\x1d\xa8\x6e\x73\x47\xc2\x47\xcf\x90\x1b\xfc\xa1\x6f\xbb\x96\xe5\x9a\x65\x62\x79\x1f\xb1\xec\x1c\x82\x56\x08\x77\x3b\xfe\x39\x94\xb0\x2d\xde\x05\x15\x62\xd1\x3a\x10\xb4\x5d\x59\xa2\xb5\x2c\xf1\xf9\x84\x3c\xb6\x6b\x36\x60\x16\x45\x11\x0e\x00\xec\xff\x51\x18\x58\x59\x5f\xf3\x65\xe6\x46\xdc\x63\x56\xae\x85\x02\x6d\x8b\x05\x1b\x69\x0a\x3f\xe7\x0c\xa1\x44\x53\x7c\xd4\x4e\xae\x76\x19\xc3\x4f\x21\xfc\xb5\x56\x2c\x6e\xde\xbe\xfb\xfc\xe9\xe8\xfd\xe6\xe3\xf2\xe8\x9d\xac\x4e\xd9\x5e\xeb\x78\x08\xa1\x84\xfa\x9f\x4b\xa6\x44\x43\xce\x21\x6a\x28\x8d\xb8\x32\xa0\xe1\x69\xde\x16\x5c\x64\x38\x37\x69\x69\x74\x8f\x72\x58\xe1\x49\xf0\xb5\x31\xde\x0c\xd7\x8d\xb6\x61\x7a\x8a\xd1\x89\xc6\x24\xe1\x82\x33\x99\x50\x30\xfc\x39\x85\x95\xff\xaf\x82\x4b\xff\x60\x25\x42\x0a\x9c\x57\xcc\xf6\x0c\xdf\xb3\x44\xc3\x9a\x92\x4d\xd2\x93\x87\x1f\xe4\x66\x0e\x2b\x77\x55\xeb\xe4\x3f\x01\x00\x00\xff\xff\xde\x72\x43\xb2\xd4\x1c\x00\x00")

func templatesApiGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesApiGotmpl,
		"templates/api.gotmpl",
	)
}

func templatesApiGotmpl() (*asset, error) {
	bytes, err := templatesApiGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/api.gotmpl", size: 7380, mode: os.FileMode(420), modTime: time.Unix(1535473472, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfigGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xc1\x8e\xdb\x36\x10\x3d\x8b\x5f\x31\xd5\x61\x2d\x15\x91\xbc\xc8\x29\x08\xe0\xc3\x36\x9b\xa0\x8b\xf5\xc1\xa8\xdd\xe6\x4c\x53\x23\x89\x30\x45\x0a\xe4\xc8\x86\xb1\xf1\xbf\x17\xa4\x29\x47\xb6\xb7\x3d\x44\x0b\x2c\xcc\x19\xf2\xbd\xe1\x9b\x37\xec\xb9\xd8\xf1\x06\xe1\xed\x0d\xca\xa7\xd5\xcb\x2a\x2e\x4f\x27\xc6\x64\xd7\x1b\x4b\x90\xb1\x24\xad\x3b\x4a\x19\x4b\xd2\xc6\xf4\xbb\xa6\x94\x7a\xce\x15\x0a\x6a\x4d\xc7\xdd\x7c\x27\x75\xd3\x4b\x5d\xee\x3f\xa6\x2c\x31\x3d\x6a\xb2\x5c\x48\xdd\x40\xda\x48\x6a\x87\x6d\x29\x4c\x37\x9f\xc4\xa7\xbf\x8b\xc6\xa4\x2c\x67\x4c\x18\xed\x02\x53\x85\x35\x1f\x14\x3d\x55\x95\x45\xe7\x60\x01\xe9\xe7\x4f\x8f\x9f\x1e\xc3\xa6\xf9\x1c\xbe\x18\x5d\xcb\x06\x2a\xac\xa5\x46\x07\xd4\x22\x88\x73\xc8\xf4\x24\x8d\x76\x50\x1b\x1b\xc2\x4f\xab\x17\x70\x68\xf7\x68\x4b\x46\xc7\x1e\xc7\xa3\x8e\xec\x20\x08\xde\x58\x32\x72\xfc\xfc\x1c\x59\xa9\x1b\x96\x3c\xe3\x76\x68\x60\xfa\x6d\x8d\x51\x2c\x79\xd1\x0e\xc5\x60\xf1\xcf\xcd\x66\x35\x8d\x3f\x0d\xd4\x3e\x4b\xc7\xb7\x0a\xab\x69\x7c\xb3\x5c\x7f\x41\x4b\xdf\xa4\xc2\x1b\x86\xcd\x72\xfd\x8a\xc7\x49\xe2\x92\xf9\x8e\x4a\xbd\x6a\x73\xd0\x17\xc0\x88\x65\x76\xa8\xff\xfe\x6b\x79\x5f\xed\xc6\x72\x81\x76\x5a\xed\x44\xe0\xf2\x9c\x65\xa7\x20\xdf\x77\x49\xed\xf3\x59\xe1\x6f\x8a\x37\x0e\x84\x45\x4e\xe8\x80\x8f\xf2\x1c\x24\xb5\x10\x9b\x00\x3c\x2a\x34\x0b\x3d\x98\xb1\x7a\xd0\x02\x32\x01\xbf\x9f\x37\xe7\x77\x78\x59\x3e\xe6\xbc\xc0\xa3\x31\x7c\x2a\x4b\x2b\x2f\x6a\xfa\x01\xd2\xaf\xda\xdf\x0b\xc2\x1a\x94\x69\x1a\xef\x15\xae\x2b\xe8\x7b\x6b\x6a\xe8\x90\xac\x14\xae\x4c\xf3\xf2\x0f\x63\xd4\x3f\xdc\x66\x0f\xa2\x0c\x2d\xc9\x6f\x31\x63\x81\x1e\x75\xec\x26\x19\x50\xd2\x11\x6a\x30\xfa\x03\x60\xd9\x94\x10\xaa\x07\x63\xe1\xb1\x0c\x7f\x61\xed\xf1\x59\x92\xc4\xea\xb3\x6b\xdf\xe5\xe5\x3a\x88\x1b\xc9\xc7\xe8\x2d\xbd\x8c\x7e\x28\x5a\xa2\xde\x17\xb1\x46\xbb\x97\x02\xc1\x68\x75\x04\xef\x92\x9b\x5b\x4c\x0d\x74\x87\x46\xca\x15\x02\x2d\x15\xb5\x54\xe8\xd1\x56\x9c\x5a\x7f\x9f\xcd\x72\x0d\xde\x47\xe0\x13\x30\x38\xac\xe0\xd0\xa2\x0e\xf6\xf6\xda\x79\xb4\x75\xbc\xcf\x55\xdd\x13\xff\xbd\xcb\xb6\xc3\xe3\xbb\x64\xaf\x78\xfc\x15\xae\xe8\xe8\x3b\xaa\xea\xec\xe4\xe2\x80\x4a\x15\x3b\xef\x6d\x4f\x18\xfd\x0d\x7c\x20\xd3\x71\x92\x02\xe6\xe5\xcf\x1d\x60\xd1\x99\xc1\x0a\x74\x91\x6c\xa2\xe2\xdd\x88\xfc\x27\x23\x1f\xa8\xbd\xe1\x6a\xc3\x0b\xc1\x95\xfa\x3f\x86\xe9\x40\xdf\x2b\xe7\x27\xb1\x18\xac\x3a\x77\x9c\xe0\x32\x9a\x41\x2e\x32\xb0\xe7\x4a\x56\x9c\x10\x8c\x27\xfc\x08\xe1\x84\x7b\x57\xb5\x78\x34\x67\x2c\xb1\x48\x83\xd5\x20\xe2\xac\xae\xb8\x75\x08\xbd\xff\xef\xe2\x23\x37\x58\xee\x1f\x39\xa8\xad\xe9\x40\x98\xae\xe3\xba\x52\x52\x23\xd4\x7e\xf8\xca\xfb\xf9\x0c\x18\x59\x0e\x68\xad\xb1\xd3\x91\x8c\x09\xc6\x12\x59\xc3\x6f\xd7\xc6\x84\x87\x07\xc8\xae\xcc\x03\x8b\x05\xa4\x29\xfc\xf8\x01\xd3\x3e\x9f\xa3\xb9\x87\x1d\x6b\xaf\x3b\x2a\xbf\x7a\xae\x3a\x4b\x67\x45\x71\xe5\xe8\x59\x98\xf1\x18\x1d\x9d\x37\x83\x6e\x70\x04\x5b\x04\xd7\xa3\x90\xb5\x1c\xfd\x36\x2b\x8a\xab\xe9\x5a\xd4\x5c\x39\x9c\xa5\x39\x4b\x4e\x17\xa9\xb4\x54\xec\xe4\xb5\xda\xcb\xee\x33\xd4\xb4\x68\x0c\xfb\x37\x00\x00\xff\xff\x12\xaa\x4c\x52\xd0\x06\x00\x00")

func templatesConfigGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigGotmpl,
		"templates/config.gotmpl",
	)
}

func templatesConfigGotmpl() (*asset, error) {
	bytes, err := templatesConfigGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config.gotmpl", size: 1744, mode: os.FileMode(420), modTime: time.Unix(1535473472, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesParameterGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5f\x93\xdb\x36\x92\x7f\x16\x3f\x45\x2f\x2f\x71\x91\x3e\x0d\x95\xdd\x4b\xe5\x41\x97\xd9\xaa\xd8\x1e\xaf\xe7\x92\xd8\x3e\x8f\xed\x7d\x48\xa5\xb2\x18\x12\x92\x10\x93\x00\x07\x80\x46\xa3\x53\xf1\xbb\x5f\xe1\x0f\x41\xf0\xef\x48\xe3\xd9\xec\x5e\x6e\xe7\xc5\x22\xd1\x68\x74\xff\xd0\x68\x74\x37\x40\x1f\x0e\x90\xe1\x15\xa1\x18\x42\x91\x93\x14\x97\x88\xa3\xe2\x16\xe5\x24\x43\x92\xf1\xb0\xaa\x82\xc3\x01\xc8\x0a\x18\x87\xe4\x47\x42\x2f\x25\x2e\x04\x24\x3f\xa2\x3b\xf3\xcb\xb4\xa7\xa8\xc0\x39\xf9\x1f\x0c\xc9\x6b\x54\x60\xa8\xaa\x2b\xf5\xb0\x3c\x07\x42\xe5\x37\x5f\x47\x39\xa6\x91\xe1\x82\x68\x06\x11\x65\x12\x92\x4b\xf1\x1d\xe7\x68\x1f\xdb\xc7\x57\x48\xbc\x20\x22\xe5\xa4\x20\x54\x0d\x1c\x3b\xb2\x4b\x2a\x31\x5f\xa1\x14\x37\xaf\xae\x24\xc7\xa8\x88\xd5\xcf\xd7\xdb\x3c\x47\xd7\xb9\x1a\xf3\xe9\xe1\x00\x98\x66\x50\x55\x87\x03\x24\x1f\x51\xbe\xc5\x17\x77\x25\xc7\x42\x10\x46\xa1\xaa\xe2\x38\x70\x14\x56\xa9\x46\xa3\xaa\x0a\xc8\x0a\x30\xe7\x4a\x6a\xab\x3e\x76\xcd\x4a\xfa\xe4\x2d\x92\x1b\xa8\xaa\x39\x1c\x0e\x50\x72\x42\xe5\x0a\xc2\x2f\x6f\x42\x48\x7e\x60\x29\x92\x66\x0c\xdd\x38\x88\x86\x6e\xf1\xc7\x8b\xff\x53\x0f\xf7\x87\x73\xa0\x24\x87\x43\x00\xc0\xb1\xdc\x72\xaa\xde\x06\xd5\x80\xa8\x1e\xe4\x43\xa2\xda\xe6\x47\x12\xd5\xf1\x3b\x5d\xd0\x0f\x94\xdc\x6c\xf1\x94\xac\x1e\xc5\x69\xe2\xfe\xa3\x2d\xe8\x44\x24\x2e\xe8\xb6\x18\x81\x40\x35\xfd\x9f\xd2\xdd\xd8\xaf\xd5\xe8\x14\x20\x9a\x5f\xb5\x9f\x49\xb7\x42\xb2\xe2\xa3\x81\x82\x30\xfa\x56\xc9\x2d\xc9\x2d\x0e\x0d\xa1\x5d\x98\x3f\x60\xba\x96\x9b\xd1\x85\x69\x9a\xad\x5f\xa9\x61\xf4\x10\x3d\x1c\x70\x2e\x70\x55\x85\xe1\xe1\x80\x69\x76\x24\xc4\x6d\x50\x22\x1f\x95\x2e\x28\x7a\xb4\x5e\x97\xd8\xc7\xd1\xb4\x3e\xd7\xfa\xbe\x64\xbc\x40\x52\x62\x0e\x55\x95\x5c\x49\x4e\xe8\x3a\x6a\x88\xcd\x9a\x6b\x94\xbe\x17\x61\xad\x52\x0d\x16\xba\x9b\x04\xab\x6e\xfe\x7d\x81\xd5\x28\x7d\x12\x58\x6f\x35\x5f\x3a\x0c\x95\x6d\xfc\xfd\x00\xf5\xb7\xc3\xa1\xd1\xf8\x6f\xa7\x59\x15\xa1\xa4\xd8\x16\xa3\x0b\x50\x35\x1a\x69\xf0\x0d\x24\x57\x3b\xb4\x5e\x63\xfe\x7e\x5f\x62\x08\x09\x95\x78\x8d\xb9\x5a\xcf\x97\x54\x3a\x71\x1e\x1b\xd6\xa9\x71\x89\x19\x37\x17\x0a\xbe\x55\xce\x50\x23\xc6\x37\x5f\x47\x43\x18\x4f\xcf\x4a\x5c\xaf\x50\x83\x89\x7e\xba\xb8\x4b\xf3\xad\x20\xb7\xd8\xbd\x3e\x75\xd9\x4e\x00\x6c\x1a\xff\xdf\x01\x5c\x63\xd2\x01\xb8\x7e\x7d\x1a\xc0\xdb\x5c\x92\x32\xc7\x6f\x56\x23\x18\xbb\xf6\xc7\x03\x4e\x23\xf1\x39\x00\x78\x32\x9f\xa4\xac\xda\x9a\xa7\x63\x8d\xc7\xb4\x8c\xcf\x0b\x45\x4e\x81\xe5\xe4\x9d\xa1\x91\xfd\xdf\x6e\xc3\x1a\x97\xa3\xa1\xac\xff\x75\xd1\x4a\xc9\x59\x89\xb9\xdc\x77\x12\x23\x2f\xcc\xbb\x14\x2e\x82\x31\xd0\x4a\x5c\x94\x39\x92\xd3\xa1\x4e\x62\x68\xdb\x11\xe3\x90\x96\x43\x53\x6a\xda\xb5\xdd\xea\x79\x54\x9a\x87\x87\x83\x9b\xa9\xaa\x0a\xcd\x0b\xbb\x7a\x0d\xbd\x7e\x3b\x84\xb1\x03\x72\x0e\x2b\x4d\x29\xa6\xf1\x1a\x90\x5b\x1b\x41\x57\xfd\xa1\x8c\xd2\x57\xfc\xac\x13\x22\xd6\x98\x5f\x13\x9a\x95\x35\x54\xba\x7f\x38\x1a\x4d\x36\x63\xa8\x5e\xc6\x37\x05\xb7\x88\x2b\x4b\xb8\x45\x9c\xaa\xbc\x26\x79\xbe\x21\x79\x36\x10\xd5\xbe\xd3\x51\xed\x5f\x98\xf6\x6e\x55\x15\xac\x18\x87\x26\xdf\x35\xbd\x5e\x21\xd1\x4c\xa0\xa8\xdf\x3e\x67\xf4\x16\x73\x35\x43\xf6\xc5\xd0\xd4\x29\xe6\x97\x34\xc3\x77\x1f\x91\x7d\xb4\xfe\xf2\x97\xb6\xc5\xde\x2b\xe7\x47\x35\xfb\x1c\xd1\x35\x3e\x8a\xfc\xb9\x9e\xb1\xae\x22\xf5\x24\x29\xd4\x75\x6b\xad\xca\x28\x97\xe5\x39\x88\x1d\x5a\x27\x57\x65\x4e\xe4\xb3\xbd\x51\x2d\x3a\x4a\xe0\xbe\x17\xa9\x71\xcb\x73\x9c\x2a\x28\x0d\x37\x15\xdc\x18\x69\x86\xcc\xa6\x9e\x52\xd3\x19\xac\xe0\x67\x06\x46\xa7\x87\xb2\xc0\xee\xac\xb8\xc6\x7b\x45\x9d\xd7\xab\xcb\x43\xc4\xb1\x31\x7b\xea\xfd\xea\x2a\x1d\xec\x3a\xf5\xd6\x8c\xbf\x6a\x18\x17\xc9\x25\xd5\xeb\x40\x59\x5b\xd4\x8c\x36\x9a\x02\x9a\xe6\x96\xf7\x0d\x9b\x6e\xce\x6a\xc3\xe3\x6c\x48\x89\xd8\xc2\xaf\x81\xad\x6f\xbb\x0f\x80\xcf\x7a\x8e\xe4\x2d\xe2\x02\xff\x7e\x51\x3b\x1e\x19\x6b\x53\xf7\xc3\x60\xf8\x59\xd7\x66\x87\x32\x0f\xdd\x95\x31\xb6\x19\xb5\xd7\xc7\x31\x7e\xef\x1c\x50\x59\x62\x9a\x1d\x35\x51\xef\x8e\xc5\xca\xf7\xd1\x25\x4a\x3f\x21\xe3\xb0\x92\xb7\xf6\xb7\x52\x69\xb1\x80\xf7\x1b\x22\x60\x45\x72\x0c\x3b\x24\x60\x8d\x29\xe6\x48\xe2\x0c\xae\xf7\x20\x37\x58\xfb\x9c\x35\xe6\x20\x19\xcb\x13\x45\x7f\x91\x11\x49\xe8\x1a\xa4\xeb\x57\x90\xf5\x46\x42\xc9\xd9\x2d\x86\xd5\x56\x6a\x56\x1b\x4c\x61\xcf\xb6\xc0\xf1\x19\xdf\xd2\x16\xa7\x7a\x08\x48\x59\x51\x20\x9a\x05\x01\x29\x4a\xc6\x25\x44\x01\x40\x48\xb1\x5c\x6c\xa4\x2c\x43\x85\x77\xb8\x26\x72\xb3\xbd\x4e\x52\x56\x2c\xd6\x84\x9e\xad\x19\x25\xa9\xfa\x15\x76\x1b\xd9\x19\x2b\x31\x45\x25\x59\x18\x43\x9d\x20\xa8\x37\xee\x09\x12\xbe\xa5\x92\x14\x47\x50\x2c\x0a\x92\x65\x39\xde\x21\x3e\x45\xac\x34\xef\x36\x17\xe4\xd3\x27\x9c\x33\x91\x22\xae\x55\xb3\xe8\x2c\x50\x49\x8e\x25\x95\x1c\xa5\x84\x6a\xce\x6a\x24\xfb\xd8\xea\xea\xbd\xf7\x7f\x9f\xad\x59\x77\x90\x71\xca\x05\xbe\x93\x61\xa0\xa6\x43\x48\xbe\x2a\xe4\xa8\x96\xba\x35\xb4\xeb\xc4\x6c\x90\xc9\x0b\xbc\x42\xdb\x5c\x5e\xea\x19\x16\x66\xdf\x6d\x79\x88\x7a\x99\x78\x2b\xce\xf6\xfd\xe2\x13\xde\xcf\xe1\x8b\x5b\x65\xda\x6a\xf9\x26\x2d\x26\xaa\x55\xad\xeb\x0e\x3f\x4b\xde\xe1\x1a\x6b\x4b\x57\xa4\x48\xa4\xa8\x55\x64\xbd\xa0\x59\xc9\x08\x95\x80\xef\x70\xba\x95\x58\x68\x5b\x4d\x19\xc7\x90\xb3\x35\x49\x81\xad\xf4\x1b\x8e\xd5\xaa\xcf\x14\x1f\xce\xb6\x12\x2b\xd6\xba\x63\x12\xac\xb6\x34\x9d\x66\x1e\x6d\x10\xcd\x72\xcc\x41\x91\x46\xa9\xbc\x83\xa7\x6b\x42\xd5\xd6\x26\xf1\x9d\x74\xe9\x00\x47\xba\x4c\x3b\x87\xd2\xfc\x7c\x3a\xc8\xd4\xd0\x39\xe5\x62\x78\x8a\x4a\x92\xbc\xc3\xa2\x64\x54\xe0\x18\x14\xe7\x57\x66\xbc\x97\x5a\xb2\x60\x66\x7d\xb8\x16\xb4\x37\x7c\xac\x28\x66\xa2\x44\x54\xa1\xec\xcd\x7d\x72\x55\x22\xfa\x92\xb3\xc2\x12\x46\xf5\xfb\xfa\x39\x95\x77\x71\x1c\x04\xb3\xd9\x62\x01\x48\x4a\x94\x6e\x40\xa2\xb5\x00\xc9\x5a\x06\xa9\x58\x07\xb3\x19\x59\xe9\x5f\xcd\xee\x32\x9b\xcd\xf0\x9d\x4c\x5e\xbd\x7f\xff\xf6\x47\x2c\x37\x2c\x4b\xae\xb0\x8c\x14\xcd\x1c\x52\x79\x97\xbc\xc3\x37\x5b\x2c\x64\x62\x1a\x63\x9f\xfe\x03\xcf\x47\x88\x3f\xbc\xfb\xc1\x45\xd0\xaa\x4b\xa5\xe4\xeb\x22\xbc\x58\x34\x5e\xc8\x62\xbd\xe2\xac\x00\x6e\x98\x04\xb3\x99\x7d\xbb\x3c\x87\xd7\x78\x37\x31\x0d\x91\x1a\xc3\x6e\xba\xa6\x4f\xc2\x31\xca\xac\x34\x1a\x21\xa3\x7a\x6b\x5f\x9d\xe9\x3e\x6f\xae\x7f\x55\xdd\x30\xe7\x49\xf4\xd4\xee\xaf\xcf\x59\x51\x32\x41\x24\xbe\x50\xcf\x5a\xe7\x92\xb3\xeb\x1c\x17\x8a\x54\x4d\xf4\x5b\xf3\xa8\x99\xcc\xde\x13\x99\xe3\x25\x40\xf8\x81\x96\x9c\xa5\x58\x08\x9d\xc7\x5d\x50\x49\xe4\x3e\x09\xe7\x9a\xe8\x4a\x22\xb9\x15\x4b\x50\x86\x68\x86\x4d\x9e\xb3\x0c\x47\x71\x6c\xda\x5f\x60\x89\x48\xbe\x04\xdb\xa6\x47\x8e\x4c\x9b\x46\xef\xb8\xe9\x1d\x9c\x5f\x37\x61\x46\x06\x35\xac\x37\x6f\x5b\x42\xe5\x1f\xbf\x89\xac\x82\x89\xa1\xd1\xb3\x66\x07\x56\x13\xfb\x57\x4e\x24\xe6\xc9\x2b\x8c\x32\xcc\xa3\x58\x77\x0f\xb5\x05\x52\x79\xa6\xe2\x07\x95\x5d\xa1\xb2\xcc\x89\x09\x32\x16\x96\xdd\xbf\xff\x2a\x18\x0d\xe3\x9a\xcd\x7f\x5d\xbd\x79\xdd\x19\x69\x0e\xf6\x59\x13\x99\x35\xa2\x2d\x46\x1b\x8c\x75\x1d\xaa\x41\x94\x0a\x7c\xbb\x86\xd5\x9c\x8e\xad\x58\x9b\xb2\x3e\xce\x9a\x98\x82\x4c\xc9\xa4\x27\xb1\xb1\x71\xb1\x23\x32\xdd\x80\x6b\xd1\xcc\x52\x24\x30\xa8\xed\xd4\xaa\xfc\x9a\x59\xe4\x96\x35\x2e\xdf\x5d\x33\x2e\xff\x4a\xe4\xc6\x10\x78\x9c\x83\xd9\x2c\x33\xde\x7b\xd9\x02\xd1\x51\xcc\xcd\x60\xcf\x58\xb6\x37\x52\xcc\xaa\xc0\xc4\x14\xd3\x6b\x06\x52\x8e\x91\xf2\xb4\x08\x28\xde\x0d\x7b\x4e\x4b\xc9\xae\x7f\xc5\xa9\x54\x2c\x77\x44\x6e\xb4\x23\xb6\x32\x81\x76\xf4\x02\x08\x25\x92\xe8\xbe\x99\xf5\xc4\xf7\x2d\xd8\x49\xb7\xaa\x83\x5e\x95\xb1\x46\xad\x6d\xcc\xcd\xb4\x9d\xf9\x57\x48\xd8\x9d\xcd\xbd\xb3\x95\x95\x97\x24\xc7\x9a\xda\x34\xb8\x38\xed\xf2\x45\x55\xd5\x5d\xce\xa1\x5f\xad\xd0\x05\x17\x13\x10\x9b\x3c\x47\x5b\x52\xaf\x82\xd2\x1d\xb5\x53\xf0\x88\x3b\x3b\xa9\xab\x53\xd7\x3f\x4c\x49\x05\x20\x6e\xaa\x08\x4f\xa6\x76\x99\x63\x31\xd0\xa9\x42\x9b\x91\xd2\x78\xf9\x1b\x9c\x96\x3d\xf1\xf5\xf4\xf0\x06\x07\xf8\x7c\x10\x0c\xa8\xad\x75\xd2\x54\x19\x95\x88\x50\x01\x28\xcf\xb5\xf9\x5d\xb3\x2d\xcd\xdc\xa6\xc1\xb8\x7e\x79\x38\xc0\x66\x5b\x20\xea\x33\x50\xcb\x9d\x6b\x8f\xa4\xc6\x90\xfb\x92\xa4\x28\xcf\x75\x50\x2d\x30\x20\x8e\x81\x5d\x2b\xd6\x38\x33\x7b\x0f\x32\xeb\xd4\xee\x1c\xc1\x62\xa1\xba\xd9\x50\x6f\xa9\xc7\xc3\x12\x73\xa1\x43\x78\x3b\x44\x20\x55\xfa\x34\x25\xbe\x90\x7c\x9b\x4a\x38\x04\xc3\xd6\xbc\x78\x6a\xe7\xf2\x05\x56\x73\x50\xda\x24\x4d\x0d\xd1\x7b\xd3\x3a\xf2\x50\x52\x12\x8e\xad\xa1\xd5\x4f\x4b\x90\x7c\x8b\xbb\xb4\xb6\x5e\x6c\x48\xed\x43\x6d\x17\xbd\xaa\x32\x54\xd5\xb7\xad\xb9\xf2\xba\xf7\x18\x9b\x4a\xbf\x65\x6c\x1e\x06\x18\x3b\xaa\x3f\x77\x18\xbb\x86\x1e\x63\x57\xf7\xb5\xbc\xed\x33\xbc\x59\x2d\xcd\xd5\x01\x9f\x60\x40\x5f\x73\x18\xe6\x34\x06\xf3\x6c\xfb\x7a\xcd\x03\x1a\xb5\xba\x12\xda\xee\xea\x35\x77\xbb\xda\x83\x25\xd3\xd1\x3e\x2c\x6d\xbe\x57\xb7\x0c\x48\xea\xae\x06\x18\x41\xf5\xa3\x93\xb3\x6e\x1c\x10\xd3\xef\x47\x68\xab\x5f\xd3\xd8\xed\xd7\xb9\x8d\x00\x60\x5e\x0c\x9b\x8d\x57\x34\x08\x00\x2e\xad\x32\xde\xdb\x6e\x87\x81\x6a\x56\x00\xd0\xbc\x05\xf3\xda\xf0\x19\x20\xee\xf2\xeb\xfa\x37\xfb\xb0\xec\x97\xb5\x5b\x4e\xd9\xb9\xdf\xa7\x0b\x57\x11\xd3\xfe\xeb\x2a\xdd\xe0\x02\xd9\x14\xc8\x77\x93\x36\x8d\x79\x34\x37\xe9\xb6\xa0\x53\xee\x19\xb8\xbd\xa7\x29\x8c\x0e\x3a\x96\x9e\xa4\x46\xad\xe4\x52\x3c\x43\x02\x2b\x16\xed\x51\x3a\x44\xb5\x20\x13\x83\xb7\xb7\xaf\xaa\xf6\xd2\x5e\x58\x0d\xd7\x4c\x6e\xe0\x9a\xd0\x4c\x68\x41\xea\x8c\x5e\x45\x14\x36\x84\x9f\x03\x91\x80\x84\xd8\x16\x3a\xa3\x43\x12\x52\x56\x94\x39\xbe\x03\xb9\x21\x74\x2d\x80\xa8\xa7\x02\x53\x09\x08\x6c\x49\x59\xc9\x1b\x99\x2c\x36\x79\x87\xd7\x44\x48\xbe\x8f\x4d\xe1\x4b\xc5\xce\x06\x66\x25\x8a\x72\xfb\x42\x33\x70\x51\x88\x84\x1d\xc9\x73\xd8\x0a\xac\x7c\x2e\xd2\x15\x91\x42\xe7\x2d\xa0\xdc\xbe\x30\xa1\x89\x2e\x99\xbd\xc3\x29\x26\xb7\x98\xd7\x80\x4e\x85\x23\x31\x74\xb2\x89\x4e\xe6\x66\x84\xab\x63\x16\x8e\x05\xfc\xf4\xb3\x7e\xd7\x14\x63\x9d\xb3\xb7\x35\x41\x5d\x48\x36\x4a\xbe\xc6\x3b\x63\xfd\xc2\x3b\xb6\x09\x9a\xae\xaf\x90\xf8\xef\x2d\xe6\x7b\xc7\xe2\x46\xf7\xb6\xb5\x0f\x53\x74\x12\x51\x37\xfb\xd2\x5d\xa2\x78\x8c\xa3\x1a\xd1\x31\x6c\x8e\x53\x7c\x2e\xba\x6a\x69\x7c\x2c\xe2\x52\x75\x88\xfe\xe3\x4f\xf0\xed\xb7\xf0\xa7\xaf\xba\x67\x21\x7e\x6e\xa5\xf7\xcf\x0b\xce\x5f\x33\xe9\x3a\xdb\x22\x66\xfd\xe7\x1d\x9c\xd4\xaf\x2a\x57\x80\x1d\x93\x44\x0b\xd0\x3f\x84\x99\xe6\x1a\xcc\xaa\xb6\xce\x1a\x2d\xa7\x78\x00\xb0\xca\xee\x43\x53\x75\x8b\x07\x63\x97\x91\xfd\xbc\xed\x70\x5a\x07\x41\x26\x56\x6c\xa6\x53\xcd\xe6\xa0\xe1\xcd\xe1\x66\xf3\x69\xa4\xe5\x17\x25\xf0\x8d\x48\xfe\x82\xe5\x9b\xef\xfd\x2b\x53\x5e\xf1\xd8\x16\xef\x3b\x56\x9e\xa8\xd5\x3a\xe0\xf8\xa2\xd3\x85\x18\x3d\x15\x03\xbd\x00\x5c\x31\x95\x63\xa1\xeb\xe1\x4d\xd5\xb8\x29\xb5\x5f\x0a\x25\x78\x0d\x04\x1f\xf6\x74\xcb\x73\xf8\xe9\x67\xa1\x6b\x09\x07\x35\x2d\x9a\xbc\xa5\x75\xf5\x60\xb5\x87\x87\x9c\xeb\x9d\xf0\xb1\x54\x34\x29\x73\xad\xe4\xc3\xe4\xf4\xcd\xd1\xf0\xfb\x49\x2f\xb3\xe7\x88\x32\xaa\x22\x5a\xf3\xf2\x7b\xbc\x6f\x01\xf3\xf3\xe3\x6a\xe2\x7c\x86\xb6\x65\xfb\xae\x9d\x68\xf5\x6e\x4f\x0e\xdf\xa9\x34\xe2\xce\x87\xd6\xba\x1a\x44\x31\x1d\xb1\xeb\xfb\x65\x67\x5c\x28\x97\x1a\x7d\xfd\xd5\x57\x73\x08\x95\xe3\x56\x99\xbe\x2e\x91\x7f\x79\x03\x2b\x44\x72\x15\x22\x7f\x79\x1b\xf6\xce\x49\xa2\xb6\x9c\x71\x7d\x94\x13\x6b\x38\x0c\x12\x87\x3a\xc5\xea\xcd\xdd\xb0\xf5\x36\x9e\x45\x29\x75\x78\x81\x24\x5a\x0e\x42\x32\x07\x03\xca\x70\xab\x69\xab\x3a\xd3\x52\x55\xab\x6c\x6c\x79\x66\xd3\xee\x63\x95\x3d\xaa\xff\x78\x88\x1c\x9f\x63\x94\x1d\x3f\xdc\xb5\xd4\x7f\x79\xdc\xe1\x45\xac\xc2\xb4\xce\x42\xfe\x97\x05\xb5\x83\x23\x0b\xd1\x33\x96\x59\x7b\x69\x32\x06\xb2\x72\x8b\xf9\x15\xd2\x14\xbe\x63\x8e\xbd\x1b\x06\xdd\x48\xdb\xe6\xba\xc7\xba\x0c\xdf\x1f\xaa\x61\x5a\x4b\xde\x13\xb3\x17\xf4\x7b\xaf\x2e\xee\x4a\xc6\xa5\xae\x0d\x5c\xb3\x6c\xdf\x3a\xd2\xff\x91\x65\x38\x17\xcd\x09\x64\xf2\x81\x16\x88\x8b\x0d\xca\x0f\x07\x15\x95\x92\xb2\x6e\xab\x0f\x87\x7b\x5d\x7a\xb7\x5c\xae\x72\x92\x36\xd9\x63\xd4\x55\x61\x6e\xce\x6b\x54\xcc\xac\xb2\x01\x3e\xe0\xd2\x61\xb0\xaa\xe1\xc8\xce\xcf\x81\xb0\xe4\xe2\xcd\x4b\x17\xf6\xe9\xb7\xb5\xcb\xaf\x7b\x1d\x7d\x85\x3c\x0e\x4c\x84\x68\xfd\xb9\x91\x7b\xd4\x6c\x1a\xfc\x55\x88\xaf\x10\xed\xdc\x93\x81\x4e\xec\xfa\x8c\xd0\x4c\x17\x6a\x9f\x28\xe2\x8e\x89\x9e\xac\xea\xe8\x4e\xe7\xab\x7d\xef\x5e\x36\x85\x86\x85\xc3\xee\x72\xad\xb3\xc0\xe9\x8d\x56\x07\xe7\xe6\xb4\xe2\x33\x65\x98\x43\x18\xda\x0d\x77\x04\x9f\xce\x6c\xf9\x4b\xb9\xbb\x3f\x0f\x6e\x04\xf5\x9d\x00\xf3\x18\x0d\xa4\xcf\x7e\x22\xdf\xba\x27\x94\x13\x24\x70\xe6\x5d\x0f\x31\x89\xec\x1b\x5d\x1a\x8f\x63\x93\xd9\xc1\x2f\xe6\xa3\x80\xf6\x0d\xa7\x5e\xae\xe9\xdf\x5c\x3a\xd2\x29\xd4\x86\xd0\x76\xae\xd3\xe3\x24\x36\xa3\xc6\xd1\x84\xa3\x9c\x74\x96\xe6\xef\x9a\x63\xf4\x29\xa8\x33\xaa\x81\x79\x68\xfd\xb0\x5b\xcd\x31\xe0\xba\x06\x87\xae\x7b\xd3\x87\xb7\xd1\x5c\x2d\xa8\x23\x75\x9b\xd0\xac\x6f\x4b\x1a\xdd\x1c\x53\x45\x18\xab\x45\xf8\x95\xe3\x73\x8a\xf7\x3e\xb1\x34\xe3\x97\xce\xaf\x8d\xb3\x37\xc2\x75\x8f\xe4\x07\xaa\xbe\xbe\xcd\xff\x36\x2e\xa2\xf2\x65\xea\xdf\x19\x70\xbf\x7d\x24\xff\xec\x80\x6c\xdf\x62\x72\xa7\xac\xcd\x65\x42\xe3\x47\x38\x16\x49\x92\xd4\xbb\xb5\xed\x44\x49\x1e\x54\x41\x70\x38\xc0\x17\x69\x8e\x84\xd0\x80\x2f\xcf\x21\xea\x4c\x42\x6c\x6f\x43\xf6\xb2\xf2\x26\x27\xd7\xc6\xd7\xda\xe3\x5b\xe5\xba\xd6\x17\x7e\xfe\xc9\xce\xf0\x65\xd4\xe9\x9a\x92\x27\x6c\x53\x4e\x1a\xcd\x47\xd1\x4e\x65\x08\x2e\xe1\x9d\xc3\x06\x89\xef\xf1\x1e\xae\x19\xcb\x5d\xbc\x03\x23\xd5\xb1\x43\x2b\x88\xa9\x8b\x8e\x2e\xc5\x8e\x5b\xa6\x43\x56\xf0\x07\xcb\x7c\x68\x6e\x1e\xb4\x9d\xb6\x8c\x40\x97\xc2\xd0\x0e\x8c\x26\x9e\x49\x18\x1d\x5b\x66\x81\x76\x2a\x51\x32\x0d\x3f\xf9\x44\x67\x7f\xfc\xb9\xe1\x7b\x8c\x62\xa6\xf1\xbb\x3c\x67\xbb\x8b\xa2\x94\x7b\x5d\xcf\x69\xbb\x0f\x77\x65\xb8\xee\x64\x2f\x2a\x1c\xff\xed\x19\x47\xbb\xe1\x98\xd3\x2b\x40\x0d\x44\xde\x11\x74\x25\x07\xe3\x08\x8d\xd0\xb5\x38\xf1\x98\xfc\x1a\xa6\x73\x08\x43\x38\xc0\x62\x01\x58\xb5\xd7\xa5\xcf\x12\x09\x73\x3a\xc6\xe4\x06\xf3\x5a\x47\xc2\xa8\xf0\x37\xc4\x56\x31\xdd\xde\x08\x6e\x7b\x81\xe6\x78\x74\xea\x3c\xb5\x89\x7b\x9a\xb8\x88\x09\x93\xe2\x9a\xc3\xcd\xdf\xea\x74\xd5\xb8\xe6\x81\x9b\x86\x7d\x3f\x7c\x7f\x31\xdf\xac\x77\xe7\x99\xa1\x5f\xb9\xef\x9c\x73\xf6\xf2\x74\x2b\x7a\x37\xa8\x74\x0e\xac\xef\xd5\xdd\xb1\x49\x73\xbf\x56\x4f\x69\xfb\x0e\xae\x7f\xfb\x56\x59\xdf\x83\x6e\x89\x1e\xfd\x61\x65\xab\xd1\x4d\xb5\xb1\xfb\x46\x85\x29\xd4\xc7\x36\x38\xad\x5a\x2f\x5d\xef\x3b\xd5\x36\x04\xbd\x7b\xb4\x2d\x01\x5b\x17\xfa\x7d\x39\xff\xa9\x11\x7a\xc0\x21\xd3\xc4\x89\x52\xfd\x5c\x83\xde\x3e\xda\x89\x34\x9c\x89\xfe\xca\xae\x91\x36\x36\x79\x8d\xf9\xd6\xe5\x81\xf3\xc9\xd1\xae\x67\xcf\xd6\xd1\xf8\x5f\x08\xdc\x57\xf7\xac\x5d\xf2\x60\x61\x60\x2a\xc9\x1f\x70\xb8\xb5\x20\xad\xb0\xc1\xa8\xe9\x27\x05\xff\x7c\x1b\x77\x27\xb6\xfb\xbb\x6f\xd0\xce\xf9\xe0\x9b\x81\x93\xd8\xb0\xd8\xe6\x92\x84\x9d\xfb\x1c\x93\x1f\x46\x58\x10\xda\xa9\xf2\xcd\xed\x70\x9c\x7c\x44\x58\x30\xd6\x75\x38\x54\x80\x33\xb0\xc1\xc2\x91\x37\xbe\xc7\x3e\xe6\x18\x19\x76\xe0\x5a\xfd\xc0\x87\x1b\xbd\xa5\xa0\x0b\x35\xf7\x87\x27\x0d\x10\x47\x89\xde\x4a\x4f\xfe\x6e\x86\xd1\x0a\x4b\x46\x8e\xe4\x33\xbc\xfa\x58\xdf\x0d\x1e\xfe\x32\xc6\xdb\xcf\x8f\xc3\xf0\x61\x58\x3c\x79\xa2\xbb\xd4\xf2\xf8\x86\x34\xea\xdc\x6a\xe2\x56\x81\xe7\x73\xe7\x81\x92\xbc\x9b\xb6\xf9\xb8\x4e\x7e\xd4\xe3\xa8\x46\x9d\xf1\xfd\x5f\x26\xb4\xce\x79\x75\x59\xee\x1f\xeb\x8a\xfb\xbe\xb8\xfd\xa9\x9c\x8a\xba\xba\x5f\x94\x0d\x4b\xfe\x10\x8f\x7d\x84\x3e\xf7\xe4\x53\x47\x7c\x6b\x32\xb8\xe3\x0c\xfe\x8f\x0f\xf6\xd7\x62\x01\xb7\xa4\x58\xc2\x4a\x9e\xaf\x59\xf0\xbf\x01\x00\x00\xff\xff\xf2\x51\x7b\xb5\x72\x46\x00\x00")

func templatesParameterGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesParameterGotmpl,
		"templates/parameter.gotmpl",
	)
}

func templatesParameterGotmpl() (*asset, error) {
	bytes, err := templatesParameterGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/parameter.gotmpl", size: 18034, mode: os.FileMode(420), modTime: time.Unix(1536181450, 0)}
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
	"templates/api.gotmpl": templatesApiGotmpl,
	"templates/config.gotmpl": templatesConfigGotmpl,
	"templates/parameter.gotmpl": templatesParameterGotmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"api.gotmpl": &bintree{templatesApiGotmpl, map[string]*bintree{}},
		"config.gotmpl": &bintree{templatesConfigGotmpl, map[string]*bintree{}},
		"parameter.gotmpl": &bintree{templatesParameterGotmpl, map[string]*bintree{}},
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

