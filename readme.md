# booky cloud native application

***application created for learning purpose in order to learn about cloud native***

## api documentation

### Book api

| Endpoint                | Method | Description                                       | body | response |
|-------------------------|--------|---------------------------------------------------| ---- | -------- |
| `/book/{id}`            | GET    | Retrieve information about a book by its ID.     | `-`   | `Book{id, isbf, title, author_id}`
| `/book`                 | POST   | Add a new book to the system.                   | `Book{isbf, title, author_id}` | `string message` or `error`
| `/book`                 | PUT    | Update an existing book's information.          | `Book{isbf, title, author_id}` |  `Book{id, isbf, title, author_id}` |
| `/book/{id}`            | DELETE | Delete a book by its ID.                        | `-` | `string message` or `error`

### Author api

| Endpoint                | Method | Description                                       | body | response |
|-------------------------|--------|---------------------------------------------------| ---- | -------- |
| `/author/{id}`            | GET    | Retrieve information about a author by its ID.     | `-`   | `Author{id, isbf, title, author_id}`
| `/author`                 | POST   | Add a new author to the system.                   | `Author{isbf, title, author_id}` | `string message` or `error`
| `/author`                 | PUT    | Update an existing author's information.          | `Author{isbf, title, author_id}` |  `Author{id, isbf, title, author_id}` |
| `/author/{id}`            | DELETE | Delete a author by its ID.                        | `-` | `string message` or `error`

### controllers & models test

you can check both `book_test.go` or `author_test.go` for testing or adding more test to the packages

### logging 

for logging system *Booky* uses **Zap** which a fast logger built for go.every logs are stored in ```/logs/api.log``` if you want to change the logging configurations please check ```./utils/log.go```