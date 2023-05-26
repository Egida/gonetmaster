package user

type User struct {
	Id  string `json:"id" validate:"uuid,required"`
	Key string `json:"key" validate:"required"`
}
