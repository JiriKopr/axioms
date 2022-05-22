# Axioms

Code generator written in Go. Generates files and directory structure based on provided templates.

It replaces any `[ variable ]` in files, file names or directory names with provided text.

![combine_images](https://user-images.githubusercontent.com/28316931/169703456-c5c627ad-c6ff-4e8c-b6bd-2ceaf1dc9193.jpg)

## Usage

You need a directory to store all your templates. It can be anywhere in your file system, but for
convenience there is a *templates* directory present in this repository.

In your templates destination, create new directory which will represent single template.
Into this new directory put your files and directories which you want to generate.

Compile the program
```
go build -o axioms
```

### List

```
./axioms list ./templates/test
```

This command will list all necesarry tags which need to be specified for generating.

### Generate

```
./axioms generate ./templates/test ~/Desktop/target axiom_value tag=value
```

To generate files and directories, run this command specifing the template, target destination and tag values.
You can supply value without specifying the key. This value will be used for all `[ axiom ]` tags.

If your template contains other tags, you have to specify their value by providing `tag=value`.

Multiword values can be passed by separating the words with _ ie. `tag=shopping_list`

## Tags

Tags have format `[ tag_name ]`

Default tag is `[ axiom ]`

Allowed names for tags have format `[a-zA-Z0-9]{1,}`

### Modifiers

You can use modifiers with your tags. They will transform passed value into provided case. `[ axiom | pascalcase ]`

Available modifiers:
- `pascalcase` -> pascalCase
- `snakecase` -> snake\_case
- `camelcase` -> CamelCase
- `kebabcase` -> kebab-case

## Reminders

- Default tag is `[ axiom ]`
- **You can use native autocompletion of your shell to select templates and target destinations**
- Tags can be placed into file and directory names
- There are modifiers which transform provided value into desired case `[ axiom | pascalcase ]`

