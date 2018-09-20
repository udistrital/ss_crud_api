package models

import (
	"github.com/astaxie/beego/orm"
)

type TrTipoUpc struct {
	TipoUpc []*TipoUpc
}

// RegistrarPagos Transancción para el registro de pagos
func RegistrarTiposUpc(m *TrTipoUpc) (alerta Alert, err error) {
	o := orm.NewOrm()
	o.Begin()
	_, err = o.InsertMulti(len(m.TipoUpc), m.TipoUpc)
	if err == nil {
		alerta = Alert{Type: "success", Code: "Ok", Body: err}
		o.Commit()
	} else {
		alerta = Alert{Type: "error", Code: "E_TIPOS_UPC", Body: err}
		o.Rollback()
	}
	return alerta, err
}
