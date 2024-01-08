## GatewayMQTT
### Description
My first experience using Go. The project was made solely for informational purposes.
Used Go 1.19.5 (now updated to 1.20)

A gateway for transmitting data measurements from equipment/meters of the microgrid model made in Matlab Simulink.
In this project are used:
  - Gin (a little exp);
  - MQTT (without security measures for data transmission (using TLS/SSL, setting a login and password for the client)).

![Image alt](https://github.com/SorokinAS)  


### Build project
For build project you need to execution next commands:
```shell
docker build -t <image_name> .
```
```shell
docker run --name <name_of_container> <image_name>
```

Upon successful start, you can see the logs of the gin starting and the connection log to the mqtt broker. Yes, the mqtt broker is required for performance. Use docker-compose for do it.
About mosquitto you can read in https://mosquitto.org/