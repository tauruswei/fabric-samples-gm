[//]: # (SPDX-License-Identifier: CC-BY-4.0)

## Hyperledger Fabric Samples(1.4.2)

- 将二进制可执行文件拷贝到bin目录

二进制文件：peer、orderer、cryptogen、idemixgen、discover、configtxgen、configtxlator、discover

      比如：cp -f cryptogen bin
    
- 导入镜像

镜像：peer.tar.gz、orderer.tar.gz、tools.tar.gz、ccenv.tar.gz、buildenv.tar.gz

      比如：docker load < peer.tar.gz
   
- 运行

    `cd firstnetwork`

    `bash byfn.sh restart`