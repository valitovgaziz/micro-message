# Microservice for processing messages
## Task:
Develop a microservice in Go that will receive messages via an HTTP API, save them in PostgreSQL, and then send them to Kafka for further processing. Processed messages should be marked. The service should also provide an API for obtaining statistics on processed messages.
## Requirements:
	1.	Use Go 1.20+
	2.	Use PostgreSQL to store messages
	3.	Implement sending and reading messages to/from Kafka
	4.	Provide the ability to run the project in Docker
 ## Result requirements:
We expect that the test task will be deployed on a server and available for testing via the Internet.

Expected output:
- Link to the project deployed on the server
- Connection instructions
- Git repository with the code


TODO LIST
done | endPoint: POST message
done | endPoint: GET statistics by userId
done | save message into psql
---- | setting kafka
---- | realize sender into kafka
---- | read from kafka
---- | run into doker

set migare on start server and DB // TODO