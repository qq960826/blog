FROM ubuntu:18.04
MAINTAINER qq960826  "qq960826@gmail.com"
ADD public /public
WORKDIR /public
ENTRYPOINT ["./server"]

