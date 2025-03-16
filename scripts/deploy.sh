rm -rf kiffari_old
tar xvf bin.tar.gz
mv bin kiffari_new
mv kiffari kiffari_old
mv kiffari_new kiffari
rm bin.tar.gz
curl -X POST --basic --user "APIKEY account=ACCOUNT:" https://api.alwaysdata.com/v1/site/SITEID/restart/