CREATE TABLE product(
    id serial primary key not null,
    name varchar(100) not null
);

INSERT INTO product (name) VALUES 
    ('Cadeira'),
    ('Mesa'),
    ('Toalha'),
    ('Fogão'),
    ('Batedeira'),
    ('Pia'),
    ('Torneira'),
    ('Forno'),
    ('Gaveta'),
    ('Copo');
