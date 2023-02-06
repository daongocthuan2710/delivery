CREATE TABLE delivery_histories
(
    id          text                                   not null,
    delivery_id text                                   not null,
    status      text                                   not null,
    created_at  timestamp with time zone default now() not null
);

ALTER TABLE delivery_histories
    ADD CONSTRAINT delivery_history_pkey PRIMARY KEY (id);
ALTER TABLE delivery_histories
    ADD CONSTRAINT delivery_history_delivery_fkey FOREIGN KEY (delivery_id) REFERENCES deliveries(id);