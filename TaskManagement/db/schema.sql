CREATE TABLE tasks (
    taskid              SERIAL PRIMARY KEY,    -- Идентификатор задачи (автоинкремент)
    task_number         VARCHAR(255) NOT NULL, -- Номер задачи
    description         TEXT,                  -- Описание задачи
    create_datetime     TIMESTAMPTZ NOT NULL,    -- Дата создания
    completion_datetime TIMESTAMPTZ,             -- Дата выполнения (может быть NULL)
    priority            INT                    -- Приоритет
);

alter table tasks
    owner to esharkova;

CREATE TABLE users (
    userid          INT PRIMARY KEY,
    firstname       VARCHAR(100) NOT NULL,
    lastname        VARCHAR(100) NOT NULL,
    passport_number VARCHAR(20) NOT NULL
);

alter table users
    owner to esharkova;