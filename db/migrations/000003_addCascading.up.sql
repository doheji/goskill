ALTER TABLE "Task"
    ADD CONSTRAINT fk_name
        FOREIGN KEY ("skill_id")
            REFERENCES "Skill" ("id")
            ON DELETE CASCADE;

ALTER TABLE "Session"
    ADD CONSTRAINT fk_session
        FOREIGN KEY ("task_id")
            REFERENCES "Task" ("id")
            ON DELETE CASCADE;