package goutil

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrorIsDir = errors.New("err: file is dir")

// MD5Sum returns MD5 checksum of filename
func MD5Sum(file string, chunk int64) (string, error) {
	if fileInfo, err := os.Stat(file); err != nil {
		return "", err
	} else if fileInfo.IsDir() {
		return "", ErrorIsDir
	}
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := md5.New()
	for buf, reader := make([]byte, chunk), bufio.NewReader(f); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hash.Write(buf[:n])
	}
	checkSum := fmt.Sprintf("%x", hash.Sum(nil))
	return checkSum, nil
}
