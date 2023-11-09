package printdoc

import (
	"github.com/cdvelop/model"
	"github.com/jung-kurt/gofpdf"
)

type PrintDoc struct {
	pdf *gofpdf.Fpdf

	App    *model.Handlers
	Object *model.Object

	text_size   float64
	margin_left float64

	StoragePath string // ej: ./pdf_files

	// static data:
	App_name           string
	App_version        string
	Business_name      string
	Business_address   string
	Business_phone     string
	Business_logo_path string //ej: ./app_files/logo.png

	//dynamic fields data:
	StaffName          string
	ServiceName        string
	Reason             string
	Diagnostic         string
	Prescription       string
	MedicalImagesPaths []string

	DayAttention string //ej: 21-05-2023

	//patient data:
	PatientName string
	PatientRun  string
	PatientAge  string
}
