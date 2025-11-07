# yup-nl

```
NAME:
   nl - number lines of files

USAGE:
   nl [OPTIONS] [FILE...]

      Write each FILE to standard output, with line numbers added.
      With no FILE, or when FILE is -, read standard input.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --body-numbering value, -b value        use STYLE for numbering body lines (a=all, t=non-empty, n=none) (default: "t")
   --header-numbering value                use STYLE for numbering header lines (default: "n")
   --footer-numbering value, -f value      use STYLE for numbering footer lines (default: "n")
   --number-separator value, -s value      add STRING after (possible) line number (default: "	")
   --number-format value, -n value         insert line numbers according to FORMAT (ln, rn, rz) (default: "rn")
   --starting-line-number value, -v value  first line number for each section (default: 1)
   --line-increment value, -i value        line number increment at each line (default: 1)
   --no-renumber, -p                       do not reset line numbers at logical pages (default: false)
   --help, -h                              show help
```
