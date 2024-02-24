if [ ! -f conf/conf_prod.go.bak ]; then
    exit
fi
mv conf/conf_prod.go.bak conf/conf_prod.go