-- Crear la base de datos
CREATE DATABASE PROVI;
USE PROVI;
-- Crear las tablas
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);
CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  price DECIMAL(10, 2) NOT NULL
);
-- Insertar datos de ejemplo
INSERT INTO users (username, email)
VALUES ('user1', 'user1@example.com'),
  ('user2', 'user2@example.com'),
  ('user3', 'user3@example.com'),
  ('user4', 'user4@example.com'),
  ('user5', 'user5@example.com'),
  ('user6', 'user6@example.com'),
  ('user7', 'user7@example.com'),
  ('user8', 'user8@example.com'),
  ('user9', 'user9@example.com');
INSERT INTO products (name, description, price)
VALUES ('Product 1', 'This is product 1', 10.99),
  ('Product 2', 'This is product 2', 19.99),
  ('Product 3', 'This is product 3', 29.99),
  ('Product 4', 'This is product 4', 39.99),
  ('Product 5', 'This is product 5', 49.99),
  ('Product 6', 'This is product 6', 59.99),
  ('Product 7', 'This is product 7', 69.99),
  ('Product 8', 'This is product 8', 79.99),
  ('Product 9', 'This is product 9', 89.99),
  ('Product 10', 'This is product 10', 99.99),
  ('Product 11', 'This is product 11', 109.99),
  ('Product 12', 'This is product 12', 119.99),
  ('Product 13', 'This is product 13', 129.99),
  ('Product 14', 'This is product 14', 139.99),
  ('Product 15', 'This is product 15', 149.99),
  ('Product 16', 'This is product 16', 159.99);
