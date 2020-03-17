# Project
Project is written using go to see how to create an api using go

This is an API designed to store food items.  Storing data including
- ID
- Name
- Description
- Cost
- Calories
- Protein
- Vegan?

#How to run
run `go build && ./food-item-api` from food-item-api directory

Set up db by
 1. `brew install mongodb`
 2. `mongod` <- starts daemon
 3. `mongo` <- starts mongo shell
 4. `use food_item_db` <- creates db (run from terminal window running mongo shell)
 5. `quit()` <- quits shell
 6. `ctrl-c` < quits daemon
