## Simple Calculator API

This calculator API supports simple math (addition, subtraction, division and multiplication) for numbers expressed in any numeral system

### How to run server
___
In order to run server using docker
```
docker build calc-api . 
docker run -d -p <host_port>:8080 --name smpl-calc-api calc-api
```
or you can do it using script
```
cd scripts && ./rundev.sh
```

also you can set your own configs in **scripts/env.sh** 

### Available endpoints

- /api/add - addition
- /api/sub - subtraction
- /api/div - division
- /api/mul - multiplication

### Request format

```
{
    "operand1": {
        "value":"35",
        // the number system of the first operand
        "base":10
    },
    "operand2": {
        "value":"3",
        // the number system of the second operand
        "base":10
    },
    // the number system in which you want to get the answer
    "toBase":10
}
```

### Work example
___

Request example: 
```json
{
  "operand1": {
    "value":"3F",
    "base":16
  },
  "operand2": {
    "value":"303",
    "base":4
  },
  "toBase":5
}
```

Response example for endpoint /api/mul:
```json
{
    "success": true,
    "result": "100323"
}
```
