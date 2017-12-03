package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrPeriodoPago struct {
	PeriodoPago PeriodoPago
	Pago        []*Pago
}

// Transancción para el registro de pagos
func RegistrarPagos(m *TrPeriodoPago) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(m.PeriodoPago)

	if err == nil {
		err = o.Rollback()
		fmt.Println("Pagos: ", m.Pago)
		for _, v := range m.Pago {
			v.PeriodoPago = &PeriodoPago{Id: int(id)}

			if _, err = o.Insert(v); err != nil {
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
