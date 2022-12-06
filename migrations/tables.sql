CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "status" (
    "id" serial PRIMARY KEY NOT NULL,
    "title" VARCHAR NOT NULL
);

CREATE TABLE "region" (
    "id" serial PRIMARY KEY NOT NULL,
    "title" VARCHAR NOT NULL
);

CREATE TABLE "district" (
    "id" serial PRIMARY KEY NOT NULL,
    "title" VARCHAR NOT NULL,
    "region_id" INTEGER NOT NULL
);

CREATE TABLE "quarter" (
    "id" serial PRIMARY KEY NOT NULL,
    "title" VARCHAR NOT NULL,
    "district_id" INTEGER NOT NULL
);

CREATE TABLE "district_budget" (
    "board_id" uuid NOT NULL,
    "district_id" INTEGER NOT NULL,
    "budget" INTEGER NOT NULL,
    "updated_by" uuid NOT NULL,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "users" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4 ()),
    "fullname" VARCHAR NOT NULL,
    "phone_number" VARCHAR UNIQUE NOT NULL,
    "role" VARCHAR DEFAULT 'client',
    "username" VARCHAR DEFAULT '',
    "password" VARCHAR DEFAULT '',
    "status" INTEGER DEFAULT 1,
    "region_id" INTEGER NOT NULL,
    "district_id" INTEGER NOT NULL,
    "quarter_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "initiatives" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4 ()),
    "title" VARCHAR NOT NULL,
    "images" VARCHAR [ ],
    "description" VARCHAR,
    "author" uuid NOT NULL,
    "board_id" uuid NOT NULL,
    "vote_count" INTEGER,
    "status" INTEGER NOT NULL,
    "requested_amount" INTEGER,
    "granted_amount" INTEGER,
    "region_id" INTEGER NOT NULL,
    "district_id" INTEGER NOT NULL,
    "quarter_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "sub_initiatives" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4 ()),
    "title" VARCHAR NOT NULL,
    "images" VARCHAR [ ],
    "vote_count" INTEGER,
    "initiative_id" uuid NOT NULL,
    "requested_amount" INTEGER,
    "granted_amount" INTEGER,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "chronology" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4 ()),
    "title" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "initiative_chronology" (
    "initiative_id" uuid NOT NULL,
    "chronology_id" uuid NOT NULL,
    "updated_by" uuid NOT NULL,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "board" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4 ()),
    "title" VARCHAR,
    "icon" VARCHAR,
    "total_amount" INTEGER,
    "accept_start_date" TIMESTAMP NOT NULL,
    "accept_end_date" TIMESTAMP NOT NULL,
    "review_start_date" TIMESTAMP NOT NULL,
    "review_end_date" TIMESTAMP NOT NULL,
    "voting_start_date" TIMESTAMP NOT NULL,
    "voting_end_date" TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "vote" (
    "phone_number" VARCHAR NOT NULL,
    "initiative_id" uuid NOT NULL,
    "board_id" uuid NOT NULL,
    "created_at" TIMESTAMP DEFAULT (now()),
    "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE UNIQUE INDEX ON "district_budget" ("board_id", "district_id");

CREATE UNIQUE INDEX ON "vote" ("phone_number", "board_id");

COMMENT ON COLUMN "users"."id" IS 'uuid.new is used';

COMMENT ON COLUMN "initiatives"."id" IS 'uuid.new is used';

COMMENT ON COLUMN "sub_initiatives"."id" IS 'uuid.new is used';

COMMENT ON COLUMN "chronology"."id" IS 'uuid.new is used';

COMMENT ON COLUMN "board"."id" IS 'uuid.new is used';

ALTER TABLE "district"
ADD FOREIGN KEY ("region_id") REFERENCES "region" ("id");

ALTER TABLE "quarter"
ADD FOREIGN KEY ("district_id") REFERENCES "district" ("id");

ALTER TABLE "district_budget"
ADD FOREIGN KEY ("board_id") REFERENCES "board" ("id");

ALTER TABLE "district_budget"
ADD FOREIGN KEY ("district_id") REFERENCES "district" ("id");

ALTER TABLE "district_budget"
ADD FOREIGN KEY ("updated_by") REFERENCES "users" ("id");

ALTER TABLE "users"
ADD FOREIGN KEY ("status") REFERENCES "status" ("id");

ALTER TABLE "users"
ADD FOREIGN KEY ("region_id") REFERENCES "region" ("id");

ALTER TABLE "users"
ADD FOREIGN KEY ("district_id") REFERENCES "district" ("id");

ALTER TABLE "users"
ADD FOREIGN KEY ("quarter_id") REFERENCES "quarter" ("id");

ALTER TABLE "initiatives"
ADD FOREIGN KEY ("author") REFERENCES "users" ("id");

ALTER TABLE "initiatives"
ADD FOREIGN KEY ("board_id") REFERENCES "board" ("id");

ALTER TABLE "initiatives"
ADD FOREIGN KEY ("status") REFERENCES "status" ("id");

ALTER TABLE "initiatives"
ADD FOREIGN KEY ("region_id") REFERENCES "region" ("id");

ALTER TABLE "initiatives"
ADD FOREIGN KEY ("district_id") REFERENCES "district" ("id");

ALTER TABLE "initiatives"
ADD FOREIGN KEY ("quarter_id") REFERENCES "quarter" ("id");

ALTER TABLE "sub_initiatives"
ADD FOREIGN KEY ("initiative_id") REFERENCES "initiatives" ("id");

ALTER TABLE "initiative_chronology"
ADD FOREIGN KEY ("initiative_id") REFERENCES "initiatives" ("id");

ALTER TABLE "initiative_chronology"
ADD FOREIGN KEY ("chronology_id") REFERENCES "chronology" ("id");

ALTER TABLE "initiative_chronology"
ADD FOREIGN KEY ("updated_by") REFERENCES "users" ("id");

ALTER TABLE "vote"
ADD FOREIGN KEY ("initiative_id") REFERENCES "initiatives" ("id");

ALTER TABLE "vote"
ADD FOREIGN KEY ("board_id") REFERENCES "board" ("id");