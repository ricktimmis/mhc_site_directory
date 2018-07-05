CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"username" TEXT NOT NULL,
"firstname" TEXT NOT NULL,
"lastname" TEXT NOT NULL,
"email" TEXT NOT NULL,
"location" TEXT NOT NULL,
"membertype" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "birthday" DATETIME, "logintype" TEXT, "loginvendorkey" TEXT, "membershipexpires" DATETIME);
