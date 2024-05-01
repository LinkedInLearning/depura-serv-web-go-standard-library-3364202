CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  uid UUID NOT NULL UNIQUE,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  email TEXT NOT NULL,
  age INTEGER,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

INSERT INTO users (uid, username, password, email, age, created_at, updated_at)
VALUES 
('6ba7b810-9dad-11d1-80b4-00c04fd430c8'::uuid, 'admin', 'admin', 'admin@example.com', 39, '2019-01-01 00:00:00', '2019-01-01 00:00:00');
