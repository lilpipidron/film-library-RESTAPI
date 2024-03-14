CREATE TYPE user_role AS ENUM ('user', 'admin');
CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE actors (
                        actor_id BIGINT PRIMARY KEY,
                        name VARCHAR,
                        surname VARCHAR,
                        gender gender,
                        date_of_birth DATE
);

CREATE TABLE actor_film (
                            actor_id BIGINT REFERENCES actors(actor_id),
                            film_id BIGINT REFERENCES film(film_id)
);

CREATE TABLE film (
                      film_id BIGINT PRIMARY KEY,
                      title VARCHAR(150),
                      description VARCHAR(1000),
                      release_date DATE,
                      rating FLOAT
);