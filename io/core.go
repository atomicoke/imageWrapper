package io

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Read(filePath string) (io.Reader, error) {
	open, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return open, nil
}

func ReaderToBytes(reader io.Reader) ([]byte, error) {
	return ioutil.ReadAll(reader)
}

func ReadFileToBytes(filePath string) ([]byte, error) {
	reader, err := Read(filePath)
	if err != nil {
		return nil, err
	}
	return ReaderToBytes(reader)
}

func ReadUrlToBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ReaderToBytes(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadUrl(url string) (io.Reader, error) {
	data, err := ReadUrlToBytes(url)
	if err != nil {
		return nil, err
	}

	return Wrap(data), nil
}

func Wrap(data []byte) *bytes.Buffer {
	return bytes.NewBuffer(data)
}
