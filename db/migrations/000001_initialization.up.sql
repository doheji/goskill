CREATE TABLE "Skill" (
                         "id" bigserial PRIMARY KEY,
                         "learner" varchar NOT NULL,
                         "name" varchar NOT NULL,
                         "score" bigint NOT NULL DEFAULT 0,
                         "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Task" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "skill_id" bigint NOT NULL,
                        "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Session" (
                           "id" bigserial PRIMARY KEY,
                           "task_id" bigint NOT NULL,
                           "name" varchar NOT NULL,
                           "description" varchar,
                           "goal" varchar,
                           "location" varchar,
                           "attended" bool NOT NULL DEFAULT false,
                           "startTime" timestamptz NOT NULL DEFAULT (now()),
                           "duration" int NOT NULL DEFAULT 60,
                           "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Skill" ("learner");

CREATE INDEX ON "Skill" ("name");

CREATE INDEX ON "Task" ("skill_id");

CREATE INDEX ON "Session" ("task_id");

COMMENT ON COLUMN "Session"."description" IS 'This is a description of the session, what you will do';

COMMENT ON COLUMN "Session"."goal" IS 'Ones goals for the session';

COMMENT ON COLUMN "Session"."location" IS 'This is where one want to carry out the session';

COMMENT ON COLUMN "Session"."duration" IS 'This is duration in minutes';

ALTER TABLE "Task" ADD FOREIGN KEY ("skill_id") REFERENCES "Skill" ("id");

ALTER TABLE "Session" ADD FOREIGN KEY ("task_id") REFERENCES "Task" ("id");

