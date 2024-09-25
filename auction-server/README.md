Auction prototype for coldplay tickets



Pg tables:
1. Users:
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,           -- Auto-incrementing primary key
    name TEXT NOT NULL,                   -- Name of the user
    salt TEXT NOT NULL,
    pw_hash TEXT NOT NULL,                -- Argon2 hashed password
    email TEXT NOT NULL UNIQUE,           -- Email address (must be unique)
    mobile TEXT NOT NULL UNIQUE,          -- Mobile number (must be unique)
    created_at TIMESTAMPTZ DEFAULT NOW()  -- Timestamp when the user was created
);

2. Tickets:
CREATE TABLE tickets (
    ticket_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id),  -- The seller of the ticket
    event_date DATE NOT NULL,                        -- Date of the event (single concert date)
    number_of_tickets INT NOT NULL CHECK (number_of_tickets > 0), -- Number of tickets being sold
    seat_info JSONB,                                 -- JSONB object to store seat details (e.g., seat numbers, block, level)
    price NUMERIC(10, 2) NOT NULL,                   -- Asking price or starting price for the ticket(s)
    best_offer NUMERIC(10, 2),                       -- Highest bid received for the ticket (nullable)
    best_bidder_id INT REFERENCES users(user_id),    -- The user ID of the highest bidder
    auction_end TIMESTAMPTZ NOT NULL,                -- End time of the auction
    created_at TIMESTAMPTZ DEFAULT NOW()             -- Timestamp when the ticket was listed
);

3. Sales:
CREATE TABLE sales (
    sale_id SERIAL PRIMARY KEY,
    ticket_id INT NOT NULL REFERENCES tickets(ticket_id),                    -- The sold ticket
    buyer_id INT NOT NULL REFERENCES users(user_id),                         -- User who bought the ticket
    sale_price NUMERIC(10, 2) NOT NULL,                                      -- Final sale price of the ticket
    sale_time TIMESTAMPTZ DEFAULT NOW(),                                     -- Time of sale
    created_at TIMESTAMPTZ DEFAULT NOW()                                     -- Timestamp when the sale record was created
);
