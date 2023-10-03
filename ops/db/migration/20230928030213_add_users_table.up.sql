CREATE TABLE "users" (
	"id" bigserial PRIMARY KEY,
	"email" varchar NOT NULL,
	"name" varchar NOT NULL,
	"password" varchar NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
	"id" bigserial PRIMARY KEY,
	"user_id" bigint NOT NULL,
	"balance" bigint NOT NULL,
	"bank" varchar NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
	"id" bigserial PRIMARY KEY,
	"account_id" bigint NOT NULL,
	"amount" bigint NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
	"id" bigserial PRIMARY KEY,
	"from_account_id" bigint NOT NULL,
	"to_account_id" bigint NOT NULL,
	"amount" bigint NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "entries" ("account_id");
CREATE INDEX ON "transfers" ("from_account_id");
CREATE INDEX ON "transfers" ("to_account_id");
CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");
CREATE INDEX ON "users" ("email");