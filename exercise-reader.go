package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	cnt := 0
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
		cnt = i + 1
	}
	return cnt, nil
}

func main() {
	reader.Validate(MyReader{})
}
