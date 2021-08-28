#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings -dbname=bookings -dbuser=postgres -dbpass=120919 -cache=false -production=false