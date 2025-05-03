-- Таблица пользователей
CREATE TABLE users (
                       id          SERIAL PRIMARY KEY,
                       username    TEXT NOT NULL UNIQUE,
                       password    TEXT NOT NULL,
                       created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица постов
CREATE TABLE posts (
                       id          SERIAL PRIMARY KEY,
                       title       TEXT NOT NULL,
                       content     TEXT NOT NULL,
                       user_id     INT NOT NULL,  -- Ссылка на пользователя
                       created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (user_id) REFERENCES users(id)
);
