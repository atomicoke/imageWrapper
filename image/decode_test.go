package image

import (
	"fmt"
	"github.com/atomicoke/imageWrapper/io"
	"testing"
)

func Test_decode_file(t *testing.T) {
	bytes, _ := io.ReadFileToBytes("branches.png")

	_, err := Decode(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_decode_url(t *testing.T) {
	url := "https://hfscrm-dev.oss-cn-beijing.aliyuncs.com/1650729600000/885ce94605ce4aadaa459f0ea519e515.png?Expires=1651912760&OSSAccessKeyId=LTAI5tQbviPENMPnqxXdtNju&Signature=Xko6de9AbJzcGp6F%2F1J0H8IUezY%3D"
	bytes, err := io.ReadUrlToBytes(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	image, err := Decode(bytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(image.ColorModel())
}
