### Create DB User
```sql
    CREATE USER '<user>'@'localhost' IDENTIFIED BY '<password>';
    GRANT SELECT, INSERT, UPDATE ON payments.* TO '<user>'@'localhost';
    FLUSH PRIVILEGES;
```