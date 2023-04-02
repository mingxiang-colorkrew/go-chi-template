CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "name" text,
  "email" text NOT NULL,
  "role" int NOT NULL,
  "tenant_id" bigint NOT NULL
);

CREATE TABLE "tenant" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "short_code" text NOT NULL
);

CREATE UNIQUE INDEX user_email_tenant_id_unique ON "user" ("email", "tenant_id");

CREATE INDEX user_role ON "user" ("role");

CREATE UNIQUE INDEX user_tenant_short_code_unique ON "tenant" ("short_code");

ALTER TABLE "user" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");

