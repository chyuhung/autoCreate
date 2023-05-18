package openstack

// import (
// 	"fmt"

// 	"github.com/gophercloud/gophercloud"
// 	"github.com/gophercloud/gophercloud/openstack"
// 	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
// 	"github.com/gophercloud/gophercloud/pagination"
// )

// // list image
// func (c *client) GetImageNames() (imageNames []string) {
// 	method := "GetImageNames"
// 	client, err := openstack.NewImageServiceV2(c.Provider, gophercloud.EndpointOpts{
// 		Region: "RegionOne",
// 	})
// 	//fmt.Println(client)
// 	if err != nil {
// 		fmt.Printf("%s : %v", method, err)
// 		return
// 	}
// 	pager := images.List(client, images.ListOpts{})
// 	//fmt.Println(pager)
// 	err = pager.EachPage(func(page pagination.Page) (bool, error) {
// 		extract, _ := images.ExtractImages(page)
// 		for _, i := range extract {
// 			imageNames = append(imageNames, i.Name)
// 			//fmt.Println(i.Name)
// 		}
// 		return true, nil
// 	})
// 	return
// }
