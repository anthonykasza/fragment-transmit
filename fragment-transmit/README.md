
1. Put some information into `data-file`

2. Start a listener on localhost
```
socat tcp-listen:9090,fork,reuseaddr OPEN:./assembled,creat,append
```

3. Run the script
```
go run main.go
```

4. Compare `data-file` and `assembled`
