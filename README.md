# Astrav
[![GolangCI](https://golangci.com/badges/github.com/squzy/golangci-lint.svg)](https://golangci.com)

## Summary
An AST traversal library to check Go code structure. It wraps Go's ast library and provides a parent - child structure and
a lot of convenience functions like searching nodes by name, or ast type within another node. Also searching in the 
call tree is supported which follows function calls.
