CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    email citext UNIQUE NOT NULL,
    password bytea, 
    created_At timestamp(0) with time zone NOT NULL DEFAULT NOW()
    )


