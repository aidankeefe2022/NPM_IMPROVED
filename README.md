# NPM_IMPROVED

Kinda a misnomer as this is a wrapper that allows for the npm audit command to be recursive. 
There is also the added feature of if the tool comes accross a package.json file and no package-lock.json file it will make one and audit it!


## Requirements

- latest version of go should do it... its a very different install process for all machines so please look it up!
- Thats it!


## Install 
### Once go is installed properly and a GOPATH is set you can install the tool super easily!
```bash 

go install github.com/aidankeefe2022/NPM_IMPROVED
```
### Thats it!!

### Try and run this 

```bash 

NPM_IMPROVED -h
```