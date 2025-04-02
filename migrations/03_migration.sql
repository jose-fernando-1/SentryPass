UPDATE Entries
SET password_encrypted = 'new_encrypted_password_1'
WHERE username = 'user1@example.com';

UPDATE Entries
SET password_encrypted = 'new_encrypted_password_2'
WHERE username = 'user2@example.com';