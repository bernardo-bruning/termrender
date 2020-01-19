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

func NewCanvas() *CanvasPixelGl {
	cfg := pixelgl.WindowConfig{
		Title:  "Term render",
		Bounds: pixel.R(0, 0, 800, 800),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)
	img := image.NewRGBA(image.Rect(0, 0, 800, 800))
	drawer := imdraw.New(pixel.MakePictureData(pixel.R(0, 0, 800, 800)))
	return &CanvasPixelGl{image: img, drawer: drawer, Win: win}
}

func Render(c *CanvasPixelGl) {
	c.Win.Clear(colornames.Black)
	pic := pixel.PictureDataFromImage(c.image)
	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(c.Win, pixel.IM.Moved(c.Win.Bounds().Center()))
	c.image = image.NewRGBA(image.Rect(0, 0, 800, 800))
	c.Win.Update()
}
