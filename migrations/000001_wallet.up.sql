create table if not exists "Wallet" (
    wallet_id serial primary key,
    number    integer        unique not null,
    code      varchar(20)    not null,
    amount    numeric(30, 2) not null default 0.00,
    status      varchar(20) not null default 'Success'
)