package api

import "github.com/labstack/echo/v4"

//routes for users
//the section that is ":id" is where the 1 id the user is placed to make the search
//the routes are divided by users and saved element, to give order to the requests

func (a *API) RegisterRoutes(e *echo.Echo) {
	users := e.Group("/user")
	users.POST("/create_user", a.Create_User)
	users.GET("/userid_by_email", a.Get_userid_Byemail)
	users.GET("/password_by_email", a.Get_password_Byemail)
	users.GET("/user_by_id", a.Read_userByid)
	users.PUT("/update_userbyid", a.Update_userByid)
	users.DELETE("/delete_userbyid", a.Delete_userByid)

	request := e.Group("/cotizacion")
	request.POST("/crear_cotizacion", a.Create_request)
	request.POST("/crear_tipo_de_estado", a.Create_requesttype)
	request.GET("/estado_por_id", a.Get_requeststatus_Byid)
	request.PUT("/actualizar_estado_por_id", a.Update_request_status_Byid)
	request.DELETE("/borrar_cotizacion_por_id", a.Delete_requests_ByUserid)
	request.PUT("/actualizar_cotizacion", a.Update_request_URL)
	request.GET("/obtener_cotizacion", a.Get_cotizacion_data)
	request.GET("/obtener_id_cotizacion", a.Get_requestId_byUserid)

	password := e.Group("/password")
	password.POST("/RevisarPassword", a.RevisarPassword)

}
