-- sources
create extension pgcrypto

-- permisions
create table permission_levels (
	lvl smallint primary key,
	name varchar(32) unique not null
);
insert into permission_levels(lvl, name) values ( 0, 'user')
                                              , ( 999, 'admin');

-- users
create table users (
	nickname varchar(128) primary key,
	password varchar(256) not null,
	status smallint 
);

-- users sessions
create table sessions (
	session_key text primary key,
	user_nickname text not null references users,
	create_time timestamp not null default(now())
);

-- Folder
-- Create table 'folders'
create table folders (
	id_uuid uuid primary key default(gen_random_uuid()),
	owner_nickname varchar(128) not null references users,
	folder varchar(32767) unique default ('/' || owner_nickname)
);
-- Insert root folder
insert into folders (folder) values('/');
-- Update, for user cycle links
alter table folders add parent varchar(240) not null references folders default ('/');


create table files(
	name varchar(240) not null check ( length(name) > 0 
                                       and name not like '.'
                                       and name not like '..'),
    id_floder uuid not null references folders
	created timestamp not null default(now()),
	is_folder bool not null default(false),
	is_shared bool not null default(false)
);