CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE users (
    id uuid primary key,
    email varchar(50) NOT NULL,
    username varchar(50) NOT NULL,
    password_hash varchar(150) NOT NULL,
    active bool default false,
    role user_role NOT NULL
);
