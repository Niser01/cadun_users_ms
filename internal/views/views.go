package views

import (
	"context"

	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/models"
	"github.com/jmoiron/sqlx"
)

// views interface has all the crud methods for the user and saved elements
type View_interface interface {
	Get_userid_Byemail(ctx context.Context, eMail string) (*models.UserId, error)
	Create_user(ctx context.Context, names string, lastNames string, alias string, password string, eMail string, phoneNumber string, country string, home_address string) error
	Read_userByid(ctx context.Context, id int) (*models.UserProfile, error)
	Get_password_Byemail(ctx context.Context, eMail string) (*models.Password, error)
	Get_status_byid(ctx context.Context, id int) (*models.Status, error)
	Update_userByid(ctx context.Context, names string, lastNames string, alias string, password string, eMail string, phoneNumber string, country string, home_address string, id int) error
	Delete_userByid(ctx context.Context, id int) error
	Create_request(ctx context.Context, idUser int) error
	Delete_requests_ByUserid(ctx context.Context, idUser int) error
	Update_request_status_Byid(ctx context.Context, id int, request_status int) error
	Create_requesttype(ctx context.Context, Status string) error
	Update_request_URL(ctx context.Context, request_status int, IAM_URL string, PDF_URL string, QUOTE_PDF_URL string, id int) error
	Get_request_status_Byid(ctx context.Context, id int) (*models.Request_Status, error)
	Get_cotizacion_data(ctx context.Context, id_request int) (*models.Get_cotizacion_data, error)
}

// view is the implementation of the views interface
type View_struct struct {
	db *sqlx.DB
}

// New returns the implementation of the views interface
func New(db *sqlx.DB) View_interface {
	return &View_struct{db: db}
}
