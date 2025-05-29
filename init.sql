-- ===== ENUM TYPES =====
CREATE TYPE task_status AS ENUM ('todo', 'done');
CREATE TYPE task_priority AS ENUM ('low', 'medium', 'high');
CREATE TYPE member_role AS ENUM ('owner', 'member');

-- ===== EXTENSIONS =====
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ===== USERS =====
CREATE TABLE users (
    user_id     UUID PRIMARY KEY,
    login       TEXT NOT NULL UNIQUE,
    email       TEXT NOT NULL UNIQUE,
    password    TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT now()
);

-- ===== WORKSPACES =====
CREATE TABLE workspace (
    workspace_id UUID PRIMARY KEY,
    name         TEXT NOT NULL,
    is_private   BOOLEAN NOT NULL DEFAULT false,
    created_at   TIMESTAMP NOT NULL DEFAULT now()
);

-- ===== WORKSPACE MEMBERS =====
CREATE TABLE workspace_member (
    user_id      UUID REFERENCES users(user_id) ON DELETE CASCADE,
    workspace_id UUID REFERENCES workspace(workspace_id) ON DELETE CASCADE,
    role         member_role NOT NULL,
    joined_at    TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, workspace_id)
);

-- ===== TASKS =====
CREATE TABLE tasks (
    task_id       UUID PRIMARY KEY,
    workspace_id  UUID REFERENCES workspace(workspace_id) ON DELETE CASCADE,
    user_id       UUID REFERENCES users(user_id) ON DELETE SET NULL,
    title         TEXT NOT NULL,
    status        task_status NOT NULL DEFAULT 'todo',
    priority      task_priority NOT NULL DEFAULT 'medium',
    created_at    TIMESTAMP NOT NULL DEFAULT now()
);
