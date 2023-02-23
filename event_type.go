package main

type ListEvent struct {
	MoveEvent struct {
		PackageID         string `json:"packageId"`
		TransactionModule string `json:"transactionModule"`
		Sender            string `json:"sender"`
		Type              string `json:"type"`
		Fields            struct {
			Ask    string `json:"ask"`
			ListID string `json:"list_id"`
			Owner  string `json:"owner"`
		} `json:"fields"`
		Bcs string `json:"bcs"`
	} `json:"moveEvent"`
}

type MarketCreate struct {
	MoveEvent struct {
		PackageID         string `json:"packageId"`
		TransactionModule string `json:"transactionModule"`
		Sender            string `json:"sender"`
		Type              string `json:"type"`
		Fields            struct {
			ID      string `json:"id"`
			OfferID string `json:"offer_id"`
		} `json:"fields"`
		Bcs string `json:"bcs"`
	} `json:"moveEvent"`
}
