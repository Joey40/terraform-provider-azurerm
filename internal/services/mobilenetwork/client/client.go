// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/datanetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/mobilenetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplane"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcoredataplane"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/service"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/simgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/simpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/site"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/slice"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	MobileNetworkClient          *mobilenetwork.MobileNetworkClient
	ServiceClient                *service.ServiceClient
	SIMGroupClient               *simgroup.SIMGroupClient
	SliceClient                  *slice.SliceClient
	SiteClient                   *site.SiteClient
	DataNetworkClient            *datanetwork.DataNetworkClient
	SIMPolicyClient              *simpolicy.SIMPolicyClient
	PacketCoreControlPlaneClient *packetcorecontrolplane.PacketCoreControlPlaneClient
	PacketCoreDataPlaneClient    *packetcoredataplane.PacketCoreDataPlaneClient
}

func NewClient(o *common.ClientOptions) *Client {
	mobileNetworkClient := mobilenetwork.NewMobileNetworkClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&mobileNetworkClient.Client, o.ResourceManagerAuthorizer)

	dataNetworkClient := datanetwork.NewDataNetworkClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&dataNetworkClient.Client, o.ResourceManagerAuthorizer)

	serviceClient := service.NewServiceClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&serviceClient.Client, o.ResourceManagerAuthorizer)

	simGroupClient := simgroup.NewSIMGroupClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&simGroupClient.Client, o.ResourceManagerAuthorizer)

	siteClient := site.NewSiteClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&siteClient.Client, o.ResourceManagerAuthorizer)

	sliceClient := slice.NewSliceClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&sliceClient.Client, o.ResourceManagerAuthorizer)

	simPolicyClient := simpolicy.NewSIMPolicyClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&simPolicyClient.Client, o.ResourceManagerAuthorizer)

	packetCoreControlPlaneClient := packetcorecontrolplane.NewPacketCoreControlPlaneClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&packetCoreControlPlaneClient.Client, o.ResourceManagerAuthorizer)

	packetCoreDataPlaneClient := packetcoredataplane.NewPacketCoreDataPlaneClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&packetCoreDataPlaneClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		MobileNetworkClient:          &mobileNetworkClient,
		DataNetworkClient:            &dataNetworkClient,
		ServiceClient:                &serviceClient,
		SIMGroupClient:               &simGroupClient,
		SiteClient:                   &siteClient,
		SliceClient:                  &sliceClient,
		SIMPolicyClient:              &simPolicyClient,
		PacketCoreControlPlaneClient: &packetCoreControlPlaneClient,
		PacketCoreDataPlaneClient:    &packetCoreDataPlaneClient,
	}
}
