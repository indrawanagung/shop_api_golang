package domain

type RatingProduct struct {
	ID     string `gorm:"primaryKey"`
	Rating int
}

func (r *RatingProduct) TableName() string {
	return "rating_products"
}
