Fabric Sample 2.1
https://hyperledger-fabric.readthedocs.io/en/release-2.1/install.html



curl -sSL https://bit.ly/2ysbOFE | bash -s


cd fabric-samples/test-network
./network.sh down
./network.sh up
docker ps -a
./network.sh createChannel

./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go