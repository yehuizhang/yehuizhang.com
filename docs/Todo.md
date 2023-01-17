# Todo

## Schema

key: `todo:<uid>`
value: a list of TodoItems or a hashmap where key is the TodoItem.ItemId and value is the actual object

## API

- POST /v1/todo
  - Add a new TodoItem
  - In DB, append new item to the list and update.
  - Drawback: we have to serialize the whole list and update the DB
  - Optimization:
    - Block user from adding new items if the list has a size of over 50
- GET /v1/todo
  - Return all TodoItems
- PUT /v1/todo/:itemId
  - Update an exiting TodoItem
  - Throw an error if TodoItem was not found
- DELETE /v1/todo/:itemId
  - Delete an exiting TodoItem
  - Throw an error if TodoItem was not found
