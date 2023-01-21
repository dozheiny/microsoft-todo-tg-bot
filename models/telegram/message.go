package telegram

type Message struct {
	ID   int64  `json:"message_id"`
	From User   `json:"from"`
	Date int64  `json:"date"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}
