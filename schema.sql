create database if not exists workshop_db;
use workshop_db;

create table if not exists attendees(
   id int auto_increment primary key,
   name varchar(225) not null,
   email varchar(225) not null,  
   index idx_at_email (email)
);

create table if not exists workshops(
   id int auto_increment primary key, 
   title varchar(225) not null,
   capacity int not null default 30
);

create table if not exists registrations(
   id int auto_increment primary key ,
   attendee_id int not null,
   workshop_id int not null,
   status varchar(30) not null default 'registered',
   unique key atttendee_worskshop_id (attendee_id,workshop_id),
   foreign key (attendee_id) references attendees (id) on delete cascade on update cascade,
   foreign key (workshop_id) references workshops (id) on delete cascade on update cascade ,
   index idx_workshop (workshop_id),
   index idx_attendee (attendee_id)

);