#!/bin/bash

# migrate database 
for file in /schemas/*; do
    psql -v ON_ERROR_STOP=1 --username $POSTGRES_USER --dbname "$POSTGRES_DB" -a -f "$file"
done

# seed database 
for file in /seeds/*; do
    psql -v ON_ERROR_STOP=1 --username $POSTGRES_USER --dbname "$POSTGRES_DB" -a -f "$file"
done
