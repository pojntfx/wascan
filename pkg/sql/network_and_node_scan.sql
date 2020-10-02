-- Remove old tables if they exist
drop table if exists nodes;
drop table if exists network_scans;
drop table if exists node_scans;
drop table if exists services;
drop table if exists network_scans_nodes;
drop table if exists node_scans_services_nodes;
drop table if exists node_node_scans_network_scans;
-- Create new tables
create table nodes (
    mac_address text not null primary key,
    ip_address text not null,
    vendor text not null,
    registry text not null,
    organization text not null,
    address text not null,
    visible integer not null
);
create table network_scans (
    id integer not null primary key,
    created_at date not null,
    done integer not null
);
create table node_scans (
    id integer not null primary key,
    created_at date not null,
    done integer not null
);
create table services (
    service_name text not null,
    port_number integer not null primary key,
    transport_protocol text not null,
    description text not null,
    assignee text not null,
    contact text not null,
    registration_date text not null,
    modification_date text not null,
    reference text not null,
    service_code text not null,
    unauthorized_use_reported text not null,
    assignment_notes text not null
);
-- Create join tables
create table network_scans_nodes (
    id integer not null primary key,
    created_at date not null,
    node_id text not null,
    network_scan_id integer not null
);
create table node_scans_services_nodes (
    id integer not null primary key,
    created_at date not null,
    service_id integer not null,
    node_id text not null,
    node_scan_id integer not null
);
create table node_node_scans_network_scans (
    id integer not null primary key,
    created_at date not null,
    node_id text not null,
    network_scan_id integer not null,
    node_scan_id integer not null
);