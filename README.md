### golocate

A custom CLI to search for files written in Go


This is a simple implementation inspired in the Unix command line tool `locate`. Currently, the CLI has three functionalities: basic search, regex search, and contains search.

The basic search matches exact names and wild cards. For example:

```bash
#matches only foo.pdf, if exists
go run locate.go foo.pdf
#matches all pdf in your PC
go run locate.go *pdf
```

The regex search can be access using the flag `-r`:

```bash
#matches all pdfs that starts with the word work, capitalized or not
go run locate.go '(W|w)ork.*.pdf' -r
```

contains search returns the files that contains the desired word:

```bash
#matches files that contains the word "capital"
go run locate.go 'capital' -c
```