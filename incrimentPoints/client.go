package incPoints

// Clients have a ticker and a base. Role of Server is to receive client, and incriment ticker
type Item struct {
	ID int
	Price int  `json:"ticker"`
}