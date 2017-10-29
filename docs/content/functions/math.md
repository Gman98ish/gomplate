---
title: math functions
menu:
  main:
    parent: functions
---

A set of basic math functions to be able to perform simple arithmetic operations with `gomplate`.

_**Note:** currently, `gomplate` supports only integer arithmetic. All functions
return 64-bit integers (`int64` type). Floating point support will be added in
later releases._

## `math.Add`

**Alias:** `add`

Adds all given operators.

### Usage
```go
math.Add n...
```
```go
x | math.Add.Add n...
```

### Example

```console
$ gomplate -i '{{ math.Add 1 2 3 4 }}
10
```

## `math.Sub`

**Alias:** `sub`

Subtract the second from the first of the given operators.

### Usage
```go
math.Sub a b
```
```go
b | math.Sub a
```

### Example

```console
$ gomplate -i '{{ math.Sub 3 1 }}'
2
```

## `math.Mul`

**Alias:** `mul`

Multiply all given operators together.

### Usage
```go
math.Mul n...
```
```go
x | math.Mul n...
```

### Example

```console
$ gomplate -i '{{ math.Mul 8 8 2 }}'
128
```

## `math.Div`

**Alias:** `div`

Divide the first number by the second.

### Usage
```go
math.Div a b
```
```go
b | math.Div a
```

### Example

```console
$ gomplate -i '{{ math.Div 8 2 }}'
4
```


## `math.Mod`

**Alias:** `mod`

### Usage
```go
math.Mod n...
```
```go
x | math.Mod n...
```

### Example

```console
$ gomplate -i '{{  }}'
```


## `math.Pow`

**Alias:** `pow`

### Usage
```go
math.Pow n...
```
```go
x | math.Pow n...
```

### Example

```console
$ gomplate -i '{{  }}'
```

