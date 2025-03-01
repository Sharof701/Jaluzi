package models

type Admin struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"delete_at,omitempty"`
}

type AdminCreate struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdminUpdate struct {
	Id       string `json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdminPrimaryKey struct {
	Id string `json:"id"`
}

type AdminGetListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type AdminGetListResponse struct {
	Admin []*Admin `json:"admin"`
	Total int64    `json:"total"`
}
