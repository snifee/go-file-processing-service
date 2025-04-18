package entity

type Product struct {
	ProductID       int     `gorm:"column:product_id" json:"product_id"`
	ProductName     string  `gorm:"column:product_name" json:"product_name"`
	SupplierID      int     `gorm:"column:supplier_id" json:"supplier_id"`
	CategoryID      int     `gorm:"column:category_id" json:"category_id"`
	QuantityPerUnit string  `gorm:"column:quantity_per_unit" json:"quantity_per_unit"`
	UnitPrice       float64 `gorm:"column:unit_price" json:"unit_price"`
	UnitsInStock    int     `gorm:"column:units_in_stock" json:"units_in_stock"`
	UnitsOnOrder    int     `gorm:"column:units_on_order" json:"units_on_order"`
	ReorderLevel    int     `gorm:"column:reorder_level" json:"reorder_level"`
	Discontinued    bool    `gorm:"column:discontinued" json:"discontinued"`
}

func (*Product) TableName() string {
	return "products"
}
