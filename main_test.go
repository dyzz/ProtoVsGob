package main

import (
	"testing"
)

func genBytes(len int) []byte {
	ret := make([]byte, len)
	for i := range ret {
		ret[i] = byte(i)
	}
	return ret
}

func genTrack() *Track {
	track := new(Track)
	track.X = 1
	track.Y = 2
	track.XTrack = genBytes(1000)
	track.YTrack = genBytes(2000)
	return track
}

func BenchmarkProtoMarshall(b *testing.B) {
	track := genTrack()
	for n := 0; n < b.N; n++ {
		protoMarshall(track)
	}
}

func BenchmarkProtoUnmarshall(b *testing.B) {
	track := genTrack()
	raw, _ := protoMarshall(track)
	for n := 0; n < b.N; n++ {
		protoUnmarshall(raw)
	}
}

func BenchmarkGobMarshall(b *testing.B) {
	track := genTrack()
	for n := 0; n < b.N; n++ {
		gobMarshall(track)
	}
}

func BenchmarkGobUnmarshall(b *testing.B) {
	track := genTrack()
	raw, _ := gobMarshall(track)
	for n := 0; n < b.N; n++ {
		gobUnmarshall(raw)
	}
}
