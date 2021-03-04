package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-30/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

const (
	tenantID       = "[tenantID]"
	subscriptionID = "[subscriptionID]"
	clientID       = "[clientID]"
	clientSecret   = "[clientSecret]"

	resourceGroupName     = "TestDisk1"
	resourceGroupLocation = "eastus"
	diskName              = "TestDisk1-Disk1"
)

var (
	ctx        = context.Background()
	authorizer autorest.Authorizer
)

func init() {
	var err error
	authorizer, err = auth.NewClientCredentialsConfig(clientID, clientSecret, tenantID).Authorizer()

	if err != nil {
		log.Fatalf("Failed to get OAuth config: %v", err)
	}
}

func main() {
	//创建资源组
	group, err := CreateGroup()
	if err != nil {
		log.Fatalf("failed to create group: %v", err)
	}
	log.Printf("Created group: %v", *group.Name)

	//创建磁盘
	log.Printf("Starting create disk: %s", diskName)
	disk, err := CreateDisk()
	if err != nil {
		log.Fatalf("Failed to create disk: %v", err)
	}
	if disk.Name != nil {
		log.Printf("Completed create disk: %v", diskName)
	} else {
		log.Printf("Completed create disk: %v (no data returned to SDK)", diskName)
	}
}

func CreateGroup() (group resources.Group, err error) {
	groupsClient := resources.NewGroupsClient(subscriptionID)
	groupsClient.Authorizer = authorizer

	return groupsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resources.Group{
			Location: to.StringPtr(resourceGroupLocation)})
}

func CreateDisk() (disk compute.Disk, err error) {
	disksClient := compute.NewDisksClient(subscriptionID)
	disksClient.Authorizer = authorizer

	future, err := disksClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		diskName,
		compute.Disk{
			Location: to.StringPtr(resourceGroupLocation),
			DiskProperties: &compute.DiskProperties{
				CreationData: &compute.CreationData{
					CreateOption: compute.Empty,
				},
				DiskSizeGB: to.Int32Ptr(64),
			},
		})
	if err != nil {
		return disk, fmt.Errorf("cannot create disk: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, disksClient.Client)
	if err != nil {
		return disk, fmt.Errorf("cannot get the disk create or update future response: %v", err)
	}

	return future.Result(disksClient)
}
