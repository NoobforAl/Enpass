package entity

type Password struct {
	ID        uint   `json:"passid"`
	UserID    uint   `json:"userid"`
	ServiceID uint   `json:"serviceid"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Note      string `json:"note"`
}
