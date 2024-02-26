SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: notebooks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.notebooks (
    ulid character varying NOT NULL,
    alias_id character varying(15) NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: notebooks_users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.notebooks_users (
    notebook_id character varying NOT NULL,
    user_id character varying NOT NULL,
    is_admin boolean DEFAULT false
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: shouts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.shouts (
    ulid character varying NOT NULL,
    name character varying NOT NULL,
    notebook_id character varying NOT NULL,
    user_id character varying NOT NULL,
    script text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    ulid character varying NOT NULL,
    uid character varying NOT NULL,
    alias_id character varying(15) NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: notebooks notebooks_alias_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notebooks
    ADD CONSTRAINT notebooks_alias_id_key UNIQUE (alias_id);


--
-- Name: notebooks notebooks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notebooks
    ADD CONSTRAINT notebooks_pkey PRIMARY KEY (ulid);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: shouts shouts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.shouts
    ADD CONSTRAINT shouts_pkey PRIMARY KEY (ulid);


--
-- Name: users users_alias_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_alias_id_key UNIQUE (alias_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (ulid);


--
-- Name: users users_uid_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_uid_key UNIQUE (uid);


--
-- Name: idx__notebooks__alias_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx__notebooks__alias_id ON public.notebooks USING btree (alias_id);


--
-- Name: idx__users__alias_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx__users__alias_id ON public.users USING btree (alias_id);


--
-- Name: idx__users__ulid; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx__users__ulid ON public.users USING btree (ulid);


--
-- Name: notebooks_users notebooks_users_notebook_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notebooks_users
    ADD CONSTRAINT notebooks_users_notebook_id_fkey FOREIGN KEY (notebook_id) REFERENCES public.notebooks(ulid);


--
-- Name: notebooks_users notebooks_users_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notebooks_users
    ADD CONSTRAINT notebooks_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(ulid);


--
-- Name: shouts shouts_notebook_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.shouts
    ADD CONSTRAINT shouts_notebook_id_fkey FOREIGN KEY (notebook_id) REFERENCES public.notebooks(ulid);


--
-- Name: shouts shouts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.shouts
    ADD CONSTRAINT shouts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(ulid);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20240221141425'),
    ('20240225151311'),
    ('20240225151435');
