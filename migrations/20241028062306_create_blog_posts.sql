-- Create "tasks" table
CREATE TABLE "public"."tasks" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "title" character varying(255) NOT NULL, "description" text NULL, "user_id" integer NOT NULL, "status" character varying(20) NOT NULL, PRIMARY KEY ("id"));
-- Create index "idx_tasks_deleted_at" to table: "tasks"
CREATE INDEX "idx_tasks_deleted_at" ON "public"."tasks" ("deleted_at");
-- Create "users" table
CREATE TABLE "public"."users" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "username" text NOT NULL, "email" text NOT NULL, "password" text NOT NULL, "role" text NOT NULL DEFAULT 'user', "contact_no" text NULL, PRIMARY KEY ("id"), CONSTRAINT "uni_users_email" UNIQUE ("email"), CONSTRAINT "uni_users_username" UNIQUE ("username"));
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "idx_users_email" to table: "users"
CREATE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create index "idx_users_username" to table: "users"
CREATE INDEX "idx_users_username" ON "public"."users" ("username");
