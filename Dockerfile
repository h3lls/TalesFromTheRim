FROM ubuntu:18.04

RUN apt-get update

# Install dev environment to compile
RUN apt-get install -y automake make gcc g++ clang libtool autoconf

ADD dockerbuild.sh /

RUN chmod +x /dockerbuild.sh

EXPOSE 5000

# Build with: docker build -t tales . 
# Start with: docker run -p 5000:5000 -v G:\Projects\TalesFromTheRim:/TalesFromTheRim -it tales "/dockerbuild.sh"
# Connect to: localhost:5000