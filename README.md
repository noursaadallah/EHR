# Blockchain-based solution to store Electronic Health Records

This tutorial will introduce to the basic steps that need to be followed in order to install and use this application which enables us to store, create and retrieve Electronic Medical Records or Electronic Health Records (next referred to as EHR).

Special thanks goes to [chainHero](https://github.com/chainHero) for sharing his knowledge and making this [amazing tutorial](https://github.com/chainHero/heroes-service).

## 1. Prerequisites

The blockchain technology used for this application is Hyperledger Fabric. In this tutorial I won't explain in detail how the technology works, just enough so that you can install this application and have a blockchain network up and running.

This application was developed on **Ubuntu 16.04**, but Hyperledger Fabric is compatible with Mac OS X, Windows and other Linux distributions. You can check the [official documentation](http://hyperledger-fabric.readthedocs.io/en/latest/) to see the functionalities of Fabric and what you can do with it.

We used the **Google Go** language to develop the whole stack of the application, because it's the initial supported language by Fabric and I find it easier to use. You can choose other supported languages if you prefer like *Java* or *Python* or *JavaScript*.

Fabric uses **Docker** to easily create a blockchain network and deploy the necessary components on its nodes.

## 2. Installation guide

All the installation steps mentioned here are for **Ubuntu 16.04**.

### 2.1. Docker

**Docker version 17.03.0-ce or greater is required.**

Start by installing the dependencies of Docker so that it may be installed correctly:

```bash
sudo apt install apt-transport-https ca-certificates curl software-properties-common
```
Once the dependencies are installed, we can install docker:

```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - && \
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" && \
sudo apt update && \
sudo apt install -y docker-ce
```

Now we need to manage the current user to avoid using administration rights (`root`) access when we will use the docker command. To do so, we need to add the current user to the `docker` group:

```bash
sudo groupadd docker ; \
sudo gpasswd -a ${USER} docker && \
sudo service docker restart
```

Do not mind if `groupadd: group 'docker' already exists` error pop up.

To apply the changes made, you need to logout/login. You can then check your version with:

```bash
docker -v
```

It should prompt a message like :

```
Docker version 17.06.0-ce, build 02c1d87
```

### 2.2. Docker compose

**Docker-compose version 1.8 or greater is required.**

Docker-compose enables the management of multiple containers at once.

The installation is pretty straightforward:

```bash
sudo curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose && \
sudo chmod +x /usr/local/bin/docker-compose
```

Logout/login to apply the changes and then check the version with:

```bash
docker-compose version
```

It should prompt a message like :

```
docker-compose version 1.14.0, build c7bdf9e
```

### 2.3. Go

**Go version 1.9.x or greater is required.**

You can either follow instructions from [golang.org](https://golang.org/dl/) or use these generics commands that will install Golang 1.9.2 and prepare your environment (generate your `GOPATH`) for Ubuntu:

```bash
wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz && \
sudo tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz && \
rm go1.9.2.linux-amd64.tar.gz && \
echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee -a /etc/profile && \
echo 'export GOPATH=$HOME/go' | tee -a $HOME/.bashrc && \
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' | tee -a $HOME/.bashrc && \
mkdir -p $HOME/go/{src,pkg,bin}
```

To make sure that the installation works, you can logout/login and run:

```bash
go version
```

### 2.4. Fabric SDK Go

Hyperledger Fabric proposes a set of SDKs that we can add to our web applications in order to communicate with the Fabric's components. In previous versions, we needed to install Fabric and Fabric-CA manually and separately. Now that's no longer an issue since the SDK-Go automatically installs the necessary components. There is a lot of version issues and incompatibility between these components, in order to avoid that, we will checkout to a specific commit with which the following tutorial (i.e. code) works.

```bash
go get -u github.com/hyperledger/fabric-sdk-go && \
cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go && \
git checkout 614551a752802488988921a730b172dada7def1d
```

Let's make sure that you have the requested dependencies:

```bash
cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go && \
make depend-install
```

Finally, we can launch the various tests of the SDK to check its proper functioning before going further:

```bash
cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go ; \
make
```

If you get the following error:

```
../fabric-sdk-go/vendor/github.com/miekg/pkcs11/pkcs11.go:29:18: fatal error: ltdl.h: No such file or directory
```

You need to install the package `libltdl-dev` and re-execute previous command (`make`):

```bash
sudo apt install libltdl-dev
```

The testing process can take some time depending on your network connection. What it does is first build a virtual network and download the docker images to do so then run some tests to check if the environment is ready and properly configured.

## 3. Download and Run the App

All you need to do now is to create a new directory in the `src` folder of your `GOPATH`, following this repository name.

```bash
mkdir -p $GOPATH/src/github.com/noursaadallah && \
cd $GOPATH/src/github.com/noursaadallah
git clone https://github.com/noursaadallah/EHR.git
cd EHR
make
```

Running the make command will run all the steps needed to run the App.
Hyperledger Fabric needs a lot of certificates to ensure encryption during the whole end to end process (TSL, authentications, signing blocks...). The creation of these files requires a little time, i copied all the crypto-material needed in the folder `fixtures` in this repository. These files are needed to have a running network. Alternatively, you can create your own files and your own network.

The steps that were ran are as follows:

1. Run the command `docker-compose up`. Which will instanciate nodes according to the configuration in `fixtures/docker-compose.yaml`. If you run `docker ps` you will see two peers, one orderer and one CA containers running. To check the different roles you can see the [official documentation](http://hyperledger-fabric.readthedocs.io/en/release-1.1/arch-deep-dive.html#client).
And to see in more detail how to create a network you can see [Building your first network](http://hyperledger-fabric.readthedocs.io/en/latest/build_network.html)
2. Initialise a client that can communicate to the nodes of the network. The corresponding configuration file is `config.yaml` and the corresponding initialisation code is in `blockchain/setup.go`.
3. Download and flatten all the dependencies (`Gopks.toml`), then build the app.
4. Install and instantiate the chaincode on the nodes.
5. Run the server.

Explaining the components of the app:
1. The `blockchain` directory comprises all that is needed by the fabric SDK go. We use the SDK to generate the transactions needed to get or create an EHR.
2. The `chaincode` directory contains the smart contract. The operations in the `blockchain` directory calls the functions of the smart contract.
3. The `fixtures` directory contains the crypto-material needed to create and run the network.
4. The `model` directory contains the model of the app, which is the structure of an EHR and the data of an appointment.
5. The `vendor` directory contains the flattened external libraries needed by the various components of the app. It is generated automatically.
6. The `web` directory contains the web application code (The controllers and The views).
7. The `main.go` file initialises the SDK, installs and instanciates the chaincode, then runs the web app.