# pow_on-_the-_knee
just simple implementation proof of work (hashcash)

server:
(under the powserver dir)

docker build -t powserver . && docker run --rm -d -p 127.0.0.1:9797:9797 --name pows powserver


client:
(under the client dir)

docker build . -t powclient && docker run --network host -it powclient

(inside client in sysout you will be able to see result of pow with amount of work)

