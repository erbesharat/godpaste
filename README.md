# godpaste

A command line tool for creating items on [dpaste](https://dpaste.com).

## Installation
```bash
go get github.com/erbesharat/godpaste
```

## Usage
Basic plain-text item creation:
```bash
godpaste --file test.txt
```
Set expiry (1–365 days):
```bash
godpaste --file test.txt --expire 5
```
Set syntax to JavaScript:
```bash
godpaste --file test.js --syntax js
```


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License
> GPLv3
