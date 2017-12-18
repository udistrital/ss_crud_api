package models

import "github.com/astaxie/beego/orm"

type TrPeriodoPago struct {
	PeriodoPago *PeriodoPago
	Pagos       []*Pago
}

// RegistrarPagos Transancción para el registro de pagos
func RegistrarPagos(m *TrPeriodoPago) (alerta Alert, err error) {
	o := orm.NewOrm()
	o.Begin()

	if id, err := o.Insert(m.PeriodoPago); err == nil {
		for _, v := range m.Pagos {
			v.PeriodoPago = &PeriodoPago{Id: int(id)}
			if _, err = o.Insert(v); err != nil {
				alerta = Alert{Type: "error", Code: "E_PAGOS", Body: err}
				o.Rollback()
				break
			}
		}
		alerta = Alert{Type: "success", Code: "Ok", Body: err}
		o.Commit()
	} else {

		alerta = Alert{Type: "error", Code: "E_PERIODO", Body: err}
		o.Rollback()
	}

	return alerta, err
}
