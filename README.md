## About

World of Wisdom DDOS Protected TCP Server with Proof of Work by hashcash

## The Task

Design and implement “Word of Wisdom” tcp server.
- TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge

## Why hashcash?

Hashcash is used for Proof of Work (PoW) because it requires a computationally expensive but easily verifiable process. This makes it effective for limiting spam, DoS attacks, and ensuring fairness in systems like Bitcoin. The key features are:

- Simple to Verify: The solution to a Hashcash challenge is easily checked by recalculating the hash.
- Costly to Compute: Generating a valid hash requires significant computational effort, deterring abuse.
- Adjustable Difficulty: The required number of leading zeros in the hash can be tuned to control the difficulty.

These properties make Hashcash ideal for establishing trust in decentralized systems.

## Build and Test

```shell
make buold-docker-client
make buold-docker-server
```

## Run

```shell
docker run --network=host wow-server -addr=":8089"
docker run --network=host wow-client -addr=":8089"
```