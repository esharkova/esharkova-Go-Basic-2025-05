-- +goose Up
CREATE TABLE tasks (
    taskid              SERIAL PRIMARY KEY,    -- Идентификатор задачи (автоинкремент)
    task_number         VARCHAR(255) NOT NULL, -- Номер задачи
    description         TEXT,                  -- Описание задачи
    create_datetime     TIMESTAMPTZ NOT NULL,    -- Дата создания
    completion_datetime TIMESTAMPTZ,             -- Дата выполнения (может быть NULL)
    priority            INT                    -- Приоритет
);

INSERT INTO tasks (task_number, description, create_datetime, completion_datetime, priority)
VALUES
    ('TASK-001', 'Описание первой задачи', NOW(), NULL, 1),
    ('TASK-002', 'Описание второй задачи', NOW(), NOW(), 2);


-- +goose Down
drop table tasks;