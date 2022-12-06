package domain

type Region struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type RegionAllResp struct {
	Regions []Region `json:"regions"`
	Count   int      `json:"count"`
}

type RegionCreate struct {
	Title string `json:"title"`
}

type District struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	RegionID int32  `json:"region_id"`
}

type DistrictAllResp struct {
	Districts []District `json:"districts"`
	Count     int        `json:"count"`
}

type DistrictCreate struct {
	Title    string `json:"title"`
	RegionID int32  `json:"region_id"`
}

type Quarter struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	DistrictID int32  `json:"district_id"`
}

type QuarterAllResp struct {
	Quarters []Quarter `json:"quarters"`
	Count    int       `json:"count"`
}

type QuarterCreate struct {
	Title      string `json:"title"`
	DistrictID int32  `json:"district_id"`
}
