package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// RegistrarDetalleNovedad transacción para realizar varios registros detallados de las novedades de seguridad social
func RegistrarDetalleNovedad(m interface{}) (alerta Alert, err error) {
	var (
		detalles []DetalleNovedadSeguridadSocial
	)
	o := orm.NewOrm()
	for _, novedad := range m.([]map[string]interface{}) {

		detalle := DetalleNovedadSeguridadSocial{
			Descripcion:              novedad["Descripcion"].(string),
			ConceptoNominaPorPersona: int(novedad["Id"].(float64)),
		}
		if novedad["Descripcion"] != nil {
			detalle.Descripcion = novedad["Descripcion"].(string)
		}

		detalles = append(detalles, detalle)
	}

	_, err = o.InsertMulti(len(detalles), detalles)

	alerta = Alert{Type: "success", Code: "1", Body: ""}

	if err != nil {
		alerta = Alert{Type: "error", Code: "0", Body: err.Error()}
	}
	return
}

func PrintMessage(message interface{}) {
	fmt.Printf("**** %s ****\n", message)
}
