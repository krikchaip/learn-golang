CREATE TABLE snippets (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created TIMESTAMP NOT NULL DEFAULT NOW(),
  expires TIMESTAMP NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
