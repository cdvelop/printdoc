package printdoc

import (
	"fmt"
)

func (p *PrintDoc) drawGrid() {
	w, h := p.pdf.GetPageSize()
	for x := 0.0; x < w; x = x + (w / 20.0) {
		p.pdf.SetFont("courier", "", 12)
		p.pdf.SetTextColor(80, 80, 80)
		p.pdf.SetDrawColor(200, 200, 200)
		p.pdf.Line(x, 0, x, h)
		_, lineHt := p.pdf.GetFontSize()
		p.pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}

	for y := 0.0; y < h; y = y + (w / 20.0) {
		p.pdf.Line(0, y, w, y)
		p.pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
