package responses

type Item struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			CustomerReturn                 string `json:"CustomerReturn"`
			CustomerReturnItem             string `json:"CustomerReturnItem"`
			HigherLevelItem                string `json:"HigherLevelItem"`
			CustomerReturnItemCategory     string `json:"CustomerReturnItemCategory"`
			CustomerReturnItemText         string `json:"CustomerReturnItemText"`
			PurchaseOrderByCustomer        string `json:"PurchaseOrderByCustomer"`
			Material                       string `json:"Material"`
			MaterialByCustomer             string `json:"MaterialByCustomer"`
			RequestedQuantity              string `json:"RequestedQuantity"`
			RequestedQuantityUnit          string `json:"RequestedQuantityUnit"`
			ItemGrossWeight                string `json:"ItemGrossWeight"`
			ItemNetWeight                  string `json:"ItemNetWeight"`
			ItemWeightUnit                 string `json:"ItemWeightUnit"`
			ItemVolume                     string `json:"ItemVolume"`
			ItemVolumeUnit                 string `json:"ItemVolumeUnit"`
			TransactionCurrency            string `json:"TransactionCurrency"`
			NetAmount                      string `json:"NetAmount"`
			MaterialGroup                  string `json:"MaterialGroup"`
			Batch                          string `json:"Batch"`
			ProductionPlant                string `json:"ProductionPlant"`
			StorageLocation                string `json:"StorageLocation"`
			ShippingPoint                  string `json:"ShippingPoint"`
			ShippingType                   string `json:"ShippingType"`
			DeliveryPriority               string `json:"DeliveryPriority"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			CustomerPaymentTerms           string `json:"CustomerPaymentTerms"`
			ProductTaxClassification1      string `json:"ProductTaxClassification1"`
			SalesDocumentRjcnReason        string `json:"SalesDocumentRjcnReason"`
			ItemBillingBlockReason         string `json:"ItemBillingBlockReason"`
			ProfitCenter                   string `json:"ProfitCenter"`
			RetsMgmtProcess                string `json:"RetsMgmtProcess"`
			RetsMgmtProcessItem            string `json:"RetsMgmtProcessItem"`
			ReturnReason                   string `json:"ReturnReason"`
			RetsMgmtProcessingBlock        string `json:"RetsMgmtProcessingBlock"`
			CustRetItmFollowUpActivity     string `json:"CustRetItmFollowUpActivity"`
			ReturnsMaterialHasBeenReceived string `json:"ReturnsMaterialHasBeenReceived"`
			ReturnsRefundType              string `json:"ReturnsRefundType"`
			ReturnsRefundProcgMode         string `json:"ReturnsRefundProcgMode"`
			ReturnsRefundExtent            string `json:"ReturnsRefundExtent"`
			PrelimRefundIsDetermined       string `json:"PrelimRefundIsDetermined"`
			ReturnsRefundRjcnReason        string `json:"ReturnsRefundRjcnReason"`
			ReplacementMaterial            string `json:"ReplacementMaterial"`
			ReplacementMaterialQuantity    string `json:"ReplacementMaterialQuantity"`
			ReplacementMaterialQtyUnit     string `json:"ReplacementMaterialQtyUnit"`
			ReplacementMaterialIsRequested string `json:"ReplacementMaterialIsRequested"`
			ReplacementMatlSupplyingPlant  string `json:"ReplacementMatlSupplyingPlant"`
			NextPlantForFollowUpActivity   string `json:"NextPlantForFollowUpActivity"`
			ReturnsTransshipmentPlant      string `json:"ReturnsTransshipmentPlant"`
			Supplier                       string `json:"Supplier"`
			SupplierRetMatlAuthzn          string `json:"SupplierRetMatlAuthzn"`
			SuplrRetMatlAuthznIsRequired   string `json:"SuplrRetMatlAuthznIsRequired"`
			CustomerRetMatlInspResultCode  string `json:"CustomerRetMatlInspResultCode"`
			NextFllwUpActivityForMatlInsp  string `json:"NextFllwUpActivityForMatlInsp"`
			RetMatlInspResultCode          string `json:"RetMatlInspResultCode"`
			ProductIsInspectedAtCustSite   string `json:"ProductIsInspectedAtCustSite"`
			CustRetMatlAuthzn              string `json:"CustRetMatlAuthzn"`
			CRMLogicalSystem               string `json:"CRMLogicalSystem"`
			CRMObjectUUID                  string `json:"CRMObjectUUID"`
			CRMObjectID                    string `json:"CRMObjectID"`
			CRMObjectType                  string `json:"CRMObjectType"`
			RetsMgmtItmLogProcgStatus      string `json:"RetsMgmtItmLogProcgStatus"`
			RetsMgmtItmCompnProcgStatus    string `json:"RetsMgmtItmCompnProcgStatus"`
			RetsMgmtItmProcgStatus         string `json:"RetsMgmtItmProcgStatus"`
			ReturnsDocumentStatus          string `json:"ReturnsDocumentStatus"`
			ReturnsDocumentApprovalStatus  string `json:"ReturnsDocumentApprovalStatus"`
			SDProcessStatus                string `json:"SDProcessStatus"`
			ReferenceSDDocument            string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem        string `json:"ReferenceSDDocumentItem"`
			ReferenceSDDocumentCategory    string `json:"ReferenceSDDocumentCategory"`
			SDDocumentRejectionStatus      string `json:"SDDocumentRejectionStatus"`
			ToItemPricingElement           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PricingElement"`
			ToItemProcessStep struct {
				Deferred struct {
				URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProcessStep"`
			ToItemScheduleLine struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ScheduleLine"`
		} `json:"results"`
	} `json:"d"`
}
