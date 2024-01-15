## GatewayMQTT
### Description
My first experience using Go. The project was made solely for informational purposes.
Used Go 1.19.5 (now updated to 1.20)

A gateway for transmitting data measurements from equipment/meters of the microgrid model made in Matlab Simulink.
In this project are used:
  - Gin (a little exp);
  - MQTT (without security measures for data transmission (using TLS/SSL, setting a login and password for the client)).

An example of interaction with the gateway is shown in the figure:

![Image alt](https://github.com/SorokinAS/GatewayMQTT/blob/master/docs/Diagram.png)  

### Run in locally
For run localy you need to enter the root directory project and run the project. Also you need to have installed make. For example:  

```shell
cd C:\GoProjects\GatewayMQTT
make run-local
```

### Run in Docker
For build and run project from docker-container you need to have installed make and to execute next command:
```shell
make build-container
```

For run project from docker-container you need to to execute next command:
```shell
make run-container
```

Upon successful start, you can see the logs of the gin starting and the connection log to the mqtt broker. Yes, the mqtt broker is required for performance. Use docker-compose for do it.
About mosquitto you can read in https://mosquitto.org/
