-- name: UpdateTask :one
UPDATE "Task"
SET "name" = CASE
    WHEN @toUpdateName::boolean = TRUE THEN @updatedName
    ELSE "name"
    END,
    "skill_id" = CASE
        WHEN @toUpdateSkillID::boolean = TRUE THEN @updatedSkillID
        ELSE "skill_id"
    END
WHERE id = @taskID
RETURNING *;

-- name: GetTasksByLearnerName :many
SELECT t.*
FROM "Task" AS t, "Skill" AS s
WHERE s.learner= $1 AND s.id= t.skill_id
LIMIT $2
OFFSET $3;

-- name: GetTasksBySkillID :many
SELECT *
FROM "Task"
WHERE skill_id=$1
LIMIT $2
OFFSET $3;

-- name: GetTaskByID :one
SELECT *
FROM "Task"
WHERE id = $1;

-- name: CreateTask :one
INSERT INTO "Task" ("name", "skill_id")
VALUES ($1, $2)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM "Task"
WHERE id = $1;




