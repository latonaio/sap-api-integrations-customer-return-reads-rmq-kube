package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			CustomerReturn                string `json:"CustomerReturn"`
			CustomerReturnType            string `json:"CustomerReturnType"`
			SalesOrganization             string `json:"SalesOrganization"`
			DistributionChannel           string `json:"DistributionChannel"`
			OrganizationDivision          string `json:"OrganizationDivision"`
			SalesGroup                    string `json:"SalesGroup"`
			SalesOffice                   string `json:"SalesOffice"`
			SalesDistrict                 string `json:"SalesDistrict"`
			SoldToParty                   string `json:"SoldToParty"`
			CreationDate                  string `json:"CreationDate"`
			LastChangeDate                string `json:"LastChangeDate"`
			SenderBusinessSystemName      string `json:"SenderBusinessSystemName"`
			LastChangeDateTime            string `json:"LastChangeDateTime"`
			PurchaseOrderByCustomer       string `json:"PurchaseOrderByCustomer"`
			CustomerPurchaseOrderType     string `json:"CustomerPurchaseOrderType"`
			CustomerPurchaseOrderDate     string `json:"CustomerPurchaseOrderDate"`
			CustomerReturnDate            string `json:"CustomerReturnDate"`
			TotalNetAmount                string `json:"TotalNetAmount"`
			TransactionCurrency           string `json:"TransactionCurrency"`
			SDDocumentReason              string `json:"SDDocumentReason"`
			PricingDate                   string `json:"PricingDate"`
			RequestedDeliveryDate         string `json:"RequestedDeliveryDate"`
			ShippingType                  string `json:"ShippingType"`
			HeaderBillingBlockReason      string `json:"HeaderBillingBlockReason"`
			DeliveryBlockReason           string `json:"DeliveryBlockReason"`
			IncotermsClassification       string `json:"IncotermsClassification"`
			CustomerPaymentTerms          string `json:"CustomerPaymentTerms"`
			PaymentMethod                 string `json:"PaymentMethod"`
			CustomerTaxClassification1    string `json:"CustomerTaxClassification1"`
			RetsMgmtProcess               string `json:"RetsMgmtProcess"`
			ReferenceSDDocument           string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentCategory   string `json:"ReferenceSDDocumentCategory"`
			CustomerReturnApprovalReason  string `json:"CustomerReturnApprovalReason"`
			SalesDocApprovalStatus        string `json:"SalesDocApprovalStatus"`
			RetsMgmtLogProcgStatus        string `json:"RetsMgmtLogProcgStatus"`
			RetsMgmtCompnProcgStatus      string `json:"RetsMgmtCompnProcgStatus"`
			RetsMgmtProcessingStatus      string `json:"RetsMgmtProcessingStatus"`
			OverallSDProcessStatus        string `json:"OverallSDProcessStatus"`
			TotalCreditCheckStatus        string `json:"TotalCreditCheckStatus"`
			OverallSDDocumentRejectionSts string `json:"OverallSDDocumentRejectionSts"`
			ToHeaderPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItem                        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
