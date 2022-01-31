CREATE TABLE "parameters"
(
    id_parameter serial not null unique,
    name varchar(255) not null,
    unit_of_product varchar(255) not null,
    PRIMARY KEY (id_parameter)
);

INSERT INTO parameters(name, unit_of_product) VALUES ('weight', 'kilograms');
INSERT INTO parameters(name, unit_of_product) VALUES ('size', 'meters');
INSERT INTO parameters(name, unit_of_product) VALUES ('price', 'dollars');

CREATE TABLE "group_of_parameters"
(
    id_group_parameter serial not null unique,
    name varchar(255) not null,
    parameters int not null,
    PRIMARY KEY (id_group_parameter),
    FOREIGN KEY (parameters) REFERENCES parameters (id_parameter)
);

INSERT INTO group_of_parameters(name, parameters) VALUES ('physical state', 1);
INSERT INTO group_of_parameters(name, parameters) VALUES ('price in the world', 3);
INSERT INTO group_of_parameters(name, parameters) VALUES ('size', 2);
INSERT INTO group_of_parameters(name, parameters) VALUES ('undefined', 1);

CREATE TABLE "group_of_products"
(
    id_group_product serial not null unique,
    name varchar(255) not null,
    group_parameters int not null,
    PRIMARY KEY (id_group_product),
    FOREIGN KEY (group_parameters) REFERENCES group_of_parameters (id_group_parameter)
);

INSERT INTO group_of_products(name, group_parameters) VALUES ('milk products', 1);
INSERT INTO group_of_products(name, group_parameters) VALUES ('cheese products', 3);
INSERT INTO group_of_products(name, group_parameters) VALUES ('chocolate products', 2);
INSERT INTO group_of_products(name, group_parameters) VALUES ('cake products', 2);


CREATE TABLE "products"
(
    id_product serial not null unique,
    name varchar(255) not null ,
    group_of_products int not null,
    description varchar(255) not null,
    release_date timestamp without time zone,
    parameters_id int not null,
    PRIMARY KEY (id_product),
    FOREIGN KEY (group_of_products) REFERENCES group_of_products(id_group_product),
    FOREIGN KEY (parameters_id) REFERENCES parameters(id_parameter)
);

INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('grandmother Any ', 1, 'very testy', current_timestamp, 2);
INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('Mozzarella', 2, 'expensive cheese', current_timestamp, 1);
INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('Alpen Gold', 3, 'the best chocolate in the world', current_timestamp, 3);
INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('Spartak', 4, 'testy cake', current_timestamp, 3);
INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('Snikers', 4, 'testy snikers cake', current_timestamp, 2);
INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('Milka', 3, 'i really like this chocolate', current_timestamp, 1);
INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('Langr', 2, 'i do not like this cheese', current_timestamp, 3);
INSERT INTO products(name, group_of_products, description, release_date, parameters_id) VALUES ('Savushkin', 1, 'the best milk', current_timestamp, 3);