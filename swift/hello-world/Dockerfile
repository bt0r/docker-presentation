FROM swift:4.2

LABEL maintainer="Thibaut.B <thibaut.bayer@evaneos.com>"

RUN mkdir /var/www

WORKDIR /var/www

RUN apt-get update
RUN apt-get install uuid-dev # Required for Perfect framework

COPY PerfectTemplate .

RUN swift build

EXPOSE 8181

CMD .build/debug/PerfectTemplate

