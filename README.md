# go-inventory-management

This is a go lang project showcasing database integration and clean architecture layering. It follows a professional folder structure.

 It implemens a GetTasks method that fetches data from a Postgres database (developmented in week-4: https://github.com/abhirup7477/inventory-management-sqlalchemy-alembic-postgresql). 

It also send a notification email, to the user on successful retreival of the record from the database, that runs in background and does not block the http request. The mail is sent even if user disconnects eraly. SMTP protocol has been used for email sending.