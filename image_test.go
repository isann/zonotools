package zonotools

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDecodeImage(t *testing.T) {
	//type args struct {
	//	reader io.Reader
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	want    image.Image
	//	want1   string
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, got1, err := DecodeImage(tt.args.reader)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("DecodeImage() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("DecodeImage() got = %v, want %v", got, tt.want)
	//		}
	//		if got1 != tt.want1 {
	//			t.Errorf("DecodeImage() got1 = %v, want %v", got1, tt.want1)
	//		}
	//	})
	//}
}

func TestGetRect(t *testing.T) {
	//type args struct {
	//	img image.Image
	//}
	//tests := []struct {
	//	name       string
	//	args       args
	//	wantWidth  int
	//	wantHeight int
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		gotWidth, gotHeight := GetRect(tt.args.img)
	//		if gotWidth != tt.wantWidth {
	//			t.Errorf("GetRect() gotWidth = %v, want %v", gotWidth, tt.wantWidth)
	//		}
	//		if gotHeight != tt.wantHeight {
	//			t.Errorf("GetRect() gotHeight = %v, want %v", gotHeight, tt.wantHeight)
	//		}
	//	})
	//}
}

func TestResizeImage(t *testing.T) {
	//type args struct {
	//	srcImg    image.Image
	//	imageType string
	//	dstWidth  int
	//	dstHeight int
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	want    image.Image
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := ResizeImage(tt.args.srcImg, tt.args.imageType, tt.args.dstWidth, tt.args.dstHeight)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("ResizeImage() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("ResizeImage() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	// TODO: テスト前に適当な a.png を用意
	t.Run("normal", func(t *testing.T) {
		file, _ := os.Open("a.png")
		img, format, _ := DecodeImage(file)
		orgWidth, orgHeight := GetRect(img)
		resizeImage, _ := ResizeImage(img, format, orgWidth/2, orgHeight/2)
		resizeWidth, resizeHeight := GetRect(resizeImage)
		assert.Equal(t, orgWidth/2, resizeWidth)
		assert.Equal(t, orgHeight/2, resizeHeight)
		create, _ := os.Create("dst.png")
		_ = EncodeImage(format, create, resizeImage)
	})
}

func TestRotation(t *testing.T) {
	//type args struct {
	//	inputImage image.Image
	//	mode       int
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want image.Image
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := RotationImage(tt.args.inputImage, tt.args.mode); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("RotationImage() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	t.Run("normal", func(t *testing.T) {
		file, _ := os.Open("a.png")
		img, format, _ := DecodeImage(file)
		rotationImage := RotationImage(img, 0)
		create, _ := os.Create("dst-rotation.png")
		// TODO: Check image manual
		_ = EncodeImage(format, create, rotationImage)
	})
}
