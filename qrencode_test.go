package qrencode

import "testing"
import "github.com/scottserok/qrencode"
import "regexp"

func TestEncode(t *testing.T) {
  qrencode.Init()
  path, err := qrencode.EncodeString("http://google.com/")
  if err != nil {
    t.Error("Error while encoding string, ", err)
  }
  qrencode.Destroy(path)
}

func TestGenUID(t *testing.T) {
  qrencode.Init()
  uid1 := qrencode.GenUID()
  uid2 := qrencode.GenUID()
  if uid1 == uid2 {
    t.Error("genUID is not uniq, ", uid1, uid2)
  }
}

func TestGenFilePath(t *testing.T) {
  qrencode.Init()
  filepath := qrencode.GenFilePath()
  matched, err := regexp.MatchString("^/tmp/.+.png$", filepath)
  if err != nil {
    t.Error("Error matching filepath", err)
  }
  if !matched {
    t.Error("filepath doesnt match expected pattern, ", filepath)
  }
}
