resource "static_data" "creation_metadata" {
  data = {
    creation_date = timestamp()
  }

  // optionally you can set the lifecycle.ignore_changes attribute to
  // prevent unnecessary resource changes
  lifecycle {
    ignore_changes = [data]
  }
}

output "creation_date" {
  value = static_data.creation_metadata.output.creation_date
}
