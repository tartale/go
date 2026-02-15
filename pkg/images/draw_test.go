package images

import (
	"image"
	"image/png"
	"os"
	"testing"

	"golang.org/x/image/colornames"
)

func loadTestImage() DrawableImage {

	imgFile, err := os.OpenFile("./test/blank.png", os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	if drawableImg, ok := img.(DrawableImage); ok {
		return drawableImg
	}

	panic("unable to load test file")
}

func saveImage(name string, img DrawableImage) {

	imgFile, err := os.Create("./test/test-" + name)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()
	err = png.Encode(imgFile, img)
	if err != nil {
		panic(err)
	}
}

func TestDrawDiagonalLine(t *testing.T) {

	testImage := loadTestImage()
	DrawLine(testImage, 100, 100, 200, 200, 10, colornames.Red)
	saveImage("diagonal-line.png", testImage)
}

func TestDrawFullVerticalLine(t *testing.T) {

	testImage := loadTestImage()
	DrawFullVerticalLine(testImage, 100, 10, colornames.Red)
	saveImage("vertical-line-full.png", testImage)
}

func TestDrawVerticalLineBetween(t *testing.T) {

	testImage := loadTestImage()
	DrawVerticalLineBetween(testImage, 100, 200, 300, 10, colornames.Red)
	saveImage("vertical-line-between.png", testImage)
}

func TestDrawFullHorizontalLine(t *testing.T) {

	testImage := loadTestImage()
	DrawFullHorizontalLine(testImage, 100, 10, colornames.Red)
	saveImage("horizontal-line-full.png", testImage)
}

func TestDrawHorizontalLineBetween(t *testing.T) {

	testImage := loadTestImage()
	DrawHorizontalLineBetween(testImage, 100, 200, 100, 10, colornames.Red)
	saveImage("horizontal-line-between.png", testImage)
}

func TestDrawRectangleBorder(t *testing.T) {

	testImage := loadTestImage()
	width := testImage.Bounds().Max.X
	height := testImage.Bounds().Max.Y
	rect := image.Rect(0, 0, width, height)
	DrawRectangle(testImage, rect, 10, colornames.Red)
	saveImage("rectangle-border.png", testImage)

}
func TestDrawRectangleInner(t *testing.T) {

	testImage := loadTestImage()
	width := testImage.Bounds().Max.X
	height := testImage.Bounds().Max.Y
	rect := image.Rect(10, 50, width-100, height-50)
	DrawRectangle(testImage, rect, 10, colornames.Red)
	saveImage("rectangle-inner.png", testImage)
}
