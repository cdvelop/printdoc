package printdoc_test

import (
	"testing"

	"github.com/cdvelop/printdoc"
)

func TestBuildPanesdoscopiaPDF(t *testing.T) {
	const test_files = "./test_files"

	new := printdoc.PrintDoc{
		StoragePath:      test_files,
		AppName:          "test app",
		AppVersion:       "0.0.1",
		BusinessName:     "Negocio Test",
		BusinessAddress:  "Calle sin nombre 5422",
		BusinessPhone:    "55-256115",
		BusinessLogoPath: test_files + "/logo_test.jpg",
		StaffName:        "Dr. KARLA ACERO SeRRANO",
		ServiceName:      "panesdoscopia",
		Reason: `Referido por      : dra. coñtreras
		Motivo               : disfagia
		Premedicacion  : dormonid-lidocainal`,
		Diagnostic: `ESOFAGO
		Calibre, peristalsis y distensibilidad conservada. Mucosa con erosiones eritematosas lineales. Union Esofagogástrica a 39cms de la arcada dentaria superior, coincide con impronta diafragmática.
	  
	  ESTOMAGO
	  Peristalsis y distensibilidad conservadña. Fundus, cuerpo y antro: mucosa con puntillado erosivo e hiperemico. Se toma muestra para yl Hepytest.
	  
	  PILORO
	   Centrico y permeable
	  
	  DUODENO
	   Peristalsis y distensibilidad conservada. Bulbo y 2da porcion: de aspecto endoscopico normal.`,
		Prescription: `1.- ESOFAGITIS PEPTICA GRADO B
		2.- GASTROPATIA EROSIVA E HIPEREMICA`,
		MedicalImagesPaths: []string{
			test_files + "/img1.jpg",
			test_files + "/img2.jpeg",
			test_files + "/img3.jpeg",
			test_files + "/img4.jpg",
			test_files + "/img5.jpg",
			test_files + "/img6.jpg",
		},
		DayAttention: "03-09-2022",
		PatientName:  "SILVIA SOTO GARRIDO GARRIDO",
		PatientRun:   "10.058.978-K",
		PatientAge:   "80",
	}

	err := new.BuildPanesdoscopiaPDF()
	if err != nil {
		t.Fatal(err)
	}

}
