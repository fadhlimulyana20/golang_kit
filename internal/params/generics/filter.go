package generics

type GenericFilter struct {
	Q         string `json:"q" schema:"q"`
	StartDate string `json:"start_date" schema:"start_date"`
	EndDate   string `json:"end_date" schema:"end_date"`
	Page      int    `json:"page" schema:"page"`
	Limit     int    `json:"limit" schema:"limit"`
}
