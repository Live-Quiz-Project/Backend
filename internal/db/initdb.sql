-- Create the "user" table
CREATE TABLE IF NOT EXISTS "user" (
  id TEXT PRIMARY KEY,
  email TEXT UNIQUE,
  "password" TEXT, 
  profile_name TEXT,
  profile_pic TEXT,
  created_date TEXT,
  account_status BOOL DEFAULT TRUE,
  suspension_start_date TEXT DEFAULT NULL,
  suspension_end_date TEXT DEFAULT NULL
);

-- Create the "quiz" table
CREATE TABLE IF NOT EXISTS quiz (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  title TEXT DEFAULT 'untitled',
  description TEXT,
  media TEXT,
  created_date TEXT,
  modified_date TEXT
);

-- Create the "question" table
CREATE TABLE IF NOT EXISTS question (
  id TEXT PRIMARY KEY,
  quiz_id TEXT NOT NULL,
  is_parent_question BOOL,
  question_id TEXT,
  type TEXT,
  "order" INT, -- Renamed "order" column
  content TEXT,
  note TEXT,
  media TEXT,
  time_limit INT,
  have_time_factor BOOL,
  time_factor INT,
  font INT,
  selected_up_to INT
);

-- Create the "option_choice" table
CREATE TABLE IF NOT EXISTS option_choice (
  id TEXT PRIMARY KEY,
  question_id TEXT NOT NULL,
  "order" INT, -- Renamed "order" column
  "content" TEXT,
  "point" REAL,
  color TEXT,
  is_correct BOOL
);

-- Create the "option_text" table
CREATE TABLE IF NOT EXISTS option_text (
  id TEXT PRIMARY KEY,
  question_id TEXT NOT NULL,
  "content" TEXT,
  "point" REAL,
  have_case_sensitive BOOL
);

-- Create the "option_matching" table
CREATE TABLE IF NOT EXISTS option_matching (
  id TEXT PRIMARY KEY,
  question_id TEXT NOT NULL,
  prompt_id TEXT,
  option_id TEXT,
  point REAL
);

-- Create the "option_matching_pool" table
CREATE TABLE IF NOT EXISTS option_matching_pool (
  id TEXT PRIMARY KEY,
  option_matching_id TEXT NOT NULL,
  "content" TEXT,
  "type" TEXT,
  "order" INT 
);

-- Create the "option_pin" table
CREATE TABLE IF NOT EXISTS option_pin (
  id TEXT PRIMARY KEY,
  question_id TEXT NOT NULL,
  x_axis INT,
  y_axis INT,
  "point" REAL
);

-- Create the "live_quiz_session" table
CREATE TABLE IF NOT EXISTS live_quiz_session (
  id TEXT PRIMARY KEY,
  quiz_id TEXT NOT NULL,
  user_id TEXT NOT NULL,
  started_at TEXT,
  ended_at TEXT,
  is_exempt TEXT[],
  status TEXT
);

-- Create the "participant" table
CREATE TABLE IF NOT EXISTS participant (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  lqs_id TEXT NOT NULL,
  name TEXT,
  point REAL,
  status TEXT
);

-- Create the "response_choice" table
CREATE TABLE IF NOT EXISTS response_choice (
  id TEXT PRIMARY KEY,
  participant_id TEXT NOT NULL,
  option_choice_id TEXT NOT NULL
);

-- Create the "response_text" table
CREATE TABLE IF NOT EXISTS response_text (
  id TEXT PRIMARY KEY,
  participant_id TEXT NOT NULL,
  option_text_id TEXT NOT NULL,
  content TEXT
);

-- Create the "response_matching" table
CREATE TABLE IF NOT EXISTS response_matching (
  id TEXT PRIMARY KEY,
  participant_id TEXT NOT NULL,
  option_matching_id TEXT NOT NULL,
  prompt_id TEXT NOT NULL,
  option_id TEXT NOT NULL
);

-- Create the "response_pin" table
CREATE TABLE IF NOT EXISTS response_pin (
  id TEXT PRIMARY KEY,
  participant_id TEXT NOT NULL,
  option_pin_id TEXT NOT NULL,
  x_axis INT NOT NULL,
  y_axis INT NOT NULL
);

-- Create the "admin" table
CREATE TABLE IF NOT EXISTS admin (
  id TEXT PRIMARY KEY,
  email TEXT UNIQUE,
  password TEXT
);

-- -- Create the foreign key constraints
-- -- Add foreign key references after all tables have been created
ALTER TABLE quiz
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES "user" (id);

ALTER TABLE live_quiz_session
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES "user" (id);

ALTER TABLE live_quiz_session
ADD CONSTRAINT fk_quiz_id FOREIGN KEY (quiz_id) REFERENCES quiz (id);

ALTER TABLE participant
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES "user" (id);

ALTER TABLE participant
ADD CONSTRAINT fk_lqs_id FOREIGN KEY (lqs_id) REFERENCES live_quiz_session (id);

ALTER TABLE question
ADD CONSTRAINT fk_quiz_id FOREIGN KEY (quiz_id) REFERENCES quiz (id);

ALTER TABLE option_choice
ADD CONSTRAINT fk_question_id FOREIGN KEY (question_id) REFERENCES question (id);

ALTER TABLE option_text
ADD CONSTRAINT fk_question_id FOREIGN KEY (question_id) REFERENCES question (id);

ALTER TABLE option_matching
ADD CONSTRAINT fk_question_id FOREIGN KEY (question_id) REFERENCES question (id);

ALTER TABLE option_matching_pool
ADD CONSTRAINT fk_option_matching_id FOREIGN KEY (option_matching_id) REFERENCES option_matching (id);

ALTER TABLE option_pin
ADD CONSTRAINT fk_question_id FOREIGN KEY (question_id) REFERENCES question (id);

ALTER TABLE response_choice
ADD CONSTRAINT fk_participant_id FOREIGN KEY (participant_id) REFERENCES participant (id);

ALTER TABLE response_text
ADD CONSTRAINT fk_participant_id FOREIGN KEY (participant_id) REFERENCES participant (id);

ALTER TABLE response_matching
ADD CONSTRAINT fk_participant_id FOREIGN KEY (participant_id) REFERENCES participant (id);

ALTER TABLE response_pin
ADD CONSTRAINT fk_participant_id FOREIGN KEY (participant_id) REFERENCES participant (id);

ALTER TABLE response_choice
ADD CONSTRAINT fk_option_choice_id FOREIGN KEY (option_choice_id) REFERENCES option_choice (id);

ALTER TABLE response_text
ADD CONSTRAINT fk_option_text_id FOREIGN KEY (option_text_id) REFERENCES option_text (id);

ALTER TABLE response_matching
ADD CONSTRAINT fk_option_matching_id FOREIGN KEY (option_matching_id) REFERENCES option_matching (id);

ALTER TABLE response_matching
ADD CONSTRAINT fk_option_matching_pool_id FOREIGN KEY (prompt_id) REFERENCES option_matching_pool (id);

ALTER TABLE response_pin
ADD CONSTRAINT fk_option_pin_id FOREIGN KEY (option_pin_id) REFERENCES option_pin (id);
