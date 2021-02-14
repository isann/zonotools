package zonotools

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
)

// Ref. https://text.baldanders.info/golang/resize-image/

func DecodeImage(reader io.Reader) (image.Image, string, error) {
	// Decode image
	return image.Decode(reader)
}

func EncodeImage(imageType string, dst io.Writer, srcImg image.Image) error {
	// Encode resized image
	var err error
	switch imageType {
	case "jpeg":
		err = jpeg.Encode(dst, srcImg, &jpeg.Options{Quality: 100})
	case "gif":
		err = gif.Encode(dst, srcImg, nil)
	case "png":
		err = png.Encode(dst, srcImg)
	default:
		err = errors.New(fmt.Sprintf("format error %s", imageType))
	}
	return err
}

func GetRect(img image.Image) (width, height int) {
	rct := img.Bounds()
	width = rct.Dx()
	height = rct.Dy()
	return
}

func ResizeImage(srcImg image.Image, imageType string, dstWidth, dstHeight int) (image.Image, error) {
	rctSrc := srcImg.Bounds()
	// Scale down
	dstImg := image.NewRGBA(image.Rect(0, 0, dstWidth, dstHeight))
	draw.CatmullRom.Scale(dstImg, dstImg.Bounds(), srcImg, rctSrc, draw.Over, nil)

	dst := new(bytes.Buffer)

	_ = EncodeImage(imageType, dst, dstImg)
	dstImage, _, err := DecodeImage(bytes.NewBuffer(dst.Bytes()))
	return dstImage, err
}

// Ref. https://qiita.com/ikeponsu/items/0beb5387882eb1f9d525
// 回転の処理
func RotationImage(inputImage image.Image, mode int) image.Image {

	// 出力画像を定義
	var outputImage image.Image

	switch mode {
	case -1:
		// 右90度回転
		outputImage = affine(inputImage, 90, inputImage.Bounds().Max.X-1, 0, 1)
	case 0:
		// 右180度回転
		outputImage = affine(inputImage, 180, inputImage.Bounds().Max.X-1, inputImage.Bounds().Max.Y-1, 1)
	case 1:
		// 右270度回転
		outputImage = affine(inputImage, 270, 0, inputImage.Bounds().Max.Y-1, 1)
	default:
		log.Fatal("angle code does not exist")
	}

	return outputImage
}

// アフィン変換の処理
func affine(inputImage image.Image, angle int, tx int, ty int, scale float64) image.Image {

	// 出力画像を定義
	size := inputImage.Bounds()
	size.Max.X = int(float64(size.Max.X) * scale)
	size.Max.Y = int(float64(size.Max.Y) * scale)

	outputImage := image.NewRGBA(size)

	// ステータスのキャスト
	theta := float64(angle) * math.Pi / 180
	cos := math.Cos(theta)
	sin := math.Sin(theta)

	matrix := [][]float64{{cos * scale, -sin * scale, float64(tx)}, {sin * scale, cos * scale, float64(ty)}, {0.0, 0.0, 1.0}}

	// 左右反転
	for y := size.Min.Y; y < size.Max.Y; y++ {
		for x := size.Min.X; x < size.Max.X; x++ {

			outputX := 0
			outputY := 0
			// 元座標を格納
			origin := []float64{float64(x), float64(y), 1.0}

			// 座標を計算
			for rowKey, rowVal := range matrix {
				var val float64

				for colIndex := 0; colIndex < len(rowVal); colIndex++ {

					val += origin[colIndex] * rowVal[colIndex]
				}

				// 座標の代入
				switch rowKey {
				case 0:
					outputX = int(math.Round(val))
					break
				case 1:
					outputY = int(math.Round(val))
					break
				default:
					break
				}

			}

			if size.Min.X <= outputX && outputX < size.Max.X && size.Min.Y <= outputY && outputY < size.Max.Y {
				outputImage.Set(outputX, outputY, inputImage.At(x, y))
			} else {
				// 何もしない
			}
		}
	}

	return outputImage
}
