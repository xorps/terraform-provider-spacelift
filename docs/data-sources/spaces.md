---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "spacelift_spaces Data Source - terraform-provider-spacelift"
subcategory: ""
description: |-
  spacelift_spaces can find all spaces in the spacelift organization.
---

# spacelift_spaces (Data Source)

`spacelift_spaces` can find all spaces in the spacelift organization.

## Example Usage

```terraform
data "spacelift_spaces" "this" {
}

output "spaces" {
  value = data.spacelift_spaces.this.spaces
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `spaces` (List of Object) (see [below for nested schema](#nestedatt--spaces))

<a id="nestedatt--spaces"></a>
### Nested Schema for `spaces`

Read-Only:

- `description` (String)
- `inherit_entities` (Boolean)
- `labels` (Set of String)
- `name` (String)
- `parent_space_id` (String)
- `space_id` (String)