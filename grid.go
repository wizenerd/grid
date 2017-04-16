package grid

import (
	"strconv"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Grid struct {
	vecty.Component
	NoSpacing bool
	Cells     []*Cell
}

func (g Grid) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-grid"] = true
	if g.NoSpacing {
		c["mdl-grid--no-spacing"] = true
	}
	var v []vecty.MarkupOrComponentOrHTML
	v = append(v, c)
	for _, cell := range g.Cells {
		v = append(v, cell)
	}
	return elem.Div(v...)
}

type Cell struct {
	vecty.Component
	Children  vecty.MarkupOrComponentOrHTML
	Mode      Mode
	Hide      bool
	Size      int
	HideMode  Mode
	Order     bool
	OrderSize int
	OrderMode Mode
	Stretch   bool
	Top       bool
	Middle    bool
	Bottom    bool
}

type Mode byte

const (
	Default Mode = 1 << iota
	Desktop
	Tablet
	Phone
)

func (m Mode) String() string {
	r := ""
	switch m {
	case Desktop:
		r = "desktop"
	case Tablet:
		r = "tablet"
	case Phone:
		r = "phone"
	}
	return r
}

func modes(m Mode) []string {
	var r []string
	if m&Default != 0 {
		r = append(r, Default.String())
	}
	if m&Desktop != 0 {
		r = append(r, Desktop.String())
	}
	if m&Tablet != 0 {
		r = append(r, Tablet.String())
	}
	if m&Phone != 0 {
		r = append(r, Phone.String())
	}
	return r
}

func (c *Cell) Render() *vecty.HTML {
	return elem.Div(
		c.Style(),
		c.Children,
	)
}

func (c *Cell) Style() vecty.ClassMap {
	s := make(vecty.ClassMap)
	s["mdl-cell"] = true
	m := modes(c.Mode)
	for _, v := range m {
		i := toString(c.Size)
		if v != "" {
			s["mdl-cell--"+i+"-col-"+v] = true
		} else {
			s["mdl-cell--"+i+"-col"] = true
		}
	}
	if c.Order {
		for _, v := range modes(c.OrderMode) {
			i := toString(c.OrderSize)
			if v != "" {
				s["mdl-cell--order"+i+"-"+v] = true
			} else {
				s["mdl-cell--order-"+i] = true
			}
		}
	}
	if c.Hide {
		for _, v := range modes(c.HideMode) {
			if v != "" {
				s["mdl-cell--hide"+"-"+v] = true
			} else {
				s["mdl-cell--hide-"] = true
			}
		}
	}
	if c.Stretch {
		s["mdl-cell-stretch"] = true
	}
	if c.Top {
		s["mdl-cell-top"] = true
	}
	if c.Middle {
		s["mdl-cell-middle"] = true
	}
	if c.Bottom {
		s["mdl-cell-bottom"] = true
	}
	return s
}

func toString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
