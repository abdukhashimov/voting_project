ALTER TABLE users
ADD constraint phone_validate CHECK(
        phone_number ~* '^998[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]$'
    );