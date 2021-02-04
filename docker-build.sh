if [[ "$1" = "NO-CACHE" ]]
then
   docker build --no-cache --tag atlas-eso:latest .
else
   docker build --tag atlas-eso:latest .
fi
