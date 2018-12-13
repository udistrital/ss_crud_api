package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/udistrital/ss_crud_api/models"

	"github.com/astaxie/beego"
)

// PagoController operations for Pago
type PagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *PagoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetPagos", c.GetPagos)
	c.Mapping("GetParamertroEntidad", c.GetParametroEntidad)
}

// GetParametroEntidad ...
// @Title GetParametroEntidad
// @Description get Pago by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Pago
// @Failure 404 not found resource
// @router /GetParametroEntidad/:id [get]
func (c *PagoController) GetParametroEntidad() {

	idStr := c.Ctx.Input.Param(":id")
	data := []byte(`[
		{
		  "id": 1,
		  "nit": 800088702,
		  "nombre": "EPS Sura",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS010"
		},
		{
		  "id": 2,
		  "nit": 800130907,
		  "nombre": "Salud Total EPS",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS002"
		},
		{
		  "id": 3,
		  "nit": 800140949,
		  "nombre": "Cafesalud EPS",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS003"
		},
		{
		  "id": 5,
		  "nit": 800250119,
		  "nombre": "SaludCoop EPS",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS013"
		},
		{
		  "id": 6,
		  "nit": 830003564,
		  "nombre": "Famisanar EPS Cafam",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS017"
		},
		{
		  "id": 7,
		  "nit": 830009783,
		  "nombre": "Cruz Blanca EPS",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS023"
		},
		{
		  "id": 8,
		  "nit": 860066942,
		  "nombre": "Compensar EPS",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS008"
		},
		{
		  "id": 9,
		  "nit": 900156264,
		  "nombre": "NuevaEPS",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "EPS037"
		},
		{
		  "id": 10,
		  "nit": 800256161,
		  "nombre": "ARL Sura",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 2,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "14-28"
		},
		{
		  "id": 11,
		  "nit": 860002184,
		  "nombre": "Axa Colpatria Seguros SA",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 2,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": ""
		},
		{
		  "id": 12,
		  "nit": 860002503,
		  "nombre": "Seguros Bolivar SA",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 2,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "14-7"
		},
		{
		  "id": 13,
		  "nit": 860011153,
		  "nombre": "Positiva Compania Seguros",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 2,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "14-23"
		},
		{
		  "id": 14,
		  "nit": 860022137,
		  "nombre": "Seguros de Vida Aurora SA",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 2,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "14-8"
		},
		{
		  "id": 15,
		  "nit": 860028415,
		  "nombre": "La Equidad Seguros GOC",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 2,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "14-29"
		},
		{
		  "id": 16,
		  "nit": 800216278,
		  "nombre": "Pensiones de Antioquia",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "25-7"
		},
		{
		  "id": 17,
		  "nit": 800224808,
		  "nombre": "Porvenir",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "230301"
		},
		{
		  "id": 4,
		  "nit": 800227940,
		  "nombre": "Colfondos",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "231001"
		},
		{
		  "id": 18,
		  "nit": 800229739,
		  "nombre": "Protección",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "230201"
		},
		{
		  "id": 19,
		  "nit": 800231967,
		  "nombre": "Horizonte",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "230501"
		},
		{
		  "id": 20,
		  "nit": 800253055,
		  "nombre": "Skandia",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "230901"
		},
		{
		  "id": 21,
		  "nit": 830125132,
		  "nombre": "Skandia Alternativo",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "230904"
		},
		{
		  "id": 22,
		  "nit": 860007379,
		  "nombre": "Caxdac",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "25-2"
		},
		{
		  "id": 23,
		  "nit": 899999734,
		  "nombre": "Fonprecon",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "25-3"
		},
		{
		  "id": 24,
		  "nit": 900336004,
		  "nombre": "Colpensiones",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "25-14"
		},
		{
		  "id": 32,
		  "nit": "",
		  "nombre": "Caprecom",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "25-4"
		},
		{
		  "id": 33,
		  "nit": "",
		  "nombre": "Capresoca",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 34,
		  "nit": "",
		  "nombre": "Colmédica",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS001"
		},
		{
		  "id": 35,
		  "nit": "",
		  "nombre": "Comfenalco Antioquia",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS009 "
		},
		{
		  "id": 36,
		  "nit": "",
		  "nombre": "Comfenalco Valle",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS012"
		},
		{
		  "id": 37,
		  "nit": "",
		  "nombre": "Convida ARS",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 38,
		  "nit": "",
		  "nombre": "Coomeva EPS",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS016"
		},
		{
		  "id": 39,
		  "nit": "",
		  "nombre": "Humana Vivir",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS014"
		},
		{
		  "id": 40,
		  "nit": "",
		  "nombre": "Instituto de los seguros sociales",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS006"
		},
		{
		  "id": 41,
		  "nit": "",
		  "nombre": "Salud colmena",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 42,
		  "nit": "",
		  "nombre": "Salud Colpatria",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS015"
		},
		{
		  "id": 43,
		  "nit": "",
		  "nombre": "Salud colombia EPS SA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 44,
		  "nit": "",
		  "nombre": "Salud vida",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS033"
		},
		{
		  "id": 45,
		  "nit": "",
		  "nombre": "Sanitas",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS005"
		},
		{
		  "id": 46,
		  "nit": "",
		  "nombre": "Selvasalud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 47,
		  "nit": "",
		  "nombre": "Solsalud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS026"
		},
		{
		  "id": 48,
		  "nit": "",
		  "nombre": "SOS EPS",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 49,
		  "nit": "",
		  "nombre": "Susalud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 50,
		  "nit": "",
		  "nombre": "Mutual ser",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 51,
		  "nit": "",
		  "nombre": "Coosalud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 52,
		  "nit": "",
		  "nombre": "Emssanar",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 53,
		  "nit": "",
		  "nombre": "Aliansalud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS001"
		},
		{
		  "id": 54,
		  "nit": "",
		  "nombre": "Asmet salud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 55,
		  "nit": "",
		  "nombre": "Servicio occidental de salud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS018"
		},
		{
		  "id": 56,
		  "nit": "",
		  "nombre": "Ambuq",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 57,
		  "nit": "",
		  "nombre": "Savia salud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 58,
		  "nit": "",
		  "nombre": "Capital salud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 59,
		  "nit": "",
		  "nombre": "Comparta",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 60,
		  "nit": "",
		  "nombre": "Calisalud",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 61,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar COMFENALCO ANTIOQUIA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF03"
		},
		{
		  "id": 62,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Antioquia COMFAMA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF04"
		},
		{
		  "id": 63,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar CAJACOPI ATLANTICO",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF05"
		},
		{
		  "id": 64,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Barranquilla COMBARRANQUILLA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF06"
		},
		{
		  "id": 65,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar COMFAMILIAR DEL ATLANTICO",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF07"
		},
		{
		  "id": 66,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Fenalco - Andi COMFENALCO CARTAGENA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF08"
		},
		{
		  "id": 67,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Cartagena",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF09"
		},
		{
		  "id": 68,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Caldas",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF11"
		},
		{
		  "id": 69,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de la Dorada - COMFAMILIAR",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF12"
		},
		{
		  "id": 70,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Caquetá - COMFACA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 71,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Cauca - COMFACAUCA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF14"
		},
		{
		  "id": 72,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Cesar COMFACESAR",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF15"
		},
		{
		  "id": 73,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Córdoba COMFACOR",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF16"
		},
		{
		  "id": 74,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar AFIDRO",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF18"
		},
		{
		  "id": 75,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar ASFAMILIAS",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF20"
		},
		{
		  "id": 76,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Fenalco COMFENALCO CUNDINAMARCA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF23"
		},
		{
		  "id": 77,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Cundinamarca - COMFACUNDI",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF26"
		},
		{
		  "id": 78,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Girardot COMGIRARDOT (en liquidación)",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF28"
		},
		{
		  "id": 79,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Chocó",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF29"
		},
		{
		  "id": 80,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de la Guajira",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF30"
		},
		{
		  "id": 81,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Huila - COMFAMILIAR",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF32"
		},
		{
		  "id": 82,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Magdalena",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF33"
		},
		{
		  "id": 83,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Nariño",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF35"
		},
		{
		  "id": 84,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Oriente Colombiano COMFAORIENTE",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF36"
		},
		{
		  "id": 85,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Norte de Santander COMFANORTE",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF37"
		},
		{
		  "id": 86,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Barrancabermeja CAFABA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF38"
		},
		{
		  "id": 87,
		  "nit": "",
		  "nombre": "Caja Santandereana de Subsidio Familiar CAJASAN",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF39"
		},
		{
		  "id": 88,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar COMFENALCO SANTANDER",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF40"
		},
		{
		  "id": 89,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Sucre",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF41"
		},
		{
		  "id": 90,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Quindío COMFAMILIAR",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF42"
		},
		{
		  "id": 91,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Fenalco COMFENALCO QUINDIO",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF43"
		},
		{
		  "id": 92,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Risaralda- COMFAMILIAR RISARALDA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF44"
		},
		{
		  "id": 93,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Sur del Tolima CAFASUR",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF46"
		},
		{
		  "id": 94,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Honda COMFAHONDA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 95,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Tolima COMFATOLIMA",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF48"
		},
		{
		  "id": 96,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Fenalco del Tolima - COMFENALCO",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF50"
		},
		{
		  "id": 97,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Buenaventura",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF51"
		},
		{
		  "id": 25,
		  "nit": 800231969,
		  "nombre": "Caja de Compensación Familiar Campesina COMCAJA",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "CCF68"
		},
		{
		  "id": 27,
		  "nit": 860013570,
		  "nombre": "Caja de Compensación Familiar CAFAM",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "CCF21"
		},
		{
		  "id": 28,
		  "nit": 860066943,
		  "nombre": "Caja de Compensación Familiar COMPENSAR",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "CCF24"
		},
		{
		  "id": 30,
		  "nit": 891800213,
		  "nombre": "Caja de Compensación Familiar de Boyacá - COMFABOY",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "CCF10"
		},
		{
		  "id": 98,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar Comfenalco del Valle del Cauca - COMFENALCO VALLE",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF56"
		},
		{
		  "id": 99,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Valle del Cauca COMFAMILIAR ANDI - COMFANDI",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF57"
		},
		{
		  "id": 100,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Cartago - Comfacartago",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF59"
		},
		{
		  "id": 101,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar Comfamiliares Unidas del Valle COMFAUNION",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF61"
		},
		{
		  "id": 102,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Tuluá - Comfamiliar de Tuluá",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF62"
		},
		{
		  "id": 103,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Putumayo - COMFAMILIAR PUTUMAYO",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF63"
		},
		{
		  "id": 104,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de San Andrés y Providencia, Islas CAJASAI",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF64"
		},
		{
		  "id": 105,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Amazonas CAFAMAZ",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF65"
		},
		{
		  "id": 106,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar de Arauca COMFIAR",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF67"
		},
		{
		  "id": 107,
		  "nit": "",
		  "nombre": "Caja de Compensación Familiar del Casanare - COMFACASANARE",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "CCF69"
		},
		{
		  "id": 26,
		  "nit": 860007336,
		  "nombre": "Caja Colombiana de Subsidio Familiar COLSUBSIDIO",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "CCF22"
		},
		{
		  "id": 29,
		  "nit": 890900840,
		  "nombre": "Caja de Compensación Familiar Camacol COMFAMILIAR CAMACOL",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "CCF02"
		},
		{
		  "id": 31,
		  "nit": 892000146,
		  "nombre": "Caja de Compensación Familiar Regional del Meta COFREM",
		  "fecha_creacion": "2016-08-24",
		  "id_tipo_entidad": 4,
		  "id_estado": 1,
		  "fecha_modificacion": "2016-08-24",
		  "Codigo": "CCF34"
		},
		{
		  "id": 108,
		  "nit": "",
		  "nombre": "Fondos de Pensiones y Cesantía Santander S.A.",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 109,
		  "nit": "",
		  "nombre": "Compañía Colombiana Administradora de Fondos de Pensiones y Cesantías S.A. COLFONDOS",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "231001"
		},
		{
		  "id": 110,
		  "nit": "",
		  "nombre": "Instituto de Seguros Sociales I.S.S. Riesgos Profesionales",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 111,
		  "nit": "",
		  "nombre": "Instituto de Seguros Sociales I.S.S. Pensiones",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "25-oct"
		},
		{
		  "id": 112,
		  "nit": "",
		  "nombre": "Caja de Auxilios y Prestaciones de la Asociación Colombiana de Aviadores Civiles Acdac \"Caxdac",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "25-2"
		},
		{
		  "id": 113,
		  "nit": "",
		  "nombre": "Caja de Previsión Social de Comunicaciones -Caprecom-",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "25-4"
		},
		{
		  "id": 114,
		  "nit": "",
		  "nombre": "Caja Nacional de Previsión Social - CAJANAL E.I.C.E",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "25-8"
		},
		{
		  "id": 115,
		  "nit": "",
		  "nombre": "Caja de Previsión Social de Los Trabajadores de la Universidad del Cauca",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": ""
		},
		{
		  "id": 116,
		  "nit": "",
		  "nombre": "Fondo de solidaridad pensional",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "ISSFSP"
		},
		{
		  "id": 118,
		  "nit": "",
		  "nombre": "Famisanar EPS Colsubsidio",
		  "fecha_creacion": "2016-01-11",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-11",
		  "Codigo": "EPS017"
		},
		{
		  "id": 119,
		  "nit": "",
		  "nombre": "No Aplica",
		  "fecha_creacion": "2017-01-30",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-30",
		  "Codigo": ""
		},
		{
		  "id": 120,
		  "nit": "",
		  "nombre": "Régimen Especial",
		  "fecha_creacion": "2017-01-30",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2017-01-30",
		  "Codigo": ""
		},
		{
		  "id": 121,
		  "nit": 901097473,
		  "nombre": "Medimás EPS",
		  "fecha_creacion": "2018-01-09",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2018-01-09",
		  "Codigo": ""
		},
		{
		  "id": 122,
		  "nit": "",
		  "nombre": "EPS Servisalud",
		  "fecha_creacion": "2018-07-20",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2018-07-20",
		  "Codigo": ""
		},
		{
		  "id": 123,
		  "nit": 830053105,
		  "nombre": "Fondo Prestacional del Magisterio FOMAG",
		  "fecha_creacion": "2018-07-20",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2018-07-20",
		  "Codigo": ""
		},
		{
		  "id": 124,
		  "nit": 900178724,
		  "nombre": "EPS MedPlus",
		  "fecha_creacion": "2018-08-20",
		  "id_tipo_entidad": 1,
		  "id_estado": 1,
		  "fecha_modificacion": "2018-08-20",
		  "Codigo": ""
		},
		{
		  "id": 125,
		  "nit": 860525148,
		  "nombre": "Fiduprevisora",
		  "fecha_creacion": "2018-08-20",
		  "id_tipo_entidad": 3,
		  "id_estado": 1,
		  "fecha_modificacion": "2018-08-20",
		  "Codigo": ""
		}
	  ]`)

	//   json.Unmarshal
	var respuesta []map[string]interface{}
	var objeto map[string]interface{}
	json.Unmarshal(data, &respuesta)
	for _, value := range respuesta {
		if idStr == fmt.Sprint(value["id"].(float64)) {
			objeto = value
		}
	}
	c.Data["json"] = objeto
	c.ServeJSON()

}

// Post ...
// @Title Post
// @Description create Pago
// @Param	body		body 	models.Pago	true		"body for Pago content"
// @Success 201 {int} models.Pago
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *PagoController) Post() {
	var v models.Pago
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddPago(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
		beego.Error(err)
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Pago by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Pago
// @Failure 404 not found resource
// @router /:id [get]
func (c *PagoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetPagoById(id)
	if err != nil {
		beego.Error(err)
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Pago
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Pago
// @Failure 404 not found resource
// @router / [get]
func (c *PagoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllPago(query, fields, sortby, order, offset, limit)
	if err != nil {
		beego.Error(err)
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Pago
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Pago	true		"body for Pago content"
// @Success 200 {object} models.Pago
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *PagoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Pago{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdatePagoById(&v); err == nil {
			c.Data["json"] = v
		} else {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
		beego.Error(err)
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Pago
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *PagoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeletePago(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		beego.Error(err)
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetPagos ...
// @Title Get Pagos
// @Description trae los pagos de seguridad de un periodo pago
// @Param	idPeriodoPago		id del periodo_pago correspondiente
// @Success 200 {object} models.PruebaPago
// @Failure 403 :idPeriodoPago is empty
// @router GetPagos/:idPeriodoPago [get]
func (c *PagoController) GetPagos() {
	idStr := c.Ctx.Input.Param(":idPeriodoPago")
	id, _ := strconv.Atoi(idStr)
	v, err := models.SumarPagos(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// PagosPorPeriodoPago ...
// @Title PagosPorPeriodoPago
// @Description lista pagos agrupados por detalle liquidacion que corresponda a un periodo de pago.
// @Param	idPeriodoPago	query	string	false	"id de periodo pago"
// @Success 201 {object} models.Alert
// @Failure 403 body is empty
// @router /PagosPorPeriodoPago [get]
func (c *PagoController) PagosPorPeriodoPago() {
	idPeriodoPago, err1 := c.GetInt("idPeriodoPago")
	if err1 == nil {
		if data, err := models.PagosPorPeriodoPago(idPeriodoPago); err == nil {
			c.Data["json"] = data
		} else {
			c.Data["json"] = models.Alert{Code: "E_0458", Body: "Not enough parameter", Type: "error"}
		}
	}
	c.ServeJSON()
}
