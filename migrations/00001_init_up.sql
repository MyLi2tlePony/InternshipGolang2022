CREATE TABLE IF NOT EXISTS Users (
    ID INTEGER PRIMARY KEY UNIQUE NOT NULL,
    Balance INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS ReservedBalance(
    ID SERIAL PRIMARY KEY UNIQUE NOT NULL,
    UserID INTEGER NOT NULL,
    OrderID INTEGER NOT NULL,
    ServiceID INTEGER NOT NULL,
    CreateDate TIMESTAMP NOT NULL,
    Amount INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS ReservedBalanceHistory(
    ID SERIAL PRIMARY KEY UNIQUE NOT NULL,
    UserID INTEGER NOT NULL,
    OrderID INTEGER NOT NULL,
    ServiceID INTEGER NOT NULL,
    CreateDate TIMESTAMP NOT NULL,
    Amount INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS CancelledBalance(
    ID SERIAL PRIMARY KEY UNIQUE NOT NULL,
    UserID INTEGER NOT NULL,
    OrderID INTEGER NOT NULL,
    ServiceID INTEGER NOT NULL,
    CreateDate TIMESTAMP NOT NULL,
    Amount INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS ConfirmedBalance(
    ID SERIAL PRIMARY KEY UNIQUE NOT NULL,
    UserID INTEGER NOT NULL,
    OrderID INTEGER NOT NULL,
    ServiceID INTEGER NOT NULL,
    CreateDate TIMESTAMP NOT NULL,
    Amount INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS ReplenishedBalance(
    ID SERIAL PRIMARY KEY UNIQUE NOT NULL,
    UserID INTEGER NOT NULL,
    CreateDate TIMESTAMP NOT NULL,
    Amount INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS TransferredBalance(
    ID SERIAL PRIMARY KEY UNIQUE NOT NULL,
    SrcUserID INTEGER NOT NULL,
    DstUserID INTEGER NOT NULL,
    CreateDate TIMESTAMP NOT NULL,
    Amount INTEGER NOT NULL
);
