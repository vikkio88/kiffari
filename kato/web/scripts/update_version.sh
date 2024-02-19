version=$(git rev-parse --short HEAD)
sed -i'.bak' -e "s/VITE_VERSION=.*/VITE_VERSION=$version/g" .env.production