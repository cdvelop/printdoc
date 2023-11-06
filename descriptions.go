package printdoc

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func (p *PrintDoc) gastroDescription() {
	tr := p.pdf.UnicodeTranslatorFromDescriptor("")

	// DESCRIPCIÓN
	const start_rectangle = 380

	p.pdf.SetFont("Arial", "BI", 10)
	p.pdf.MoveTo(p.margin_left, start_rectangle)
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.MultiCell(0, lineHt, tr("Descripción"), gofpdf.BorderNone, gofpdf.AlignCenter, false)

	// RECTÁNGULO DESCRIPCIÓN
	p.pdf.Rect(p.margin_left, start_rectangle, 550, 270, "D")
	p.pdf.Line(p.margin_left, start_rectangle+lineHt, 580, start_rectangle+lineHt)

	words := strings.Split(p.Diagnostic, "\n")
	var margin_left_text_description = p.margin_left + 2

	var star_text_description = start_rectangle + lineHt*2

	for _, word_line := range words {
		word_line = strings.TrimSpace(word_line)
		if word_line != "" {
			// fmt.Printf("PALABRA: [%v]\n", word_line)
			if strings.ToUpper(word_line) == word_line {
				// fmt.Printf("PALABRA: [%v] ESTA EN MAYÚSCULA\n", word_line)
				//cambio el estilo de letra a negrita
				p.pdf.SetFont("Arial", "B", p.text_size)
				_, lineHt = p.pdf.GetFontSize()

				// escribo titulo
				p.pdf.Text(margin_left_text_description, star_text_description, tr(word_line))

			} else {
				// cambio el estilo de letra a normal
				p.pdf.SetFont("Arial", "", p.text_size)
				_, lineHt = p.pdf.GetFontSize()
				p.pdf.MoveTo(margin_left_text_description, star_text_description)
				p.pdf.MultiCell(0, lineHt, tr(word_line), gofpdf.BorderNone, gofpdf.AlignLeft, false)

				// set posición
				star_text_description += lineHt * 2
			}

		} else {
			// aca hay un salto de linea
			star_text_description += lineHt * 2
		}
	}

	// CONCLUSION
	_, lineHt = p.pdf.GetFontSize()
	star_text_description += lineHt
	p.pdf.Line(p.margin_left, star_text_description-1, 580, star_text_description-1)

	p.pdf.Line(p.margin_left, star_text_description+lineHt, 580, star_text_description+lineHt)
	p.pdf.SetFont("Arial", "BI", 10)
	p.pdf.MoveTo(p.margin_left, star_text_description)
	p.pdf.MultiCell(0, lineHt, tr("Conclusión"), gofpdf.BorderNone, gofpdf.AlignCenter, false)

	// cambio el estilo de letra a normal
	p.pdf.SetFont("Arial", "", p.text_size)
	_, lineHt = p.pdf.GetFontSize()
	p.pdf.MoveTo(margin_left_text_description, star_text_description+lineHt*2)
	p.pdf.MultiCell(0, lineHt, tr(p.Prescription), gofpdf.BorderNone, gofpdf.AlignLeft, false)

}
