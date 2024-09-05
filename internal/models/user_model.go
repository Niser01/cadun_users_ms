package models

// UserProfile db model
type UserProfile struct {
	ID           int    `db:"id"`
	Names        string `db:"names"`
	LastNames    string `db:"lastnames"`
	Alias        string `db:"alias"`
	Password     string `db:"password"`
	EMail        string `db:"email"`
	PhoneNumber  string `db:"phonenumber"`
	Country      string `db:"country"`
	Home_address string `db:"home_address"`
}

type Password struct {
	Password string `db:"password"`
}

type UserId struct {
	ID int `db:"id"`
}

// RequestType db model
type RequestType struct {
	ID     int    `db:"id"`
	Status string `db:"status"`
}

// Request db model
type Request struct {
	ID            int `db:"id"`
	IDUser        int `db:"idUser"`
	RequestStatus int `db:"request_status"`
}

type Request_Status struct {
	RequestStatus int `db:"request_status"`
}

type Status struct {
	Status string `db:"Status"`
}

type Get_cotizacion_data struct {
	Names         string `db:"names"`
	LastNames     string `db:"lastnames"`
	EMail         string `db:"email"`
	PhoneNumber   string `db:"phonenumber"`
	Home_address  string `db:"home_address"`
	IAM_URL       string `db:"iam_url"`
	PDF_URL       string `db:"pdf_url"`
	QUOTE_PDF_URL string `db:"quote_pdf_url"`
}

type Get_requestId_byUserid struct {
	ID int `db:"id"`
}
