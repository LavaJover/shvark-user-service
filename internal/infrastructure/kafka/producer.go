package kafka

type UserCreated struct {
	UserID		string `json:"user_id"`
	Username 	string `json:"username"`
}