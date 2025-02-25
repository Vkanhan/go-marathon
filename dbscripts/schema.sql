-- Enable necessary extension for UUIDs
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create runners table
CREATE TABLE runners (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INTEGER,
    is_active BOOLEAN DEFAULT TRUE,
    country TEXT NOT NULL,
    personal_best INTERVAL,
    season_best INTERVAL,
    CONSTRAINT runners_pk PRIMARY KEY (id)
);

-- Create results table
CREATE TABLE results (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    runner_id UUID NOT NULL,
    race_result INTERVAL NOT NULL,
    location TEXT NOT NULL,
    position INTEGER,
    year INTEGER NOT NULL,
    CONSTRAINT results_pk PRIMARY KEY (id),
    CONSTRAINT fk_results_runner_id FOREIGN KEY (runner_id)
        REFERENCES runners (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

