java -jar $OPENAPI_GEN_HOME/openapi-generator-cli-4.3.1.jar  generate -g go-server -i ./openapi.yaml /
  -o ../gen -c config.json
