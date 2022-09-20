package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchasing-source-list-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToList(raw []byte, l *logger.Logger) ([]List, error) {
	pm := &responses.List{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to List. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}

	list := make([]List, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		list = append(list, List{
			Material:                   data.Material,
			Plant:                      data.Plant,
			SourceListRecord:           data.SourceListRecord,
			ValidityStartDate:          data.ValidityStartDate,
			ValidityEndDate:            data.ValidityEndDate,
			Supplier:                   data.Supplier,
			PurchasingOrganization:     data.PurchasingOrganization,
			SupplyingPlant:             data.SupplyingPlant,
			OrderQuantityUnit:          data.OrderQuantityUnit,
			PurchaseOutlineAgreement:   data.PurchaseOutlineAgreement,
			SupplierIsFixed:            data.SupplierIsFixed,
			SourceOfSupplyIsBlocked:    data.SourceOfSupplyIsBlocked,
			MRPSourcingControl:         data.MRPSourcingControl,
			LastChangeDateTime:         data.LastChangeDateTime,
			IssgPlantIsFixed:           data.IssgPlantIsFixed,
			PurOutlineAgreementIsFixed: data.PurOutlineAgreementIsFixed,
			SourceOfSupplyIsFixed:      data.SourceOfSupplyIsFixed,
		})
	}

	return list, nil
}
