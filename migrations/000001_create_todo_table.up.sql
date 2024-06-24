-- Enable uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Create enum type "status"
CREATE TYPE "public"."status" AS ENUM ('IN_PROGRESS', 'COMPLETED');
-- Create "tbl_todos" table
CREATE TABLE "public"."tbl_todos" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "status" character varying(12) NOT NULL,
  "title" character varying(100) NOT NULL,
  "description" text NOT NULL,
  "image" bytea NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);
-- Create index "idx_title" to table: "tbl_todos"
CREATE INDEX "idx_title" ON "public"."tbl_todos" ("title");
-- Create index "idx_status" to table: "tbl_todos"
CREATE INDEX "idx_status" ON "public"."tbl_todos" ("status");
-- Create index "idx_created_at" to table: "tbl_todos"
CREATE INDEX "idx_created_at" ON "public"."tbl_todos" ("created_at");