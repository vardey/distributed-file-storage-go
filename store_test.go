package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	pathString := CASPathTransformFunc("myKeyForPath")
	fmt.Println(pathString)
}

func TestStore(t *testing.T) {
	myStoreOpts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	myStore := NewStore(myStoreOpts)

	myData := bytes.NewReader([]byte("my file data"))

	if err := myStore.WriteStream("myPath", myData); err != nil {
		t.Error(err)
	}
}
