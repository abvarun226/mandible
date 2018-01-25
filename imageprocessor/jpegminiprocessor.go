package imageprocessor

import (
	"github.com/Imgur/mandible/imageprocessor/processorcommand"
	"github.com/Imgur/mandible/uploadedfile"
)

type JpegMiniOptimizer struct{}

func (this *JpegMiniOptimizer) Process(image *uploadedfile.UploadedFile) error {
	if !image.IsJpeg() {
		return nil
	}

	jpegMiniQuality := s.Config.JpegMiniConfig["quality"]
	jpegMiniSHC := s.Config.JpegMiniConfig["shc"]
	thumbPath := image.GetThumbs()[0].GetPath()

	err := processorcommand.JpegMini(jpegMiniQuality, jpegMiniSHC, thumbPath)

	if err != nil {
		return err
	}

	return nil
}

func (this *JpegMiniOptimizer) String() string {
	return "Optimize JPEG thumbs using JPEGMini"
}
