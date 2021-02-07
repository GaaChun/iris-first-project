package Models

type User struct {
	ID       uint64 `json:"id"`
	UserName string `json:"user_name"`
	Age      uint   `json:"age"`
}
