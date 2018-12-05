// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/ss_crud_api/controllers"
)

func init() {
	// auditoria.InitMiddleware()
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_pago",
			beego.NSInclude(
				&controllers.TipoPagoController{},
			),
		),

		beego.NSNamespace("/periodo_pago",
			beego.NSInclude(
				&controllers.PeriodoPagoController{},
			),
		),

		beego.NSNamespace("/pago",
			beego.NSInclude(
				&controllers.PagoController{},
			),
		),

		beego.NSNamespace("/upc_adicional",
			beego.NSInclude(
				&controllers.UpcAdicionalController{},
			),
		),

		beego.NSNamespace("/zona_upc",
			beego.NSInclude(
				&controllers.ZonaUpcController{},
			),
		),

		beego.NSNamespace("/rango_edad_upc",
			beego.NSInclude(
				&controllers.RangoEdadUpcController{},
			),
		),

		beego.NSNamespace("/tipo_upc",
			beego.NSInclude(
				&controllers.TipoUpcController{},
			),
		),

		beego.NSNamespace("/aportante",
			beego.NSInclude(
				&controllers.AportanteController{},
			),
		),

		beego.NSNamespace("/tr_periodo_pago",
			beego.NSInclude(
				&controllers.TrPeriodoPagoController{},
			),
		),
		beego.NSNamespace("/estado_seguridad_social",
			beego.NSInclude(
				&controllers.EstadoSeguridadSocialController{},
			),
		),
		beego.NSNamespace("/beneficiarios",
			beego.NSInclude(
				&controllers.BeneficiariosController{},
			),
		),
		beego.NSNamespace("/detalle_novedad_seguridad_social",
			beego.NSInclude(
				&controllers.DetalleNovedadSeguridadSocialController{},
			),
		),
		beego.NSNamespace("/tipo_novedad_seguridad_social",
			beego.NSInclude(
				&controllers.TipoNovedadSeguridadSocialController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
