package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-purchasing-source-list-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchasingSourceList(material, plant, sourceListRecord, supplier, supplyingPlant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "List":
			func() {
				c.List(material, plant, sourceListRecord)
				wg.Done()
			}()
		case "Supplier":
			func() {
				c.Supplier(material, plant, sourceListRecord, supplier)
				wg.Done()
			}()
		case "SupplyingPlant":
			func() {
				c.SupplyingPlant(material, plant, sourceListRecord, supplyingPlant)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) List(material, plant, sourceListRecord string) {
	data, err := c.callPurchasingSourceListSrvAPIRequirementList("A_PurchasingSource", material, plant, sourceListRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementList(api, material, plant, sourceListRecord string) ([]sap_api_output_formatter.List, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	param := c.getQueryWithList(map[string]string{}, material, plant, sourceListRecord)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToList(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Supplier(material, plant, sourceListRecord, supplier string) {
	data, err := c.callPurchasingSourceListSrvAPIRequirementSupplier("A_PurchasingSource", material, plant, sourceListRecord, supplier)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementSupplier(api, material, plant, sourceListRecord, supplier string) ([]sap_api_output_formatter.List, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	param := c.getQueryWithSupplier(map[string]string{}, material, plant, sourceListRecord, supplier)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToList(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SupplyingPlant(material, plant, sourceListRecord, supplyingPlant string) {
	data, err := c.callPurchasingSourceListSrvAPIRequirementSupplyingPlant("A_PurchasingSource", material, plant, sourceListRecord, supplyingPlant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementSupplyingPlant(api, material, plant, sourceListRecord, supplyingPlant string) ([]sap_api_output_formatter.List, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	param := c.getQueryWithSupplyingPlant(map[string]string{}, material, plant, sourceListRecord, supplyingPlant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToList(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithList(params map[string]string, material, plant, sourceListRecord string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and Plant eq '%s' and SourceListRecord eq '%s'", material, plant, sourceListRecord)
	return params
}

func (c *SAPAPICaller) getQueryWithSupplier(params map[string]string, material, plant, sourceListRecord, supplier string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and Plant eq '%s' and SourceListRecord eq '%s' and Supplier eq '%s'", material, plant, sourceListRecord, supplier)
	return params
}

func (c *SAPAPICaller) getQueryWithSupplyingPlant(params map[string]string, material, plant, sourceListRecord, supplyingPlant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and Plant eq '%s' and SourceListRecord eq '%s' and SupplyingPlant eq '%s'", material, plant, sourceListRecord, supplyingPlant)
	return params
}
