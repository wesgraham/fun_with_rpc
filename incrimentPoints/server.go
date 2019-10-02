package main


//type Item struct {
//	ID int
//	Price int  `json:"ticker"`
//}

type Computer struct {}

func (computer *Computer) AddFive(item *Item, reply *int) error {
	*reply = item.Price + 5
	return nil
}

func (computer *Computer) SubtractFive(item *Item, reply *int) error {
	*reply = item.Price - 5
	return nil
}