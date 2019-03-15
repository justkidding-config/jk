package std

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/unicode"

	"github.com/jkcfg/jk/pkg/__std"

	"github.com/ghodss/yaml"
	yamlclassic "gopkg.in/yaml.v2"
)

func readYAML(in io.Reader) ([]byte, error) {
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	return yaml.YAMLToJSON(bytes)
}

func readYAMLStream(in io.Reader) ([]byte, error) {
	// This specific one is tricky: since YAML can have more than just
	// strings as map keys, the decoder may return a
	// `map[interface{}]interface{}`, which can't be directly encoded
	// in JSON.
	//
	// ghodss/yaml has a conversion function, from YAML bytes to JSON
	// butes -- so we have to _un_parse the value we just got, to put
	// it through the conversion.
	decoder := yamlclassic.NewDecoder(in)
	var yamlvalues []interface{}
	for {
		var v interface{}
		err := decoder.Decode(&v)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		yamlvalues = append(yamlvalues, v)
	}
	yamlbytes, err := yamlclassic.Marshal(yamlvalues)
	if err != nil {
		return nil, err
	}
	return yaml.YAMLToJSON(yamlbytes)
}

func readJSON(in io.Reader) ([]byte, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(in, &buf)
	var obj interface{}
	if err := json.NewDecoder(tee).Decode(&obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func readJSONStream(in io.Reader) ([]byte, error) {
	decoder := json.NewDecoder(in)
	var items []interface{}
	for {
		var v interface{}
		err := decoder.Decode(&v)
		if err == nil {
			items = append(items, v)
			continue
		}
		if err == io.EOF {
			break
		}
		return nil, err
	}
	return json.Marshal(items)
}

func readRaw(in io.Reader) ([]byte, error) {
	return ioutil.ReadAll(in)
}

type readFunc func(io.Reader) ([]byte, error)

func readerByPath(path string) readFunc {
	switch filepath.Ext(path) {
	case ".yml", ".yaml":
		return readYAML
	case ".json":
		return readJSON
	}
	return readJSON
}

// Read returns the contents of the file at `path`, relative to the
// root path or if given, the module directory identified by `module`.
func (r ReadBase) Read(path string, format __std.Format, encoding __std.Encoding, module string) ([]byte, error) {
	// Special case for reading from stdin
	if path == "" {
		return read("", format, encoding)
	}

	base, path, err := r.getPath(path, module)
	if err != nil {
		return nil, err
	}
	return read(filepath.Join(base, path), format, encoding)
}

func read(path string, format __std.Format, encoding __std.Encoding) ([]byte, error) {
	var reader readFunc = readRaw

	if encoding == __std.EncodingJSON {
		switch format {
		case __std.FormatAuto:
			reader = readerByPath(path)
		case __std.FormatYAML:
			reader = readYAML
		case __std.FormatYAMLStream:
			reader = readYAMLStream
		case __std.FormatJSON:
			reader = readJSON
		case __std.FormatJSONStream:
			reader = readJSONStream
		default:
			reader = readJSON
		}
	}

	var in io.Reader
	if path == "" {
		in = os.Stdin
	} else {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		in = f
	}

	bytes, err := reader(in)
	if err != nil {
		return nil, err
	}

	switch encoding {
	case __std.EncodingBytes:
		break
	case __std.EncodingString, __std.EncodingJSON:
		encoder := unicode.UTF16(NativeEndian, unicode.IgnoreBOM).NewEncoder()
		bytes, err = encoder.Bytes(bytes)
	}
	return bytes, err
}
