CREATE OR REPLACE PROCEDURE CreateUser(user_id INTEGER, user_balance INTEGER)
AS $$
BEGIN
INSERT INTO Users (ID, Balance)
VALUES (user_id, user_balance);
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE CreateReservedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER)
AS $$
BEGIN
INSERT INTO ReservedBalance (UserID, OrderID, ServiceID, CreateDate, Amount)
VALUES (user_id, order_id, service_id, create_date, amount);
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE CreateReservedBalanceHistory(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER)
AS $$
BEGIN
INSERT INTO ReservedBalanceHistory (UserID, OrderID, ServiceID, CreateDate, Amount)
VALUES (user_id, order_id, service_id, create_date, amount);
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE CreateCancelledBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER)
AS $$
BEGIN
INSERT INTO CancelledBalance (UserID, OrderID, ServiceID, CreateDate, Amount)
VALUES (user_id, order_id, service_id, create_date, amount);
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE CreateConfirmedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER)
AS $$
BEGIN
INSERT INTO ConfirmedBalance (UserID, OrderID, ServiceID, CreateDate, Amount)
VALUES (user_id, order_id, service_id, create_date, amount);
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE CreateReplenishedBalance(user_id INTEGER, create_date TIMESTAMP, amount INTEGER)
AS $$
BEGIN
INSERT INTO ReplenishedBalance (UserID, CreateDate, Amount)
VALUES (user_id, create_date, amount);
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE CreateTransferredBalance(src_user_id INTEGER, dst_user_id INTEGER, create_date TIMESTAMP, amount INTEGER)
AS $$
BEGIN
INSERT INTO TransferredBalance (SrcUserID, DstUserID, CreateDate, Amount)
VALUES (src_user_id, dst_user_id, create_date, amount);
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE UpdateUserBalance(user_id INTEGER, amount INTEGER)
AS $$
BEGIN
    UPDATE Users SET Balance = Balance + amount WHERE id = user_id;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE DeleteUser(user_id INTEGER)
AS $$
BEGIN
    DELETE FROM Users WHERE ID = user_id;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE DeleteReservedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER)
AS $$
BEGIN
    DELETE FROM ReservedBalance
    WHERE UserID = user_id AND OrderID = order_id AND ServiceID = service_id;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE DeleteReservedBalanceHistory(user_id INTEGER, order_id INTEGER, service_id INTEGER)
AS $$
BEGIN
    DELETE FROM ReservedBalanceHistory
    WHERE UserID = user_id AND OrderID = order_id AND ServiceID = service_id;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE DeleteCancelledBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER)
AS $$
BEGIN
    DELETE FROM CancelledBalance
    WHERE UserID = user_id AND OrderID = order_id AND ServiceID = service_id;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE DeleteConfirmedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER)
AS $$
BEGIN
    DELETE FROM ConfirmedBalance
    WHERE UserID = user_id AND OrderID = order_id AND ServiceID = service_id;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE DeleteReplenishedBalance(user_id INTEGER, create_date TIMESTAMP)
AS $$
BEGIN
    DELETE FROM ReplenishedBalance
    WHERE UserID = user_id AND CreateDate = create_date;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE PROCEDURE DeleteTransferredBalance(src_user_id INTEGER, dst_user_id INTEGER, create_date TIMESTAMP)
AS $$
BEGIN
    DELETE FROM TransferredBalance
    WHERE SrcUserID = src_user_id AND DstUserID = dst_user_id AND CreateDate = create_date;
END;
$$ LANGUAGE PLPGSQL;
