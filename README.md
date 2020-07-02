# KLFGroupAssignment

To re-style this file later.

languages and tools used: 
Core:
docker latest stable
docker-compose (installed manually only if using linux) latest stable

Backend (using docker):
postgres 12.x stable
golang 1.14.3 latest stable

Frontend (using docker):
yarn 1.x latest stable
node v12.x lts latest stable
react latest stable

libraries golang:
github.com/joho/godotenv access .env files in go app
github.com/lib/pq add postgres dialect for go app
golang.org/x/crypto add cryto library primarily for salt & hash passwords stored in database.
github.com/gorilla/mux add HTTP router and URL matcher for go web app.


libraries for react:
react-redux state management for react
redux-persist add-on to persist state after reload for react-redux

Core Notes:
The frontend and backend normally should be in different repositories.
I will add both of them here for simplicity.
This whole project combines all 3 parts of the assignment.

Normally *.env files are excluded and *.env.template files are provided with a list of used variable names.
For simplicity's sake, I will add the .env so that it can be executed quickly without much set up.
Feel free to change any of the values in the *.env files.


Backend notes:
I will not be using an ORM because I feel that will not show my SQL knowledge.
Which is why I will only be using raw SQL queries when performing transactions with the database.
Otherwise, I would be using a library called gorm for golang ORM.

The default postgres user and database should not be used and should create a new user and database for the app.
For simplicity's sake I will be using the default ones.

Frontend notes:
-


Running instructions:
*Make sure your ports 3000, 5432, 8080 are not being occupied locally.
*Make sure your computer cpu architecture is amd64 (64-bit intel/amd). (With a little bit of change, this could work on other archs too)
*My computer is Windows 10 using WSL2 for running docker, so this should work on linux as well.
Navigate to the project root directory.
Navigate inside the backend folder and execute the start.sh script.
Execute the dbsetup.sh once the postgres docker is started.
Navigate back to the project root directory, then navigate inside the frontend directory.
Execute the build.sh script once only, then execute the start.sh script.

Answers:
Then you can head to localhost:3000 for the company home page (Assignment 2).
Logging in is on the same webpage (Assignment 1).

Assignment 3 answer is here:

	SELECT users.name AS user_name, activities.name AS activity_name, x.amount, x.first_occurrence, x.last_occurrence
	FROM (
	SELECT user_id, activity_id, COUNT(*) AS amount, MIN(occurrence) AS first_occurrence, MAX(occurrence) AS last_occurrence
		FROM user_activities
		WHERE occurrence BETWEEN '2019-10-01 00:00:00' AND '2019-10-31 23:59:59'
		GROUP BY activity_id, user_id
	) AS x
	INNER JOIN users ON x.user_id = users.id
	INNER JOIN activities ON x.activity_id = activities.id
	ORDER BY activity_name

