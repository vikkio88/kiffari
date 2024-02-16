version=$(git rev-parse --short HEAD)
sed -i'.bak' -e "s/dev/$version/g" .env.production