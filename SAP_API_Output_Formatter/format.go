package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-customer-return-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			CustomerReturn:                data.CustomerReturn,
			CustomerReturnType:            data.CustomerReturnType,
			SalesOrganization:             data.SalesOrganization,
			DistributionChannel:           data.DistributionChannel,
			OrganizationDivision:          data.OrganizationDivision,
			SalesGroup:                    data.SalesGroup,
			SalesOffice:                   data.SalesOffice,
			SalesDistrict:                 data.SalesDistrict,
			SoldToParty:                   data.SoldToParty,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			SenderBusinessSystemName:      data.SenderBusinessSystemName,
			LastChangeDateTime:            data.LastChangeDateTime,
			PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
			CustomerPurchaseOrderType:     data.CustomerPurchaseOrderType,
			CustomerPurchaseOrderDate:     data.CustomerPurchaseOrderDate,
			CustomerReturnDate:            data.CustomerReturnDate,
			TotalNetAmount:                data.TotalNetAmount,
			TransactionCurrency:           data.TransactionCurrency,
			SDDocumentReason:              data.SDDocumentReason,
			PricingDate:                   data.PricingDate,
			RequestedDeliveryDate:         data.RequestedDeliveryDate,
			ShippingType:                  data.ShippingType,
			HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
			DeliveryBlockReason:           data.DeliveryBlockReason,
			IncotermsClassification:       data.IncotermsClassification,
			CustomerPaymentTerms:          data.CustomerPaymentTerms,
			PaymentMethod:                 data.PaymentMethod,
			CustomerTaxClassification1:    data.CustomerTaxClassification1,
			RetsMgmtProcess:               data.RetsMgmtProcess,
			ReferenceSDDocument:           data.ReferenceSDDocument,
			ReferenceSDDocumentCategory:   data.ReferenceSDDocumentCategory,
			CustomerReturnApprovalReason:  data.CustomerReturnApprovalReason,
			SalesDocApprovalStatus:        data.SalesDocApprovalStatus,
			RetsMgmtLogProcgStatus:        data.RetsMgmtLogProcgStatus,
			RetsMgmtCompnProcgStatus:      data.RetsMgmtCompnProcgStatus,
			RetsMgmtProcessingStatus:      data.RetsMgmtProcessingStatus,
			OverallSDProcessStatus:        data.OverallSDProcessStatus,
			TotalCreditCheckStatus:        data.TotalCreditCheckStatus,
			OverallSDDocumentRejectionSts: data.OverallSDDocumentRejectionSts,
			ToHeaderPartner:               data.ToHeaderPartner.Deferred.URI,
			ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			CustomerReturn:                 data.CustomerReturn,
			CustomerReturnItem:             data.CustomerReturnItem,
			HigherLevelItem:                data.HigherLevelItem,
			CustomerReturnItemCategory:     data.CustomerReturnItemCategory,
			CustomerReturnItemText:         data.CustomerReturnItemText,
			PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			RequestedQuantity:              data.RequestedQuantity,
			RequestedQuantityUnit:          data.RequestedQuantityUnit,
			ItemGrossWeight:                data.ItemGrossWeight,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemWeightUnit:                 data.ItemWeightUnit,
			ItemVolume:                     data.ItemVolume,
			ItemVolumeUnit:                 data.ItemVolumeUnit,
			TransactionCurrency:            data.TransactionCurrency,
			NetAmount:                      data.NetAmount,
			MaterialGroup:                  data.MaterialGroup,
			Batch:                          data.Batch,
			ProductionPlant:                data.ProductionPlant,
			StorageLocation:                data.StorageLocation,
			ShippingPoint:                  data.ShippingPoint,
			ShippingType:                   data.ShippingType,
			DeliveryPriority:               data.DeliveryPriority,
			IncotermsClassification:        data.IncotermsClassification,
			CustomerPaymentTerms:           data.CustomerPaymentTerms,
			ProductTaxClassification1:      data.ProductTaxClassification1,
			SalesDocumentRjcnReason:        data.SalesDocumentRjcnReason,
			ItemBillingBlockReason:         data.ItemBillingBlockReason,
			ProfitCenter:                   data.ProfitCenter,
			RetsMgmtProcess:                data.RetsMgmtProcess,
			RetsMgmtProcessItem:            data.RetsMgmtProcessItem,
			ReturnReason:                   data.ReturnReason,
			RetsMgmtProcessingBlock:        data.RetsMgmtProcessingBlock,
			CustRetItmFollowUpActivity:     data.CustRetItmFollowUpActivity,
			ReturnsMaterialHasBeenReceived: data.ReturnsMaterialHasBeenReceived,
			ReturnsRefundType:              data.ReturnsRefundType,
			ReturnsRefundProcgMode:         data.ReturnsRefundProcgMode,
			ReturnsRefundExtent:            data.ReturnsRefundExtent,
			PrelimRefundIsDetermined:       data.PrelimRefundIsDetermined,
			ReturnsRefundRjcnReason:        data.ReturnsRefundRjcnReason,
			ReplacementMaterial:            data.ReplacementMaterial,
			ReplacementMaterialQuantity:    data.ReplacementMaterialQuantity,
			ReplacementMaterialQtyUnit:     data.ReplacementMaterialQtyUnit,
			ReplacementMaterialIsRequested: data.ReplacementMaterialIsRequested,
			ReplacementMatlSupplyingPlant:  data.ReplacementMatlSupplyingPlant,
			NextPlantForFollowUpActivity:   data.NextPlantForFollowUpActivity,
			ReturnsTransshipmentPlant:      data.ReturnsTransshipmentPlant,
			Supplier:                       data.Supplier,
			SupplierRetMatlAuthzn:          data.SupplierRetMatlAuthzn,
			SuplrRetMatlAuthznIsRequired:   data.SuplrRetMatlAuthznIsRequired,
			CustomerRetMatlInspResultCode:  data.CustomerRetMatlInspResultCode,
			NextFllwUpActivityForMatlInsp:  data.NextFllwUpActivityForMatlInsp,
			RetMatlInspResultCode:          data.RetMatlInspResultCode,
			ProductIsInspectedAtCustSite:   data.ProductIsInspectedAtCustSite,
			CustRetMatlAuthzn:              data.CustRetMatlAuthzn,
			CRMLogicalSystem:               data.CRMLogicalSystem,
			CRMObjectUUID:                  data.CRMObjectUUID,
			CRMObjectID:                    data.CRMObjectID,
			CRMObjectType:                  data.CRMObjectType,
			RetsMgmtItmLogProcgStatus:      data.RetsMgmtItmLogProcgStatus,
			RetsMgmtItmCompnProcgStatus:    data.RetsMgmtItmCompnProcgStatus,
			RetsMgmtItmProcgStatus:         data.RetsMgmtItmProcgStatus,
			ReturnsDocumentStatus:          data.ReturnsDocumentStatus,
			ReturnsDocumentApprovalStatus:  data.ReturnsDocumentApprovalStatus,
			SDProcessStatus:                data.SDProcessStatus,
			ReferenceSDDocument:            data.ReferenceSDDocument,
			ReferenceSDDocumentItem:        data.ReferenceSDDocumentItem,
			ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
			SDDocumentRejectionStatus:      data.SDDocumentRejectionStatus,
			ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
			ToItemProcessStep:              data.ToItemProcessStep.Deferred.URI,
			ToItemScheduleLine:             data.ToItemScheduleLine.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
			CustomerReturn:               data.CustomerReturn,
			PartnerFunction:              data.PartnerFunction,
			Customer:                     data.Customer,
			Supplier:                     data.Supplier,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			CustomerReturn:                 data.CustomerReturn,
			CustomerReturnItem:             data.CustomerReturnItem,
			HigherLevelItem:                data.HigherLevelItem,
			CustomerReturnItemCategory:     data.CustomerReturnItemCategory,
			CustomerReturnItemText:         data.CustomerReturnItemText,
			PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			RequestedQuantity:              data.RequestedQuantity,
			RequestedQuantityUnit:          data.RequestedQuantityUnit,
			ItemGrossWeight:                data.ItemGrossWeight,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemWeightUnit:                 data.ItemWeightUnit,
			ItemVolume:                     data.ItemVolume,
			ItemVolumeUnit:                 data.ItemVolumeUnit,
			TransactionCurrency:            data.TransactionCurrency,
			NetAmount:                      data.NetAmount,
			MaterialGroup:                  data.MaterialGroup,
			Batch:                          data.Batch,
			ProductionPlant:                data.ProductionPlant,
			ShippingPoint:                  data.ShippingPoint,
			ShippingType:                   data.ShippingType,
			DeliveryPriority:               data.DeliveryPriority,
			IncotermsClassification:        data.IncotermsClassification,
			CustomerPaymentTerms:           data.CustomerPaymentTerms,
			ProductTaxClassification1:      data.ProductTaxClassification1,
			SalesDocumentRjcnReason:        data.SalesDocumentRjcnReason,
			ItemBillingBlockReason:         data.ItemBillingBlockReason,
			ProfitCenter:                   data.ProfitCenter,
			RetsMgmtProcess:                data.RetsMgmtProcess,
			RetsMgmtProcessItem:            data.RetsMgmtProcessItem,
			ReturnReason:                   data.ReturnReason,
			RetsMgmtProcessingBlock:        data.RetsMgmtProcessingBlock,
			CustRetItmFollowUpActivity:     data.CustRetItmFollowUpActivity,
			ReturnsMaterialHasBeenReceived: data.ReturnsMaterialHasBeenReceived,
			ReturnsRefundType:              data.ReturnsRefundType,
			ReturnsRefundProcgMode:         data.ReturnsRefundProcgMode,
			ReturnsRefundExtent:            data.ReturnsRefundExtent,
			PrelimRefundIsDetermined:       data.PrelimRefundIsDetermined,
			ReturnsRefundRjcnReason:        data.ReturnsRefundRjcnReason,
			ReplacementMaterial:            data.ReplacementMaterial,
			ReplacementMaterialQuantity:    data.ReplacementMaterialQuantity,
			ReplacementMaterialQtyUnit:     data.ReplacementMaterialQtyUnit,
			ReplacementMaterialIsRequested: data.ReplacementMaterialIsRequested,
			ReplacementMatlSupplyingPlant:  data.ReplacementMatlSupplyingPlant,
			NextPlantForFollowUpActivity:   data.NextPlantForFollowUpActivity,
			ReturnsTransshipmentPlant:      data.ReturnsTransshipmentPlant,
			Supplier:                       data.Supplier,
			SupplierRetMatlAuthzn:          data.SupplierRetMatlAuthzn,
			SuplrRetMatlAuthznIsRequired:   data.SuplrRetMatlAuthznIsRequired,
			CustomerRetMatlInspResultCode:  data.CustomerRetMatlInspResultCode,
			NextFllwUpActivityForMatlInsp:  data.NextFllwUpActivityForMatlInsp,
			RetMatlInspResultCode:          data.RetMatlInspResultCode,
			ProductIsInspectedAtCustSite:   data.ProductIsInspectedAtCustSite,
			CustRetMatlAuthzn:              data.CustRetMatlAuthzn,
			CRMLogicalSystem:               data.CRMLogicalSystem,
			CRMObjectUUID:                  data.CRMObjectUUID,
			CRMObjectID:                    data.CRMObjectID,
			CRMObjectType:                  data.CRMObjectType,
			RetsMgmtItmLogProcgStatus:      data.RetsMgmtItmLogProcgStatus,
			RetsMgmtItmCompnProcgStatus:    data.RetsMgmtItmCompnProcgStatus,
			RetsMgmtItmProcgStatus:         data.RetsMgmtItmProcgStatus,
			ReturnsDocumentStatus:          data.ReturnsDocumentStatus,
			ReturnsDocumentApprovalStatus:  data.ReturnsDocumentApprovalStatus,
			SDProcessStatus:                data.SDProcessStatus,
			ReferenceSDDocument:            data.ReferenceSDDocument,
			ReferenceSDDocumentItem:        data.ReferenceSDDocumentItem,
			ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
			SDDocumentRejectionStatus:      data.SDDocumentRejectionStatus,
			ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
	        ToItemProcessStep:              data.ToItemProcessStep.Deferred.URI,
	        ToItemScheduleLine:             data.ToItemScheduleLine.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
			CustomerReturn:                data.CustomerReturn,
			CustomerReturnItem:            data.CustomerReturnItem,
			PricingProcedureStep:          data.PricingProcedureStep,
			PricingProcedureCounter:       data.PricingProcedureCounter,
			ConditionType:                 data.ConditionType,
			PricingDateTime:               data.PricingDateTime,
			ConditionCalculationType:      data.ConditionCalculationType,
			ConditionBaseValue:            data.ConditionBaseValue,
			ConditionRateValue:            data.ConditionRateValue,
			ConditionCurrency:             data.ConditionCurrency,
			ConditionQuantity:             data.ConditionQuantity,
			ConditionQuantityUnit:         data.ConditionQuantityUnit,
			ConditionCategory:             data.ConditionCategory,
			PricingScaleType:              data.PricingScaleType,
			ConditionRecord:               data.ConditionRecord,
			ConditionSequentialNumber:     data.ConditionSequentialNumber,
			TaxCode:                       data.TaxCode,
			ConditionAmount:               data.ConditionAmount,
			TransactionCurrency:           data.TransactionCurrency,
			PricingScaleBasis:             data.PricingScaleBasis,
			ConditionScaleBasisValue:      data.ConditionScaleBasisValue,
			ConditionScaleBasisUnit:       data.ConditionScaleBasisUnit,
			ConditionScaleBasisCurrency:   data.ConditionScaleBasisCurrency,
			ConditionIsManuallyChanged:    data.ConditionIsManuallyChanged,
		})
	}

	return toItemPricingElement, nil
}

func ConvertToToItemProcessStep(raw []byte, l *logger.Logger) ([]ToItemProcessStep, error) {
	pm := &responses.ToItemProcessStep{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProcessStep. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemProcessStep := make([]ToItemProcessStep, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemProcessStep = append(toItemProcessStep, ToItemProcessStep{
			RetsMgmtProcess:               data.RetsMgmtProcess,
			RetsMgmtProcessItem:           data.RetsMgmtProcessItem,
			RetsMgmtProcItmQtySplit:       data.RetsMgmtProcItmQtySplit,
			RetsMgmtProcessStep:           data.RetsMgmtProcessStep,
			ReturnsDocumentType:           data.ReturnsDocumentType,
			ReturnsDocument:               data.ReturnsDocument,
			ReturnsDocumentItem:           data.ReturnsDocumentItem,
			ReturnsDocumentStatus:         data.ReturnsDocumentStatus,
			ReturnsDocumentApprovalStatus: data.ReturnsDocumentApprovalStatus,
			ReturnsReferenceDocumentType:  data.ReturnsReferenceDocumentType,
			ReturnsReferenceDocument:      data.ReturnsReferenceDocument,
			ReturnsReferenceDocumentItem:  data.ReturnsReferenceDocumentItem,
			RetsMgmtProcStepExecStatus:    data.RetsMgmtProcStepExecStatus,
		})
	}

	return toItemProcessStep, nil
}

func ConvertToToItemScheduleLine(raw []byte, l *logger.Logger) ([]ToItemScheduleLine, error) {
	pm := &responses.ToItemScheduleLine{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemScheduleLine. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemScheduleLine := make([]ToItemScheduleLine, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemScheduleLine = append(toItemScheduleLine, ToItemScheduleLine{
			CustomerReturn:                data.CustomerReturn,
			CustomerReturnItem:            data.CustomerReturnItem,
			ScheduleLine:                  data.ScheduleLine,
			OrderQuantityUnit:             data.OrderQuantityUnit,
			ConfdOrderQtyByMatlAvailCheck: data.ConfdOrderQtyByMatlAvailCheck,
			DeliveredQtyInOrderQtyUnit:    data.DeliveredQtyInOrderQtyUnit,
			OpenConfdDelivQtyInOrdQtyUnit: data.OpenConfdDelivQtyInOrdQtyUnit,
			DelivBlockReasonForSchedLine:  data.DelivBlockReasonForSchedLine,
		})
	}

	return toItemScheduleLine, nil
}
