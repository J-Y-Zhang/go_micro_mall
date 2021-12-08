package model

type ProductImage struct {
	Id        int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ImageName string `json:"image_name"`
	ImageCode string `gorm:"unique_index;not_null" json:"image_code"`
	ImageUrl  string `json:"image_url"`
	//连接Image和Product
	ImageProductID int64 `json:"image_product_id"`
}
