CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    room_name VARCHAR(20) NOT NULL,
    room_name VARCHAR(20) NOT NULL,
    start_at DATE NOT NULL,
    end_at DATE NOT NULL,
    created DATE NOT NULL DEFAULT NOW() 
);