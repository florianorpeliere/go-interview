CREATE TABLE "public".currencies ( 
    id                   bigserial  NOT NULL,
    name                 varchar(3)  NOT NULL,
    dollars_rate         numeric(8,3)  NOT NULL,
	created_at           timestamp DEFAULT current_timestamp NOT NULL,
    CONSTRAINT pk_currencies PRIMARY KEY ( id )
);