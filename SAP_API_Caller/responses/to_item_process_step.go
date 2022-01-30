package responses

type ToItemProcessStep struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			RetsMgmtProcess               string `json:"RetsMgmtProcess"`
			RetsMgmtProcessItem           string `json:"RetsMgmtProcessItem"`
			RetsMgmtProcItmQtySplit       string `json:"RetsMgmtProcItmQtySplit"`
			RetsMgmtProcessStep           string `json:"RetsMgmtProcessStep"`
			ReturnsDocumentType           string `json:"ReturnsDocumentType"`
			ReturnsDocument               string `json:"ReturnsDocument"`
			ReturnsDocumentItem           string `json:"ReturnsDocumentItem"`
			ReturnsDocumentStatus         string `json:"ReturnsDocumentStatus"`
			ReturnsDocumentApprovalStatus string `json:"ReturnsDocumentApprovalStatus"`
			ReturnsReferenceDocumentType  string `json:"ReturnsReferenceDocumentType"`
			ReturnsReferenceDocument      string `json:"ReturnsReferenceDocument"`
			ReturnsReferenceDocumentItem  string `json:"ReturnsReferenceDocumentItem"`
			RetsMgmtProcStepExecStatus    string `json:"RetsMgmtProcStepExecStatus"`
		} `json:"results"`
	} `json:"d"`
}
