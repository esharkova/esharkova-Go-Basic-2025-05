-- name: GetUserByID :one
SELECT userid, firstname, lastname, passport_number FROM users WHERE userid = $1;

-- name: ListUsers :many
SELECT userid, firstname, lastname, passport_number FROM users ORDER BY lastname;

-- name: CreateUser :exec
INSERT INTO users (firstname, lastname, passport_number) VALUES ($1, $2, $3) RETURNING userid;

-- name: DeleteUser :exec
DELETE FROM users WHERE userid = $1;

-- name: GetTaskByID :one
SELECT taskid, task_number, description, create_datetime, completion_datetime, priority FROM tasks WHERE taskid = $1;

-- name: ListTasks :many
SELECT taskid, task_number, description, create_datetime, completion_datetime, priority FROM tasks ORDER BY task_number;

-- name: CreateTask :exec
INSERT INTO tasks (task_number, description, create_datetime, completion_datetime, priority) VALUES ($1, $2, $3, $4, $5) RETURNING taskid;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE taskid = $1;

-- name: UpdateTask :one
UPDATE tasks
SET
  task_number = $2,
  description = $3,
  priority = $4
WHERE taskid = $1
RETURNING
  taskid,
  task_number,
  description,
  create_datetime,
  completion_datetime,
  priority;

-- name: UpdateUser :one
UPDATE users
SET
  firstname = $2,
  lastname = $3
WHERE userid = $1
RETURNING
  userid,
  firstname,
  lastname,
  passport_number;  