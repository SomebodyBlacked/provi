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
  ('user2', 'user2@example.com');
INSERT INTO products (name, description, price)
VALUES ('Product 1', 'This is product 1', 10.99),
  ('Product 2', 'This is product 2', 19.99);