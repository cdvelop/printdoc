package printdoc

import "github.com/cdvelop/model"

func (p PrintDoc) Create(u *model.User, data ...map[string]string) error {

	p.App.Log(" BACKEND IMPRIMIR: ", data)

	return nil
}
