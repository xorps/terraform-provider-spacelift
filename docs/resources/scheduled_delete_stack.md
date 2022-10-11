---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "spacelift_scheduled_delete_stack Resource - terraform-provider-spacelift"
subcategory: ""
description: |-
  spacelift_scheduled_delete_stack represents a scheduling configuration for a Stack. It will trigger a stack deletion task at the given timestamp.
---

# spacelift_scheduled_delete_stack (Resource)

`spacelift_scheduled_delete_stack` represents a scheduling configuration for a Stack. It will trigger a stack deletion task at the given timestamp.

## Example Usage

```terraform
resource "spacelift_stack" "k8s-core" {
  // ...
}

// at a given timestamp (unix)
resource "spacelift_scheduled_delete_stack" "k8s-core-delete" {
  stack_id = spacelift_stack.k8s-core.id

  at               = "1663336895"
  delete_resources = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `at` (Number) Timestamp (unix timestamp) at which time the scheduling should happen.
- `stack_id` (String) ID of the stack for which to set up scheduling

### Optional

- `delete_resources` (Boolean) Indicates whether the resources of the stack should be deleted.
- `schedule_id` (String) ID of the schedule

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import spacelift_scheduled_delete_stack.ireland-kubeconfig $STACK_ID/$SCHEDULED_DELETE_STACK_ID
```