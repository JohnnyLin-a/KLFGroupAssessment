CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE activities(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE user_activity(
    activity_id INT,
    user_id INT,
    occurrence timestamp,
    FOREIGN KEY(activity_id) REFERENCES activities(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);
