package domain

type GetAllParams struct {
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Search     string `json:"search"`
	RegionID   int    `json:"region_id"`
	DistrictID int    `json:"district_id"`
	Status     int    `json:"status"`
}
