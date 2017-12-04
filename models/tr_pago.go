package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrPeriodoPago struct {
	PeriodoPago *PeriodoPago
	Pagos       []*Pago
}

// Transancción para el registro de pagos
func RegistrarPagos(m *TrPeriodoPago) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()

	if id, err := o.Insert(m.PeriodoPago); err == nil {
		for _, v := range m.Pagos {
			v.PeriodoPago = &PeriodoPago{Id: int(id)}
			if _, err = o.Insert(v); err != nil {
				fmt.Println("rollback")
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Error registrado cada pago")
				break
			}
		}
		o.Commit()
		alerta = append(alerta, "Ok")
	} else {
		o.Rollback()
		alerta[0] = "error"
		alerta = append(alerta, "Error grande")
	}

	return alerta, err
}
