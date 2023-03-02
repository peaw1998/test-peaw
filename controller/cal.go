package controller

import (
	"calculate/service"

	"github.com/gin-gonic/gin"
)

type Area interface {
	cal()
}

func calculateArea(cp Area) {
	cp.cal()
}

type Triangle struct {
	h    int
	w    int
	area float32
}

func (t *Triangle) cal() {
	t.area = (float32(0.5) * float32(t.h) * float32(t.w))
}

type Square struct {
	h    int
	w    int
	area float32
}

func (s *Square) cal() {
	s.area = float32(s.h * s.w)
}

type ResultArea struct {
	TriangleArea float32
	SquareArea   float32
}

type InputData struct {
	H int `json:"Height"`
	W int `json:"Width"`
}

func Test() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data InputData

		if err := c.BindJSON(&data); err != nil {
			return
		}

		t := Triangle{h: data.H, w: data.W}
		s := Square{h: data.H, w: data.W}

		calculateArea(&t)
		calculateArea(&s)

		result := ResultArea{TriangleArea: t.area, SquareArea: s.area}
		service.Resp(c, 1, "OK", result)
	}
}
