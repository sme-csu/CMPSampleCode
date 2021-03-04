package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-30/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	tenantID       = "[tenantID]"
	subscriptionID = "[subscriptionID]"
	clientID       = "[clientID]"
	clientSecret   = "[clientSecret]"
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

	vms, err := listAllVM(ctx)
	if err != nil {
		log.Fatalf("Failed to list vm: %v", err)
	}

	for i := 0; i < len(vms); i++ {
		vm := vms[i]
		fmt.Printf("Index=%v ID=%s\n", i, *vm.ID)
	}

}

func listAllVM(ctx context.Context) (vms []compute.VirtualMachine, err error) {
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)
	vmClient.Authorizer = authorizer

	for page, err := vmClient.ListAll(ctx, ""); page.NotDone(); err = page.NextWithContext(ctx) {
		if err != nil {
			panic(err)
		}
		vms = append(vms, page.Values()...)
	}
	return vms, err
}
