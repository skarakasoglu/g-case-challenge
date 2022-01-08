#!/bin/bash

$APP_MODE_PROD=PROD
$APP_MODE_TEST=TEST

if [[ $APP_MODE == $APP_MODE_PROD ]]; then
  make build && make run
else
  make test
fi