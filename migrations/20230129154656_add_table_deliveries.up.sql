CREATE TABLE deliveries
(
    id                    text                                   not null,
    code                  text                     default ''    not null,
    tracking_code         text                     default 0     not null,
    note                  text                     default ''    not null,
    status                text                     default ''    not null,
    created_at            timestamp with time zone default now() not null,
    updated_at            timestamp with time zone default now() not null,
    value                 numeric                  default 0     not null,
    cod                   numeric                  default 0     not null,
    weight                numeric                  default 0     not null,
    service_code          text                     default ''    not null,
    partner_status        text                     default ''    not null,
    partner_identity_code text                     default ''    not null,
    from_name             text                     default ''    not null,
    from_phone            text                     default ''    not null,
    from_address          text                     default ''    not null,
    from_province_code    int                      default 0     not null,
    from_district_code    int                      default 0     not null,
    from_ward_code        int                      default 0     not null,
    to_name               text                     default ''    not null,
    to_phone              text                     default ''    not null,
    to_address            text                     default ''    not null,
    to_province_code      int                      default 0     not null,
    to_district_code      int                      default 0     not null,
    to_ward_code          int                      default 0     not null
);
ALTER TABLE deliveries
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);

CREATE INDEX orders_code_index
    ON deliveries (code);
CREATE INDEX orders_tracking_code_index
    ON deliveries (tracking_code);

