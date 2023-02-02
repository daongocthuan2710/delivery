CREATE TABLE provinces
(
    id       int  not null,
    name     text not null,
    ghn_id   int,
    ghn_code text
);
ALTER TABLE provinces
    ADD CONSTRAINT province_pkey PRIMARY KEY (id);

CREATE TABLE districts
(
    id          int  not null,
    name        text not null,
    province_id int  not null,
    ghn_id      int,
    ghn_code    text
);
ALTER TABLE districts
    ADD CONSTRAINT district_pkey PRIMARY KEY (id);
ALTER TABLE districts
    ADD CONSTRAINT district_province_fkey FOREIGN KEY (province_id) REFERENCES provinces (id);

CREATE TABLE wards
(
    id          int  not null,
    name        text not null,
    district_id int  not null,
    ghn_code    text
);
ALTER TABLE wards
    ADD CONSTRAINT ward_pkey PRIMARY KEY (id);
ALTER TABLE wards
    ADD CONSTRAINT ward_district_fkey FOREIGN KEY (district_id) REFERENCES districts (id);
