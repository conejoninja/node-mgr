## Node Manager

Tool to create and manage Ethereum nodes.

### Usage

*node-mgr* will look for a config.yml file in the current directory.

#### Commands

* **build:** creates two docker images called {base_name}:dev and {base_name}:alltoolls-dev based on the code given (local folder or remote repository) 
* **bootnode:** starts a bootnode (inside a docker container), needed to make nodes find each others  
* **node:** starts a regular Ethereum node (inside a docker container)
* **miner:** starts a miner (inside a docker container)
* **all:** starts a bootnode, a regular node and a miner (inside a docker container)
* **wipe:** stops and remove all containers, it also deletes all files and folders used by those containers


#### Configuration

*node-mgr* will look for a config.yml file in the current directory.

* **data_path:** folder to save required data by the different nodes
* **source_code:** local folder where the code is located, if left empty it will use the remote git repository specified at git_repo
* **git_repo:** remote git repository where the code is located, if left empty (and source_code too), it will clone the ethereum/go-ethereum repository
* **etherbase:** address where the reward from mining is sent to


### License

The MIT License

Copyright (c) 2018 Marketpay.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.