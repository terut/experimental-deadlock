CREATE TABLE test (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    uniquekey VARCHAR(20) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (uniquekey)
);
