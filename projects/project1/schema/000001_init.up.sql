create table users
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

create table catalogs
(
    id serial primary key,
    name varchar(255) not null,
    description text not null
);

create table books
(
    id serial primary key,
    name varchar(255) not null,
    description text,
    price double precision not null
);

create table user_catalogs
(
    id serial primary key,
    user_id int not null,
    catalog_id int not null,
    constraint user_id_fk foreign key(user_id) references users(id),
    constraint catalog_id_fk foreign key(catalog_id) references catalogs(id)
);



create table book_catalogs
(
    id serial primary key,
    book_id int not null,
    catalog_id int not null,
    constraint book_id_fk foreign key(book_id) references books(id),
    constraint catalog_id_fk foreign key(catalog_id) references catalogs(id)
)