DROP INDEX IF EXISTS demo_user_username_idx;
DROP INDEX IF EXISTS demo_user_status_idx;
DROP INDEX IF EXISTS demo_user_password_idx;
ALTER TABLE IF EXISTS demo_user ALTER COLUMN tid DROP DEFAULT;
ALTER TABLE IF EXISTS demo_article ALTER COLUMN tid DROP DEFAULT;
DROP SEQUENCE IF EXISTS demo_user_tid_seq;
DROP TABLE IF EXISTS demo_user;
DROP SEQUENCE IF EXISTS demo_article_tid_seq;
DROP TABLE IF EXISTS demo_article;
