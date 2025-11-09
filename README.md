# gormery

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-0.4.1-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

[document](https://pkg.go.dev/github.com/myyrakle/gormery)

Boilerplate generator for gorm

## install

```
go install github.com/myyrakle/gormery@v0.4.1
```

## Confiuration

The `.gormery.yaml` file must exist in the project root path.

Here is an example of a config file.

```
basedir: example
output-suffix: "_gorm.go"
runner-path: "cmd/gormery"
features:
  - SLICE
```

It means that all files in the example directory will be read, and the output file will be created with the name "\*\_gorm.go".

## How to use?

Usage is very simple. Just run the following command in your project root path:

```
gormery
```

gormery only generates structures with `// @Gorm` comments. t reads structures and fields and produces a list of methods and constants.

If you have a struct like

```
// @Gorm
type Person struct {
	ID   string
	Name string
}
```

mongery produces a list of constants like this:

```
func (t Person) TableName() string {
	return "people"
}
func (t Person) StructName() string {
	return "Person"
}
const Person_ID = "id"
const Person_Name = "name"
func (t Person) Columns() []string {
	return []string{
		Person_ID,
		Person_Name,
	}
}
```

## Features - Slice

```yaml
features:
  - SLICE
```

If you enable Slice among the features flags, it creates a basic boilerplate for Slice.

It is as follows:

```go
type PersonSoManies []PersonSoMany

func (t PersonSoManies) Len() int {
	return len(t)
}

func (t PersonSoManies) IsEmpty() bool {
	return len(t) == 0
}

func (t PersonSoManies) First() PersonSoMany {
	if t.IsEmpty() {
		return PersonSoMany{}
	}
	return t[0]
}
```
