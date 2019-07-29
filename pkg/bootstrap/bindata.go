// Code generated for package bootstrap by go-bindata DO NOT EDIT. (@generated)
// sources:
// data/cluster_role.yaml
// data/cluster_role_binding.yaml
// data/crds/crd_v1alpha1_crdconfig_cr-allowall.yaml
// data/crds/crd_v1alpha1_crdconfig_cr-denyall.yaml
// data/crds/crd_v1alpha1_crdconfig_crd.yaml
// data/crds/crd_v1alpha1_crdrequest_cr.yaml
// data/crds/crd_v1alpha1_crdrequest_crd.yaml
// data/namespace.yaml
// data/operator.yaml
// data/role.yaml
// data/role_binding.yaml
// data/service_account.yaml
package bootstrap

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cluster_roleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x4b\x04\x31\x0c\x85\xef\xfd\x15\xb9\x2d\x08\x5b\xf1\x26\xbd\x7a\xf0\x2e\xe2\x3d\xd3\xc9\x6a\x98\xb6\xa9\x49\xb3\x88\xbf\x5e\x46\x66\x60\x61\xf6\xf8\x1e\x5f\xfb\xf1\x82\x9d\x3f\x48\x8d\xa5\x25\xd0\x09\x73\x44\x1f\x5f\xa2\xfc\x8b\x83\xa5\xc5\xe5\xd9\x22\xcb\xe3\xf5\x29\x2c\xdc\xe6\x04\x2f\xc5\x6d\x90\xbe\x49\xa1\x50\x69\xe0\x8c\x03\x53\x00\xc8\x4a\xff\x0f\xde\xb9\x92\x0d\xac\x3d\x41\xf3\x52\x02\x40\xc3\x4a\x09\x16\xd2\x89\xd4\x6d\x2b\xac\x63\xbe\x6d\xd5\x0b\x59\x0a\x67\xc0\xce\xaf\x2a\xde\x6d\xfd\xf4\x0c\x59\xe7\xb8\x43\x91\x25\x00\x28\x99\xb8\x66\xda\x80\xd3\xc3\x69\x07\xb3\xb4\x0b\x7f\xda\x1e\x95\xbe\x9d\x6c\xac\xf9\x4a\x3a\xdd\xf0\x07\x0b\x76\xa6\x9f\x41\x6d\xbd\x82\x6d\x93\x8f\xaa\xec\x36\xa4\xee\xe5\x4c\x17\x6e\xbc\x4e\xbe\x63\xf8\x0b\x00\x00\xff\xff\xce\x06\xed\x45\x55\x01\x00\x00")

func cluster_roleYamlBytes() ([]byte, error) {
	return bindataRead(
		_cluster_roleYaml,
		"cluster_role.yaml",
	)
}

func cluster_roleYaml() (*asset, error) {
	bytes, err := cluster_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cluster_role_bindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x0a\xc2\x40\x0c\x86\xf7\x7b\x8a\xbc\x80\x15\x37\xb9\x4d\x1d\xdc\x2b\xb8\xa7\xd7\xa8\xb1\xed\xe5\x48\x72\x1d\x7c\x7a\x29\x22\x08\x05\xd7\x2f\xe4\xff\xbe\x81\x73\x1f\xe1\x34\x56\x73\xd2\x56\x46\x3a\x72\xee\x39\xdf\x03\x16\xbe\x92\x1a\x4b\x8e\xa0\x1d\xa6\x06\xab\x3f\x44\xf9\x85\xce\x92\x9b\x61\x6f\x0d\xcb\x76\xde\x85\x89\x1c\x7b\x74\x8c\x01\x20\xe3\x44\x11\x06\xd2\x8e\xb4\x5a\xb0\xda\x3d\x29\xb9\x2d\xa7\x0d\x7c\x54\x17\xd2\x99\x13\x1d\x52\x92\x9a\x3d\x00\xac\xbe\xbe\xc8\x0a\xa6\x5f\xae\x32\x52\x4b\xb7\x65\x6c\x55\xbd\x76\x03\x60\xe1\xb3\x4a\x2d\x7f\xfa\xc3\x3b\x00\x00\xff\xff\x50\xe1\xe8\xab\xff\x00\x00\x00")

func cluster_role_bindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_cluster_role_bindingYaml,
		"cluster_role_binding.yaml",
	)
}

func cluster_role_bindingYaml() (*asset, error) {
	bytes, err := cluster_role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster_role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _crdsCrd_v1alpha1_crdconfig_crAllowallYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcb\xb1\x0e\xc2\x20\x10\xc6\xf1\x9d\xa7\xb8\x74\x07\xd3\xcd\xb0\x35\xf5\x09\x1c\x1c\x34\x0e\x67\xb9\xda\x0b\x57\x20\x40\x35\xfa\xf4\xa6\x8c\xae\xff\xdf\xf7\x61\xe2\x0b\xe5\xc2\x31\x58\x98\xb2\x33\x9e\xf2\x83\xf2\x56\x0c\xc7\xc3\xab\x47\x49\x0b\xf6\xca\x73\x70\x16\xc6\xf3\x69\x8c\x61\xe6\xa7\x5a\xa9\xa2\xc3\x8a\x56\x01\x04\x5c\xa9\x5d\xa7\x66\x1a\x45\xe2\x1b\x45\x54\x49\x34\xed\x83\x16\x06\x11\x0b\x35\x6f\xa4\x00\x1c\x85\xcf\x0e\x1a\xba\x1b\xea\xef\xa0\xaf\x77\x13\x13\x85\xb2\xf0\x5c\x0d\xc7\xee\xcf\xfc\xb1\xb4\xaa\x7e\x01\x00\x00\xff\xff\x05\xb5\xd0\xca\xac\x00\x00\x00")

func crdsCrd_v1alpha1_crdconfig_crAllowallYamlBytes() ([]byte, error) {
	return bindataRead(
		_crdsCrd_v1alpha1_crdconfig_crAllowallYaml,
		"crds/crd_v1alpha1_crdconfig_cr-allowall.yaml",
	)
}

func crdsCrd_v1alpha1_crdconfig_crAllowallYaml() (*asset, error) {
	bytes, err := crdsCrd_v1alpha1_crdconfig_crAllowallYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "crds/crd_v1alpha1_crdconfig_cr-allowall.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _crdsCrd_v1alpha1_crdconfig_crDenyallYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcb\xb1\xae\xc2\x30\x0c\x85\xe1\x3d\x4f\x61\x75\x4f\xae\xba\x5d\x65\xab\xca\x13\x30\x30\x80\x18\x4c\xe3\x52\xab\xae\x13\x25\x29\xa8\x3c\x3d\x6a\x47\xd6\xf3\x9d\x1f\x13\x5f\x28\x17\x8e\xea\x61\xc8\xc1\xcd\x94\x1f\x94\xd7\xe2\x38\xfe\xbd\x5a\x94\x34\x61\x6b\x66\xd6\xe0\xa1\x3f\x9f\xfa\xa8\x23\x3f\xcd\x42\x15\x03\x56\xf4\x06\x40\x71\xa1\x23\x1d\x0e\xb3\x81\x54\x37\x14\x31\x25\xd1\xb0\x1f\x02\xe9\xd6\x89\x78\xa8\x79\x25\x03\x80\x22\xf1\xbd\x83\x85\xe6\x86\xf6\xd3\xd9\xeb\xdd\xc5\x44\x5a\x26\x1e\xab\xe3\xd8\xfc\xd8\xfc\x5f\x8e\xf5\x1b\x00\x00\xff\xff\xd5\xae\x47\xdf\xab\x00\x00\x00")

func crdsCrd_v1alpha1_crdconfig_crDenyallYamlBytes() ([]byte, error) {
	return bindataRead(
		_crdsCrd_v1alpha1_crdconfig_crDenyallYaml,
		"crds/crd_v1alpha1_crdconfig_cr-denyall.yaml",
	)
}

func crdsCrd_v1alpha1_crdconfig_crDenyallYaml() (*asset, error) {
	bytes, err := crdsCrd_v1alpha1_crdconfig_crDenyallYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "crds/crd_v1alpha1_crdconfig_cr-denyall.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _crdsCrd_v1alpha1_crdconfig_crdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xcb\x8e\x23\xb7\x0e\xdd\xfb\x2b\x08\xdc\xc5\x6c\xc6\x36\x1a\x77\x13\xd4\x6e\xe0\x49\x80\x41\x32\x41\xd0\xdd\x98\x3d\x2d\xd1\xb6\xd2\x2a\xa9\x42\x52\xee\x38\x41\xfe\x3d\xa0\xca\x8f\xb2\x5d\xee\x71\x5e\x5a\x18\x30\x49\x1d\x92\x87\x0f\x15\x76\xe1\x0b\xb1\x84\x9c\x1a\xc0\x2e\xd0\xaf\x4a\xc9\xfe\xc9\xec\xe5\x1b\x99\x85\x3c\xdf\x3e\x2c\x49\xf1\x61\xf2\x12\x92\x6f\x60\x51\x44\x73\xfb\x48\x92\x0b\x3b\xfa\x48\xab\x90\x82\x86\x9c\x26\x2d\x29\x7a\x54\x6c\x26\x00\x09\x5b\x6a\xc0\xb1\x77\x39\xad\xc2\x5a\x66\x8e\xfd\xec\x85\x78\x49\x5c\x0c\x73\x22\x1d\x39\x33\x5c\x73\x2e\x5d\xb5\x3c\x53\xf7\x08\x62\x16\x00\x7b\xbf\x8f\x1f\x17\x15\xac\xca\x62\x10\xfd\xfe\x5c\xfe\x43\x10\xad\xba\x2e\x16\xc6\x38\x74\x5f\xc5\x12\xd2\xba\x44\xe4\x81\x62\x02\x20\x2e\x77\xd4\xc0\x8f\xe6\xae\x43\x47\xde\x64\x65\xc9\xfb\xfc\xf6\x21\x88\xa2\x16\x69\xe0\xf7\x3f\x26\x00\x5b\x8c\xc1\xa3\xe5\xdc\x2b\x73\x47\xe9\xc3\x4f\x9f\xbe\xfc\xff\xc9\x6d\xa8\xc5\x5e\x08\xd0\x71\xee\x88\x35\x1c\x30\xec\x0c\xb8\x3e\xca\x00\x3c\x89\xe3\xd0\x55\x44\x78\x67\x50\xbd\x0d\x78\x63\x97\x04\x74\x43\xb0\xed\x65\xe4\x41\xaa\x1b\xc8\x2b\xd0\x4d\x10\x60\xea\x98\x84\x92\xd6\x90\x06\xb0\x60\x26\x98\x20\x2f\x7f\x26\xa7\x33\x78\x22\x36\x10\x90\x4d\x2e\xd1\x83\xcb\x69\x4b\xac\xc0\xe4\xf2\x3a\x85\xdf\x8e\xc8\x02\x9a\xab\xcb\x88\x4a\x7b\x4e\x0f\x27\x24\x25\x4e\x18\x8d\x84\x42\xef\x01\x93\x87\x16\x77\xc0\x64\x3e\xa0\xa4\x01\x5a\x35\x91\x19\x7c\xce\x4c\x10\xd2\x2a\x37\xb0\x51\xed\xa4\x99\xcf\xd7\x41\x0f\xdd\xe5\x72\xdb\x96\x14\x74\x37\x77\x39\x29\x87\x65\xd1\xcc\x32\xf7\xb4\xa5\x38\xc7\x2e\x4c\x6b\x9c\x49\x6b\x47\xb6\xfe\x7f\xc7\xca\xbc\x1b\x04\xa6\x3b\x2b\xa2\x28\x87\xb4\x3e\x8a\x6b\xdf\xdc\xa4\xd9\xba\x07\x82\x00\xee\xaf\xf5\xe1\x9e\xd8\x34\x91\x91\xf0\xf8\xed\xd3\x33\x1c\x9c\x56\xc6\xcf\x29\xae\xe4\x9e\xae\xc9\x89\x67\xe3\x25\xa4\x15\x71\x5f\xa7\x15\xe7\xb6\x22\x52\xf2\x5d\x0e\x49\xeb\x1f\x17\x03\xa5\x73\x8e\xa5\x2c\xdb\xa0\x56\xd8\x5f\x0a\x89\x5a\x39\x66\xb0\xc0\x94\xb2\xc2\x92\xa0\x74\x1e\x95\xfc\x0c\x3e\x25\x58\x60\x4b\x71\x81\x42\xff\x36\xcb\x46\xa8\x4c\x8d\xc1\xaf\xf3\x3c\x1c\xfc\x73\xc3\x9e\x9c\xa3\xf8\x30\xf4\x87\x33\x36\x21\x75\x4a\x62\xcc\xaf\xe7\x22\x80\xa0\xd4\xca\xa5\xf0\x46\x44\x43\x15\x32\xe3\xee\x1a\xfe\x43\x8c\x97\x60\x67\xfd\x11\x56\x47\x3b\xeb\x12\x21\x7d\x0f\xaf\xd4\xcb\xec\x17\x96\x45\xc1\x53\xda\xd5\x55\x34\xea\x79\x99\x73\x24\x3c\x9f\x48\xbb\xf1\xdf\x65\x66\xe8\x77\x24\xb6\x37\x1b\xe6\x55\x33\x39\xa4\xd5\x27\xf9\x17\xf2\x1a\xaf\x76\xbf\x33\xef\xa8\xb7\xcb\xc9\xd7\x37\xe4\x8a\x85\xb3\xc8\x6d\xc1\xdb\x4a\x3b\x99\xc3\x2a\x73\x9d\xa2\xc5\xe3\x7d\x9c\xde\x8a\xa0\x3f\x11\x45\x9f\x19\x93\x54\xf4\xe7\xd0\xd2\x98\xd5\x65\x54\x57\x97\x8c\xd7\x7e\x7d\x8a\x82\x9a\xa0\xce\xff\x31\xec\x51\x4c\x80\x57\x14\x40\xa7\x61\x4b\xb7\xd3\xea\xcf\x2a\x73\x8b\xda\x80\xed\x81\xa9\x39\x18\xb5\x7a\xa3\x7f\xa0\x4e\xad\x08\xae\xef\xc9\xf0\x73\x6f\x09\x4c\xf6\x20\xf8\xc3\xeb\xf0\x76\x3e\x5f\xf1\x6e\xcd\x71\x8f\xef\x27\xb3\x03\xdd\xa0\x5e\x90\x38\xd8\xb8\x7f\xd7\x7f\x19\x6d\x82\x8b\x00\xbe\x8b\xb8\xb6\x8c\x43\xf2\xc1\x59\x2c\x61\xd0\x80\x7b\x18\x2b\xb8\x2b\xcc\x94\x34\xee\x6e\x54\xb7\xaf\xec\x1b\xa1\x8e\x6d\x0b\x3b\xf6\x0a\x04\x26\x7f\x1d\xeb\x74\xa4\x61\x47\x8c\xfa\x18\x6f\x28\xae\x2f\x8c\xce\xf2\x50\x75\xbd\x76\x46\x8b\xf9\x4f\xcb\x78\xa3\x80\x63\x74\x5c\xa7\x72\x91\xc4\xf6\xf0\x7d\xbb\x7d\xc0\xd8\x6d\xf0\xe1\x24\xab\x3d\x30\xdd\x7f\xb0\x0e\xd4\x00\x62\xcf\xb8\x6f\x40\xb9\xd0\xfe\x23\x30\xb3\x4d\x4c\x2f\xf9\x33\x00\x00\xff\xff\xb4\x7b\xa8\x81\x37\x0b\x00\x00")

func crdsCrd_v1alpha1_crdconfig_crdYamlBytes() ([]byte, error) {
	return bindataRead(
		_crdsCrd_v1alpha1_crdconfig_crdYaml,
		"crds/crd_v1alpha1_crdconfig_crd.yaml",
	)
}

func crdsCrd_v1alpha1_crdconfig_crdYaml() (*asset, error) {
	bytes, err := crdsCrd_v1alpha1_crdconfig_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "crds/crd_v1alpha1_crdconfig_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _crdsCrd_v1alpha1_crdrequest_crYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x90\xc1\x6e\xea\x40\x0c\x45\xf7\xf9\x0a\xff\x00\x79\xca\xee\x29\xbb\x0a\x96\x15\x0b\x54\x75\xef\x4c\x0c\x8c\x48\xc6\xae\xed\xa1\x7c\x7e\x15\x08\x21\x54\xd0\xee\x3a\xbb\xf8\x5e\xd9\x27\x07\x25\xbe\x93\x5a\xe4\x54\x43\xd0\xb6\x3c\x90\x36\xa4\xd9\xca\xc8\xff\x8e\x15\x76\xb2\xc7\xaa\x38\xc4\xd4\xd6\xb0\xdc\xac\x36\xf4\x91\xc9\xbc\xe8\xc9\xb1\x45\xc7\xba\x00\x48\xd8\x53\x0d\x74\xc2\x5e\x3a\x5a\x04\x6d\x75\x2c\x99\x50\x18\x0a\xfc\x99\xa8\xad\xc1\x35\x53\x01\xc3\x91\xd7\x68\x3e\x04\x0b\x98\x5f\x47\x89\x74\x72\x4a\xc3\x97\x95\x87\xff\x23\x42\x43\x8e\x55\x01\x00\x30\x52\x64\x73\xee\x37\x64\x9c\x35\xd0\x8a\xb6\x31\x45\x8f\x9c\xce\x95\x39\xd7\xf0\x2e\x6c\x41\x39\x39\x36\x56\x9a\x63\xd3\x51\xc9\x42\xc9\xf6\x71\xeb\x65\xe4\x73\xf1\x4a\x3a\xbc\x9d\x72\x96\x1a\x9e\x55\x01\x8e\x57\xe0\x63\x05\xe3\xc8\x02\x0b\xd5\xb0\xc6\x9e\x4c\x30\x50\x0b\xb3\xfb\x76\xdd\x0c\x20\x5d\x56\xec\x6e\x40\x30\x25\x16\xd3\x2e\x77\xa8\x53\x76\x8b\xc6\xdf\x56\x4e\x6f\xf3\xb1\xed\x59\x7d\x7d\xbf\x7f\x01\xc1\xe1\x0f\xc5\xa2\x48\xe0\xb4\x8d\xbb\x49\xed\x65\xfb\x2f\x52\x67\xa5\x3b\x9d\xcf\x6c\xfe\x2c\xf3\x06\xf1\xc0\xe6\x14\x7e\xd3\xf9\x22\xb2\xbc\x9f\x3f\xf6\x89\xa1\xf8\x0a\x00\x00\xff\xff\x15\x87\x63\x26\x22\x03\x00\x00")

func crdsCrd_v1alpha1_crdrequest_crYamlBytes() ([]byte, error) {
	return bindataRead(
		_crdsCrd_v1alpha1_crdrequest_crYaml,
		"crds/crd_v1alpha1_crdrequest_cr.yaml",
	)
}

func crdsCrd_v1alpha1_crdrequest_crYaml() (*asset, error) {
	bytes, err := crdsCrd_v1alpha1_crdrequest_crYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "crds/crd_v1alpha1_crdrequest_cr.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _crdsCrd_v1alpha1_crdrequest_crdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xc1\x8e\x23\x45\x0c\xbd\xe7\x2b\x2c\x71\xd8\x0b\x93\x68\xc4\x05\xf5\x0d\x65\x41\x5a\xc1\x22\x94\x19\xed\xdd\xa9\x72\x92\x62\xba\xab\x1a\xdb\x15\x08\x88\x7f\x47\xae\x4a\x27\x9d\xa4\xb3\x33\xc0\xd6\x2d\xb6\xcb\x7e\x7e\x7e\xae\x34\xf6\xe1\x13\xb1\x84\x14\x1b\xc0\x3e\xd0\x1f\x4a\xd1\x7e\xc9\xfc\xe5\x5b\x99\x87\xb4\xd8\x3f\xae\x49\xf1\x71\xf6\x12\xa2\x6f\x60\x99\x45\x53\xb7\x22\x49\x99\x1d\xbd\xa7\x4d\x88\x41\x43\x8a\xb3\x8e\x14\x3d\x2a\x36\x33\x80\x88\x1d\x35\xe0\xd8\x33\xfd\x96\x49\x54\xe6\x8e\xfd\xfc\x85\x78\x4d\x9c\x2d\xe9\x4c\x7a\x72\x16\xb9\xe5\x94\xfb\x12\x7a\xe1\xae\x29\xc4\x22\x00\x8e\x85\x57\xef\x57\x35\x5b\x31\xb6\x41\xf4\xc7\x2b\xc7\x4f\xe1\xe8\xec\xdb\xcc\xd8\x5e\x20\x28\x76\x09\x71\x9b\x5b\xe4\xb1\x67\x06\x20\x2e\xf5\xd4\xc0\xcf\x56\xb2\x47\x47\xde\x6c\x79\xcd\xc7\x26\x8f\x30\x44\x51\xb3\x34\xf0\xd7\xdf\x33\x80\x3d\xb6\xc1\xa3\x35\x5e\x9d\xa9\xa7\xf8\xdd\x2f\x1f\x3e\x7d\xf3\xe4\x76\xd4\x61\x35\x02\xf4\x9c\x7a\x62\x0d\x43\x0e\x3b\x23\xc2\x4f\x36\x00\x4f\xe2\x38\xf4\x25\x23\xbc\xb3\x54\x35\x06\xbc\x51\x4c\x02\xba\x23\xd8\x57\x1b\x79\x90\x52\x06\xd2\x06\x74\x17\x04\x98\x7a\x26\xa1\xa8\x05\xd2\x28\x2d\x58\x08\x46\x48\xeb\x5f\xc9\xe9\x1c\x9e\x88\x2d\x09\xc8\x2e\xe5\xd6\x83\x4b\x71\x4f\xac\xc0\xe4\xd2\x36\x86\x3f\x4f\x99\x05\x34\x95\x92\x2d\xea\x40\xf9\x70\x42\x54\xe2\x88\xad\x91\x90\xe9\x6b\xc0\xe8\xa1\xc3\x03\x30\x59\x0d\xc8\x71\x94\xad\x84\xc8\x1c\x3e\x26\x26\x08\x71\x93\x1a\xd8\xa9\xf6\xd2\x2c\x16\xdb\xa0\x83\xc4\x5c\xea\xba\x1c\x83\x1e\x16\x2e\x45\xe5\xb0\xce\x9a\x58\x16\x9e\xf6\xd4\x2e\xb0\x0f\x0f\x05\x67\xd4\x22\xcb\xce\x7f\x75\x9a\xcc\xbb\x11\x30\x3d\xd8\x10\x45\x39\xc4\xed\xc9\x5c\xb4\x73\x97\x66\x13\x10\x04\x01\x3c\x5e\xab\x70\xcf\x6c\x9a\xc9\x48\x58\x7d\xff\xf4\x0c\x43\xd1\xc2\xf8\x25\xc5\x85\xdc\xf3\x35\x39\xf3\x6c\xbc\x84\xb8\x21\xae\x73\xda\x70\xea\x4a\x46\x8a\xbe\x4f\x21\x6a\xf9\xe1\xda\x40\xf1\x92\x63\xc9\xeb\x2e\xa8\x0d\xb6\xca\x17\x34\xcd\x61\x89\x31\x26\x85\x35\x41\xee\x3d\x2a\xf9\x39\x7c\x88\xb0\xc4\x8e\xda\x25\x0a\x7d\x69\x96\x8d\x50\x79\x30\x06\x5f\xe7\x79\xbc\xfd\x97\x81\x95\x9c\x93\x79\x58\xfc\xe1\x4c\x6d\x88\x1d\xc7\xde\xd6\xf9\xd2\x08\x10\x94\x3a\xb9\x36\xde\x29\x35\x76\x21\x33\x1e\x5e\xc5\x56\x37\xfc\x2d\xe8\x52\xf4\xe5\xd9\xbb\xc1\x72\x21\x30\xeb\xc0\x16\xf0\x1c\x0e\x9b\xc4\x65\xe6\xcb\xd5\xdb\x3a\xbb\x87\xa0\x9e\x16\x45\x9f\x19\xa3\x94\xec\xcf\xa1\xa3\xa9\xa8\x6b\x54\x37\x97\x6c\x07\xea\xb2\x8b\x82\x9a\xa1\xa8\xf5\x04\x7b\x32\x27\xc0\xef\x28\x80\x4e\xc3\x9e\xee\xb7\x55\xcf\x26\x71\x87\xda\x80\xa9\xf6\xc1\x0a\x4c\x46\x4d\x2a\xeb\x7c\x3a\x12\xc1\xed\x5b\x3a\xfc\x58\x23\x81\xc9\x9e\x2f\x3f\xbc\x65\x9f\xef\xe7\x95\xea\x26\x8e\xb7\xd4\x7e\xb2\x38\xd0\x1d\xea\x15\x89\xa3\xf7\xe1\xbf\xd6\xcf\x93\x22\xb8\x02\xf0\x43\x8b\x5b\xeb\x38\x44\x1f\x9c\x61\x09\x23\x01\x1e\xd3\xd8\xc0\x5d\x66\xa6\xa8\xed\xe1\xce\x74\xeb\x64\x3f\x03\x75\x9d\x52\x4b\x78\x4b\xa6\xbd\x59\x81\xc9\xdf\x62\x7d\x98\x10\xec\x44\x50\xc5\x78\xc7\x71\x7b\xe1\x5f\x2f\xff\x9d\x61\xfe\xdf\x31\xde\x19\xe0\x14\x1d\xb7\xad\x5c\x35\xb1\x1f\x3e\xc9\xf6\x8f\xd8\xf6\x3b\x7c\x3c\xdb\x8a\x06\x1e\x8e\xdf\x58\x23\x37\x80\xd8\x9f\x8e\x6f\x40\x39\xd3\xf1\x93\x25\xb1\x6d\x4c\xb5\xfc\x13\x00\x00\xff\xff\xef\xec\x6c\xd1\xea\x09\x00\x00")

func crdsCrd_v1alpha1_crdrequest_crdYamlBytes() ([]byte, error) {
	return bindataRead(
		_crdsCrd_v1alpha1_crdrequest_crdYaml,
		"crds/crd_v1alpha1_crdrequest_crd.yaml",
	)
}

func crdsCrd_v1alpha1_crdrequest_crdYaml() (*asset, error) {
	bytes, err := crdsCrd_v1alpha1_crdrequest_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "crds/crd_v1alpha1_crdrequest_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc8\xb1\x0d\x02\x31\x0c\x05\xd0\xde\x53\x78\x01\x0a\x5a\x0f\x41\x49\xff\xef\xee\x23\x59\x21\x26\xb2\x13\x0a\xa6\x47\x29\xdf\xc3\xf0\x27\xb3\xfc\x13\xa6\xdf\xbb\x34\x8f\xcb\xf4\x81\xce\x1a\x38\x29\x9d\x13\x17\x26\x4c\x54\x03\x9d\xa6\x8d\x79\x30\x57\x49\x0d\x9e\xbb\x5f\x1e\x78\xfb\x8f\x59\x5b\x37\x6d\xeb\x60\x06\x27\x4b\xfe\x01\x00\x00\xff\xff\x51\xf1\x1e\xc1\x5d\x00\x00\x00")

func namespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_namespaceYaml,
		"namespace.yaml",
	)
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _operatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xcd\x6e\xc2\x30\x10\x84\xef\x79\x8a\x15\x3d\x17\xc4\xd5\xb7\x54\xd0\x53\x4b\xa2\x14\xf5\x5a\x2d\xce\xd2\x58\xf8\xaf\xf6\x26\x28\x6f\x5f\x25\xfc\xc8\x40\xf0\xc9\x9a\x9d\x6f\x3c\xab\x04\xbd\xfa\xa6\x10\x95\xb3\x02\xd0\xfb\xb8\xe8\x96\xd9\x41\xd9\x5a\xc0\x8a\xbc\x76\xbd\x21\xcb\x99\x21\xc6\x1a\x19\x45\x06\x60\xd1\x90\x80\x03\x85\x1d\x85\x36\x9e\x85\xe8\x51\xa6\x6a\xf4\x24\x07\x73\x20\xaf\x95\xc4\x28\x60\x99\x01\x44\xd2\x24\xd9\x85\x61\x02\x60\x90\x65\xf3\x81\x3b\xd2\xf1\x24\x4c\x64\x33\x19\xaf\x91\xe9\x4c\x24\x35\x86\xa3\x6f\xe0\x09\x1c\xe0\x52\x64\xbc\x53\xe8\x94\xa4\x5c\x4a\xd7\x5a\xde\x3c\x9a\x01\xa4\xb3\x8c\xca\x52\x48\x52\x5f\xa7\x72\x4f\xe7\x05\x2a\xf2\x1a\x25\x01\x37\x2a\xc2\x51\x71\x03\xdc\x10\xec\x5a\xa5\x19\x94\xc1\x5f\x1a\xe1\x04\x19\x45\x01\x7f\x2d\xf6\x73\xe5\x16\x97\xd0\xeb\x45\x0c\xeb\x46\x4e\x08\xe9\x8c\x41\x5b\x8b\x44\x1a\x4a\x4d\xd4\x19\xb3\xcb\x56\xeb\xd2\x69\x25\x7b\x01\xb9\x3e\x62\x9f\x3a\xc8\x76\xf7\x39\xa7\xe5\x8a\x72\x5d\xe5\xdb\xa2\xfa\xd9\xe4\x9f\xeb\x1b\x07\x40\x87\xba\x25\x01\xb3\xcb\x8b\xb3\xc9\x84\xb7\xa2\xd8\x7e\x6d\xab\xbc\x7c\x42\x73\x68\x69\x9a\x2c\x8b\xd5\xd3\x67\xdf\x83\x33\xe2\x6e\x00\xb0\x57\xa4\xeb\x8a\xf6\x8f\x93\xf3\xac\x44\x6e\xc4\xf5\x7f\x99\x8f\xdf\xe0\x3f\x00\x00\xff\xff\xdc\xed\xb9\x22\xea\x02\x00\x00")

func operatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_operatorYaml,
		"operator.yaml",
	)
}

func operatorYaml() (*asset, error) {
	bytes, err := operatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "operator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _roleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\xc1\x0a\xc2\x30\x0c\xc6\xf1\x7b\x9f\x22\xec\x32\x10\x36\xf1\x26\x7d\x01\xef\x22\xde\xb3\x2e\x6a\x58\xdb\x94\xa4\xdd\xc1\xa7\x97\x8d\x1d\xbc\x25\x3f\xfe\x7c\x58\xf8\x49\x6a\x2c\xd9\x83\x4e\x18\x46\x6c\xf5\x23\xca\x5f\xac\x2c\x79\x5c\xae\x36\xb2\x9c\xd7\x8b\x5b\x38\xcf\x1e\xee\x12\xc9\x25\xaa\x38\x63\x45\xef\x00\x82\xd2\x5e\x3e\x38\x91\x55\x4c\xc5\x43\x6e\x31\x3a\x80\x8c\x89\x3c\x2c\xa4\x13\x69\xb3\x03\xac\x60\xf8\x57\x6d\x91\xcc\xbb\x01\xb0\xf0\x4d\xa5\x15\xdb\x46\x07\xe8\x3a\x07\xa0\x64\xd2\x34\xd0\x61\x45\x66\xdb\x8f\x20\xf9\xc5\xef\x84\x65\x7b\x57\xd2\xe9\x08\xfa\x53\xef\x7e\x01\x00\x00\xff\xff\xcd\x85\xf1\xed\xd0\x00\x00\x00")

func roleYamlBytes() ([]byte, error) {
	return bindataRead(
		_roleYaml,
		"role.yaml",
	)
}

func roleYaml() (*asset, error) {
	bytes, err := roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _role_bindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x0a\xc2\x40\x0c\x86\xf7\x7b\x8a\xbc\x80\x15\x37\xb9\x4d\x17\xf7\x0a\xee\xe9\x35\x6a\x6c\x9b\x1c\xb9\x5c\x07\x9f\x5e\x8a\x88\x42\xc1\xf5\x0b\xf9\xbf\x6f\x60\xe9\x23\xb4\x3a\xd2\x91\xa5\x67\xb9\x05\xcc\x7c\x21\x2b\xac\x12\xc1\x3a\x4c\x0d\x56\xbf\xab\xf1\x13\x9d\x55\x9a\x61\x5f\x1a\xd6\xed\xbc\x0b\x13\x39\xf6\xe8\x18\x03\x80\xe0\x44\x11\x06\xb2\x8e\xac\x96\x50\x6a\xf7\xa0\xe4\x65\x39\x6d\xe0\xed\x38\x93\xcd\x9c\xe8\x90\x92\x56\xf1\x00\xb0\xfa\xfa\xa0\x92\x31\xfd\x72\xd3\x91\x5a\xba\x2e\x63\xdf\xdc\xb5\x14\x00\x33\x9f\x4c\x6b\xfe\x13\x1e\x5e\x01\x00\x00\xff\xff\x8e\xd6\x11\xff\xf1\x00\x00\x00")

func role_bindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_role_bindingYaml,
		"role_binding.yaml",
	)
}

func role_bindingYaml() (*asset, error) {
	bytes, err := role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _service_accountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xc8\x2b\x0e\x80\x30\x0c\x06\x60\xbf\x53\xf4\x0a\xd8\x3a\xce\x40\x82\x2f\xdd\x2f\x9a\x65\xdd\xd2\x3d\xce\x8f\x41\x60\x3f\xe9\x76\x23\x86\x35\x67\xda\x47\x2a\xe6\x99\xe9\x42\x6c\x53\x9c\xaa\x6d\xf9\x4c\x15\x53\xb2\x4c\xe1\x44\xe4\x52\xc1\x54\x10\x0f\x62\x8d\x0f\x46\x17\xfd\xeb\x1b\x00\x00\xff\xff\x1a\xd4\xba\xd2\x55\x00\x00\x00")

func service_accountYamlBytes() ([]byte, error) {
	return bindataRead(
		_service_accountYaml,
		"service_account.yaml",
	)
}

func service_accountYaml() (*asset, error) {
	bytes, err := service_accountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service_account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"cluster_role.yaml":                            cluster_roleYaml,
	"cluster_role_binding.yaml":                    cluster_role_bindingYaml,
	"crds/crd_v1alpha1_crdconfig_cr-allowall.yaml": crdsCrd_v1alpha1_crdconfig_crAllowallYaml,
	"crds/crd_v1alpha1_crdconfig_cr-denyall.yaml":  crdsCrd_v1alpha1_crdconfig_crDenyallYaml,
	"crds/crd_v1alpha1_crdconfig_crd.yaml":         crdsCrd_v1alpha1_crdconfig_crdYaml,
	"crds/crd_v1alpha1_crdrequest_cr.yaml":         crdsCrd_v1alpha1_crdrequest_crYaml,
	"crds/crd_v1alpha1_crdrequest_crd.yaml":        crdsCrd_v1alpha1_crdrequest_crdYaml,
	"namespace.yaml":                               namespaceYaml,
	"operator.yaml":                                operatorYaml,
	"role.yaml":                                    roleYaml,
	"role_binding.yaml":                            role_bindingYaml,
	"service_account.yaml":                         service_accountYaml,
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
	"cluster_role.yaml":         {cluster_roleYaml, map[string]*bintree{}},
	"cluster_role_binding.yaml": {cluster_role_bindingYaml, map[string]*bintree{}},
	"crds": {nil, map[string]*bintree{
		"crd_v1alpha1_crdconfig_cr-allowall.yaml": {crdsCrd_v1alpha1_crdconfig_crAllowallYaml, map[string]*bintree{}},
		"crd_v1alpha1_crdconfig_cr-denyall.yaml":  {crdsCrd_v1alpha1_crdconfig_crDenyallYaml, map[string]*bintree{}},
		"crd_v1alpha1_crdconfig_crd.yaml":         {crdsCrd_v1alpha1_crdconfig_crdYaml, map[string]*bintree{}},
		"crd_v1alpha1_crdrequest_cr.yaml":         {crdsCrd_v1alpha1_crdrequest_crYaml, map[string]*bintree{}},
		"crd_v1alpha1_crdrequest_crd.yaml":        {crdsCrd_v1alpha1_crdrequest_crdYaml, map[string]*bintree{}},
	}},
	"namespace.yaml":       {namespaceYaml, map[string]*bintree{}},
	"operator.yaml":        {operatorYaml, map[string]*bintree{}},
	"role.yaml":            {roleYaml, map[string]*bintree{}},
	"role_binding.yaml":    {role_bindingYaml, map[string]*bintree{}},
	"service_account.yaml": {service_accountYaml, map[string]*bintree{}},
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
