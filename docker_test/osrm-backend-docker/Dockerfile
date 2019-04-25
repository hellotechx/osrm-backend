FROM debian:stretch-slim as builder

# Let the container know that there is no TTY
ENV DEBIAN_FRONTEND noninteractive

# Install necessary packages for proper system state
RUN apt-get -y update && apt-get install -y \
    build-essential \
    cmake \
    curl \
    git \
    libboost-all-dev \
    libbz2-dev \
    libtbb-dev \
    libxml2-dev \
    libzip-dev \
    lua5.1 \
    liblua5.1-0-dev \
    libluabind-dev \
    libluajit-5.1-dev \
    pkg-config

RUN mkdir -p /osrm-build \
 && mkdir -p /osrm-data \
 && mkdir -p /osrm-src
COPY . /osrm-src
WORKDIR /osrm-build

RUN cd /osrm-src \
 && echo "list files in osrm-src" \
 && ls -l \
 && mkdir -p build \
 && cd build \ 
 && cmake --version \
 && cmake .. \
 && make -j4\
 && mv /osrm-src/profiles/car.lua /osrm-build/profile.lua \
 && mv /osrm-src/profiles/lib/ /osrm-build/lib \
 && echo "disk=/tmp/stxxl,25000,syscall" > .stxxl \
 && rm -rf /osrm-src

# Cleanup --------------------------------

RUN apt-get clean \
 && rm -rf /var/lib/apt/lists/*

# Publish --------------------------------

COPY docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 5000
