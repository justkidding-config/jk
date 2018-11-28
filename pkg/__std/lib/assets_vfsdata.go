// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package lib

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 11, 21, 15, 57, 26, 637456488, time.UTC),
		},
		"/std.js": &vfsgen۰CompressedFileInfo{
			name:             "std.js",
			modTime:          time.Date(2018, 11, 28, 18, 21, 17, 100371192, time.UTC),
			uncompressedSize: 20438,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x7c\x5b\x73\xa3\x38\xda\xf0\xfd\xf7\x2b\x48\x2e\x52\x68\xc0\x14\x10\xb7\x27\x6d\x2c\x77\x75\xcf\x24\xbb\xd9\xea\x4d\x4f\x25\xe9\xd9\xda\x4d\xa5\x5c\xc4\x88\x44\xdf\x38\xc2\x23\x44\x67\x7b\x6d\xbf\xbf\xfd\x2d\x49\x08\x04\x08\x1f\xb2\x99\x77\xf7\xa6\xd3\x80\x9e\x83\x1e\x3d\x67\x49\xfe\x16\x53\x8b\x41\xb6\x5e\xaf\x36\x11\xf3\x2e\x32\xfa\x1c\x33\xb8\xfa\x58\xb0\x6c\xec\xbb\x7f\xb9\xf9\x72\x35\x0e\xdc\xbf\x7f\xfc\xeb\xe7\x71\xc8\xbf\xff\x8d\x62\x86\x3e\xd2\xc7\x1c\xa6\x05\x99\x33\x9c\x11\x1b\xac\xd8\x13\xce\xbd\x87\x07\x48\x8a\xc5\x22\x2a\x1f\x66\xcb\x2c\x87\x7e\x03\xc4\x5b\xd2\x8c\x65\xec\xfb\x12\x79\xb3\x19\x26\x98\xd5\x38\x98\x8b\x2a\x34\x02\x92\x29\x3c\x10\x45\x14\xb1\x82\x12\x8b\xbf\x68\xe2\x7b\x44\xec\x3a\xcb\xd8\xc7\xdc\xc0\x15\x72\x29\x58\x49\x48\x9b\xae\xd7\x04\xbd\x58\x1a\x28\x28\x39\xb0\x91\x47\x51\x9c\x5c\x12\x76\x1a\xda\xc8\x5b\x66\x39\x96\x73\x02\x8e\xfe\xe4\x22\xd0\x37\x93\x65\xcc\x9e\xb4\x79\x80\x15\x97\x27\x82\x25\xf7\xde\x6c\x96\xa5\x69\x8e\x98\xad\x4d\xce\x1d\x02\x35\x27\xf4\xa1\x1e\x98\x33\x8a\xc9\xa3\x3e\xd0\x41\x2e\x03\x63\x2e\xd5\x3e\xea\xdf\xe2\x45\x81\x0e\x25\x3f\x7a\x33\xf2\xfc\x1f\x5d\x11\xf6\x20\x7e\x66\x20\x5e\x2e\xc1\x59\x93\x38\x18\x2b\x75\xf4\xb8\x36\xf6\xf1\x80\x49\x82\x08\x6b\x73\xc1\xb6\x73\x11\xf8\x15\x1b\x6c\x3b\x1b\x0c\x8c\x5b\x5a\x9c\xb3\x98\x32\x83\xc6\x31\xb0\x62\xf2\xe3\x97\x87\xff\x8f\xe6\xcc\x1e\xb6\x94\x26\x4e\x92\x5f\x9a\xda\x22\xb4\x9e\xbf\xbf\xc0\x68\x91\x7c\x91\x7c\xfa\x2e\x72\xfd\x2e\xe8\xaf\xad\xa5\x36\xc2\x06\x66\xd8\xdb\xc6\x3a\x09\xd3\x40\x15\xa8\x98\x72\xe8\x52\xb7\x21\xee\x2e\x92\xcb\x96\xa0\x5b\x1c\x08\x34\xa7\x06\xfa\x88\x24\x66\x69\x95\xda\xc2\x07\x94\x12\xab\x75\x63\x13\xc9\xaf\x48\xf8\x25\xe4\x09\xe0\xd5\xd5\x97\xab\xf3\xb1\xef\x56\xe8\xc6\x81\x7b\x8d\xe2\x44\xfc\x37\x74\x7f\x8a\xc9\x1c\x2d\xc4\xc3\x29\x87\xf9\x2b\xca\xf3\xf8\x11\xed\xeb\xa9\x2a\x80\x37\xf1\x53\x35\xb6\xca\x4b\x75\xf8\x61\x5d\x1f\x55\x81\x55\x1e\x8a\x69\x1e\x8a\x35\x3c\x94\xfe\xe4\x32\x60\x9e\x41\x4c\x1f\xf3\x5b\x83\x99\xee\x30\x90\x61\x8f\x7d\x7c\xc5\x06\x03\x91\xcb\xe3\xf1\xd5\xe9\x67\xe2\x6d\xbc\x54\x41\xa4\xe0\x5a\x9e\x42\x3a\xa9\x9a\xb4\x30\xc3\xae\xbc\xdb\x16\x1a\x36\x84\x16\x27\x42\x93\x9a\xd2\x12\x6b\xd4\xd2\x72\xdf\xa5\xae\x36\x67\x13\x92\x7d\x0d\xb5\x06\x44\x24\x31\x31\xbc\xdb\x48\x28\x5c\x6d\x22\xea\xdd\x5c\xfe\xe3\xfc\xcb\xc5\xec\xe6\xcf\x5f\xae\x6f\x61\x58\xbf\xb8\xbc\xba\x85\xc3\x88\x7a\x17\x97\x9f\xcf\x67\x97\x3f\x9f\x5f\xdd\x5e\x5e\x5c\x9e\x5f\xcf\x3e\x9f\x5f\xfd\xe9\xf6\xcf\xe2\xd3\x39\x99\x67\x09\x26\x8f\x70\xf5\xf5\xf6\xe2\x6c\xf6\xe9\xef\xb7\xe7\x37\xe3\xc0\xfd\x7a\x7b\x11\x8c\x66\x37\xb7\xd7\x97\x57\x7f\xe2\xd1\x9f\x7a\x98\x6b\x21\xe4\x7a\x2a\xf4\xf1\x23\xa5\xf1\x77\x3b\x04\x11\xf5\xd2\x45\x16\xab\x6f\x17\xf2\xff\xf2\x6b\x09\xe4\x3d\x14\x69\x8a\x68\x35\x74\x34\xac\x87\x8e\x86\x7d\x43\x71\xfe\x19\x33\xb6\x40\xe7\x24\xc1\x31\x11\x10\x5c\x01\x83\x91\x04\x50\xcf\x67\xf2\xf1\x2e\x70\xfd\x7b\xa0\xc0\xef\xfc\x7b\x08\x61\x10\x51\xef\x73\x46\x1e\x4d\x26\xbc\xc8\x5e\x20\x5b\xfb\xd2\x82\x9f\xf0\xe3\x13\x44\x6b\x7f\x53\x02\x78\x73\x8a\x62\xd6\x76\xb8\xca\x26\x20\xf4\x4f\x4e\x10\x84\xfe\x87\x72\xf4\x3f\xce\xaf\xbf\x8c\x39\x3f\xf2\x59\x8c\xae\x50\x69\x01\x33\x2b\x67\xac\x9b\x63\xe9\x01\x14\x4f\xd3\xe9\xd4\x07\x4e\xc5\xd4\x0f\xc3\xf0\xfd\xf0\xfd\xe8\xc7\xf0\xfd\xc8\x80\x10\xfd\x5e\xc4\x8b\xa6\x6d\x69\x6e\x48\x4c\x11\x32\xfe\xe7\xe4\xa4\x9e\x26\x64\xe2\x6f\x85\x8e\x33\x0f\x35\xe6\x7d\x97\x87\x47\xef\x53\x81\x17\x09\xa2\x0d\xe4\x38\xb5\x8f\x2a\xad\x0c\xfc\x70\xb8\x41\x8b\x1c\x29\x2d\xdd\x28\x67\x48\xbd\x4f\xdf\x19\xfa\x24\x56\xc2\x8b\x17\x8b\x6c\x1e\x33\x64\x23\x20\x65\x9d\x2f\xe3\x39\x82\x48\x3e\x3c\x63\x12\x2f\xf0\x23\x81\x81\x7c\xfe\xc6\xe2\x87\x05\xd2\x3c\xb4\x7c\x31\xc3\x64\x56\xe4\x08\x96\xcb\x85\xf3\x2b\x94\x33\x94\xc0\x34\x5e\xe4\x25\xa6\x4c\xd8\xc7\x4c\x18\xb8\x1a\x27\x81\x73\x78\x77\x5f\x3e\xa3\x39\xcb\xe8\x8c\x14\xcf\x33\xb4\x40\xcf\xb9\x1a\x97\x66\x74\x8e\x66\x09\x4a\xe3\x62\xc1\x72\x89\x75\x53\xcb\x40\x93\xb8\x18\xf9\x73\x35\x50\xf7\x2d\x06\x44\xcc\x8c\x24\x89\x59\x2c\xa5\xd3\x55\x04\xab\x14\xa2\x19\x32\xce\x6b\x8d\xef\x87\xf5\x1e\xbe\x33\x94\xdb\xc0\xcb\x8b\x87\x58\x18\x87\xfa\xa0\x87\x8c\xce\x2b\xa9\x75\xa5\x33\x06\xa0\x8f\x03\x41\x7c\x17\xfb\x15\x0b\xd2\x1e\xbd\x7c\x81\xe7\xe8\x0d\xd9\x58\x52\xb4\x6c\x59\x27\x4e\x6d\x36\x6d\x28\x55\xb9\x28\x95\x8e\xb1\x0d\xd7\x54\x0c\xff\xa7\x62\x64\x1e\x2f\xe3\x39\x66\xdf\x6d\x30\xa8\x75\xd3\x41\xc0\x09\x4e\xd8\x20\x88\x5e\x9e\xf0\xa2\xe4\x5a\x7c\x99\x60\x87\x39\x48\x1a\x00\x81\x5d\x24\x91\x66\x01\x25\xcb\x8f\x34\x7b\xa9\xad\x41\x11\xd6\x4d\xc1\x31\x20\x1a\x10\x69\x4c\xcb\x38\xb1\x71\x9f\x08\xe2\xa4\xa1\x80\x69\x46\x6d\x69\x89\x7e\x84\x26\x2c\x42\x8e\x53\x65\x2b\xde\x0b\x4f\x9a\x44\xf8\x1a\x68\x33\xe5\x41\xc8\x8c\xbc\x1a\xdf\xd5\xf1\x06\xb6\x1a\xd7\x00\x06\x22\x0d\xd9\x86\x2d\x18\x6d\x45\x17\x8c\x1a\xf8\xc2\x9d\xf8\x4e\xc3\xad\xf8\x78\xce\xa4\xe1\x1b\xee\xc4\xa7\x7b\x66\x03\xbe\xd1\xb0\x81\xef\x6c\x3b\xbe\x32\x10\xf6\x63\x2c\x07\x1c\xc2\x63\x27\x7e\x18\x71\xee\xcd\xa7\xc8\xeb\x4d\x6b\xcc\xcd\xcb\x0e\x78\x20\x10\x8f\xda\x7a\x6f\xc5\x64\x5a\x5f\x81\x2a\xec\xa0\xe2\x6b\xbd\x15\x97\x49\x72\x02\xd7\xb0\x83\x8b\xcb\x70\x2b\x2e\x93\xc4\x04\xae\xb3\x0e\x2e\x2e\xbb\x7e\x5c\xbd\x6b\x6a\xe2\xac\x5a\xdf\x1d\xf8\xf6\xe5\xae\x5a\xdb\x2d\xf8\x54\x96\xda\xf0\x8d\x3c\x89\xe5\xde\xb1\x1b\xa1\xd6\x6b\x74\x04\x69\x49\xb3\xd4\x86\x3a\x50\x2f\x32\xc6\x89\xed\xa4\xd6\x58\xf6\x03\xc9\x05\xa3\xc3\xe9\x35\x16\xe0\x40\x7a\xa7\xe1\xe1\xf4\x1a\x0b\xb4\x9d\xde\x91\x4a\xca\x6c\x0a\x1a\x74\x47\xc3\x03\xe9\x76\x55\xed\x90\x99\x2a\xdd\x7b\x05\xcd\x03\x66\x6b\xa0\x79\xf0\x3c\x65\x49\xf4\x4a\x92\x65\x3d\x75\x18\xc5\x1b\x46\x8b\xb9\x91\xa2\x86\x9c\x88\x04\x73\x4f\xcc\xa4\xcc\x46\x9b\xa9\x32\x3b\x82\xcd\x54\x66\xc5\x9e\x68\xf6\x62\xf1\x5c\xfb\x9c\xd2\x8c\xda\xc7\x17\x8b\x98\xc9\xcc\x20\x1f\x5b\xb9\xe0\xcb\x7a\x2e\x72\x66\x3d\x20\x2b\x47\x14\xc7\x0b\xfc\x2f\x94\x58\x98\x2c\x30\x41\xde\x71\x2f\xfd\x8c\x5d\xb5\x58\xa8\x25\xa8\x92\xe5\x5d\xe4\x65\x02\x5d\x91\x8d\x39\x1a\xc9\x0c\xc9\x04\x43\x72\x96\xfd\x5c\x70\x19\x75\x3d\x99\xcc\xc1\xef\xd8\x7d\x53\x18\x66\x14\x59\x4b\x19\x3a\x69\xa5\x31\x61\xd3\x71\x35\xf3\x2d\x63\x31\xad\xe5\x6b\x7c\xcd\x4f\x4e\xc3\x30\x08\xc3\x77\xc3\x1f\xc3\x5d\x22\x9a\xc7\x84\xcb\x82\xd3\xb0\x64\x62\x6b\x3d\xa0\xef\x19\x49\xac\xd0\x7a\xc4\x8f\xb1\x48\x7a\xb9\x80\x64\x9e\x89\x26\x93\x20\x92\xb9\xa2\xb9\x24\xc2\x20\x22\x5e\x8e\xd8\x2f\x2a\xf9\xc5\x03\xc4\x5f\x55\xf9\x3b\x62\x36\x53\x4f\xae\xf8\x58\xca\x83\xf4\xaa\x78\xc7\x9e\xf4\x88\xa2\xf7\x05\x8c\x61\x54\x5f\xa2\x01\x73\xf4\xf1\x3d\x4b\xa6\x75\x54\xba\x44\x2b\xc5\x94\xb2\xd6\x4b\x3c\x51\xe3\x35\x74\x04\xde\xdd\x6f\x0c\x35\x1f\x8b\xfa\x33\xdb\x52\xb9\xd0\x3d\xf4\x37\xcd\xca\x90\xd1\xc2\x54\x18\xee\xa1\x83\x55\xaf\xc5\x64\x4c\x3a\xf3\xeb\xf5\xd1\x41\x06\x56\x21\xb6\xe6\xf1\x62\x81\x12\xeb\x05\xb3\xa7\xac\x60\x96\x26\xc2\x63\xb0\x69\x06\x2b\x1f\x44\x5a\xc3\x4e\x71\x1e\x69\x2d\xb4\x86\xb4\x06\x81\x90\x56\x84\xa6\xd0\x2f\x8b\xfe\x5a\x46\x5c\x7c\x83\x01\x58\x29\xf5\x74\xb4\xc1\xf2\x4b\x33\x30\x37\xc1\x8f\xa0\xff\x81\x0d\x9a\xef\xc6\xbe\xd4\x75\x02\xc3\xa8\x05\x3b\xe8\x08\x5f\x72\x9d\x43\x1b\x3b\x04\xfc\xd0\xec\x59\xb5\xa0\x73\x39\x36\x83\xbe\xf8\x1b\xc3\xda\xd8\x23\x36\xe6\x4c\x97\xba\xa0\x55\xf9\xde\x02\x91\x47\xf6\x24\xb5\x83\x43\xa5\xa6\xc2\x4a\x87\xb8\x43\xf7\x5c\x2b\x73\x08\xf5\x26\xa7\xe4\x20\x05\x75\x45\x55\xc0\x16\xb3\xc5\x24\x8f\x0a\xa7\xf5\xb6\x56\x92\x06\xa2\xd8\x29\xc0\x91\x89\x80\x53\x00\xb0\x9a\x67\x84\x61\x52\x20\x8b\x6d\x36\x19\x6c\x33\xf7\x40\x51\xfc\xdb\x66\x83\x53\x3b\x2b\xd7\x46\x36\x4f\x4c\xf3\x8a\xb6\x56\x3f\x6e\x36\x60\x40\x76\x6b\x1a\x32\x5b\x16\xf9\x53\xd3\xea\x41\x2f\xa2\x16\x41\xb7\xe5\x2c\xc0\xc6\xd4\x9b\x51\xfe\xbb\xa7\x9d\x82\x09\xce\xdb\x7b\x22\xdc\x29\xcb\x15\xc4\xaa\x4d\x24\x9c\x57\xa3\x98\x77\x75\xd7\xe4\xf4\x35\x38\x85\xd7\xc1\xa5\x66\x1c\xc1\xde\x61\x3b\x2c\x37\xc5\x0b\x64\xe1\x04\x11\x86\x53\x8c\x68\x15\xa2\x25\x5e\xeb\xb8\x9f\xfe\x46\xe9\x10\xe9\x25\x3e\x08\x22\xc2\x0d\x90\x54\x06\x58\x97\x5b\xd8\x9b\x3f\xc5\xf4\xa7\x2c\x41\x1f\x99\x4d\x00\xd8\x6c\xf6\x11\x07\x88\x5a\xf9\x11\xab\x17\x55\x8f\x36\xb5\x7a\xf4\xb8\x42\x8a\x7e\x2f\x30\x45\x32\x6d\x6a\xad\x92\x6c\x3e\x1b\x35\x51\x2e\x1d\x1d\xb4\x94\xfe\x34\xb4\x29\x88\x9a\xbd\x93\xda\x1e\xb0\x83\xc0\x11\xf4\xf9\x82\x1d\x91\xdd\x0b\x82\x16\x89\x75\xec\x20\xe7\x58\xcb\x97\x58\x7f\x6a\xc2\xfd\xcf\xaf\xa2\xfb\xd7\xcc\xfb\xb0\x21\x50\x49\x03\x69\xb7\x0a\x75\x4d\x6c\x84\x51\xf6\x83\x4a\x12\xc5\x37\x2c\x5e\xf4\x86\x96\x36\x13\xad\x25\x57\xb6\xd6\x26\x0f\xf4\xfd\xa5\x1d\x01\x4c\xf6\xae\x6f\xc4\x8e\x6e\x27\x27\xb5\x30\xc9\x59\x4c\xe6\x28\x4b\xb5\xee\x79\x95\x1c\xe9\x0d\xdd\xbb\xfb\x72\x87\xc1\x2f\x1b\x61\x74\xc2\x4a\x53\x2a\xcd\x53\x2d\xa6\xae\xa5\xd4\x71\x84\xd5\x91\xc9\xbb\x77\xe1\xfb\xd1\x7a\x4d\xa6\xf0\xdd\xe8\x34\xf4\xc1\x0a\x43\x52\xe3\xcf\x4d\x60\xd0\x26\x93\x49\xe0\x03\x27\x77\xec\xd1\xbb\x77\xa7\xa3\x81\x2d\xb0\x88\x97\x03\x89\x86\xbb\x43\x3c\x09\xc2\x33\xb0\x42\xd2\x7f\xe1\xd2\xb3\x89\x0f\xa1\x3f\xd4\xbe\x4c\xa7\xa3\x93\xd3\x60\x1d\xbc\x0f\xf5\x31\x02\xb5\x3e\x28\x08\x4f\x82\x77\xeb\x30\x1c\x96\xa3\xb4\x2f\x67\x27\x3f\xae\xc3\xa1\xef\xca\x51\xa3\xd3\x35\xa7\xbc\xd1\xf1\xb7\xdf\x55\x2f\x36\x9b\x46\x5d\xad\xf2\x2d\x4d\x19\xed\xc0\x45\xa5\x48\xdd\x60\x97\x91\x0e\xa0\x1a\x0b\xaa\x94\x88\x42\xdf\xcd\xb4\xe8\xe8\xc6\x95\x69\x95\x69\x63\x44\x27\x0a\x2c\xe2\x42\x5e\xc5\x77\x99\xe3\xdc\x43\x74\x47\xef\x37\xba\x52\x55\xca\xb9\x5d\xaf\x0c\xfb\x29\x25\x96\xc6\xd6\x49\xb5\x07\xd2\x93\x8c\x4b\x2e\x39\x8b\x33\xb5\x8d\xaa\x5a\xc0\x33\xe8\x37\x01\xab\x84\xd9\xb4\xd3\x21\x77\x2e\xb4\x16\x6b\x6b\x5f\x88\x81\x16\x1b\xda\x8c\x04\xfd\xde\x6a\x43\x30\xd7\x0b\xab\xb8\xed\x03\xaf\x66\xd3\x8b\x41\x5b\x65\x43\xba\x5e\x09\x83\xf5\x22\x50\x4e\x77\xfb\x0c\xca\xc5\xef\xc5\xa2\x0e\x3e\xf4\x6e\x23\x69\x3b\xbf\x60\x32\x09\x87\xd3\x69\x38\xdc\x8a\x4d\x0c\xee\x45\x27\xb9\xba\x63\xf7\xbb\x38\x6a\x75\x13\x4d\x2c\x89\x2e\xe2\x64\x12\x8c\xa6\xd3\x60\xb4\x93\xa7\x2d\x08\x2b\xa6\xd6\x8d\x47\x27\xb8\x9f\x4c\xce\x76\x31\xda\x6a\x08\xee\x8d\xb7\xf5\x2e\xbc\xe7\x33\x69\xbd\x3c\xbd\xe7\x22\xdf\x39\xb5\x2d\x2c\x68\x07\x08\xc0\x74\x3a\x6d\xdb\x56\x67\x36\xad\x76\x64\xc3\xcc\xe4\xee\x66\x1b\xab\xdb\x7e\xe3\x0c\xb7\x18\x9d\x62\x79\x7f\x3a\x72\x82\x0d\x42\xea\xd5\x4e\x4a\xa6\x8e\x6d\xb9\xd5\x7c\xe7\x97\x3d\x09\x7d\x2e\x51\xe5\xcc\xca\x4d\xed\x3b\x7f\xbb\xa2\x9a\x5a\xb8\x8a\x40\x7b\x0f\xfb\x83\x3f\x0e\x8c\x34\xfb\xc6\x07\x63\x03\x8f\xce\xb0\xcd\xe5\x68\xb8\x8d\x4b\xd3\x5e\x4e\x7d\x92\x45\x29\x28\x44\xdb\x11\xb4\x8d\xfa\x70\x0c\xdd\x4e\x71\x17\x43\xd4\x36\x13\x88\xa6\xd3\x7e\x03\xac\x38\xfb\x43\x10\x77\x5b\xcd\xfb\xe2\x8d\xda\x96\xcd\xdf\x06\xa3\xa8\x6d\xdb\xfc\xf5\x16\xeb\xae\xa6\xf7\x5f\xc0\x46\xb7\x0f\x6e\x48\x5d\x79\x36\x93\xbd\x18\xfa\x4a\xce\xd0\x45\xe2\xd8\x41\xbf\xb5\x56\x93\xdd\x42\x46\x99\x7d\x97\x4e\xed\x10\xf6\x23\x64\x6a\xaf\x73\xc3\xad\xad\x5e\x09\xb5\x31\xbf\xda\x75\xec\x83\xbf\x33\x13\xdd\x5e\xb7\xe2\x37\x7a\x0e\xb3\x5c\xb7\x3a\x8f\x7e\x2e\x1f\x51\x59\x4e\x5d\x56\x75\xad\xa9\xe1\xd5\xc8\x21\x26\xcd\xd4\xc4\xd9\xaf\x06\xdf\x51\xcb\xd5\xdc\x59\x38\xb7\x58\x96\x59\xf9\x53\x46\x99\xc5\x32\x6b\x9e\x11\x16\x63\x62\xc5\x44\xab\xbe\x55\x7f\x95\xc1\xe3\xe3\x56\x6b\xb0\x8f\x85\xb2\x63\xe8\x40\x59\x14\x79\x29\xcd\x9e\x7f\x2a\x6b\x8f\x46\x38\x2b\x37\xb4\x7b\x66\x88\x00\xd8\x34\xfa\x19\x26\xb9\xaa\xe3\x6c\xe6\x72\x79\xd0\x1b\x70\xca\x66\x56\x5d\x10\x53\xf0\xa1\xfd\xc6\x41\xe2\x3c\x6a\x2f\x65\x71\x3e\xae\x73\xf8\xac\x3c\xac\x88\x9c\x16\x6d\x5e\xb6\x7a\x0f\x0f\x22\xbe\x44\x7b\xcc\x2b\x6f\x57\x94\x02\xbd\x63\x88\x67\xb2\x01\x60\x7e\x4f\xf8\xb2\xc9\x02\xd0\x8f\x98\xd6\x44\xbb\xbc\xba\x15\xfd\x78\x08\x61\x7d\x28\xcd\xab\xcf\xa4\x19\x73\xdb\xfa\x84\x8b\xcb\x1c\x0c\x36\xb2\x50\xcd\x27\x58\x4a\x3c\xd3\xdb\x87\x5a\x1e\xeb\xe4\x65\x95\x1a\x4f\x78\x69\xb8\xca\x60\x5c\x57\xa6\xe9\xb6\xe1\xbc\x46\x5c\x65\xd0\x8e\x4f\x4e\x03\x30\x99\x8c\xd6\xe9\xc9\xe8\xb4\x86\x2d\xb6\xc2\x0e\xfd\x12\x36\x78\xc7\xb3\xd6\x70\x6d\x73\x68\x81\xa6\x68\xa0\x49\x7a\xd0\x08\xe0\x1f\x39\xec\x59\x05\xcb\xd1\x14\x0a\x4d\xc2\xd1\xc8\x8e\xa1\xaa\x72\x89\x59\xed\xb3\xb2\xd4\xcd\x06\x50\x0c\x8c\x7a\xc6\xd9\xd9\x74\xca\x2b\x72\x51\x87\xbb\x76\x76\x62\x07\xb2\x1a\x0f\x80\x53\x16\xe4\x9b\xc6\xa6\x84\x59\x79\x30\x49\x30\x6d\xef\x11\xa8\x05\x6d\x6b\x26\xeb\x77\x5b\xb3\xd9\xb7\x76\x1b\x67\x1b\x9e\x86\x01\xef\x44\x3a\x5b\x20\xb2\x4f\x1e\xdd\xa5\xb3\x8d\xe1\xa7\x38\x9f\x61\x83\x93\x2d\x3b\x32\xff\x0d\xdd\xc9\x7d\xdc\xa7\xe0\x55\xeb\xd8\x20\xd5\xd6\xde\xd7\x77\x2a\x69\xca\xe3\x75\x95\x27\xa5\x45\x7f\xde\xf8\x9a\x96\xc3\x3c\x23\x39\xb3\x30\x94\x1b\x2a\x5e\x4a\x11\xfa\x17\xb2\xd5\xf1\x76\x10\x29\x44\x16\xb1\x89\x9b\xc3\xe3\x63\x77\x95\x8a\x4f\xe3\x0c\x62\x71\xfa\xdd\x95\xb7\x0a\xc6\x31\x0c\x37\x70\xb5\x11\x5d\xfa\x9c\x59\x69\x79\x4a\xb2\xec\x8c\xd8\x81\x1f\x0e\x41\x49\xae\x80\x7f\xb9\xf9\x72\xe5\x49\x17\x89\xd3\xef\x36\x51\x5f\x12\x98\x36\x3a\x72\x76\xa1\xbe\x3c\xb5\xbf\xe4\x60\xcb\x55\x03\x3b\x05\xc6\x4b\x01\x76\xea\x26\x9d\x2f\xbf\xc4\xec\xc9\x4e\xdd\xa7\xce\x87\xdb\xef\x4b\x0e\x91\x75\x3e\xc8\xf3\xfd\x76\xea\xc6\x8a\xbd\x07\xd8\x77\x8e\x9f\xb3\x62\x3e\x6d\xdd\xfc\xa2\x1d\xa9\xb6\x53\x75\x66\xba\xbe\x85\xd3\x1d\x69\xa7\xee\x83\x22\xbf\x84\xa6\xd3\xd1\x9c\x40\x5a\x6e\x1b\xd8\x4b\x10\xfd\x7a\xf6\xb7\x8c\xfe\x86\x68\xe8\xe5\x88\x24\x76\xda\x3c\x9f\x68\x03\xb0\xa9\x96\x3b\x97\x9a\xc3\x15\xe8\xf8\x98\x6b\x8a\xdc\x61\xca\xc4\x1d\x83\xcc\xbb\x46\xec\x5b\xbc\xa8\x6e\x19\xfc\x8c\x52\x44\x29\x4a\xc6\x81\x2b\x6c\x6e\x1c\xca\x41\x79\xb6\x28\x38\x3a\x79\x21\xa3\x1a\x1d\xb3\x98\x8f\x24\xc9\x97\xf4\x86\x51\x14\x3f\x8f\xc3\x12\xee\x94\xc3\xd5\x97\x13\xf6\xbd\x8f\xa0\xc3\xbc\xc9\x95\x84\x06\xc2\xea\x56\x82\x89\x31\xcd\xc0\x6c\x24\x2f\x26\xe8\xc0\xaf\xb9\x9b\xd0\x33\x1b\x79\xd6\xe0\x0d\x2f\x27\xa8\x93\x6f\xf5\xed\x84\x6a\x83\xa2\xf2\x24\xe2\x80\x73\x8b\x25\xa1\xc3\x46\x59\xb4\x2f\x0d\x04\x6d\xd0\x38\x49\x6e\x5a\xd3\xe8\xde\x8d\x19\x0d\xc5\xc5\x1e\xd6\x66\xa3\x8d\x0c\x91\xa4\x87\x8b\xed\x37\x01\x32\x4f\xe9\xeb\xfe\xfa\xa5\x20\xde\x48\xbb\x2a\x74\x95\x6e\x75\x59\x32\x6a\x96\x1a\xf6\x3a\xbd\x32\xcc\xe2\x3f\xab\x55\x15\x43\x42\x6b\x0c\x32\x30\x6b\x54\x05\xf6\xef\xeb\x53\x85\x0a\x91\xc4\x48\x7f\x97\x2e\x09\xc7\xb5\xbf\x22\x89\xe1\x6f\xa4\x45\x12\x57\xa5\x42\x2d\x4e\x8c\xfa\x23\xc6\xbc\x4e\x79\xda\x9c\x3f\xf7\x5f\xc1\xd9\x4f\x77\xf6\xbf\x52\xa9\x88\x0b\x4d\x68\x4f\xd3\xac\x22\x12\x20\x4e\x0c\x37\x85\xb6\x5e\x25\x54\x90\x88\x24\x5d\x42\xbb\x74\xe1\x1a\xe5\xcb\x8c\xe4\x7b\xdf\xa3\xab\x21\xde\x48\x23\x2a\x74\x95\x52\x74\x59\x32\xea\x85\x1a\xf6\x3a\xd5\x30\xcc\x82\x8a\x0c\xe1\x8f\xbe\x4e\xa7\x32\x91\xf2\x42\xdd\x16\x4e\xfe\x0f\xee\xd4\x69\xd4\x85\x3e\x1a\x44\x6f\xba\x55\xa7\x81\xc5\x49\x72\x6d\x90\x9b\xe9\xfa\x28\xd7\xd7\xc6\xec\x7b\x30\xed\x7b\xbb\x4e\x83\x45\x24\x31\xb2\xbe\x33\xaa\xc6\x2c\x3e\x20\xa2\xc6\x2c\x7e\xab\x68\xca\x51\xd5\x91\xb4\xc1\x86\x39\x8a\xc6\x2c\x7e\x65\x04\x6d\x72\xdd\xda\x13\xfd\x23\x3d\xa0\xa0\x2c\xc3\x64\x73\x82\x3d\x21\x92\x0f\x8f\x93\xe4\x53\x8b\xc3\x1d\xce\x4f\x80\xf1\x50\xd8\xa6\xb1\x33\x0c\xd6\xd9\xfc\x01\xc1\xb0\x06\x7a\xab\x90\xa8\x61\xac\x03\xa3\x89\x37\x73\x78\xac\x47\xbe\x32\x48\x6a\xf4\x65\xb4\x32\xd2\x6e\xaf\x99\xdf\x01\xe6\xf1\xa7\x07\x74\x8f\x28\x54\xd6\x5d\x07\xc5\xa1\x12\xe6\xed\x22\x91\x42\xa8\xc7\xa2\x0e\x63\x7d\xd1\xa8\x1c\xf8\xea\x78\xd4\x9d\xcd\x7f\x36\xd3\x35\xb2\x24\x7e\x93\xe2\x15\x31\x72\x74\x68\x8c\x6c\x14\xe2\x7a\xb0\xec\xe1\xe9\x50\x97\x66\xfa\xa9\x8a\x7d\xa2\xa5\x22\xaf\xe2\x65\x57\x3d\xda\x96\x72\xda\x96\xe6\xbf\x5f\x02\x34\x91\xfd\xda\x5d\x13\x53\xfc\x0d\xca\xf8\xdb\x95\x6c\x1f\xce\x5d\x4e\x38\xd4\x23\xb1\x02\x96\xb1\xd8\x24\x96\xdd\xb7\xdd\x63\xb8\xda\xd4\x0d\xbc\x94\x03\xc9\x6e\x11\x82\x7b\x9c\x03\xaa\x5a\x83\xbb\x6c\xd9\x46\x6a\x2c\x81\xb8\xb4\x32\x1b\xd4\xf7\xb7\xcb\x03\xd8\xb9\x9b\x46\xf9\x0b\x66\xf3\x27\x1b\xd7\x8a\x6f\x03\xb0\x9a\xc7\x39\xb2\xba\xb2\x14\x6d\x22\x7b\x95\xf0\x3f\xf9\x06\xc6\x77\xe4\x5e\x11\x62\xb0\x8e\xe3\x51\x89\xcd\x66\x20\x4a\x61\x75\x01\x40\x1e\x07\x8e\x7a\x70\xcb\x6e\x93\xbd\x42\xe2\x6f\x0b\x3b\x82\x5a\xad\x54\xa1\x47\x1c\x7d\xdd\x55\x46\x9e\x80\xb5\xc1\x2e\x4a\x5a\x97\xcb\x5e\x21\x92\xd4\xd4\x24\x5c\x79\x5b\x67\xdc\x69\x5b\x7f\x25\xbf\x91\xec\x85\x58\x65\x9d\x65\x51\x34\xc7\xe8\x1b\x4a\xac\x94\x66\xcf\x16\x2d\x08\xc3\xcf\xe8\x58\x9c\xf0\xcb\x21\x84\x05\x49\x50\x8a\x09\x4a\x94\x3b\xdd\xe4\x76\xaa\x75\xf4\x0a\x79\x75\xc7\xc5\x60\x15\xdf\xb1\x7b\x28\x05\x8b\x5c\x29\x02\xea\x72\xd6\xf0\xa6\x1e\x9f\x68\x2d\xfd\x46\x18\x69\xb2\xc9\x80\x06\xf3\xa4\xeb\x18\x57\x48\xa5\x43\xbb\xf4\x0d\x81\x5a\x87\xb6\x55\x31\x36\x06\x4a\x89\x88\x56\x62\x34\xb4\x48\x64\xc4\x72\x81\x9b\xda\x22\xd7\x53\xc1\x69\xbb\x89\xbf\xd0\xec\x19\xe7\xbc\x52\x10\x76\xa4\x4d\xae\x5a\x64\x10\x35\xf1\x57\x1d\xcf\xa6\xca\xa8\xd7\x35\x95\xca\x38\x28\x44\x66\xe3\xd0\xce\xf2\x94\x8c\xd8\x4d\x4f\x51\x89\x17\x73\xf1\x26\x68\x81\x18\xb2\xe2\x3b\x7a\xaf\xc7\x5c\xbe\x10\x85\x4d\x5d\xec\x3d\x60\xc2\xd7\x4e\xfd\x0f\x55\xff\x4b\xb8\x4e\xa1\x7f\x2e\xd1\x9c\xa1\xc4\xd2\x14\xd3\x4a\x33\x6a\x2d\x05\x6d\x9c\x62\x94\x58\x49\x39\x8d\x63\x00\xc0\x06\x54\x2a\xba\x4b\x5c\xc7\x17\x31\x5e\xa0\xc4\x62\x99\x95\xa0\x79\x96\x70\x9d\x95\xeb\x56\xea\x2c\xfa\xbd\x40\x39\x3b\x06\x60\xb3\xa9\x7b\xd1\x14\xcd\xbf\xd9\xa9\x74\x12\x0f\xf0\x41\x74\x98\x1f\x3c\xf5\x63\x35\xfb\xa6\x30\x35\xc4\x9b\x24\x30\x1a\x3a\x4d\x09\xdb\x2c\x19\x92\x97\x1a\xf0\x35\xa9\x8b\x71\x16\x05\x3d\xb8\x7a\x7d\x4d\x89\x61\xa4\xcd\x7d\x4c\x56\x1c\xfa\x9b\x51\xdb\xf2\x93\xfa\x82\x84\xf6\xa3\x51\x1a\xed\x32\x15\xe8\x88\xda\x54\x3a\x6b\x60\x71\x92\x7c\xa5\x3b\xab\x5d\x55\xe8\x34\x01\x6f\xdb\x93\xec\x06\xfc\xd3\xb0\xaa\x94\x35\x58\x11\x9d\x0d\x9c\x6e\x8f\xcd\x95\x3d\x2f\x6b\x77\x89\x5b\xdb\x65\xef\x82\x50\x8f\xaa\x8d\x9d\x2f\x06\x7a\xe5\xc5\xfd\x63\x47\x28\x36\x76\xab\x0d\xb6\x1c\x9a\xd9\xe7\x80\x3d\xbb\x54\xb8\x6f\x97\x0a\xab\x5d\xaa\x4a\xe3\x0d\x9b\x54\xd8\xcd\x15\xed\xcc\xbc\x49\x85\x41\x84\xd5\x26\x55\x56\xc9\xe9\xc9\xb6\x01\x9c\xaa\x00\xd4\xda\xb9\xc2\xdd\x9d\x2b\xb9\x43\x35\x87\xab\x45\xf6\x38\xce\x5d\xb9\x7f\x39\xc6\xae\x38\xff\x33\x26\x2e\x57\xbf\xf1\x72\x13\xa1\x7f\x2e\x33\xca\xac\xd2\xa5\x59\xf3\xe8\xff\xfd\x6f\x00\x00\x00\xff\xff\x29\xf0\x06\xb5\xd6\x4f\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/std.js"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
