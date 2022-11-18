DROP PROCEDURE IF EXISTS CreateUser(user_id INTEGER, user_balance INTEGER);

DROP PROCEDURE IF EXISTS CreateReservedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER);

DROP PROCEDURE IF EXISTS CreateReservedBalanceHistory(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER);

DROP PROCEDURE IF EXISTS CreateCancelledBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER);

DROP PROCEDURE IF EXISTS CreateConfirmedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER, create_date TIMESTAMP, amount INTEGER);

DROP PROCEDURE IF EXISTS CreateReplenishedBalance(user_id INTEGER, create_date TIMESTAMP, amount INTEGER);

DROP PROCEDURE IF EXISTS CreateTransferredBalance(src_user_id INTEGER, dst_user_id INTEGER, create_date TIMESTAMP, amount INTEGER);

DROP PROCEDURE IF EXISTS UpdateUserBalance(user_id INTEGER, amount INTEGER);

DROP PROCEDURE IF EXISTS DeleteUser(user_id INTEGER);

DROP PROCEDURE IF EXISTS DeleteReservedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER);

DROP PROCEDURE IF EXISTS DeleteReservedBalanceHistory(user_id INTEGER, order_id INTEGER, service_id INTEGER);

DROP PROCEDURE IF EXISTS DeleteCancelledBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER);

DROP PROCEDURE IF EXISTS DeleteConfirmedBalance(user_id INTEGER, order_id INTEGER, service_id INTEGER);

DROP PROCEDURE IF EXISTS DeleteReplenishedBalance(row_id INTEGER);

DROP PROCEDURE IF EXISTS DeleteTransferredBalance(row_id INTEGER);


