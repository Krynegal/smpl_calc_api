Simple api calculator

You can run server using scripts/rundev.sh

Request example for endpoint /api/add: 
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

Response example:
```json
{
    "success": true,
    "value": 38
}
```
