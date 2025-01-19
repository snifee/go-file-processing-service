package entity

type Order struct {
}

func (*Order) TableName() string {
	return "order_detail"
}
