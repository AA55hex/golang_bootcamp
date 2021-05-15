create database if not exists book_store

use book_store;

create table if not exists genre
(
    id int not null,
    name varchar(20) not null,
    primary key (id)
);

create table if not exists book
(
    id int auto_increment not null,
    name varchar(100) not null,
    price float default 0 not null,
    genre int not null,
    amount int default 0 not null,
    primary key (id),
    foreign key (genre) references genre(id) on delete cascade
);
COMMIT;
