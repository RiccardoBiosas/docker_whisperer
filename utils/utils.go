package utils

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"os"
)

func CreateTarball(tarballFilePath string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	if _, err := os.Stat(tarballFilePath); err != nil {
		return nil, err
	}

	reader, err := os.Open(tarballFilePath)
	if err != nil {
		return nil, err
	}

	dataReader, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	tarHeader := &tar.Header{
		Name: tarballFilePath,
		Size: int64(len(dataReader)),
	}

	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return nil, err
	}

	_, err = tw.Write(dataReader)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
