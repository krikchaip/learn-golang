CREATE TABLE IF NOT EXISTS snippets (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created TIMESTAMP NOT NULL DEFAULT NOW(),
  expires TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO users (name, email, hashed_password, created) VALUES (
  'Alice Jones',
  'alice@example.com',
  '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
  '2022-01-01 09:18:24'
);
