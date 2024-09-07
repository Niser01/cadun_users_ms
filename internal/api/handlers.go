package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/api/dtos"
	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/views"
	"github.com/labstack/echo/v4"
)

//Handlers for the API
//There is a handdler for each route that was created, each handler calls the respective view that is defined
//If there is an error with the request each handler throws an exception
//The handlers that recevies a parametes through the URL, that parameter is recived by the echo context and used in the query
//All the other handlers use the dtos parameters since the parameters are given by a JSON format.
//param es el parametro tomado desde la URL /users/123

func (a *API) Create_User(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Create_User{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Create_user(ctx, parametros.Names, parametros.LastNames, parametros.Alias, parametros.Password, parametros.EMail, parametros.PhoneNumber, parametros.Country, parametros.Home_address)
	if err != nil {
		if err == views.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (a *API) Get_userid_Byemail(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_userid_Byemail{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := a.view.Get_userid_Byemail(ctx, parametros.EMail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, id)
}

func (a *API) Get_password_Byemail(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_password_Byemail{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password, err := a.view.Get_password_Byemail(ctx, parametros.EMail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, password)
}

func (a *API) Read_userByid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Read_userByid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := a.view.Read_userByid(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (a *API) Update_userByid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Update_userByid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Update_userByid(ctx, parametros.Names, parametros.LastNames, parametros.Alias, parametros.Password, parametros.EMail, parametros.PhoneNumber, parametros.Country, parametros.Home_address, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (a *API) Delete_userByid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Delete_userByid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Delete_userByid(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (a *API) Get_requeststatus_Byid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_requeststatus_Byid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request, err := a.view.Get_request_status_Byid(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	estado, err := a.view.Get_status_byid(ctx, request.RequestStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, estado)
}

func (a *API) Update_request_status_Byid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Update_requeststatus_Byid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Update_request_status_Byid(ctx, parametros.Id, parametros.RequestStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (a *API) Delete_requests_ByUserid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Delete_requests_ByUserid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Delete_requests_ByUserid(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (a *API) Create_requesttype(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Create_requesttype{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Create_requesttype(ctx, parametros.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (a *API) Create_request(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Create_request{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Llamada a la función para crear la solicitud y obtener el ID generado
	id, err := a.view.Create_request(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Devolver el ID generado
	return c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (a *API) Update_request_URL(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Update_cotizacion{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Update_request_URL(ctx, parametros.RequestStatus, parametros.IAM_URL, parametros.PDF_URL, parametros.QUOTE_PDF_URL, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func checkPassword(providedPassword, storedPassword string) bool {
	hash := sha256.New()
	hash.Write([]byte(providedPassword))
	hashedText := hash.Sum(nil)

	hashedTextHex := hex.EncodeToString(hashedText)

	return hashedTextHex == storedPassword
}

func (a *API) RevisarPassword(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.RevisarPassword{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "err")
	}

	u, _ := a.view.Get_userid_Byemail(ctx, parametros.EMail)
	print(u)
	if u == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	storedPassword, error_Pass_almacenado := a.view.Get_password_Byemail(ctx, parametros.EMail)

	if error_Pass_almacenado != nil {
		fmt.Println("Error fetching password:", error_Pass_almacenado.Error())
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error fetching password"})
	}

	isValid := checkPassword(parametros.Password, storedPassword.Password)
	if !isValid {
		return c.JSON(http.StatusUnauthorized, "Incorrect Password")
	}

	return c.JSON(http.StatusOK, "Password is correct")
}

func (a *API) Get_cotizacion_data(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_cotizacion{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cotizacion, err := a.view.Get_cotizacion_data(ctx, parametros.Id)
	if err != nil {
		fmt.Println("Error fetching cotizacion data:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch cotizacion data"})
	}
	return c.JSON(http.StatusOK, cotizacion)
}

func (a *API) Get_requestId_byUserid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_requestId_byUserid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	requestids, err := a.view.Get_requestId_byUserid(ctx, parametros.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Manejar el caso de un array vacío
	if len(requestids) == 0 {
		return c.JSON(http.StatusOK, []dtos.Get_requestId_byUserid{}) // Devuelve un array vacío en lugar de null
	}

	return c.JSON(http.StatusOK, requestids) // Enviar la lista completa de cotizaciones
}
