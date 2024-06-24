schema "public" {
  comment = "standard public schema"
}

enum "status" {
  schema = schema.public
  values = ["IN_PROGRESS", "COMPLETED"]
}

table "tbl_todos" {
  schema = schema.public
  column "id" {
    type    = uuid
    null    = false
    default = sql("gen_random_uuid()")
  }

  column "status" {
    type = varchar(10)
    null = false
  }

  column "title" {
    type = varchar(100)
    null = false
  }

  column "description" {
    type = text
    null = false
  }

  column "image" {
    type = bytea
    null = false
  }

  column "created_at" {
    type    = timestamptz
    null    = false
    default = sql("now()")
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_title" {
    columns = [column.title]
  }

  index "idx_status" {
    columns = [column.status]
  }

  index "idx_created_at" {
    columns = [column.created_at]
  }
}