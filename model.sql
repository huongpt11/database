CREATE TABLE movies (
    id SERIAL,
    movieID varchar(50) NOT NULL,
    movieName varchar(50) NOT NULL,
    PRIMARY KEY (id)
)
INSERT INTO movies (
    movieID,
    movieName
)
VALUES
    ('1', 'movie1'),
    ('2 ', 'movie2'),
    ('3', 'movie3'),
    ('4', 'movie4');
