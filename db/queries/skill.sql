-- name: CreateSkill :one
INSERT INTO "Skill" ("learner", "name")
VALUES ($1, $2)
RETURNING *;

-- name: GetSkillsByLearner :many
SELECT *
FROM "Skill"
WHERE learner=$1
LIMIT $2
OFFSET $3;

-- name: GetAllSkills :many
SELECT *
FROM "Skill"
LIMIT $1
OFFSET $2;

-- name: GetSkillByID :one
SELECT *
FROM "Skill"
WHERE id=$1;

-- name: DeleteSkillByID :exec
DELETE FROM "Skill"
WHERE id=$1;

-- name: GetSkillBySession :one
SELECT s.*
FROM "Skill" as s, "Session" AS Se, "Task" as T
WHERE Se.id = $1 AND Se.task_id = t.id AND t.skill_id=s.id
FOR NO KEY UPDATE;

-- name: IncreaseSkillScore :one
UPDATE "Skill"
SET "score"="score" + $1
WHERE id = $2
RETURNING *;