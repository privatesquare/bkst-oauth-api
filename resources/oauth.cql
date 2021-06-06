CREATE KEYSPACE oauth WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
CREATE TABLE oauth.access_tokens(access_token VARCHAR PRIMARY KEY,user_id BIGINT,client_id BIGINT,expires DATE);
