package cloudstack

import (
	"fmt"
	"github.com/xanzy/go-cloudstack/cloudstack"
	"log"
	"net/url"
)


type Artifact struct {
	// The name of the template
	templateName string

	// The ID of the image
	templateId string

	// The URL of the cloudstack API endpoint
	providerUrl string

	// The client for making API calls
	client *cloudstack.CloudStackClient
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Files() []string {
	// No local files created with Cloudstack.
	return nil
}

func (a *Artifact) Id() string {
	values := url.Values{}
	values.Set("templateid", a.templateId)
	return a.providerUrl + "?" + values.Encode()
}

func (a *Artifact) String() string {
	return fmt.Sprintf("A template was created: UUID: %v - Name: %v",
		a.templateId, a.templateName)
}

func (a *Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	log.Printf("Delete template: %s", a.templateId)
	templateService := cloudstack.NewTemplateService(a.client)
	_, err := templateService.DeleteTemplate(templateService.NewDeleteTemplateParams(a.templateId))
	return err
}
