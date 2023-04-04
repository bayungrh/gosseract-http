package lib

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"

	"github.com/otiai10/gosseract/v2"
)

// return bytes
func CropImageRect(imageByte []byte, rect []int) []byte {
  img, _, _ := image.Decode(bytes.NewReader(imageByte))
  x := rect[0]
  y := rect[1]
  w := rect[2]
  h := rect[3]

  cropSize := image.Rect(0, 0, w, h)
  cropSize = cropSize.Add(image.Point{x, y})

  grayImg := image.NewGray(cropSize.Bounds())
  draw.Draw(grayImg, grayImg.Bounds(), img, cropSize.Min, draw.Src)

  buf := &bytes.Buffer{}
  png.Encode(buf, grayImg)
  return buf.Bytes()
}

func ParseText(imageByte []byte, rect []int) (string, error) {
  var instance gosseract.Client

  tesseract := TesseractEngine{
    Name:     "tesseract",
    Language: "eng",
    Variables: map[string]string{
      "tessedit_pageseg_mode": "6",
    },
    Client: nil,
    ImageBytes: nil,
  }
  defer tesseract.Close()

  instance = *gosseract.NewClient()
  tesseract.Client = &instance
  tesseract.ImageBytes = CropImageRect(imageByte, rect)
  tesseract.TesseractSettings()

  object, err := tesseract.ExtractText()
  return object, err
}
