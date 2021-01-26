package request


type AddTODO struct {
	Title string `json:"title" binding:"required"`
	Status bool `json:"status"`
}
