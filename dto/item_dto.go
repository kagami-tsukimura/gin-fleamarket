package dto

type CreateItemInput struct {
	// validator: ``でvalidation check
	Name        string `json:"name" binding:"requires,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}
