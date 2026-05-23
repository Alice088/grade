CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    room_name VARCHAR(20) NOT NULL CHECK(length(room_name) >= 5 AND length(room_name) <= 20),
    user_name VARCHAR(100) NOT NULL CHECK(length(user_name) >= 2 AND length(user_name) <= 100),
    start_at TIMESTAMPTZ NOT NULL,
    end_at TIMESTAMPTZ NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT now()
);