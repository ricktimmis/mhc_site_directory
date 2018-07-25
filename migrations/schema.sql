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
"birthday" DATETIME,
"location" TEXT NOT NULL,
"membertype" TEXT NOT NULL,
"loginprovider" TEXT NOT NULL,
"loginprovider_id" TEXT NOT NULL,
"membershipexpires" DATETIME,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "images" (
"id" TEXT PRIMARY KEY,
"related_id" TEXT NOT NULL,
"original" TEXT NOT NULL,
"resized" TEXT NOT NULL,
"thumbnail" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
