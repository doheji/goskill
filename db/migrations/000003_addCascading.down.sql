ALTER TABLE "Task" DROP CONSTRAINT fk_name;
ALTER TABLE "Session" DROP CONSTRAINT fk_session;

ALTER TABLE "Task" ADD
        FOREIGN KEY ("skill_id")
            REFERENCES "Skill" ("id");

ALTER TABLE "Session" ADD
        FOREIGN KEY ("task_id")
            REFERENCES "Task" ("id");