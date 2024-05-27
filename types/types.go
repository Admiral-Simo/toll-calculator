package types

type Distance struct {
	OBUID int     `json:"obuID"`
	Value float64 `json:"value"`
	Unix  int64   `json:"unix"`
}

type OBUData struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}
