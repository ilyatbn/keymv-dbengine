CREATE KEYSPACE keymv
WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

-- primary tables
USE keymv

CREATE TABLE organizations (
    orgid uuid,
    orgname text,
    PRIMARY KEY (orgid)
);

CREATE TABLE users (
	userid uuid,
	email text,
	org uuid,
	role uuid,
	password text,
    PRIMARY KEY (email)
);

CREATE TABLE roles (
	roleid uuid,
	rolename text,
	roletype text,
	PRIMARY KEY (roleid)
);

-- in the future, can separate each of these to separate keyspaces per organization

CREATE TABLE user_settings (
	userid uuid,
	enabled boolean,
	registration_date date,
	last_logon_time date,
	firstname text,
	lastname text,
	PRIMARY KEY (userid)
);

CREATE TABLE user_groups (
	userid uuid,
	group_id uuid,
	PRIMARY KEY (userid)
);

CREATE TABLE groups (
	group_id uuid,
	group_name text,
	orgid uuid,
	PRIMARY KEY (group_id)
);

