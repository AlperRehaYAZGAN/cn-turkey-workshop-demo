### Cloud Native Turkey Demo - Simple Golang Application (K8S & ArgoCD)

[Live Link](https://www.youtube.com/watch?v=zjfcn2BfCFo)

This is a simple Golang application that is used for demo purposes in Cloud Native Turkey live session.  


### Build  
```bash
# build
docker buildx build --platform linux/amd64 --push -t alperreha/cn-turkey-demo:1.0.0 .
# push
docker push alperreha/cn-turkey-demo:1.0.0
```