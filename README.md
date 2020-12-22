![](https://i.imgur.com/ONTIffX.png)

[![Build Status](https://cloud.drone.io/api/badges/mi-24v/parakeet/status.svg)](https://cloud.drone.io/mi-24v/parakeet)
[![Go Report Card](https://goreportcard.com/badge/github.com/mi-24v/parakeet)](https://goreportcard.com/report/github.com/mi-24v/parakeet)
[![codebeat badge](https://codebeat.co/badges/8817e250-699a-46ad-ad78-d77d4e88545f)](https://codebeat.co/projects/github-com-mi-24v-parakeet-master)
[![GitHub](https://img.shields.io/github/license/mohemohe/parakeet.svg)](https://github.com/mohemohe/parakeet/blob/master/LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmi-24v%2Fparakeet.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmi-24v%2Fparakeet?ref=badge_shield)
[![Docker Image Size (tag)](https://img.shields.io/docker/image-size/mi-24v/parakeet/latest)](https://hub.docker.com/r/miwpayou0808/parakeet)

Fast weblog built in golang and top of echo.  
Supports React SSR and hydrate.

## Require

- [MongoDB](https://www.mongodb.com)

## Quickstart

```
wget https://raw.githubusercontent.com/mohemohe/parakeet/master/docker-compose.yml
docker-compose up -d
```

You can see parakeet on [0.0.0.0:1323](http://127.0.0.1:1323).  
Admin page is [/admin](http://127.0.0.1:1323/admin) and default root password is 'root'.

## Production usage

1. Create MongoDB via [Atlas](https://cloud.mongodb.com)
2. Deploy [mohemohe/parakeet](https://hub.docker.com/r/mohemohe/parakeet) in any Docker (e.g. swarm, k8s, Amazon ECS ...) and set below environment variables

### Environment variables

| key           | value                                                                                                                     |
| :------------ | :------------------------------------------------------------------------------------------------------------------------ |
| ECHO_ENV      | production                                                                                                                |
| MONGO_ADDRESS | mongodb://{user}:{password}@{replset-1},{replset-2},{replset-3}/{database}?ssl=true&replicaSet={replset}&authSource=admin |
| MONGO_SSL     | true                                                                                                                      |
| ROOT_PASSWORD | {initial root password}                                                                                                   |
| SIGN_SECRET   | {jwt secret}                                                                                                              |

## License

Icons made by [Zlatko Najdenovski](https://www.flaticon.com/authors/zlatko-najdenovski) from [www.flaticon.com](https://www.flaticon.com/)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmohemohe%2Fparakeet.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmohemohe%2Fparakeet?ref=badge_large)
