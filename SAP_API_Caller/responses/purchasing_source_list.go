package responses

type List struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Material                   string `json:"Material"`
			Plant                      string `json:"Plant"`
			SourceListRecord           string `json:"SourceListRecord"`
			ValidityStartDate          string `json:"ValidityStartDate"`
			ValidityEndDate            string `json:"ValidityEndDate"`
			Supplier                   string `json:"Supplier"`
			PurchasingOrganization     string `json:"PurchasingOrganization"`
			SupplyingPlant             string `json:"SupplyingPlant"`
			OrderQuantityUnit          string `json:"OrderQuantityUnit"`
			PurchaseOutlineAgreement   string `json:"PurchaseOutlineAgreement"`
			SupplierIsFixed            bool   `json:"SupplierIsFixed"`
			SourceOfSupplyIsBlocked    bool   `json:"SourceOfSupplyIsBlocked"`
			MRPSourcingControl         string `json:"MRPSourcingControl"`
			LastChangeDateTime         string `json:"LastChangeDateTime"`
			IssgPlantIsFixed           bool   `json:"IssgPlantIsFixed"`
			PurOutlineAgreementIsFixed bool   `json:"PurOutlineAgreementIsFixed"`
			SourceOfSupplyIsFixed      bool   `json:"SourceOfSupplyIsFixed"`
		} `json:"results"`
	} `json:"d"`
}
