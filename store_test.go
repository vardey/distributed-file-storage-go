package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	myStoreOpts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	myStore := NewStore(myStoreOpts)

	myData := bytes.NewReader([]byte("my file data"))

	if err := myStore.WriteStream("myPath", myData); err != nil {
		t.Error(err)
	}
}
