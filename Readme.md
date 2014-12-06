# nflag

Utility around the `flag` package  to set the default value from an ENV variable, which can be overridden by setting the flag. 

## Example

    var port String
    nflag.StringVar(&port, "PORT", "bind-address", ":3000", "default server address")

Use default

    main

Use ENV

    PORT=:4000 main

Override with flag

    main --bind-address=:5000

## License

MIT
