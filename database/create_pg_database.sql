-- sources
create extension pgcrypto;

-- permisions
drop table if exists permission_levels; 
create table permission_levels (
	lvl smallint primary key,
	name varchar(32) unique not null
);
insert into permission_levels(lvl, name) values ( 0, 'user')
                                              , ( 999, 'admin');

-- users
drop table if exists users; 
create table users (
	nickname varchar(128) primary key,
	password varchar(256) not null,
	status smallint not null references permission_levels
);


-- users sessions
drop table if exists sessions; 
create table sessions (
	session_key text primary key default(gen_random_bytes(64)),
	user_nickname varchar(128) not null references users,
	create_time timestamp not null default(now())
);


drop table if exists folders; 
create table folders (
	id_uuid uuid primary key default(gen_random_uuid()),
	owner_nickname varchar(128) not null references users,
	folder varchar(32767) unique default ('/')
);


drop table if exists files; 
create table files(
	name varchar(240) not null check ( length(name) > 0 
                                       and name not like '.'
                                       and name not like '..'),
    id_folder uuid not null references folders,
	created timestamp not null default(now()),
	is_folder bool not null default(false),
	is_shared bool not null default(false),
	unique (id_folder, name)
);

-- triggers
-------------------------
-- Users
-------------------------
create or replace function onAddNewUser_function()
returns trigger as 
$body$
	begin
		insert into folders(owner_nickname, folder) values(new.nickname, '/' || new.nickname );
		return new;
	end
$body$
language plpgsql;

create trigger onAddNewUser_trigger
after insert on users for each row
execute procedure onAddNewUser_function();

create or replace function onDeleteUser_function()
returns trigger as 
$body$
	begin
		delete from folders
		where folders.owner_nickname like old.nickname;

		return old;
	end
$body$
language plpgsql;

create trigger onDeleteUser_trigger
before delete on users for each row
execute procedure onDeleteUser_function();


-------------------------
-- folders
-------------------------
create or replace function onDeleteFolder_function()
returns trigger as 
$body$
	begin
		delete from files
		where files.id_folder = old.id_uuid;

		return old;
	end
$body$
language plpgsql;

create trigger onDeleteFolder_trigger
before delete on folders for each row
execute procedure onDeleteFolder_function();


-------------------------
-- files
-------------------------
create or replace function onAddFile_function()
returns trigger as 
$body$
	begin
		if new.is_folder then
			insert into folders(owner_nickname, folder)
				select 	owner_nickname, folder || '/' || new.name
				from 	folders
				where 	folders.id_uuid = new.id_folder;
		end if;

		return new;
	end
$body$
language plpgsql;

create trigger onAddFile_trigger
after insert on files for each row
execute procedure onAddFile_function();


create or replace function onDeleteFile_function()
returns trigger as 
$body$
	begin
		if old.is_folder then
			delete from folders
			where folders.folder = ( select folder from folders where id_uuid = old.id_folder ) || '/' || old.name;
		end if;

		return old;
	end
$body$
language plpgsql;

create trigger onDeleteFile_trigger
before delete on files for each row
execute procedure onDeleteFile_function();


-- functions
-------------------------
-- Validation user key. 
-- Return null, if key key not exists.
-- Return user nickname, if key exists.
-------------------------
create or replace function test_user_session_key(  in userKey text
                                                , out out_nickname varchar(128)
                                                , out out_status smallint )
returns record as $body$
	begin
		select 	user_nickname, status
		into 	out_nickname, out_status
		from 	sessions inner join users on nickname = user_nickname
		where	session_key = userKey
		group by user_nickname, status;	
	
		if out_nickname is not null then
			update sessions
				set create_time = now()
				where userKey = session_key;
		end if;
	end
$body$
language plpgsql;


-------------------------
-- folder is exist
-------------------------
create or replace function test_folder_exists(  in target_folder varchar(32767)
                                        	 , out folder_exists bool )
returns bool as $body$
	begin
		folder_exists := exists (
			select * from folders where folder = target_folder
		);
	end
$body$
language plpgsql;


-------------------------
-- file is exist
-------------------------
create or replace function test_file_exists(  in target_folder varchar(32767)
                                           ,  in target_file varchar(240)
                                           , out file_exists bool )
returns bool as $body$
	begin
		file_exists := exists (
			select * 
			from 	folders inner join files on folders.id_uuid = files.id_folder
			where 	folders.folder = target_folder and files.name = target_file
		);
	end
$body$
language plpgsql;


---------------------------
-- return list of files in folder
---------------------------
create or replace function get_folder_content( in target_folder varchar(32767) )
returns table( name varchar(240)
             , created timestamp
             , is_folder bool
			 )
as $body$
	select 	files.name, files.created, files.is_folder
	from	files inner join folders on files.id_folder = folders.id_uuid
	where	folders.folder = target_folder;
$body$
language sql;


---------------------------
-- create file in folder
---------------------------
create or replace function create_file_in_folder( in target_folder varchar(32767)
                                                , in target_file_name varchar(240)
                                                , in target_file_is_folder bool)
returns void
as $body$
	insert into files(name, is_folder, id_folder) 
		select 	target_file_name, target_file_is_folder, id_uuid
	  	from	folders
	  	where 	folders.folder = target_folder;
$body$
language sql;


---------------------------
-- delete file in folder
---------------------------
create or replace function delete_file_in_folder( in target_folder varchar(32767)
                                                , in target_file_name varchar(240))
returns void
as $body$
	delete 	from files
	where	files.name = target_file_name
		and files.id_folder = ( 
				select 	id_uuid 
				from 	folders 
				where 	folders.folder = target_folder
			);
$body$
language sql;

--select from delete_file_in_folder('/', 'test.txt');


---------------------------
-- Register new user
---------------------------
create or replace function registration_new_user( in target_login varchar(128)
                                                , in target_password varchar(256))
returns void
as $body$
	insert into users(nickname, password) values(target_login, target_password);
$body$
language sql;


---------------------------
-- Authorization user
---------------------------
create or replace function authorization_user( in target_login varchar(128)
                                             , in target_password varchar(256)
                                             , out user_access_key text
                                             , out user_status smallint)
returns record
as $body$	
	begin
		if not exists(
			select * from users
			where nickname = target_login and password = target_password
		)
		then
			raise exception 'Invalid login or password!';
		end if;
	
		insert into sessions(user_nickname) values(target_login) 
			returning session_key into user_access_key;
		
		select status into user_status from users where nickname = target_login;
	end
$body$
language plpgsql;


---------------------------
-- Logout user
---------------------------
create or replace function logout_user( in target_session_key text)
returns void
as $body$
	delete from sessions where session_key = target_session_key;
$body$
language sql;


---------------------------
-- Close all user sessions
---------------------------
create or replace function close_all_user_sessions( in nickname varchar(128))
returns void
as $body$
	delete from sessions where user_nickname = nickname;
$body$
language sql;
