CREATE TABLE user(
    id SERIAL,
    name TEXT
);

CREATE TABLE activity(
    id SERIAL,
    name TEXT
);

CREATE TABLE user_activity(
    activity_id INT,
    user_id INT,
    occurrence timestamp,
    FOREIGN KEY(activity_id) REFERENCES activity(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
);
