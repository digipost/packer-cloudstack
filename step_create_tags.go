package cloudstack

import (
	"fmt"
	"github.com/xanzy/go-cloudstack/cloudstack"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepCreateTags struct{}

func (s *stepCreateTags) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*cloudstack.CloudStackClient)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("config").(config)
	template := state.Get("template_id").(string)

	if len(c.TemplateTags) > 0 {
		ui.Say(fmt.Sprintf("Adding tags to template (%s)...", template))

		tags := make(map[string]string)
		for key, value := range c.TemplateTags {
			ui.Message(fmt.Sprintf("Adding tag: \"%s\": \"%s\"", key, value))
			tags[key] = value
		}

		resourcetagService := cloudstack.NewResourcetagsService(client)
		createTagParams := resourcetagService.NewCreateTagsParams([]string{template}, "Template", tags)

		_, err := resourcetagService.CreateTags(createTagParams)

		if err != nil {
			err := fmt.Errorf("Error creating tags: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	return multistep.ActionContinue
}

func (s *stepCreateTags) Cleanup(state multistep.StateBag) {
	// No cleanup...
}
