package main

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

type BuyEvent struct {
	MoveEvent struct {
		PackageID         string `json:"packageId"`
		TransactionModule string `json:"transactionModule"`
		Sender            string `json:"sender"`
		Type              string `json:"type"`
		Fields            struct {
			Ask    string `json:"ask"`
			ListID string `json:"list_id"`
			Owner  string `json:"owner"`
			Buyer  string `json:"buyer"`
		} `json:"fields"`
		Bcs string `json:"bcs"`
	} `json:"moveEvent"`
}

type OfferToNftEvent struct {
	MoveEvent struct {
		PackageID         string `json:"packageId"`
		TransactionModule string `json:"transactionModule"`
		Sender            string `json:"sender"`
		Type              string `json:"type"`
		Fields            struct {
			ExpireTime string `json:"expire_time"`
			OfferID    string `json:"offer_id"`
			Owner      string `json:"owner"`
			ListID     string `json:"list_id"`
		} `json:"fields"`
		Bcs string `json:"bcs"`
	} `json:"moveEvent"`
}

type CancelOfferEvent struct {
	MoveEvent struct {
		PackageID         string `json:"packageId"`
		TransactionModule string `json:"transactionModule"`
		Sender            string `json:"sender"`
		Type              string `json:"type"`
		Fields            struct {
			OfferID string `json:"offer_id"`
			ListID  string `json:"list_id"`
			Owner   string `json:"owner"`
		} `json:"fields"`
		Bcs string `json:"bcs"`
	} `json:"moveEvent"`
}

type AcceptOfferEvent struct {
	MoveEvent struct {
		PackageID         string `json:"packageId"`
		TransactionModule string `json:"transactionModule"`
		Sender            string `json:"sender"`
		Type              string `json:"type"`
		Fields            struct {
			OfferID string `json:"offer_id"`
			ListID  string `json:"list_id"`
			Owner   string `json:"owner"`
			Buyer   string `json:"buyer"`
		} `json:"fields"`
		Bcs string `json:"bcs"`
	} `json:"moveEvent"`
}
