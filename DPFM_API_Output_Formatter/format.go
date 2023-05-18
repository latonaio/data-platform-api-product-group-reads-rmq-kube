package dpfm_api_output_formatter

import (
	"data-platform-api-product-group-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-product-group-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToProductGroup(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ProductGroup, error) {
	defer rows.Close()
	productGroup := make([]ProductGroup, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductGroup{}

		err := rows.Scan(
			&pm.ProductGroup,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
		data := pm
		productGroup = append(productGroup, ProductGroup{
			ProductGroup: data.ProductGroup,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
	}

	return &productGroup, nil
}

func ConvertToProductGroupText(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ProductGroupText, error) {
	defer rows.Close()
	productGroupText := make([]ProductGroupText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductGroupText{}
		err := rows.Scan(
			&pm.ProductGroup,
			&pm.Language,
			&pm.ProductGroupName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
		data := pm
		productGroupText = append(productGroupText, ProductGroupText{
			ProductGroup:     data.ProductGroup,
			Language:         data.Language,
			ProductGroupName: data.ProductGroupName,
		})

	}
	if i == 0 {
		return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
	}

	return &productGroupText, nil
}
