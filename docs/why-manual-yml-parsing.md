# Why Manual YAML Parsing
***
Manually iterating through a YAML file in Go instead of using a library gave us greater control over parsing and transformation processes, while also leaving key aspects of .yml file intact (like comments and indentations).

## Why Not Use Existing YAML Libraries?
### go-yaml
- Provides only a **Node Graph representation**, without access to an **Event Tree**.
- Doesn't properly preserve comments.
- Basically deprecated, hasn't been updated in years.

### go-yaml-edit
- Designed for **in-place modification** rather than structured reading.
- **Deprecated**, unreliable for consistent YAML handling.

### goccy/go-yaml
- Is being regularly updated and maintained thankfully.
- Sadly **Deletes comments**, which are crucial for preserving context.

***

## Flaws with Manual Parsing
- **Time Consuming**: 
  - Manually parsing YAML files is a time-consuming process, especially for large files.
  - It requires a lot of manual effort to parse and transform the YAML file properly without any erros.
- **Error Prone**:
  - Manual parsing is prone to errors, especially when dealing with complex YAML structures. 
  - We wouldn't have these problems with YML libraries, but it's crucial for us to have comments saved, because GitHub Actions Importer leaves comments for whatever it couldn't convert properly.

For the sake of our project and the need to preserve comments and indentation, we decided to go with manual parsing. I am completely aware of it's flaws, but none of the libraries I tried were able to do what I was asking for. 
