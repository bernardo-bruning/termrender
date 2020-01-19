package pixel

import (
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type CanvasPixelGl struct {
	image  *image.RGBA
	drawer *imdraw.IMDraw
	Win    *pixelgl.Window
}

func (c *CanvasPixelGl) Set(x, y int, color color.Color) {
	c.image.Set(x, y, color)
}

func Render(window *pixelgl.Window, image *image.RGBA) {
	window.Clear(colornames.Black)
	pic := pixel.PictureDataFromImage(image)
	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()))
	window.Update()
}
