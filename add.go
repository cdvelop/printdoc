package printdoc

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

func New(o *model.Object, h *model.Handlers) (*PrintDoc, error) {
	p := PrintDoc{}

	err := object.New(&p, o.Module, h)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
