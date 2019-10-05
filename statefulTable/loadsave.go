package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)


// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func() (io.Reader, error) {
	b, err := json.MarshalIndent(PTable, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (potatoTable PotatoTable) Save(path string) error {
		potatoTable.Lock.Lock()
		defer potatoTable.Lock.Unlock()

		// Create File
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		r, err := Marshal()
		if err != nil {
			return err
		}
		_, err = io.Copy(f, r)
		return err
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
	PTable.Lock.Lock()
	defer PTable.Lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}



