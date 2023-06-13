DROP TABLE IF EXISTS pareto CASCADE;
DROP TABLE IF EXISTS users CASCADE;

SET timezone TO '+03';

create table method (
    id      integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id integer NOT NULL,
    method_name  varchar(256) NOT NULL,
    name    VARCHAR(256) NOT NULL,
    data    bytea
);

CREATE TABLE users (
   id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
   username VARCHAR(256) NOT NULL,
   password VARCHAR(256) NOT NULL
);