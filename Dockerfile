FROM golang:1.8-stretch
RUN apt-get update && apt-get install -yqq aspell aspell-en libaspell-dev tesseract-ocr tesseract-ocr-eng libc6 optipng exiftool libjpeg-progs webp
ADD docker/build_gm.sh /tmp/build_gm.sh
ADD jpegmini_3.14.14.72670860-ubuntu0.1_amd64.deb /tmp/jpegmini.deb
RUN \
    bash /tmp/build_gm.sh && \
    dpkg -i /tmp/jpegmini.deb
ADD docker/meme.traineddata /usr/share/tesseract-ocr/tessdata/meme.traineddata
RUN mkdir -p /etc/mandible /tmp/imagestore
ENV MANDIBLE_CONF /etc/mandible/conf.json
ADD . /go/src/github.com/Imgur/mandible
WORKDIR /go/src/github.com/Imgur/mandible
RUN go get github.com/mattn/goveralls
RUN go get github.com/tools/godep
RUN godep restore
RUN godep go install -v .
CMD ["mandible"]
