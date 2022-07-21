# Tokoin Test - command-line tool for searching

This program is built to search specific data from provided JSON files, namely `organizations.json`, `tickets.json`, `users.json`.

## Motivation

The requirement is quite... interesting (a little weird, to be honest).

In the sample program, users must go through several steps with minimal instruction. So I used [promptui](https://github.com/manifoldco/promptui) to make this CLI looks more like a wizard to help them navigate.

Initially, there were two ways that I can think about to implement this idea:

### As a hybrid CLI-webservice

This is the ideal way because: We (the developer) can manage our data remotely (just change the record on the database if needed). In fact, this challenge is implemented in several places like this [repository](https://github.com/shubham-pathak22/zendesk-coding-challenge).

Generally, the way to go is:

- Creating a database (locally or just boot a free [Postgres instance at Heroku](https://www.heroku.com/postgres))
- Write the schema or a seed SQL script, run it (with proper constraints, because the data in 3 json files have relationships).
- With above things prepared, import the data into database (using a json parser, then feed to database through scripting).
- Create APIs to search. Make sure the query input matches requirement. Returned result also must include relevant (constrained) info as specified in the PDF.
- On the CLI side, use promptui to make a wizard to collect the input data. Then use that data to search by querying created APIs.

### As an offline CLI app

This is the least favored one because we have to:

- CLI side: Prepare the wizard to collect input data.
- Read data from JSON files: This is the bottleneck if number of records becomes big (10000+ ?).
- Unmarshal (deserialize) items in read data.
- Search using provided input and get indexes of matched records.
- Display those records.

---

## implementation and usage

Due to various personal issues, I chose the second way but couldn't complete it nicely.
You can download a sample (incomplete) release in the "Release" section to have a look at it.

- Run the program (`tokoin-test.exe`)
- Type `search` (or just `s` and press Tab)
- Follow the instruction.

## Tools and libraries used

This program is powered by:

- [promptui](https://github.com/manifoldco/promptui)
- [cobra-prompt](cobra-prompt)
- [easyjson](jsonparser) - For fast JSON parsing from data loaded from disk

## Known issues

- The prompt line is duplicated due to race condition. There is a fix at the end but didn't get merged yet: [Link](https://github.com/manifoldco/promptui/issues/129)
- The only field that can be searched for now is `_id`, more fields will be added later.
