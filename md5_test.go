package md5

import (
	"github.com/Congenital/log/v0.2/log"
	"testing"
)

func TestMD5(t *testing.T) {
	log.Info(MD5CMD("/home/andy/Downloads/STAX人头录音专辑.zip"))
}
