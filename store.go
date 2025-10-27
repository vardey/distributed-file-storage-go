package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	Original string
}

type StoreOpts struct{ PathTransformFunc PathTransformFunc }

type Store struct{ StoreOpts }

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hexHash := hex.EncodeToString(hash[:])

	blockSize := 5
	noOfDir := len(hexHash) / blockSize
	paths := make([]string, noOfDir)

	for i := range noOfDir {
		from, to := (i * blockSize), (i*blockSize)+blockSize
		paths[i] = hexHash[from:to]

	}

	return PathKey{PathName: strings.Join(paths, "/"), Original: hexHash}
}

func (p *PathKey) PathAndFilename() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Original)
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

func (s *Store) WriteStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	//create path
	if err := os.MkdirAll(pathKey.PathName, os.ModePerm); err != nil {
		fmt.Println("make dir err: ", err)
		return err
	}

	pathAndFileName := pathKey.PathAndFilename()

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
