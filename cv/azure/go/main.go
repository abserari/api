package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
)

// Declare global so don't have to pass it to all of the tasks.
var computerVisionContext context.Context

func main() {
	/*
	 * Configure the Computer Vision client
	 * Set environment variables for COMPUTER_VISION_SUBSCRIPTION_KEY and COMPUTER_VISION_ENDPOINT,
	 * then restart your command shell or your IDE for changes to take effect.
	 */
	computerVisionKey := "eb90cedc20e64323bd61d1b518a33704"

	if computerVisionKey == "" {
		log.Fatal("\n\nPlease set a COMPUTER_VISION_SUBSCRIPTION_KEY environment variable.\n" +
			"**You may need to restart your shell or IDE after it's set.**\n")
	}

	endpointURL := "https://free-abser.cognitiveservices.azure.com/"
	if endpointURL == "" {
		log.Fatal("\n\nPlease set a COMPUTER_VISION_ENDPOINT environment variable.\n" +
			"**You may need to restart your shell or IDE after it's set.**")
	}

	computerVisionClient := computervision.New(endpointURL)
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)

	computerVisionContext = context.Background()
	/*
	 * END - Configure the Computer Vision client
	 */
	landmarkImageURL := "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/landmark.jpg"
	DescribeRemoteImage(computerVisionClient, landmarkImageURL)
}

func DescribeRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	fmt.Println("-----------------------------------------")
	fmt.Println("DESCRIBE IMAGE - remote")
	fmt.Println()
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	maxNumberDescriptionCandidates := new(int32)
	*maxNumberDescriptionCandidates = 1

	remoteImageDescription, err := client.DescribeImage(
		computerVisionContext,
		remoteImage,
		maxNumberDescriptionCandidates,
		"") // language
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Captions from remote image: ")
	if len(*remoteImageDescription.Captions) == 0 {
		fmt.Println("No captions detected.")
	} else {
		for _, caption := range *remoteImageDescription.Captions {
			fmt.Printf("'%v' with confidence %.2f%%\n", *caption.Text, *caption.Confidence*100)
		}
	}
	fmt.Println()
}
