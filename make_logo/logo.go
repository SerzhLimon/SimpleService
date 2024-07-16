package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	width  = 300
	height = 300
)

var (
	red   = color.RGBA{245, 15, 15, 0xff}
	blue  = color.RGBA{13, 23, 66, 0xff}
	black = color.Black
	white = color.White
	s21   = color.RGBA{15, 245, 195, 0xff}
)

func PrintSector(picture *image.RGBA, xmin, xmax, ymin, ymax int, c color.Color) {
	for x := xmin; x < xmax; x++ {
		for y := ymin; y < ymax; y++ {
			picture.Set(x, y, c)
		}
	}
}

type Coord struct {
	x int
	y int
}

func DrawCircle(img *image.RGBA, x, y, r int, c color.Color) {
	for i := x - r; i < x+r; i++ {
		for j := y - r; j < y+r; j++ {
			lenght := math.Pow(float64(i-x), 2) + math.Pow(float64(j-y), 2)
			if int(lenght) < r*r {
				img.Set(i, j, c)
			}
		}
	}
}

func PrintRound(img *image.RGBA, x0, y0, r float64, c color.Color) float64 {
	res := r - 0.5
	x, y, dx, dy := r-1.0, 0.0, 1.0, 1.0
	err := dx - (r * 2.0)

	for x >= y {
		img.Set(int(x0+x), int(y0+y), c)
		img.Set(int(x0+y), int(y0+x), c)
		img.Set(int(x0-y), int(y0+x), c)
		img.Set(int(x0-y), int(y0-x), c)
		img.Set(int(x0-x), int(y0+y), c)
		img.Set(int(x0-x), int(y0-y), c)
		img.Set(int(x0+y), int(y0-x), c)
		img.Set(int(x0+x), int(y0-y), c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
	if res == 0 {
		return res
	}
	res = PrintRound(img, x0, y0, res, c)
	return res
}

func Print21(img *image.RGBA) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, blue)
		}
	}
	//top
	PrintSector(img, 55, 120, 75, 100, s21)
	DrawCircle(img, 117, 82, 8, s21)
	DrawCircle(img, 117, 92, 8, s21)
	DrawCircle(img, 57, 92, 8, s21)
	DrawCircle(img, 57, 82, 8, s21)
	DrawCircle(img, 57, 87, 8, s21)
	DrawCircle(img, 117, 87, 8, s21)
	//right top block
	DrawCircle(img, 132, 107, 8, s21)
	DrawCircle(img, 142, 107, 8, s21)
	DrawCircle(img, 142, 117, 8, s21)
	DrawCircle(img, 132, 117, 8, s21)
	PrintSector(img, 125, 150, 104, 115, s21)
	PrintSector(img, 135, 140, 100, 125, s21)

	//middle block
	PrintSector(img, 75+5, 125-5, 125, 150, s21)
	DrawCircle(img, 57+25, 92+50, 8, s21)
	DrawCircle(img, 57+25, 87+50, 8, s21)
	DrawCircle(img, 57+25, 82+50, 8, s21)
	DrawCircle(img, 57+60, 92+50, 8, s21)
	DrawCircle(img, 57+60, 87+50, 8, s21)
	DrawCircle(img, 57+60, 82+50, 8, s21)

	//left bot block
	DrawCircle(img, 132-75, 107+50, 8, s21)
	DrawCircle(img, 142-75, 107+50, 8, s21)
	DrawCircle(img, 142-75, 117+50, 8, s21)
	DrawCircle(img, 132-75, 117+50, 8, s21)
	PrintSector(img, 125-75, 150-75, 104+50, 115+50, s21)
	PrintSector(img, 135-75, 140-75, 100+50, 125+50, s21)

	//bot
	PrintSector(img, 75+5, 145, 175, 200, s21)
	DrawCircle(img, 57+25, 92+100, 8, s21)
	DrawCircle(img, 57+25, 87+100, 8, s21)
	DrawCircle(img, 57+25, 82+100, 8, s21)
	DrawCircle(img, 57+60+25, 92+100, 8, s21)
	DrawCircle(img, 57+60+25, 87+100, 8, s21)
	DrawCircle(img, 57+60+25, 82+100, 8, s21)

	//one
	PrintSector(img, 175, 200, 75+10, 90, s21)
	PrintSector(img, 175+10, 200-10, 75, 100, s21)

	DrawCircle(img, 132+50, 107-25, 8, s21)
	DrawCircle(img, 142+50, 107-25, 8, s21)
	DrawCircle(img, 142+50, 117-25, 8, s21)
	DrawCircle(img, 132+50, 117-25, 8, s21)

	PrintSector(img, 200+5, 225-5, 100, 200, s21)
	PrintSector(img, 200, 225, 100+5, 200-5, s21)
	DrawCircle(img, 132+75, 107, 8, s21)
	DrawCircle(img, 142+75, 107, 8, s21)
	DrawCircle(img, 132+75, 192, 8, s21)
	DrawCircle(img, 142+75, 192, 8, s21)

	for x := 151; x < 158; x++ {
		Print45Line(img, 70, x, 10, s21)
	}
	for x := 126; x < 133; x++ {
		Print45Line(img, 120, x, 10, s21)
	}

	for i := 93; i < 100; i++ {
		Print135Line(img, 121, i, 10, s21)
	}

	for i := 93; i < 100; i++ {
		Print135Line(img, 121+75, i, 10, s21)
	}

	for i := 93 + 75; i < 100+75; i++ {
		Print135Line(img, 71, i, 10, s21)
	}
}

func Print45Line(img *image.RGBA, x0, y0, width int, c color.Color) {
	thickness := 3
	startX, startY := x0, y0
	endX := x0 + width

	for i := 0; i < thickness; i++ {
		if i%2 != 0 {
			x, y := startX+i, startY-i
			for x <= endX {
				img.Set(x, y, c)
				x++
				y--
			}
		}
	}
}

func Print135Line(img *image.RGBA, x0, y0, width int, c color.Color) {
	thickness := 3
	startX, startY := x0, y0
	endX := x0 + width

	for i := 0; i < thickness; i++ {
		if i%2 == 0 {
			x, y := startX+i, startY+i
			for x <= endX {
				img.Set(x, y, c)
				x++
				y++
			}
		}
	}
}

func PrintLogo(img *image.RGBA) {
	// DrawCircle(img, 100, 250, 25, white)
	// DrawCircle(img, 100, 250, 20, blue)
	DrawCircle(img, 100, 250, 25, white)
	DrawCircle(img, 100, 250, 20, blue)

	Print45Line(img, 100, 250+5, 20, white)
	Print45Line(img, 100, 249+5, 20, white)
	Print45Line(img, 100, 248+5, 20, white)
	Print45Line(img, 100, 247+5, 20, white)
	Print45Line(img, 100, 246+5, 20, white)

	Print45Line(img, 100, 245+5, 20, blue)
	Print45Line(img, 100, 244+5, 20, blue)
	Print45Line(img, 100, 243+5, 20, blue)
	Print45Line(img, 100, 242+5, 20, blue)
	Print45Line(img, 100, 241+5, 20, blue)

	Print45Line(img, 100, 256, 25, blue)
	Print45Line(img, 100, 257, 25, blue)
	Print45Line(img, 100, 258, 25, blue)
	Print45Line(img, 100, 259, 25, blue)
	Print45Line(img, 100, 260, 25, blue)

	Print135Line(img, 90, 240+5, 10, white)
	Print135Line(img, 90, 239+5, 10, white)
	Print135Line(img, 90, 238+5, 10, white)
	Print135Line(img, 90, 237+5, 10, white)
	Print135Line(img, 90, 236+5, 10, white)
}

func PrintSchool(img *image.RGBA) {
	PrintSector(img, 210, 218, 105, 107, black)
	PrintSector(img, 208, 210, 107, 110, black)
	PrintSector(img, 210, 216, 107+3, 109+3, black)
	PrintSector(img, 216, 218, 109+3, 111+5, black)
	PrintSector(img, 208, 216, 111+5, 113+5, black)

	PrintSector(img, 210, 216, 115+5, 117+5, black)
	PrintSector(img, 216, 218, 117+5, 119+5, black)
	PrintSector(img, 208, 210, 117+5, 125+5, black)
	PrintSector(img, 210, 216, 125+5, 127+5, black)
	PrintSector(img, 216, 218, 123+5, 125+5, black)

	PrintSector(img, 216, 218, 129+5, 141+5, black)
	PrintSector(img, 208, 210, 129+5, 141+5, black)
	PrintSector(img, 208, 218, 134+5, 136+5, black)

	PrintSector(img, 208, 210, 145+5, 154+5, black)
	PrintSector(img, 210, 216, 143+5, 145+5, black)
	PrintSector(img, 210, 216, 154+5, 156+5, black)
	PrintSector(img, 216, 218, 145+5, 154+5, black)

	PrintSector(img, 208, 210, 160+5, 169+5, black)
	PrintSector(img, 210, 216, 158+5, 160+5, black)
	PrintSector(img, 210, 216, 169+5, 171+5, black)
	PrintSector(img, 216, 218, 160+5, 169+5, black)

	PrintSector(img, 208, 210, 173+5, 185+5, black)
	PrintSector(img, 208, 218, 183+5, 185+5, black)
}

func PrintSber(img *image.RGBA) {
	DrawCircle(img, 154, 257, 15, white)
	PrintSector(img, 140, 154, 227, 257, blue)

	DrawCircle(img, 154, 240, 15, white)
	DrawCircle(img, 154, 238, 6, blue)

	PrintSector(img, 150, 160, 233, 244, blue)
	PrintSector(img, 160, 175, 236, 244, blue)

	DrawCircle(img, 154, 257, 7, blue)
	PrintSector(img, 140, 152, 251, 257, blue)

	for i := 235; i < 242; i++ {
		Print135Line(img, 160, i, 10, blue)
	}
	PrintSector(img, 142, 148, 231, 249, white)
	PrintSector(img, 161, 166, 230, 236, white)
	PrintSector(img, 160, 166, 262, 269, white)
	PrintSector(img, 144, 149, 261, 269, white)

	//BB
	PrintSector(img, 180, 190, 227, 272, white)
	// DrawCircle(img, 194, 257, 15, white)
	// DrawCircle(img, 194, 257, 7, blue)
	// PrintSector(img, 200, 206, 262, 269, white)
	// PrintSector(img, 200, 206, 247, 254, white)
	DrawCircle(img, 194, 257, 15, white)
	DrawCircle(img, 194, 257, 7, blue)
	DrawCircle(img, 194, 237, 11, white)
	DrawCircle(img, 193, 237, 5, blue)
	//EE
	PrintSector(img, 220, 250, 227, 237, white)
	PrintSector(img, 220, 250, 244, 254, white)
	PrintSector(img, 220, 250, 262, 272, white)
	PrintSector(img, 220, 230, 227, 272, white)

	//RRRR
	DrawCircle(img, 274, 241, 15, white)
	PrintSector(img, 260, 270, 227, 272, white)
	for i := 251; i < 261; i++ {
		Print135Line(img, 270, i, 19, white)
	}
	PrintSector(img, 260, 290, 272, 280, blue)
	DrawCircle(img, 274, 241, 7, blue)
	PrintSector(img, 279, 286, 246, 253, white)
	PrintSector(img, 279, 286, 231, 237, white)
}

func main() {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	Print21(img)
	PrintLogo(img)
	PrintSchool(img)
	PrintSber(img)

	// DrawCircle(img, 100, 100, 500, red)

	f, _ := os.Create("amazing_logo.png")
	png.Encode(f, img)

}
