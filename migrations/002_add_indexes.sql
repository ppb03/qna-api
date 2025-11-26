-- +goose Up
CREATE INDEX idx_questions_created_at ON questions(created_at);
CREATE INDEX idx_answers_created_at ON answers(created_at);

-- +goose Down  
DROP INDEX idx_questions_created_at;
DROP INDEX idx_answers_created_at;