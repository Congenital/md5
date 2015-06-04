package md5

import (
	"encoding/hex"
	"github.com/Congenital/log/v0.2/log"
	"os"
	"testing"
)

func TestMD5(t *testing.T) {
	h, err := MD5([]byte("fdsafdsa"))

	log.Info(hex.EncodeToString(h), err)

	file, err := os.Open("md5.go")
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	he, err := IoMD5(file)

	log.Info(hex.EncodeToString(he), err)
}
