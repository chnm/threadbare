package main

// CooperItems describes the collected items from the Cooper Hewitt API.
type CooperItem struct {
	Objects []struct {
		ID                  string      `json:"id"`
		TmsID               string      `json:"tms:id"`
		AccessionNumber     string      `json:"accession_number"`
		Title               string      `json:"title"`
		TitleRaw            string      `json:"title_raw"`
		URL                 string      `json:"url"`
		HasNoKnownCopyright interface{} `json:"has_no_known_copyright"`
		DepartmentID        string      `json:"department_id"`
		PeriodID            interface{} `json:"period_id"`
		MediaID             string      `json:"media_id"`
		TypeID              string      `json:"type_id"`
		Date                string      `json:"date"`
		YearStart           interface{} `json:"year_start"`
		YearEnd             interface{} `json:"year_end"`
		YearAcquired        string      `json:"year_acquired"`
		Decade              interface{} `json:"decade"`
		WoeCountryID        string      `json:"woe:country_id"`
		Medium              string      `json:"medium"`
		Markings            interface{} `json:"markings"`
		Signed              interface{} `json:"signed"`
		Inscribed           interface{} `json:"inscribed"`
		Provenance          string      `json:"provenance"`
		Dimensions          string      `json:"dimensions"`
		DimensionsRaw       struct {
			Warp []string `json:"warp"`
			Weft []string `json:"weft"`
		} `json:"dimensions_raw"`
		Creditline    string      `json:"creditline"`
		Description   string      `json:"description"`
		Justification interface{} `json:"justification"`
		GalleryText   interface{} `json:"gallery_text"`
		LabelText     interface{} `json:"label_text"`
		Videos        interface{} `json:"videos"`
		OnDisplay     interface{} `json:"on_display"`
		Country       string      `json:"woe:country"`
		Type          string      `json:"type"`
		Images        []struct {
			X struct {
				URL       string `json:"url"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsPrimary string `json:"is_primary"`
				ImageID   string `json:"image_id"`
			} `json:"x,omitempty"`
			B struct {
				URL       string `json:"url"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsPrimary string `json:"is_primary"`
				ImageID   string `json:"image_id"`
			} `json:"b"`
			Z struct {
				URL       string `json:"url"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsPrimary string `json:"is_primary"`
				ImageID   string `json:"image_id"`
			} `json:"z"`
			N struct {
				URL       string `json:"url"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsPrimary string `json:"is_primary"`
				ImageID   string `json:"image_id"`
			} `json:"n"`
			D struct {
				URL       string `json:"url"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsPrimary string `json:"is_primary"`
				ImageID   string `json:"image_id"`
			} `json:"d"`
			Sq struct {
				URL       string `json:"url"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsPrimary string `json:"is_primary"`
				ImageID   string `json:"image_id"`
			} `json:"sq"`
			O struct {
				URL       string `json:"url"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsPrimary string `json:"is_primary"`
				ImageID   string `json:"image_id"`
			} `json:"o,omitempty"`
		} `json:"images"`
	} `json:"objects"`
	Participants         []interface{} `json:"participants"`
	WoeCountryName       string        `json:"woe:country_name"`
	IsLoanObject         int           `json:"is_loan_object"`
	Total                int           `json:"total"`
	Page                 int           `json:"page"`
	PerPage              int           `json:"per_page"`
	Pages                int           `json:"pages"`
	Stat                 string        `json:"stat"`
	EventPublishingState string        `json:"event_publishing_state"`
}

// type ItemList struct {
// 	Results []string `json:"results"`
// }
