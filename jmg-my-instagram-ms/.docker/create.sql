CREATE TABLE hashtags (
                          hashtag_id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE locations (
                           location_id SERIAL PRIMARY KEY,
                           country VARCHAR(255) NOT NULL,
                           city VARCHAR(255) NOT NULL,
                           info_aditional VARCHAR(255)
);

CREATE TABLE multimedia_contents (
                                     multimedia_content_id SERIAL PRIMARY KEY,
                                     content_url VARCHAR(255) NOT NULL,
                                     content_type VARCHAR(255) NOT NULL CHECK (content_type IN ('IMAGE', 'VIDEO'))
);

CREATE TABLE user_instas (
                             user_id SERIAL PRIMARY KEY,
                             username VARCHAR(255) NOT NULL UNIQUE,
                             email VARCHAR(255) NOT NULL UNIQUE,
                             profile_multimedia_content_id INTEGER REFERENCES multimedia_contents(multimedia_content_id),
                             description VARCHAR(255)
);

CREATE TABLE post (
                       post_id SERIAL PRIMARY KEY,
                       user_id INTEGER NOT NULL REFERENCES user_instas(user_id),
                       title VARCHAR(255),
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       post_type VARCHAR(255) NOT NULL CHECK (post_type IN ('POST', 'REEL'))
);

CREATE TABLE post_content (
                               post_content_id SERIAL PRIMARY KEY,
                               post_id INTEGER NOT NULL REFERENCES post(post_id),
                               multimedia_content_id INTEGER NOT NULL REFERENCES multimedia_contents(multimedia_content_id),
                               location_id INTEGER REFERENCES locations(location_id),
                               description VARCHAR(255),
                               orden INTEGER,
                               enabled BOOLEAN DEFAULT TRUE
);

CREATE TABLE post_content_hashtags (
                                       post_content_hashtag_id SERIAL PRIMARY KEY,
                                       post_content_id INTEGER NOT NULL REFERENCES post_content(post_content_id),
                                       hashtag_id INTEGER NOT NULL REFERENCES hashtags(hashtag_id)
);

CREATE TABLE post_content_users (
                                    post_content_users_id SERIAL PRIMARY KEY,
                                    post_content_id INTEGER NOT NULL REFERENCES post_content(post_content_id),
                                    user_id INTEGER NOT NULL REFERENCES user_instas(user_id)
);

CREATE TABLE stories (
                         story_id SERIAL PRIMARY KEY,
                         title VARCHAR(255),
                         multimedia_content_id INTEGER NOT NULL REFERENCES multimedia_contents(multimedia_content_id),
                         orden INTEGER NOT NULL
);

CREATE TABLE story_contents (
                                story_content_id SERIAL PRIMARY KEY,
                                story_id INTEGER NOT NULL REFERENCES stories(story_id),
                                multimedia_content_id INTEGER NOT NULL REFERENCES multimedia_contents(multimedia_content_id),
                                location_id INTEGER REFERENCES locations(location_id),
                                description VARCHAR(255),
                                orden INTEGER,
                                enabled BOOLEAN DEFAULT TRUE
);

CREATE TABLE story_content_hashtags (
                                        story_content_hashtag_id SERIAL PRIMARY KEY,
                                        story_content_id INTEGER NOT NULL REFERENCES story_contents(story_content_id),
                                        hashtag_id INTEGER NOT NULL REFERENCES hashtags(hashtag_id)
);

CREATE TABLE story_content_users (
                                     story_content_users_id SERIAL PRIMARY KEY,
                                     story_content_id INTEGER NOT NULL REFERENCES story_contents(story_content_id),
                                     user_id INTEGER NOT NULL REFERENCES user_instas(user_id)
);
