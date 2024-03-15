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

CREATE VIEW actor_film_view AS
SELECT a.actor_id AS actor_id, a.name AS actor_name, a.surname AS actor_surname, a.gender AS actor_gender, a.date_of_birth AS actor_date_of_birth,
       f.film_id AS film_id, f.title AS film_title, f.description AS film_description, f.release_date AS film_release_date, f.rating AS film_rating
FROM actors a
         JOIN actor_film af ON a.actor_id = af.actor_id
         JOIN films f ON af.film_id = f.film_id;

