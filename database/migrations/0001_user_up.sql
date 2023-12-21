CREATE TABLE Companies (
                           Id INT AUTO_INCREMENT PRIMARY KEY,
                           Name VARCHAR(255) NOT NULL UNIQUE,
                           AddressStreet VARCHAR(255),
                           AddressCity VARCHAR(255),
                           AddressState VARCHAR(255),
                           AddressCountry VARCHAR(255),
                           PostalCode VARCHAR(20),
                           Industry VARCHAR(255),
                           Size INT,
                           ContactPhone VARCHAR(20),
                           ContactEmail VARCHAR(255),
                           Logo BLOB
);

CREATE TABLE Users (
                       Id INT AUTO_INCREMENT PRIMARY KEY,
                       FullName VARCHAR(255),
                       Email VARCHAR(255) NOT NULL UNIQUE,
                       Phone VARCHAR(20) NOT NULL,
                       CompanyId INT NOT NULL,
                       Username VARCHAR(255) NOT NULL UNIQUE,
                       Password VARCHAR(255) NOT NULL,
                       PasswordSalt VARCHAR(50) NOT NULL,
                       IsEnabled Boolean,
                       FOREIGN KEY (CompanyId) REFERENCES Companies(Id)
);