# Based on: https://stackoverflow.com/a/67126761

echo "This script is taken from: https://stackoverflow.com/a/67126761"
echo "Converting key..."
xxd -r -p pub.hex pub.bin
echo "Generateing temp key..."
openssl ecparam -name secp256r1 -genkey -noout -out tempkey.pem
echo "Converting temp key to .der..."
openssl ec -in tempkey.pem -pubout -outform der -out temppub.der
echo "Extracting bytes..."
head -c 26 temppub.der > public-header.der
echo "Creating .der file of pub key...."
cat public-header.der pub.bin > publickey.der
#echo "Converting to .pem file"
openssl ec -in publickey.der -pubin -inform der -out publickey.pem -pubout

echo "Removeing files...."
rm tempKey.pem
rm temppub.der
rm pub.bin
rm public-header.der
rm publickey.der
