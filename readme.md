<div style="text-align:center">
    <img src="./media/siphon.jpg">
</div>

# Siphon
Canry Testing with granular control. Test reqeuests based on REST Verb used and a % of request coming through. 

The pre generation phase will ensure that at run time siphon will not have evaluate much at run time. Making the canary proxy as light weight as possible


### How to use

The generation of the proxy is dependent on the configurations. An example config can be found in the root of the project (`siphon.config.yaml`). We use template strings to generate `.go` files based on your unique settings.

Generate the files
```
make gen
```

Build
```
make build
```

Run
```
./siphon.exe
```