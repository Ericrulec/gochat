CREATE TYPE "UserStatusCode" AS ENUM (
  'UNKNOWN_STATUS',
  'OFFLINE',
  'ONLINE'
);

CREATE TYPE "ChatMessageStatus" AS ENUM (
  'UNKNOWN_STATUS',
  'NOT_DELIVERED',
  'DELIVERED'
);

CREATE TABLE "Users" (
  "user_id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "status" UserStatusCode NOT NULL
);

CREATE TABLE "ChatMessage" (
  "chat_id" SERIAL PRIMARY KEY,
  "from_user" int NOT NULL,
  "to_user" int NOT NULL,
  "body" varchar NOT NULL,
  "status" ChatMessageStatus,
  "time_stamp" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Users" ("user_id");

CREATE INDEX ON "ChatMessage" ("from_user");

CREATE INDEX ON "ChatMessage" ("to_user");

CREATE INDEX ON "ChatMessage" ("from_user", "to_user");

ALTER TABLE "ChatMessage" ADD FOREIGN KEY ("from_user") REFERENCES "Users" ("user_id");

ALTER TABLE "ChatMessage" ADD FOREIGN KEY ("to_user") REFERENCES "Users" ("user_id");
