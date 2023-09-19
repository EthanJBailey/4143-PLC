// Package Declaration:
package imageManipulator

// Import required packages
import (
	"github.com/fogleman/gg"
)

// ImageManipulator represents an image manipulation tool.
type ImageManipulator struct {
	Image     *gg.Context
	ImagePath string // Add a field to store the image path
}

// NewImageManipulator creates a new ImageManipulator instance.
func NewImageManipulatorWithImage(imagePath string) (*ImageManipulator, error) {
	img, err := gg.LoadImage(imagePath)
	if err != nil {
		return nil, err
	}

	ctx := gg.NewContextForImage(img)
	return &ImageManipulator{Image: ctx, ImagePath: imagePath}, nil
}

// SaveToFile saves the manipulated image to a file.
func (im *ImageManipulator) SaveToFile(filename string) error {
	return im.Image.SavePNG(filename)
}

// DrawRectangle draws a rectangle on the image.
func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
	im.Image.DrawRectangle(x, y, width, height)
	im.Image.Stroke()
}
