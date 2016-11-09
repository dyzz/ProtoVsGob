//go:generate protoc --go_out=. *.proto
package main

import (
	"bytes"
	"encoding/gob"
	"github.com/golang/protobuf/proto"
)

func protoUnmarshall(raw []byte) (*Track, error) {
	track := new(Track)
	err := proto.Unmarshal(raw, track)
	if err != nil {
		return nil, err
	} else {
		return track, nil
	}
}

func protoMarshall(track *Track) ([]byte, error) {
	return proto.Marshal(track)
}

func gobUnmarshall(raw []byte) (*Track, error) {
	buf := bytes.NewReader(raw)
	track := new(Track)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(track)
	if err != nil {
		return nil, err
	} else {
		return track, nil
	}
}

func gobMarshall(track *Track) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(track)
	if err != nil {
		return nil, err
	} else {
		return buf.Bytes(), nil
	}
}

func main() {
}
