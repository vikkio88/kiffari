go install github.com/air-verse/air@latest
cp web/.env.example web/.env
. ${NVM_DIR}/nvm.sh
cd web && nvm use && npm i