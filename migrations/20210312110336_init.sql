-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";
create type user_role as ENUM('client', 'provider');
create type user_status as ENUM('normal', 'blocked');
create table client
(
    id uuid primary key default uuid_generate_v4(),
    login text not null,
    key text not null,
    role user_role default 'client',
    rolled_id uuid not null,
    status user_status,
    last_login timestamptz,
    created_at timestamptz default now() not null
);

create type history_status as ENUM('planned', 'expired', 'succesfull', 'broken');
create type mark as ENUM('1', '2', '3', '4', '5');
create table history (
     id uuid primary key default uuid_generate_v4(),
     user_id uuid not null,
     created_at timestamptz default now() not null,
     uptime timestamptz,
     endtime timestamptz,
     latitude double precision default null,
     longitude double precision default null,
     geofile text default null,
     service uuid default null,
     feedbaack text default null,
     user_mark mark default '1'
);

create table provider (
    id uuid primary key default uuid_generate_v4(),
    public_name text not null,
    created_at timestamptz default now() not null,
    latitude double precision default null,
    longitude double precision default null,
    geofile text default null,
    info text default null,
    media text default null, --тут лежац фсе медиа файлики, по определенной логике внутри сервиса. в идеале должен быть cdn, но нец его пока
    worktime text default null,
    feedbaack text default null,
    mrkdown text default null, --mrkdwn text
    user_mark mark default '1'
);

create type service_type as ENUM('online', 'offlane', 'selfhosted', 'walk');
create table service (
    id uuid primary key default uuid_generate_v4(),
    public_name text not null,
    created_at timestamptz default now() not null,
    latitude double precision default null,
    longitude double precision default null,
    geofile text default null,
    info text default null,
    media text default null, --тут лежац фсе медиа файлики, по определенной логике внутри сервиса. в идеале должен быть cdn, но нец его пока
    worktime text default null,
    avaible service_type default 'offlane',
    user_mark mark default '1',
    user_feedback text default null,
    mrkdown text default null, --mrkdwn text
    provider_id uuid not null
);

create table tag (
    id uuid primary key default uuid_generate_v4(),
    public_name text not null,
    service uuid,
    provider uuid
);

create index auth on client (login, key);
create index tagged on tag (public_name);
create index provide on service (provider_id);
create index story on history (user_id, created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tag;
drop table service;
drop table provider;
drop table history;
drop table client;

drop type user_role;
drop type user_status;
drop type history_status;
drop type mark;
drop type service_type;
-- +goose StatementEnd
