-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE customers (
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(256) NOT NULL,
    password        VARCHAR(256) NOT NULL,
    fullname       VARCHAR(256),
    email           VARCHAR(256),
    phone           VARCHAR(256),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

CREATE TABLE cinema_halls (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(256),
    capacity        INTEGER,
    location        VARCHAR(256),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

CREATE TABLE movies (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(256),
    genre           VARCHAR(256),
    duration        VARCHAR(256),
    rating          VARCHAR(256),
    release_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description     VARCHAR(256),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

CREATE TABLE showtimes (
    id              SERIAL PRIMARY KEY,
    movie_id        INTEGER,
    cinema_hall_id  INTEGER,
    showtime_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    price           INTEGER,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM',
    CONSTRAINT fk_movies
        FOREIGN KEY (movie_id) REFERENCES movies(id),
    CONSTRAINT fk_cinema_halls
        FOREIGN KEY (cinema_hall_id) REFERENCES cinema_halls(id)
);

CREATE TABLE bookings (
    id              SERIAL PRIMARY KEY,
    customer_id     INTEGER,
    showtime_id     INTEGER,
    booking_date    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    seat_number     INTEGER,
    status          VARCHAR(256),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM',
    CONSTRAINT fk_customers
        FOREIGN KEY (customer_id) REFERENCES customers(id),
    CONSTRAINT fk_showtimes
        FOREIGN KEY (showtime_id) REFERENCES showtimes(id)
);

ALTER TABLE customers
    ADD CONSTRAINT unique_username UNIQUE (username);
ALTER TABLE cinema_halls
    ADD CONSTRAINT unique_name UNIQUE (name);

CREATE INDEX idx_username ON customers(username);
CREATE INDEX idx_movie ON movies(title);

-- +migrate StatementEnd