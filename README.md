## Hyperledger Fabric Sample Application

This application demonstrates the creation and transfer of container shipments between actors leveraging Hyperledger Fabric in the supply chain. In this demo app we will set up the minimum number of nodes required to develop chaincode. It has a single peer and a single organization.

if getting error about running ./startFabric.sh permission 

chmod a+x startFabric.sh

This code is based on code written by the Hyperledger Fabric community. Source code can be found here: (https://github.com/hyperledger/fabric-samples). 


## Cleanup

Remove any pre-existing containers, as it may conflict with commands in this project

```
docker rm -f $(docker ps -aq)

docker rmi -f $(docker images -a -q)

rm -rf ~/Library/Containers/com.docker.docker/Data/*
```

Remove key store contents

```
cd ~ && rm -rf .hfc-key-store/
```

## Starting the Application

1. Start the Hyperledger Fabric network

```
./startFabric.sh
```

2. Install the required libraries from the package.json file

```
yarn
```

3. Register the Admin and User components of our network

```
node registerAdmin.js

node registerUser.js
```

4. Start the client application

```
yarn dev
```

Load the client simply by opening localhost:3000 in any browser window of your choice, and you should see the user interface for our simple application at this URL.