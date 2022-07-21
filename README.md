# Tokoin Test - command-line tool for searching

This program is built to search specific data from provided JSON files, namely `organizations.json`, `tickets.json` and more.

## Tools and libraries used

This program is powered by:

[promptui](https://github.com/manifoldco/promptui)
[cobra-prompt](cobra-prompt)
[easyjson](jsonparser) - For fast JSON parsing from data loaded from disk

## Known issue

- The prompt line is duplicated due to race condition. There is a fix at the end but didn't get merged yet: [Link](https://github.com/manifoldco/promptui/issues/129)

## Reference links (will remove later)
