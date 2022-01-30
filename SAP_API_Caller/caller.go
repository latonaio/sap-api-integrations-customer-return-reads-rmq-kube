package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-customer-return-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetCustomerReturn(customerReturn, customerReturnItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(customerReturn)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(customerReturn, customerReturnItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(customerReturn string) {
	headerData, err := c.callCustomerReturnSrvAPIRequirementHeader("A_CustomerReturn", customerReturn)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": headerData, "function": "CustomerReturnHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	headerPartnerData, err := c.callToHeaderPartner(headerData[0].ToHeaderPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": headerPartnerData, "function": "CustomerReturnHeaderPartner"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerPartnerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemData, "function": "CustomerReturnItem"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemPricingElementData, "function": "CustomerReturnItemPricingElement"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)

	itemProcessStepData, err := c.callToItemProcessStep(itemData[0].ToItemProcessStep)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemProcessStepData, "function": "CustomerReturnItemProcessStep"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemProcessStepData)

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemScheduleLineData, "function": "CustomerReturnItemScheduleLine"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemScheduleLineData)

}

func (c *SAPAPICaller) callCustomerReturnSrvAPIRequirementHeader(api, customerReturn string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_CUSTOMER_RETURN_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, customerReturn)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToHeaderPartner(url string) ([]sap_api_output_formatter.ToHeaderPartner, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPricingElement(url string) ([]sap_api_output_formatter.ToItemPricingElement, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemProcessStep(url string) ([]sap_api_output_formatter.ToItemProcessStep, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemProcessStep(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemScheduleLine(url string) ([]sap_api_output_formatter.ToItemScheduleLine, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(customerReturn, customerReturnItem string) {
	itemData, err := c.callCustomerReturnSrvAPIRequirementItem("A_CustomerReturnItem", customerReturn, customerReturnItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemData, "function": "CustomerReturnItem"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemPricingElementData, "function": "CustomerReturnItemPricingElement"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)

	itemProcessStepData, err := c.callToItemProcessStep(itemData[0].ToItemProcessStep)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemProcessStepData, "function": "CustomerReturnItemProcessStep"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemProcessStepData)

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemScheduleLineData, "function": "CustomerReturnItemScheduleLine"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemScheduleLineData)
}

func (c *SAPAPICaller) callCustomerReturnSrvAPIRequirementItem(api, customerReturn, customerReturnItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_CUSTOMER_RETURN_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, customerReturn, customerReturnItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, customerReturn string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CustomerReturn eq '%s'", customerReturn))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, customerReturn, customerReturnItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CustomerReturn eq '%s' and CustomerReturnItem eq '%s'", customerReturn, customerReturnItem))
	req.URL.RawQuery = params.Encode()
}
