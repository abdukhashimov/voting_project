package domain

type Initiative struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Author      int32    `json:"author"`
	BoardID     int32    `json:"board_id"`
	RegionID    int32    `json:"region_id"`
	DistrictID  int32    `json:"district_id"`
}

type InitiativeAllResp struct {
	Initiatives []Initiative `json:"initiatives"`
	Count       int          `json:"count"`
}

type InitiativeCreate struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Author      int32    `json:"author"`
	BoardID     int32    `json:"board_id"`
	RegionID    int32    `json:"region_id"`
	DistrictID  int32    `json:"district_id"`
}
