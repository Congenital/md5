package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5(buff []byte) ([]byte, error) {
	h := md5.New()

	_, err := h.Write(buff)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func IoMD5(read io.Reader) ([]byte, error) {
	h := md5.New()

	buff := make([]byte, h.BlockSize())
	for {
		_, err := read.Read(buff)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		_, err = h.Write(buff)
		if err != nil {
			return nil, err
		}
	}

	return h.Sum(nil), nil
}

func FileMD5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h, err := IoMD5(file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h), nil
}
