package demoupgrade

var LATEST = `
--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.3 (Debian 13.3-1.pgdg100+1)




--
-- Name: demo_article; Type: TABLE; Schema: public;
--

CREATE TABLE demo_article (
    tid bigint NOT NULL,
    user_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    create_time timestamp with time zone NOT NULL,
    update_time timestamp with time zone NOT NULL,
    status integer NOT NULL
);


--
-- Name: COLUMN demo_article.status; Type: COMMENT; Schema: public;
--

COMMENT ON COLUMN demo_article.status IS '100 is normal, -1 delete';


--
-- Name: demo_article_tid_seq; Type: SEQUENCE; Schema: public;
--

CREATE SEQUENCE demo_article_tid_seq
    START WITH 1000
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: demo_article_tid_seq; Type: SEQUENCE OWNED BY; Schema: public;
--

ALTER SEQUENCE demo_article_tid_seq OWNED BY demo_article.tid;


--
-- Name: demo_user; Type: TABLE; Schema: public;
--

CREATE TABLE demo_user (
    tid bigint NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    create_time timestamp with time zone NOT NULL,
    update_time timestamp with time zone NOT NULL,
    status integer NOT NULL
);


--
-- Name: COLUMN demo_user.tid; Type: COMMENT; Schema: public;
--

COMMENT ON COLUMN demo_user.tid IS 'the primary key';


--
-- Name: COLUMN demo_user.username; Type: COMMENT; Schema: public;
--

COMMENT ON COLUMN demo_user.username IS 'the login username';


--
-- Name: COLUMN demo_user.password; Type: COMMENT; Schema: public;
--

COMMENT ON COLUMN demo_user.password IS 'the login password';


--
-- Name: COLUMN demo_user.create_time; Type: COMMENT; Schema: public;
--

COMMENT ON COLUMN demo_user.create_time IS 'the user create time';


--
-- Name: COLUMN demo_user.update_time; Type: COMMENT; Schema: public;
--

COMMENT ON COLUMN demo_user.update_time IS 'the user last update time';


--
-- Name: COLUMN demo_user.status; Type: COMMENT; Schema: public;
--

COMMENT ON COLUMN demo_user.status IS 'the user status, 100 is normal, 200 disabled, -1 is deleted';


--
-- Name: demo_user_tid_seq; Type: SEQUENCE; Schema: public;
--

CREATE SEQUENCE demo_user_tid_seq
    START WITH 1000
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: demo_user_tid_seq; Type: SEQUENCE OWNED BY; Schema: public;
--

ALTER SEQUENCE demo_user_tid_seq OWNED BY demo_user.tid;


--
-- Name: demo_article tid; Type: DEFAULT; Schema: public;
--

ALTER TABLE IF EXISTS ONLY demo_article ALTER COLUMN tid SET DEFAULT nextval('demo_article_tid_seq'::regclass);


--
-- Name: demo_user tid; Type: DEFAULT; Schema: public;
--

ALTER TABLE IF EXISTS ONLY demo_user ALTER COLUMN tid SET DEFAULT nextval('demo_user_tid_seq'::regclass);


--
-- Name: demo_article demo_article_pkey; Type: CONSTRAINT; Schema: public;
--

ALTER TABLE IF EXISTS ONLY demo_article
    ADD CONSTRAINT demo_article_pkey PRIMARY KEY (tid);


--
-- Name: demo_user demo_user_pkey; Type: CONSTRAINT; Schema: public;
--

ALTER TABLE IF EXISTS ONLY demo_user
    ADD CONSTRAINT demo_user_pkey PRIMARY KEY (tid);


--
-- Name: demo_user_password_idx; Type: INDEX; Schema: public;
--

CREATE INDEX demo_user_password_idx ON demo_user USING btree (password);


--
-- Name: demo_user_status_idx; Type: INDEX; Schema: public;
--

CREATE INDEX demo_user_status_idx ON demo_user USING btree (status);


--
-- Name: demo_user_username_idx; Type: INDEX; Schema: public;
--

CREATE UNIQUE INDEX demo_user_username_idx ON demo_user USING btree (username);


--
-- PostgreSQL database dump complete
--

` + INIT

var DROP = `
DROP INDEX IF EXISTS demo_user_username_idx;
DROP INDEX IF EXISTS demo_user_status_idx;
DROP INDEX IF EXISTS demo_user_password_idx;
ALTER TABLE IF EXISTS demo_user ALTER COLUMN tid DROP DEFAULT;
ALTER TABLE IF EXISTS demo_article ALTER COLUMN tid DROP DEFAULT;
DROP SEQUENCE IF EXISTS demo_user_tid_seq;
DROP TABLE IF EXISTS demo_user;
DROP SEQUENCE IF EXISTS demo_article_tid_seq;
DROP TABLE IF EXISTS demo_article;
`

const CLEAR = `
DELETE FROM demo_user;
DELETE FROM demo_article;
`
