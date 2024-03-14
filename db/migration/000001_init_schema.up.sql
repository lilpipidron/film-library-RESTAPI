CREATE TYPE user_role AS ENUM ('user', 'admin');
CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS actors
(
    actor_id      BIGSERIAL PRIMARY KEY,
    name          VARCHAR,
    surname       VARCHAR,
    gender        gender,
    date_of_birth DATE
);

CREATE TABLE  IF NOT EXISTS films
(
    film_id      BIGSERIAL PRIMARY KEY,
    title        VARCHAR(150),
    description  VARCHAR(1000),
    release_date DATE,
    rating       FLOAT
);

CREATE TABLE  IF NOT EXISTS actor_film
(
    actor_id BIGINT REFERENCES actors (actor_id),
    film_id  BIGINT REFERENCES films (film_id)
);