CREATE TABLE credit_cards (
    id uuid NOT NULL,
    name VARCHAR NOT NULL,
    number VARCHAR NOT NULL,
    expmonth VARCHAR NOT NULL,
    expyear VARCHAR,
	CVV VARCHAR NOT NULL,
	balance float not null,
	balance_limit float not null,
    PRIMARY KEY (id)
);

CREATE TABLE transactions (
    id uuid NOT NULL,
	credit_card_id uuid NOT NULL references credit_cards(id),
    amount float NOT NULL,
    status VARCHAR NOT NULL,
    description VARCHAR,
	store VARCHAR NOT NULL,
	created_at timestamp not null,
    PRIMARY KEY (id)
);