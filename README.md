## Simple Calculator API

This calculator API supports simple math (addition, subtraction, division and multiplication) for numbers expressed in any numeral system

### How to run server
___
In order to run server
```
cd scripts && ./rundev.sh
```

also you can set your own configs in **scripts/env.sh** 

### Work example
___

Request example: 
```json
{
    "operand1": {
        "value":"35",
        "base":10
    },
    "operand2": {
        "value":"3",
        "base":10
    },
    "toBase":10
}
```

Response example for endpoint /api/add:
```json
{
    "success": true,
    "result": 38
}
```
