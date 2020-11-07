create table meetings (
  id serial PRIMARY KEY,
  title VARCHAR ( 255 )  NOT NULL,
  participant_name VARCHAR ( 255 ) NOT NULL,
  participant_email VARCHAR ( 255 )  NOT NULL,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  timestamp TIMESTAMP  NOT NULL default NOW(),
  rsvp  VARCHAR ( 100 ) NOT NULL default 'no'
);