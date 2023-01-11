# Go code generation

This directory is where I put contract files for testing w3utils

Download `.abi` text files to your local directory,
then use the provided shell command template to generate Go code.

```shell
abigen \
    --abi=1inchlimitorder.abi \
    --pkg=oneinchlimitorder \
    --type=oneInchLimitOrder \
    --out=1inchlimitorder.go
```
