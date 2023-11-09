module github.com/cdvelop/printdoc

go 1.20

require (
	github.com/cdvelop/model v0.0.67
	github.com/cdvelop/object v0.0.23
	github.com/jung-kurt/gofpdf v1.16.2
)

require github.com/cdvelop/strings v0.0.3 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/object => ../object
