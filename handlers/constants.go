package handlers

type Message struct {
	ID   int64  `json:"message_id"`
	From User   `json:"from"`
	Date int64  `json:"date"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
}
