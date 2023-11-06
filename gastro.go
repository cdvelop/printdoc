package printdoc

import (
	"fmt"
	"strings"

	"github.com/cdvelop/model"
	"github.com/jung-kurt/gofpdf"
)

func (p *PrintDoc) BuildPanesdoscopiaPDF() error {
	fmt.Printf("CREando pdf ruta: %v", p.StoragePath)
	p.margin_left = 30
	p.text_size = 9

	p.pdf = gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	tr := p.pdf.UnicodeTranslatorFromDescriptor("")
	p.pdf.AddPage()

	//LOGO
	p.pdf.ImageOptions(p.BusinessLogoPath, p.margin_left, 20, 55, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	const start_patient_info = 97
	//FECHA
	p.pdf.SetFont("Arial", "", p.text_size)
	// pdf.MoveTo(p.margin_left, 30)
	p.pdf.Text(505, start_patient_info, "Fecha: "+p.DayAttention)

	//TITULO
	const title_line = 36
	p.pdf.SetFont("Arial", "B", 14)
	_, lineHt := p.pdf.GetFontSize()
	p.pdf.MoveTo(p.margin_left, title_line)
	p.pdf.MultiCell(0, lineHt, "UNIDAD DE GASTROSCOPIA", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	// LINEA TITULO
	p.pdf.SetDrawColor(10, 76, 10)
	p.pdf.Line(91, title_line+lineHt, 520, title_line+lineHt)

	// SERVICIO TITULO
	_, lineHt = p.pdf.GetFontSize()
	p.pdf.SetFont("Arial", "", 13)
	p.pdf.MoveTo(p.margin_left, title_line+lineHt+2)
	p.pdf.MultiCell(0, lineHt, tr(strings.ToUpper(p.ServiceName)), gofpdf.BorderNone, gofpdf.AlignCenter, false)

	// PACIENTE INFO
	p.pdf.SetFont("Arial", "B", p.text_size)
	p.pdf.Text(p.margin_left, start_patient_info, "Paciente:")

	//datos paciente
	p.pdf.SetFont("Arial", "", p.text_size)
	_, lineHt = p.pdf.GetFontSize()
	p.pdf.Text(p.margin_left+2, start_patient_info+lineHt+2, tr(fmt.Sprintf("%v, Run: %v, %v años.", strings.ToUpper(p.PatientName), p.PatientRun, p.PatientAge)))

	//procedimiento
	p.pdf.SetFont("Arial", "B", p.text_size)
	_, lineHt = p.pdf.GetFontSize()
	p.pdf.Text(p.margin_left, start_patient_info+lineHt*3, "Procedimiento:")

	p.pdf.SetFont("Arial", "", p.text_size)
	_, lineHt = p.pdf.GetFontSize()
	p.pdf.MoveTo(p.margin_left, start_patient_info+lineHt*3+2)
	p.pdf.MultiCell(0, lineHt, tr(p.Reason), gofpdf.BorderNone, gofpdf.AlignLeft, false)

	// CHECK TEST UREASA
	p.pdf.SetFillColor(10, 76, 10)
	p.pdf.Polygon([]gofpdf.PointType{
		{X: 554, Y: 143.15},
		{X: 558, Y: 139.15},
		{X: 563.6, Y: 144.75},
		{X: 575.6, Y: 132.75},
		{X: 579.6, Y: 136.75},
		{X: 563.6, Y: 152.75},
	}, "D")

	p.pdf.Text(485, start_patient_info+lineHt*5, "TEST UREASA")

	p.insertImages()

	p.gastroDescription()

	// FIRMA
	p.pdf.SetFont("Arial", "B", 10)
	// _, lineHt = pdf.GetFontSize()
	p.pdf.MoveTo(p.margin_left, 700)
	p.pdf.MultiCell(0, 0, tr(strings.ToUpper(p.StaffName)), gofpdf.BorderNone, gofpdf.AlignCenter, false)

	p.footerPdfDocument()

	// drawGrid(pdf)

	err := p.pdf.OutputFileAndClose(p.StoragePath + "/test.pdf")
	if err != nil {
		return model.Error("error al crear pdf", err)
	}

	return nil
}
