-- Drop indices
DROP INDEX IF EXISTS "public"."idx_created_at";
DROP INDEX IF EXISTS "public"."idx_status";
DROP INDEX IF EXISTS "public"."idx_title";

-- Drop table
DROP TABLE IF EXISTS "public"."tbl_todos";

-- Drop enum type
DROP TYPE IF EXISTS "public"."status";
