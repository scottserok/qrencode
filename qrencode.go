package qrencode

import "os"
import "log"
import "bytes"
import "strconv"
import "image/png"
import "github.com/sony/sonyflake"
import qre "github.com/qpliu/qrencode-go/qrencode"

var sf *sonyflake.Sonyflake
var fileDirectory string

func Init() {
  sf = sonyflake.NewSonyflake(sonyflake.Settings{})
  fileDirectory = "/tmp/"
}

func EncodeString(content string) (string, error) {
  grid, err := qre.Encode(content, qre.ECLevelQ)
	if err != nil {
		panic(err)
	}
  filepath := GenFilePath()
  f, err := os.Create(filepath)
  if err != nil {
    log.Fatal(err)
  }
  png.Encode(f, grid.Image(8))
  if err := f.Close(); err != nil {
    log.Fatal(err)
  }
  return filepath, err
}

func GenFilePath() (string) {
  var buffer bytes.Buffer
  buffer.WriteString(fileDirectory)
  buffer.WriteString("qrencoded")
  buffer.WriteString(strconv.Itoa(int(GenUID())))
  buffer.WriteString(".png")
  return buffer.String()
}

func GenUID() (uint64) {
  id, err := sf.NextID()
  if err != nil {
      log.Fatalf("flake.NextID() failed with %s\n", err)
  }
  return id
}

func Destroy(filepath string) {
  os.Remove(filepath)
}
