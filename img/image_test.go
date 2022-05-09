package img

import (
	"fmt"
	"github.com/disintegration/imaging"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"net/http"
	"os"
	"strings"
	"testing"
)

func Test_cut(t *testing.T) {
	//resp, _ := http.GetFromCache("https://hfscrm-dev.oss-cn-beijing.aliyuncs.com/1650729600000/fa48652f25284bc3973a668ba189d940.png?Expires=1651896149&OSSAccessKeyId=LTAI5tQbviPENMPnqxXdtNju&Signature=wEegNcjBE5RSi9kiHpYATU1P4Ok%3D")
	source, _ := os.Open("fa48652f25284bc3973a668ba189d940.png")

	image, _ := imaging.Decode(source)

	lanczos := imaging.Resize(image, 400, 0, imaging.NearestNeighbor)
	crop := imaging.CropAnchor(lanczos, 400, 700, imaging.Top)
	f, _ := os.OpenFile("lanczos.jpg", os.O_RDWR|os.O_CREATE, 0755)
	_ = imaging.Encode(f, crop, imaging.JPEG)
}

func Test_fuzz_image(t *testing.T) {
	resp, _ := http.Get("http://114.132.249.192:9000/chat/1651676158000/312362229276450816blob")
	//source, _ := os.Open("branches.png")
	image, _ := imaging.Decode(resp.Body)

	lanczos := imaging.Resize(image, 400, 0, imaging.Box)
	lanczos = imaging.Blur(lanczos, 5)
	f, _ := os.OpenFile("lanczos.jpg", os.O_RDWR|os.O_CREATE, 0755)
	_ = imaging.Encode(f, lanczos, imaging.JPEG)
}

func Test_Create(t *testing.T) {
	i := imaging.New(200, 200, color.White)

	draw.Draw(i, i.Bounds(), &image.Uniform{C: color.White}, image.Point{
		X: 40,
		Y: 40,
	}, draw.Over)
	imaging.Save(i, "test.jpg")
}

func Test_split_url(t *testing.T) {
	s := "http://localhost:8888/fuzz?url=http://114.132.249.192/a/500?url=http://114.132.249.192:9000/chat/1651676158000/312362229276450816blob"

	split := strings.Split(s, "url=")
	fmt.Println(split)
}
