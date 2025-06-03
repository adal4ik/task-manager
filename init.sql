-- ===== ENUM TYPES =====
CREATE TYPE task_status AS ENUM ('todo', 'done', 'overdue');
CREATE TYPE task_priority AS ENUM ('low', 'medium', 'high');

-- ===== EXTENSIONS =====
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ===== USERS =====
CREATE TABLE users (
    user_id     UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login       TEXT NOT NULL UNIQUE,
    email       TEXT NOT NULL UNIQUE,
    password    TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT now()
);

-- ===== TASKS =====
CREATE TABLE tasks (
    task_id      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id      UUID REFERENCES users(user_id) ON DELETE CASCADE,
    title        TEXT NOT NULL,
    status       task_status NOT NULL DEFAULT 'todo',
    priority     task_priority NOT NULL DEFAULT 'medium',
    due_date     TIMESTAMP, -- дедлайн
    created_at   TIMESTAMP NOT NULL DEFAULT now()
);
