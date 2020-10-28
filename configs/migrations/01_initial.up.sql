create table items
(
    id             SERIAL PRIMARY KEY,
    name           TEXT                     NOT NULL,
    description    TEXT,
    date_purchased TIMESTAMP WITH TIME ZONE NOT NULL,
    price          INT                      NOT NULL
)