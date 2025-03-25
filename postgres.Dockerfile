FROM postgres:16.2

# https://docs.timescale.com/self-hosted/latest/install/installation-linux/ 

# set timezone
ENV TZ=Asia/Tehran 
RUN ls /usr/share/zoneinfo && cp /usr/share/zoneinfo/Iran /etc/localtime && echo "Asia/Tehran" >  /etc/timezone

# copy postgres config file 
COPY ./.docker/postgres/postgresql.conf /var/lib/pgsql/16/data/postgresql.conf

# copy init shell
COPY ./.docker/postgres/init.sh /docker-entrypoint-initdb.d
 
# copy SQL schema files
COPY ./sql/schemas /schemas

# copy SQL seed files
COPY ./sql/seeds /seeds
