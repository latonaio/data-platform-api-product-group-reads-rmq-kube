package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-group-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-group-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
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
	var productGroup *[]dpfm_api_output_formatter.ProductGroup
	var productGroupText *[]dpfm_api_output_formatter.ProductGroupText
	for _, fn := range accepter {
		switch fn {
		case "ProductGroup":
			func() {
				productGroup = c.ProductGroup(mtx, input, output, errs, log)
			}()
		case "ProductGroups":
			func() {
				productGroup = c.ProductGroups(mtx, input, output, errs, log)
			}()
		case "ProductGroupText":
			func() {
				productGroupText = c.ProductGroupText(mtx, input, output, errs, log)
			}()
		case "ProductGroupTexts":
			func() {
				productGroupText = c.ProductGroupTexts(mtx, input, output, errs, log)
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
) *[]dpfm_api_output_formatter.ProductGroup {
	where := fmt.Sprintf("WHERE ProductGroup = '%s'", input.ProductGroup.ProductGroup)

	if input.ProductGroup.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.ProductGroup.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_group_product_group_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, ProductGroup DESC;`,
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
func (c *DPFMAPICaller) ProductGroups(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductGroup {

	if input.ProductGroup.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.ProductGroup.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_group_product_group_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, ProductGroup DESC;`,
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
) *[]dpfm_api_output_formatter.ProductGroupText {
	productGroup := input.ProductGroup[0].ProductGroup
	language := input.ProductGroup[0].ProductGroupText.Language

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

func (c *DPFMAPICaller) ProductGroupTexts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductGroupText {
	productGroup := input.ProductGroup
	where := "WHERE 1 = 1"

	for _, v := range productGroup {
		where = fmt.Sprintf("%s OR (ProductGroup, Language) = ('%s', '%s') ", where, v.ProductGroup, v.ProductGroupText.Language)
	}

	rows, err := c.db.Query(
		`SELECT ProductGroup, Language, ProductGroupName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_group_product_group_text_data
		` + where + `;`,
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
