# go-reloaded

> A text completion, editing, and auto-correction tool written in Go — processes an input file and writes the cleaned, transformed result to an output file.

---

## What Is This Project?

**go-reloaded** is a command-line text processing tool that applies a set of transformation rules to a text file. Think of it as a lightweight auto-formatter: it converts number bases, fixes punctuation spacing, handles capitalization, and corrects grammar — all in one pass.

---

## Usage

```bash
go run . input.txt output.txt
```

The program reads from `input.txt`, applies all transformations, and writes the result to `output.txt`.

---

## Transformations

### Number Base Conversion

| Tag | Behavior | Example |
|-----|----------|---------|
| `(hex)` | Converts preceding word from hexadecimal to decimal | `1E (hex)` → `30` |
| `(bin)` | Converts preceding word from binary to decimal | `10 (bin)` → `2` |

### Case Modifiers

| Tag | Behavior | Example |
|-----|----------|---------|
| `(up)` | Uppercases the preceding word | `go (up)` → `GO` |
| `(low)` | Lowercases the preceding word | `SHOUTING (low)` → `shouting` |
| `(cap)` | Capitalizes the preceding word | `bridge (cap)` → `Bridge` |
| `(up, N)` | Uppercases the previous N words | `so excited (up, 2)` → `SO EXCITED` |
| `(low, N)` | Lowercases the previous N words | `IT WAS THE (low, 3)` → `it was the` |
| `(cap, N)` | Capitalizes the previous N words | `age of foolishness (cap, 3)` → `Age Of Foolishness` |

### Punctuation Formatting

- Punctuation marks (`. , ! ? : ;`) are placed **directly after** the preceding word with a space before the next word
- Groups like `...` or `!?` are kept together and treated as one unit
- Single quotes `'` wrap their content with **no inner spaces**

```
Before: "I was sitting over there ,and then BAMM !!"
After:  "I was sitting over there, and then BAMM!!"

Before: "' awesome '"
After:  "'awesome'"

Before: "I was thinking ... You were right"
After:  "I was thinking... You were right"
```

### Grammar Correction

- `a` is replaced with `an` when the next word starts with a vowel (`a, e, i, o, u`) or `h`

```
Before: "A amazing rock"
After:  "An amazing rock"
```

---

## Full Example

```bash
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of foolishness (cap, 6) , IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt
It was the best of times, it was the worst of TIMES, It Was The Age Of Foolishness, it was the winter of despair.
```

---

## Project Structure

```
go-reloaded/
├── main.go
├── transformations.go   # All transformation logic
├── transformations_test.go
└── sample.txt
```

---

## Concepts Practiced

- Go file system (`fs`) API
- String and number manipulation
- Regular expressions / text parsing
- Unit testing in Go

---

*Project by Bennacer Douirat — Zone01 Oujda*
