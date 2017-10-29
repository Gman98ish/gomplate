---
title: math functions
menu:
  main:
    parent: functions
---

A set of basic math functions to be able to perform simple arithmetic operations with `gomplate`.

## `math.Add`

**Alias:** `add`

Adds a set of numbers together

### Usage
```go
math.Add n...
```
```go
x | math.Add n...
```

### Example

```console
$ gomplate -i '{{ math.Add 1 2 3 4 }}
10
```

## `math.Sub`

**Alias:** `sub`

Subtracts

