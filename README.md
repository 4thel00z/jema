# jema

![jema-tests](https://github.com/4thel00z/jema/workflows/Test/badge.svg)
![jema-logo](https://github.com/4thel00z/jema/raw/assets/logo.svg)

jema is a cli application that validates json data using a json schema and prints the errors on the console.

## Example usage

This demonstrates how you can validate a json document. 

```
cat sample.json | jema -schema schema.json
```

## Todo

- Add reading from stdin
- Add loading the json schema
- Return non zero exit code on validation issue
- Print out on stderr what went wrong

## License

This project is licensed under the GPL-3 license.
