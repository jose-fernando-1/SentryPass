# 📜 SentryPass — Requirements Analysis

## 📖 1. Project Overview
SentryPass is a personal project developed in Go, focused on creating a secure, practical, and lightweight password manager, similar to Kaspersky Password Manager. The system aims for encrypted local storage, with future possibilities for cloud synchronization and mobile device integration.

---

## ✅ 2. Functional Requirements (FR)

| Code  | Description                                                                                      |
|-------|--------------------------------------------------------------------------------------------------|
| FR01  | Secure user authentication using hashed master password.                                                 |
| FR02  | CRUD for encrypted credentials (add, view, edit, and delete passwords).                                    |
| FR03  | Organization of passwords into categories (e.g., website, email, other).               |
| FR04  | Locally encrypted vault (secure storage).                                                        |
| FR05  | Generator of strong, customizable passwords.     |
| FR06  | Search and filter entries.                                                                     |
| FR07  | Automatic vault lock after a period of inactivity.                                               |

---

## ✅ 3. Non-Functional Requirements (NFR)

| Code  | Description                                                                                      |
|-------|--------------------------------------------------------------------------------------------------|
| NFR01 | Security with AES-256 encryption and key derivation via PBKDF2 or Argon2.                        |
| NFR02 | Portability with standalone executable distribution for Windows.                                 |
| NFR03 | Simple, user-friendly, and responsive interface.               |
| NFR04 | High performance in vault opening and operation execution.                                       |
| NFR05 | Modular and documented code, with automated tests.                                               |
| NFR06 | Planning for secure automatic updates in the future.                                             |

---

## ✅ 4. Technical Requirements (TR)

| Code  | Description                                                                                                   |
|-------|-------------------------------------------------------------------------------------------------------------|
| TR01  | Main language: Go.                                                                                           |
| TR02  | Local database: Encrypted relational database.                                                             |
| TR03  | Encryption using Go's native libraries (`crypto/aes`, `crypto/cipher`, `golang.org/x/crypto`).               |
| TR04  | Future support for cloud synchronization (e.g., Firebase, Supabase, or AWS).                                 |
| TR05  | Future possibility of integration with browser extensions and mobile apps.                                   |

---

## 🚀 5. Future Features (roadmap)

- Synchronization with cloud services (end-to-end encryption).  
- Browser extension for autofill.  
- Complementary mobile app (React Native or Flutter).  
- Multi-factor authentication (MFA).  
- Automatic updates and in-app notifications.  
- Fuzzy search
- Import passwords from CSV and export encrypted backup.
- Display activity logs (locally stored, optionally exportable) with timestamped events (login, vault unlock, add/edit/delete entry).
- Backup and restore vault manually with encrypted archive files. Automatic backup feature.

---

## Detailing

| Code  | Description                                                                                      |
|-------|--------------------------------------------------------------------------------------------------|
| FR01  | Using Argon2 for key derivation and secure storage.                                |
| FR02  | Each entry includes: title, username/email, password, URL (optional), details (optional), and creation/update timestamps. Only password and details stored encrypted with AES-256, so it is not needed to decrypt other fields when searching for an entry.              |           |
| FR04  | Locally encrypted vault with strong encryption (AES-256 GCM). Encryption key derived from master password (Argon2). Vault file protected against tampering.                                                   |
| FR05  | Strong password generator with the following features: Checkboxes for uppercase, lowercase and numbers. Dropdown for selecting password length (8 to 64 characters). Show password intensity. Textbox for custom special characters input. Buttons for peeking, regenerating and copying to clipboard. Presets: Default(strong) and Custom.    |
| FR06  | Search by title or any of the optional fields words using a searchbox. Dropdown category selection for filtering (independent from string search).                                                             |
| FR07  | Automatic vault lock after 15 minutes of inactivity. Requires master password again.                                              |

---
