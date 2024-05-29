# toll-calculator

// microservices needed to run in the background
```
docker-compute up -d
```

// means simulate obus -> 10 obus sending real time location every (timeInteval : default is 1 second)
```
make obu
```

// this microservice receives from obu and puts it into a kafka pipeline "obudata"
```
make receiver
```

// the main functionality of this microservice is to calculate geographical distance from the beginning of the streaming to the end
```
make calculator
```

made by `admiral simo`
