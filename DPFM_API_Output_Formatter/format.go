package dpfm_api_output_formatter

import (
	"data-platform-api-product-group-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-product-group-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToProductGroup(sdc *api_input_reader.SDC, rows *sql.Rows) (*ProductGroup, error) {
	pm := &requests.ProductGroup{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ProductGroup,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	productGroup := &ProductGroup{
		ProductGroup: data.ProductGroup,
	}
	return productGroup, nil
}

func ConvertToProductGroupText(sdc *api_input_reader.SDC, rows *sql.Rows) (*ProductGroupText, error) {
	pm := &requests.ProductGroupText{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ProductGroup,
			&pm.Language,
			&pm.ProductGroupName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	productGroupText := &ProductGroupText{
		ProductGroup:     data.ProductGroup,
		Language:         data.Language,
		ProductGroupName: data.ProductGroupName,
	}
	return productGroupText, nil
}
