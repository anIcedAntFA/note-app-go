package notemodel

type Filter struct {
	Category string `json:"category" form:"category"`
	Status   string `json:"status" form:"status"`
}
