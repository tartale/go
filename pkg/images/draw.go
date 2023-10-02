package images

import (
	"image"
	"image/color"

	"github.com/tartale/go/pkg/mathx"
)

type DrawableImage interface {
	Set(x, y int, clr color.Color)
	image.Image
}

func IsVertical(startX, startY, endX, endY int) bool {

	return mathx.Abs(endX-startX) == 0
}

func IsHorizontal(startX, startY, endX, endY int) bool {

	return mathx.Abs(endY-startY) == 0
}

func NormalizePoints(img DrawableImage, startX, startY, endX, endY *int, thickness int) {

	if *startX > *endX {
		*startX, *endX = *endX, *startX
	}
	if *startY > *endY {
		*startY, *endY = *endY, *startY
	}
	width := img.Bounds().Max.X
	if IsVertical(*startX, *startY, *endX, *endY) {
		*startX = mathx.Min(*startX, width-thickness)
		*endX = mathx.Min(*endX, width-thickness)
	}
	height := img.Bounds().Max.Y
	if IsHorizontal(*startX, *startY, *endX, *endY) {
		*startY = mathx.Min(*startY, height-thickness)
		*endY = mathx.Min(*endY, height-thickness)
	}
}

func DrawPoint(img DrawableImage, x, y int, clr color.Color) {

	img.Set(x, y, clr)
}

func DrawLine(img DrawableImage, startX, startY, endX, endY, thickness int, clr color.Color) {

	NormalizePoints(img, &startX, &startY, &endX, &endY, thickness)
	dx := mathx.Abs(endX - startX)
	dy := mathx.Abs(endY - startY)

	// Check if the line is horizontal or vertical
	if IsVertical(startX, startY, endX, endY) {
		// Vertical line
		for y := startY; y <= endY; y++ {
			for w := 0; w < thickness; w++ {
				img.Set(startX+w, y, clr)
			}
		}
	} else if IsHorizontal(startX, startY, endX, endY) {
		// Horizontal line
		for x := startX; x <= endX; x++ {
			for w := 0; w < thickness; w++ {
				img.Set(x, startY+w, clr)
			}
		}
	} else {
		// Diagonal line (Bresenham's algorithm)
		var sx, sy int
		if startX < endX {
			sx = 1
		} else {
			sx = -1
		}
		if startY < endY {
			sy = 1
		} else {
			sy = -1
		}

		err := dx - dy
		for {
			for w := 0; w < thickness; w++ {
				img.Set(startX, startY+w, clr)
			}

			if startX == endX && startY == endY {
				break
			}

			e2 := 2 * err
			if e2 > -dy {
				err -= dy
				startX += sx
			}
			if e2 < dx {
				err += dx
				startY += sy
			}
		}
	}
}

func DrawFullVerticalLine(img DrawableImage, x, thickness int, clr color.Color) {

	height := img.Bounds().Max.Y
	DrawLine(img, x, 0, x, height, thickness, clr)
}

func DrawVerticalLineBetween(img DrawableImage, x, startY, endY, thickness int, clr color.Color) {

	DrawLine(img, x, startY, x, endY, thickness, clr)
}

func DrawFullHorizontalLine(img DrawableImage, y, thickness int, clr color.Color) {

	width := img.Bounds().Max.X
	DrawLine(img, 0, y, width, y, thickness, clr)
}

func DrawHorizontalLineBetween(img DrawableImage, startX, endX, y, thickness int, clr color.Color) {

	DrawLine(img, startX, y, endX, y, thickness, clr)
}

func DrawRectangle(img DrawableImage, r image.Rectangle, thickness int, clr color.Color) {

	/*
		Rectangle points look like this:

		minX,minY

										maxX,maxY
	*/

	// top border
	DrawHorizontalLineBetween(img, r.Min.X, r.Max.X, r.Min.Y, thickness, clr)
	// left border
	DrawVerticalLineBetween(img, r.Min.X, r.Min.Y, r.Max.Y, thickness, clr)
	// bottom border
	DrawHorizontalLineBetween(img, r.Min.X, r.Max.X+thickness-1, r.Max.Y, thickness, clr)
	// right border
	DrawVerticalLineBetween(img, r.Max.X, r.Min.Y, r.Max.Y+thickness-1, thickness, clr)
}
