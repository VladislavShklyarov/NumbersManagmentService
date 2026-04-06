CREATE USER numbers_admin WITH PASSWORD 'admin_password';

GRANT ALL PRIVILEGES ON DATABASE numbers_db TO numbers_admin;

ALTER USER numbers_admin CREATEDB;
ALTER USER numbers_admin CREATEROLE;