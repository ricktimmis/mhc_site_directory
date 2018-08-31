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
CREATE TABLE IF NOT EXISTS "campsites" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"description" text,
"attractions" text,
"email" TEXT NOT NULL,
"telephone" TEXT NOT NULL,
"website" TEXT,
"facilities" char(36) NOT NULL,
"owninguser" char(36) NOT NULL,
"images" char(36) NOT NULL,
"longitude" float,
"latitude" float,
"addressstreet" TEXT,
"locality" TEXT,
"town" TEXT,
"county" TEXT,
"postcode" TEXT,
"listingexpires" DATETIME NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "facilities" (
"id" TEXT PRIMARY KEY,
"name" TEXT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
