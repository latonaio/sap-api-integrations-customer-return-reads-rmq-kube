package responses

type ToItemScheduleLine struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			CustomerReturn                string `json:"CustomerReturn"`
			CustomerReturnItem            string `json:"CustomerReturnItem"`
			ScheduleLine                  string `json:"ScheduleLine"`
			OrderQuantityUnit             string `json:"OrderQuantityUnit"`
			ConfdOrderQtyByMatlAvailCheck string `json:"ConfdOrderQtyByMatlAvailCheck"`
			DeliveredQtyInOrderQtyUnit    string `json:"DeliveredQtyInOrderQtyUnit"`
			OpenConfdDelivQtyInOrdQtyUnit string `json:"OpenConfdDelivQtyInOrdQtyUnit"`
			DelivBlockReasonForSchedLine  string `json:"DelivBlockReasonForSchedLine"`
		} `json:"results"`
	} `json:"d"`
}  
