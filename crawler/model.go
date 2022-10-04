package main

// CooperItems describes the collected items from the Cooper Hewitt API.
type CooperItem struct {
	Objects []struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Date        string `json:"date"`
		Description string `json:"description"`
		Type        string `json:"type,omitempty"`
		Medium      string `json:"medium,omitempty"`
		URL         string `json:"url"`
		Country     string `json:"country"`
	} `json:"objects"`
}

type VAItem struct {
	Records []struct {
		SystemNumber    string `json:"systemNumber,omitempty"`
		AccessionNumber string `json:"accessionNumber,omitempty"`
		ObjectType      string `json:"objectType,omitempty"`
		PrimaryTitle    string `json:"_primaryTitle,omitempty"`
		PrimaryMaker    struct {
			Name        string `json:"name,omitempty"`
			Association string `json:"association,omitempty"`
		} `json:"_primaryMaker,omitempty"`
		PrimaryImageID string `json:"_primaryImageId,omitempty"`
		PrimaryDate    string `json:"_primaryDate,omitempty"`
		PrimaryPlace   string `json:"_primaryPlace,omitempty"`
		Images         struct {
			PrimaryThumbnail    string      `json:"_primary_thumbnail,omitempty"`
			IiifImageBaseURL    string      `json:"_iiif_image_base_url,omitempty"`
			IiifPresentationURL interface{} `json:"_iiif_presentation_url,omitempty"`
			ImageResolution     string      `json:"imageResolution,omitempty"`
		} `json:"_images,omitempty"`
		Clusters struct {
			Category struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"category,omitempty"`
			Person struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"person,omitempty"`
			Organisation struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"organisation,omitempty"`
			Collection struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"collection,omitempty"`
			Gallery struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"gallery,omitempty"`
			Style struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"style,omitempty"`
			Place struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"place,omitempty"`
			ObjectType struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"object_type,omitempty"`
			Technique struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"technique,omitempty"`
			Material struct {
				OtherTermsRecordCount int `json:"other_terms_record_count,omitempty"`
				Terms                 []struct {
					ID            string `json:"id,omitempty"`
					Value         string `json:"value,omitempty"`
					Count         int    `json:"count,omitempty"`
					CountMaxError int    `json:"count_max_error,omitempty"`
				} `json:"terms,omitempty"`
			} `json:"material,omitempty"`
		} `json:"clusters,omitempty"`
	}
}
