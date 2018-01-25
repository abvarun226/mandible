package imageprocessor

import (
	"github.com/Imgur/mandible/imageprocessor/processorcommand"
	"github.com/Imgur/mandible/uploadedfile"
)

type JpegMiniOptimizer struct {
	quality string
	shc     string
}

func (this *JpegMiniOptimizer) Process(image *uploadedfile.UploadedFile) error {
	if !image.IsJpeg() {
		return nil
	}

	thumbPath := image.GetThumbs()[0].GetPath()
	err := processorcommand.JpegMini(this.quality, this.shc, thumbPath)

	if err != nil {
		return err
	}

	return nil
}

func (this *JpegMiniOptimizer) String() string {
	return "Optimize JPEG thumbs using JPEGMini"
}
