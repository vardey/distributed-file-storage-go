package main

import (
	"fmt"
	"io"
	"os"
)

type PathTransformFunc func(string) string

type StoreOpts struct{ PathTransformFunc PathTransformFunc }

type Store struct{ StoreOpts }

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

func (s *Store) WriteStream(key string, r io.Reader) error {
	path := s.PathTransformFunc(key)
	//create path
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println("make dir err: ", err)
		return err
	}

	fileName := "myFile"
	pathAndFileName := path + "/" + fileName

	//create path
	file, err := os.Create(pathAndFileName)
	if err != nil {
		fmt.Println("create file err: ", err)
	}

	writtenN, err := io.Copy(file, r)
	if err != nil {
		fmt.Println("write file err: ", err)
		return err
	}
	fmt.Println("written bytes to file: ", writtenN, pathAndFileName)
	return nil
}
