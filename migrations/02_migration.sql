INSERT INTO User (master_password_hash, master_password_salt)
VALUES 
    ('hashed_password_1', 'salt_1');

INSERT INTO Categories (type)
VALUES 
    ('Site'),
    ('Email'),
    ('Other');

INSERT INTO Entries (title, username, password_encrypted, details_encrypted, id_category, id_user) 
VALUES 
    ('Entry 1', 'user1@example.com', 'encrypted_password_1', 'First user notes', 1, 1),
    ('Entry 2', 'user2@example.com', 'encrypted_password_2', 'Second user notes', 2, 1),
    ('Entry 3', 'user3@example.com', 'encrypted_password_3', NULL, 3, 1);