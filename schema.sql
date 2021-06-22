DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    published_at BIGINT NOT NULL
);

INSERT INTO authors (id, name) VALUES (0, 'Дмитрий');
INSERT INTO posts(
	author_id,
	title,
	content,
	created_at,
	published_at
) VALUES
(
	0,
	'Effective Go',
	'Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, its important to understand its properties and idioms. Its also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.',
	0,
	0
);
INSERT INTO posts(
	author_id,
	title,
	content,
	created_at,
	published_at
) VALUES
(
	0,
	'The Go Memory Model',
	'The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.',
	0,
	0
);