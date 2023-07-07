package model

type Currencies struct {
	BaseCode       string  `json:"base_code"`
	TargetCode     string  `json:"target_code"`
	ConversionRate float64 `json:"conversion_rate"`
	CreatedAt      int64   `json:"create_at"`
}
