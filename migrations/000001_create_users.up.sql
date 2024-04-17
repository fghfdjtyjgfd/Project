create table users
(
	id serial
		constraint users_pk
			primary key,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null,
	deleted_at timestamp with time zone,
	email varchar(135),
	phone_number varchar(10) not null,
	first_name varchar(256),
	last_name varchar(256)
);

create index users_deleted_at_index
	on users (deleted_at);

create unique index users_email_uindex
	on users (email);

create unique index users_phone_number_uindex
	on users (phone_number);

