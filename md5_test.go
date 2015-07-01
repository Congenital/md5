package md5

import (
	"github.com/Congenital/log/v0.2/log"
	"testing"
)

func TestMD5(t *testing.T) {
	log.Info(FileMD5("/home/andy/Godin/src/git.sudoteam.com/gogroup/gos-signalling/bin/rom/Hammerhead/Hammerhead_1.0.0_1.0.5_68334334eb250279d524d8ae8845c123.zip"))
	log.Info(MD5CMD("/home/andy/Godin/src/git.sudoteam.com/gogroup/gos-signalling/bin/rom/Hammerhead/Hammerhead_1.0.0_1.0.5_68334334eb250279d524d8ae8845c123.zip"))
}
