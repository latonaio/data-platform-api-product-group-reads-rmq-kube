package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-group-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-group-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var productGroup *dpfm_api_output_formatter.ProductGroup
	var productGroupText *dpfm_api_output_formatter.ProductGroupText
	for _, fn := range accepter {
		switch fn {
		case "ProductGroup":
			func() {
				productGroup = c.ProductGroup(mtx, input, output, errs, log)
			}()
		case "ProductGroupText":
			func() {
				productGroupText = c.ProductGroupText(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProductGroup:     productGroup,
		ProductGroupText: productGroupText,
	}

	return data
}

func (c *DPFMAPICaller) ProductGroup(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductGroup {
	productGroup := input.ProductGroup.ProductGroup

	rows, err := c.db.Query(
		`SELECT ProductGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_group_product_group_data
		WHERE ProductGroup = ?;`, productGroup,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductGroup(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductGroupText(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductGroupText {
	productGroup := input.ProductGroup.ProductGroup
	language := input.ProductGroup.ProductGroupText.Language

	rows, err := c.db.Query(
		`SELECT ProductGroup, Language, ProductGroupName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_group_product_group_text_data
		WHERE (ProductGroup, Language) = (?, ?);`, productGroup, language,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductGroupText(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
