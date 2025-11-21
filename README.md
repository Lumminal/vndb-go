# vndb-go

vndb-go is an api wrapper for [VNDB API](https://api.vndb.org/kana).
Zero dependencies are required to run this package.

Check out the examples to see how it works.

### Status 

This project was made as part of learning backend development and golang,
therefore **it's experimental and unstable for actual use**.

Some features like UList PATCH/DELETE have not been implemented.

# Installation

Run
```go get github.com/Lumminal/vndb-go```

# Examples

Check the examples folder for how to use the library!

It includes a basic rundown for creating a client and running a query.

# ULists

Ulists have not been implemented fully, but they support POST/GET methods.
They have their own query and you have to specify the user's id before making
a query by calling:

```yourUlistQuery.SetUser(USER_ID_HERE)```

# Issues

If you find any problems with the code, please open an issue and I'm gonna try
to fix it.

# [TODO]

[] Add better tests

[] Add patching/delete for ulists

[] Properly test all features

# License

The code is licensed under GPL-3.0