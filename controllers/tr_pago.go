package controllers

import (
	"encoding/json"

	"github.com/udistrital/ss_crud_api/models"

	"github.com/astaxie/beego"
)

// TrPeriodoPagoController ...
type TrPeriodoPagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrPeriodoPagoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title TrPeriodoPago
// @Description Registra los pagos asociados a un nuevo periodo_pago
// @Param	body		body 	models.TrPeriodoPago	true	"body for TrPeriodoPago content"
// @Success 201 {object} msg
// @Failure 403 :id is not int
// @router / [post]
func (c *TrPeriodoPagoController) Post() {
	var v models.TrPeriodoPago
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		alerta, err := models.RegistrarPagos(&v)
		if err == nil {
			c.Ctx.Output.SetStatus(201)
		}
		c.Data["json"] = alerta
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
