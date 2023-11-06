package printdoc

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func (p *PrintDoc) insertImages() {
	width_page, _ := p.pdf.GetPageSize()
	const separation = 5
	const img_start = 160
	const height_img = 100
	var width_img = (width_page - ((separation * 2) + (p.margin_left * 2))) / 3

	var margin_left_start float64
	var margin_top_start float64

	for i, img_path := range p.MedicalImagesPaths {

		fmt.Printf("IMÁGENES: %v", img_path)
		// for i := 0; i < len(new_directory); i++ {

		if i < 3 { // primera fila de 3
			if i == 0 { // inicio
				margin_top_start = img_start
				margin_left_start = p.margin_left
			} else {
				margin_left_start += width_img + separation
			}

		} else if i < 6 { // segunda fila de 3
			if i == 3 { // reset fila 2
				margin_top_start = img_start + height_img + separation
				margin_left_start = p.margin_left
			} else {
				margin_left_start += width_img + separation
			}

		}

		p.pdf.ImageOptions(img_path, margin_left_start, margin_top_start, width_img, height_img, false, gofpdf.ImageOptions{
			ReadDpi: true,
		}, 0, "")

	}

}
