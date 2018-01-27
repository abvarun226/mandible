package processorcommand

import (
	"fmt"
	"log"
	"os"
)

const JPEGMINI_COMMAND = "/usr/bin/jpegmini"

func JpegMini(quality, shc, path string) error {
	shc_arg := fmt.Sprintf("-shc=%s", shc)
	quality_arg := fmt.Sprintf("-qual=%s", quality)
	in_file := fmt.Sprintf("-f=%s", path)
	out_file := fmt.Sprintf("-o=%s_mini", path)

	args := []string{
		shc_arg,
		quality_arg,
		in_file,
		out_file,
	}

	err := runProcessorCommand(JPEGMINI_COMMAND, args)

	// if jpegmini cmd was successful
	if err == nil {
		// Overwrite thumb with jpegmini optimized image
		err = os.Rename(path+"_mini", path)
		if err != nil {
			log.Printf("JPEGMini error: error renaming file: %s", err.Error())
		}
	} else {
		// Dont fail the request even if jpegmini cmd fails
		log.Printf("JPEGMini error while processing: %s", err.Error())
	}

	return nil
}
