## LoadBalancer With GoLang
---
> Serving remote server traffic as a loadbalancer reverse proxy using golang net http package


#### Experiment:
`The remote server here will be API server built using python3 FLASK module, so make sure flask and python3 installed`

##### 1. Spin test server
*
```sh
./script.sh
```

 - The above script.sh will execute 3 API server, served at 5000, 6001, 8000
 - WARN: this runs a background process. if incase to stop all there , ps aux and kill PID
 
##### 2. Spin LoadBalancer
*
```sh
go run main.go
```

> VISIT : http://127.0.0.1:9091/

`you will see traffic loadbalanced from different server:port (Reload and see !!)`
