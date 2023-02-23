-- name: GetUpcomingSessionsBySkill :many
SELECT Se.*
FROM "Session" AS Se, "Skill" AS S
WHERE S.learner = $1 AND S.id = $2 AND Se.startTime > now()
LIMIT $3
OFFSET $4;

-- name: GetUpcomingSessionsByTask :many
SELECT Se.*
FROM "Session" AS Se, "Skill" AS S, "Task" as T
WHERE S.learner = $1 AND T.id = $2 AND Se.startTime > now()
LIMIT $3
OFFSET $4;

-- name: GetUpcomingSessions :many
SELECT Se.*
FROM "Session" AS Se, "Skill" AS S
WHERE S.learner = $1 AND Se.startTime > now()
LIMIT $2
OFFSET $3;

-- name: GetAllSessionsBySkill :many
SELECT Se.*
FROM "Session" AS Se, "Skill" AS S
WHERE S.learner = $1 AND S.id = $2
LIMIT $3
OFFSET $4;

-- name: GetAllSessionsByTask :many
SELECT Se.*
FROM "Session" AS Se, "Skill" AS S, "Task" as T
WHERE S.learner = $1 AND T.id = $2
LIMIT $3
OFFSET $4;

-- name: GetAllSessions :many
SELECT Se.*
FROM "Session" AS Se, "Skill" AS S
WHERE S.learner = $1
LIMIT $2
OFFSET $3;

-- name: GetSessionByID :one
SELECT *
FROM "Session"
WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM "Session"
WHERE id = $1;

-- name: CreateSession :one
INSERT INTO "Session" (
    "task_id",
    "name",
    "description",
    "goal",
    "location",
    "startTime",
    "duration"
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateSession :one
UPDATE "Session"
SET "task_id" = CASE
                 WHEN @toUpdateTaskID::boolean = TRUE THEN @updatedTaskID
    ELSE "taskID"
END,
"name" = CASE
                 WHEN @toUpdateName::boolean = TRUE THEN @updatedName
    ELSE "name"
END,
"description" = CASE
                 WHEN @toUpdateDescription::boolean = TRUE THEN @updatedDescription
    ELSE "description"
END,
"goal" = CASE
                 WHEN @toUpdateGoal::boolean = TRUE THEN @updatedGoal
    ELSE "goal"
END,
"location" = CASE
                 WHEN @toUpdateLocation::boolean = TRUE THEN @updatedLocation
    ELSE "location"
END,
"attended" = CASE
                 WHEN @toUpdateAttended::boolean = TRUE THEN @updatedAttended
    ELSE "attended"
END,
"startTime" = CASE
                 WHEN @toUpdateStartTime::boolean = TRUE THEN @updatedStartTime
    ELSE "startTime"
END,
"duration" = CASE
                 WHEN @toUpdateDuration::boolean = TRUE THEN @updatedDuration
    ELSE "duration"
END
WHERE id = @SessionID
RETURNING *;

