package printdoc

import "github.com/jung-kurt/gofpdf"

type PrintDoc struct {
	pdf         *gofpdf.Fpdf
	text_size   float64
	margin_left float64

	StoragePath string // ej: ./pdf_files

	// static data:
	AppName          string
	AppVersion       string
	BusinessName     string
	BusinessAddress  string
	BusinessPhone    string
	BusinessLogoPath string //ej: ./app_files/logo.png

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
