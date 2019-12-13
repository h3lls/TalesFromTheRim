FROM ubuntu:18.04

RUN apt-get update

# Install dev environment to compile
RUN apt-get install -y automake make gcc g++ clang libtool autoconf

ADD dockerstartmud.sh /

RUN chmod +x /dockerbuild.sh

EXPOSE 3306/TCP
EXPOSE 4000

# Build with: docker build -t tales . 
# Start with: docker run -p 4000:4000 -v G:\Projects\TalesFromTheRim:/TalesFromTheRim -it tales "/dockerbuild.sh"
# Connect to: localhost:4000