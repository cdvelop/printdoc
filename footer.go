package printdoc

import (
	"fmt"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func (p *PrintDoc) footerPdfDocument() {
	_, height_page := p.pdf.GetPageSize()

	p.pdf.SetFont("Arial", "", 8)
	p.pdf.SetTextColor(169, 169, 169)
	// _, lineHt := pdf.GetFontSize()
	p.pdf.MoveTo(p.margin_left, height_page-60)

	info := fmt.Sprintf("Plataforma %v-%v - %v - %v - Fono: %v", p.App_name, p.App_version, p.Business_name, p.Business_address, p.Business_phone)

	tr := p.pdf.UnicodeTranslatorFromDescriptor("")

	p.pdf.MultiCell(0, 0, tr(strings.ToUpper(info)), gofpdf.BorderNone, gofpdf.AlignCenter, false)

}
