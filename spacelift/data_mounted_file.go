package spacelift

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/spacelift-io/terraform-provider-spacelift/spacelift/structs"
)

func dataMountedFile() *schema.Resource {
	return &schema.Resource{
		Read: dataMountedFileRead,

		Schema: map[string]*schema.Schema{
			"checksum": &schema.Schema{
				Type:        schema.TypeString,
				Description: "SHA-256 checksum of the value",
				Computed:    true,
			},
			"content": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Content of the mounted file encoded using Base-64",
				Sensitive:   true,
				Computed:    true,
			},
			"context_id": &schema.Schema{
				Type:          schema.TypeString,
				Description:   "ID of the context where the mounted file is stored",
				Optional:      true,
				ConflictsWith: []string{"stack_id"},
			},
			"relative_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Relative path to the mounted file",
				Required:    true,
			},
			"stack_id": &schema.Schema{
				Type:          schema.TypeString,
				Description:   "ID of the stack where the mounted file is stored",
				Optional:      true,
				ConflictsWith: []string{"context_id"},
			},
			"write_only": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Indicates whether the value can be read back outside a Run",
				Computed:    true,
			},
		},
	}
}

func dataMountedFileRead(d *schema.ResourceData, meta interface{}) error {
	_, contextOK := d.GetOk("context_id")
	_, stackOK := d.GetOk("stack_id")

	if contextOK == stackOK {
		return errors.New("either context_id or stack_id must be provided")
	}

	if contextOK {
		return dataMountedFileReadContext(d, meta)
	}

	return dataMountedFileReadStack(d, meta)
}

func dataMountedFileReadContext(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		Context *struct {
			ConfigElement *structs.ConfigElement `graphql:"configElement(id: $id)"`
		} `graphql:"context(id: $context)"`
	}

	contextID := d.Get("context_id")
	variableName := d.Get("relative_path")

	variables := map[string]interface{}{
		"context": toID(contextID),
		"id":      toID(variableName),
	}

	if err := meta.(*Client).Query(&query, variables); err != nil {
		return errors.Wrap(err, "could not query for context mounted file")
	}

	if query.Context == nil {
		return errors.New("context not found")
	}

	configElement := query.Context.ConfigElement
	if configElement == nil {
		return errors.New("mounted file not found")
	}

	if configElement.Type != "FILE_MOUNT" {
		return errors.New("config element is not a mounted file")
	}

	d.SetId(fmt.Sprintf("context/%s/%s", contextID, variableName))

	populateMountedFile(d, query.Context.ConfigElement)

	return nil
}

func dataMountedFileReadStack(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		Stack *struct {
			ConfigElement *structs.ConfigElement `graphql:"configElement(id: $id)"`
		} `graphql:"stack(id: $stack)"`
	}

	stackID := d.Get("stack_id")
	variableName := d.Get("relative_path")

	variables := map[string]interface{}{
		"stack": toID(stackID),
		"id":    toID(variableName),
	}

	if err := meta.(*Client).Query(&query, variables); err != nil {
		return errors.Wrap(err, "could not query for stack mounted file")
	}

	if query.Stack == nil {
		return errors.New("stack not found")
	}

	if query.Stack.ConfigElement == nil {
		return errors.New("mounted file not found")
	}

	d.SetId(fmt.Sprintf("stack/%s/%s", stackID, variableName))

	populateMountedFile(d, query.Stack.ConfigElement)

	return nil
}

func populateMountedFile(d *schema.ResourceData, el *structs.ConfigElement) {
	d.Set("checksum", el.Checksum)
	d.Set("write_only", el.WriteOnly)

	if el.Value != nil {
		d.Set("content", *el.Value)
	} else {
		d.Set("content", nil)
	}
}