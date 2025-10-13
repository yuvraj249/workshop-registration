create database if not exists workshop_db;
use workshop_db;

create table if not exists attendees(
   id int auto_increment primary key,
   AName varchar(225) not null,
   email varchar(225) not null,  
   index idx_at_email (email)
);

create table if not exists workshops(
    
);

create table if not exists registrations();