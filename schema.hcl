table "tasks" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "title" {
    null = false
    type = character_varying(255)
  }
  column "description" {
    null = true
    type = text
  }
  column "user_id" {
    null = false
    type = integer
  }
  column "status" {
    null = false
    type = character_varying(20)
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_tasks_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "username" {
    null = false
    type = text
  }
  column "email" {
    null = false
    type = text
  }
  column "password" {
    null = false
    type = text
  }
  column "role" {
    null    = false
    type    = text
    default = "user"
  }  
  primary_key {
    columns = [column.id]
  }
  index "idx_users_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_users_email" {
    columns = [column.email]
  }
  index "idx_users_username" {
    columns = [column.username]
  }
  unique "uni_users_email" {
    columns = [column.email]
  }
  unique "uni_users_username" {
    columns = [column.username]
  }
}
schema "public" {
  comment = "standard public schema"
}
