package image

import (
	"github.com/disintegration/imaging"
	"os"
	"testing"
)

func Test_fuzz_image(t *testing.T) {
	//resp, _ := http.Get("https://pic2.zhimg.com/v2-471f8aa91487ac3c073ab5c5b42361ca_400x224.jpg?source=7e7ef6e2")
	source, _ := os.Open("branches.png")
	image, _ := imaging.Decode(source)

	lanczos := imaging.Resize(image, 400, 0, imaging.NearestNeighbor)
	f, _ := os.OpenFile("lanczos.jpg", os.O_RDWR|os.O_CREATE, 0755)
	_ = imaging.Encode(f, lanczos, imaging.JPEG)
}
