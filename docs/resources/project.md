---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zeabur_project Resource - zeabur"
subcategory: ""
description: |-
  
---

# zeabur_project (Resource)



## Example Usage

```terraform
resource "zeabur_project" "test" {
  name   = "test_project"
  region = "hkg1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `region` (String)

### Optional

- `description` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)
