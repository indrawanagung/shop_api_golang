package domain

type UserReview struct {
	ID              string `gorm:"primaryKey"`
	UserID          string
	OrderProductID  string
	RatingProductID string
	Comment         string
	CreatedAt       string
	User            User          `gorm:"foreignKey:user_id;references:id"`
	ShopOrderItem   ShopOrderItem `gorm:"foreignKey:order_product_id;references:id"`
	RatingProduct   RatingProduct `gorm:"foreignKey:rating_product_id;references:id"`
}

func (u *UserReview) TableName() string {
	return "user_reviews"
}
