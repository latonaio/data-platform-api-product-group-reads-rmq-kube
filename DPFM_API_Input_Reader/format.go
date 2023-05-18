package dpfm_api_input_reader

import (
	"data-platform-api-product-group-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToProductGroup() []requests.ProductGroup {
	data := sdc.ProductGroup[0]
	return []requests.ProductGroup{
		{ProductGroup: data.ProductGroup},
	}
}

func (sdc *SDC) ConvertToProductGroupText() []requests.ProductGroupText {
	dataProductGroup := sdc.ProductGroup[0]
	data := sdc.ProductGroup[0].ProductGroupText
	return []requests.ProductGroupText{
		{ProductGroup: dataProductGroup.ProductGroup,
			Language:         data.Language,
			ProductGroupName: data.ProductGroupName},
	}
}
