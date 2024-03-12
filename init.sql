DROP TABLE IF EXISTS pets;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       first_name VARCHAR(255),
                       last_name VARCHAR(255),
                       email VARCHAR(255),
                       phone VARCHAR(255),
                       password VARCHAR(255),
                       user_status int
);

CREATE TABLE sessions (
                      id SERIAL PRIMARY KEY,
                      user_id INTEGER REFERENCES users (id)
);


CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE tags (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE pets (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(255),
                      status  VARCHAR(255),
                      tags INTEGER REFERENCES tags (id)
                      category INTEGER REFERENCES categories (id)
);


CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        pet_id INTEGER REFERENCES pets (id),
                        ship_date TIMESTAMP,
                        status VARCHAR(255),
                        complete BOOLEAN DEFAULT false
);