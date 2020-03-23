package app

type Document struct {
	Availablecars []struct {
		Charges struct {
			PAYLATER struct {
				ChargeType        string `json:"charge_type"`
				TotalPricePayment struct {
					Amount string `json:"amount"`
					Symbol string  `json:"symbol"`
				} `json:"total_price_payment"`
			} `json:"PAYLATER"`
		} `json:"charges"`
		Code                   string `json:"code"`
		MakeModelOrSimilarText string `json:"make_model_or_similar_text"`
	} `json:"availablecars"`
}
