version=$(git rev-parse --short HEAD)
mv .env.production.bak .env.production
sed -i'.bak' -e "s/dev/$version/g" .env.production