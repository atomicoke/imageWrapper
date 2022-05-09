package img

import (
	"github.com/disintegration/imaging"
	"os"
	"testing"
)

func Test_cut(t *testing.T) {
	//resp, _ := http.Get("https://hfscrm-dev.oss-cn-beijing.aliyuncs.com/1650729600000/fa48652f25284bc3973a668ba189d940.png?Expires=1651896149&OSSAccessKeyId=LTAI5tQbviPENMPnqxXdtNju&Signature=wEegNcjBE5RSi9kiHpYATU1P4Ok%3D")
	source, _ := os.Open("fa48652f25284bc3973a668ba189d940.png")

	image, _ := imaging.Decode(source)

	lanczos := imaging.Resize(image, 400, 0, imaging.NearestNeighbor)
	crop := imaging.CropAnchor(lanczos, 400, 700, imaging.Top)
	f, _ := os.OpenFile("lanczos.jpg", os.O_RDWR|os.O_CREATE, 0755)
	_ = imaging.Encode(f, crop, imaging.JPEG)
}

func Test_fuzz_image(t *testing.T) {
	//resp, _ := http.Get("https://pic2.zhimg.com/v2-471f8aa91487ac3c073ab5c5b42361ca_400x224.jpg?source=7e7ef6e2")
	source, _ := os.Open("branches.png")
	image, _ := imaging.Decode(source)

	lanczos := imaging.Resize(image, 400, 0, imaging.NearestNeighbor)
	f, _ := os.OpenFile("lanczos.jpg", os.O_RDWR|os.O_CREATE, 0755)
	_ = imaging.Encode(f, lanczos, imaging.JPEG)
}
