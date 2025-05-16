package bot

type UserState struct {
	Step       int
	SellerName string
	BuyerName  string
	VIN        string
	BrandModel string
	Year       string
	Color      string
	Price      string
	Date       string
	City       string
}

var userStates = make(map[int64]*UserState)
