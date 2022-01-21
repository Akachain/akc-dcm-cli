## Environmnet
- CLUSTER_NAME=cluster-akcsandbox
- ORDERER_NAME=orderer
- ORDERER_DOMAIN=akcsandbox
- ROOT_CA_NAME=rca-akc
- ORG_NAME=akc
- ORG_DOMAIN=akcsandbox
##
I. Mamba version 1

1. Root CA certs
- Expire in: 15 years
- Signcert Location:
In rca pod: `/etc/hyperledger/fabric-ca/ca-cert.pem`
In efs pod:
    ```bash
    # Cert
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/${ROOT_CA_NAME}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/msp/tlscacerts/ica-${ORDERER_NAME}-${ORDERER_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/msp/cacerts/ica-${ORDERER_NAME}-${ORDERER_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/admin/msp/cacerts/ica-${ORDERER_NAME}-${ORDERER_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/msp/tlscacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/admin/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem

    # Chain
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/ca/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/tlsca/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/ca/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/tlsca/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/ica-${ORG_NAME}-ca-chain.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/ica-${ORDERER_NAME}-ca-chain.pem
    ```

- Key Location:
In rca pod: `/etc/hyperledger/fabric-ca/msp/keystore/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_sk`

2. Intermediate CA
- Expire in: 5 years
- Signcert Location:
In ica pod: `/etc/hyperledger/fabric-ca/ca-chain.pem`
In efs pod:
    ```bash
    # PEER ORG
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/ca/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/tlsca/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/ica-${ORG_NAME}-ca-chain.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/msp/tlsintermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/admin/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem

    # ORDERER ORG
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/ca/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/msp/tlscacerts/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/msp/cacerts/ca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/tls/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/tlsca/tlsca.${ORDERER_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/ica-${ORDERER_NAME}-ca-chain.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/msp/intermediatecerts/ica-${ORDERER_NAME}-${ORDERER_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/msp/tlsintermediatecerts/ica-${ORDERER_NAME}-${ORDERER_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/admin/msp/intermediatecerts/ica-${ORDERER_NAME}-${ORDERER_DOMAIN}-7054.pem
    ```
- Key Location:
In ica pod: `/etc/hyperledger/fabric-ca/msp/keystore/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_sk`

3. Org MSP
- None
4. Peer MSP
- Expire in: 1 year
- Signcert Location:
In efs pod:
    ```bash
    # Peer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/msp/signcerts/cert.pem
    # Orderer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/msp/signcerts/cert.pem
    ```
- Key Location:
In efs pod:
    ```bash
    # Peer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/msp/keystore/key.pem
    # Orderer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/msp/keystore/key.pem
    ```
5. Peer TLS
- Expire in: 1 year
- Signcert Location:
In efs pod:
    ```bash
    # Peer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/tls/server.crt
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/peer0-akc/tls/server.crt

    # Orderer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/tls/server.crt
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/tls/server.crt
    ```
- Key Location:
In efs pod:
    ```bash
    # Peer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/tls/server.key
    # Orderer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/tls/server.key
    ```
6. User is used by peer/cli
- Expire in: 1 year
- Location:
    ```
    # Peer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/peers/peer0.${ORG_DOMAIN}/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/admin/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORG_NAME}/admin/msp/admincerts/cert.pem

    # Orderer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/orderers/orderer0-${ORDERER_NAME}.${ORDERER_DOMAIN}/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/msp/admincerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/admin/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/orgs/${ORDERER_NAME}/admin/msp/admincerts/cert.pem
    ```
- Key Location:
In efs pod:
    ```bash
    # Peer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/peerOrganizations/${ORG_DOMAIN}/users/admin/msp/keystore/key.pem

    # Orderer
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/ordererOrganizations/${ORDERER_DOMAIN}/users/admin/msp/keystore/key.pem
    ```

7. User is used by application
- Expire in: 1 year
- Location:
    ```
    /tmp/artifact/${CLUSTER_NAME}/admin/crypto-path/fabric-client-kv-${ORG_NAME}
    /tmp/artifact/${CLUSTER_NAME}/admin/crypto-store/fabric-client-kv-${ORG_NAME}
    ```
- How to get private key:
    - Read json file of user in crypto-path folder: ```cat admin/crypto-path/fabric-client-kv-org1/org1```. Find the value of the "signingIdentity". Example: ```"signingIdentity":"233446076089220cecaae13366ddc8d9f8d66b9f58fa6569e3bd0f544331eb6b"```

    - Private key locate at folder crypto-store. Name is same with value of signingIdentity. Example: ```cat admin/crypto-store/fabric-client-kv-org1/233446076089220cecaae13366ddc8d9f8d66b9f58fa6569e3bd0f544331eb6b-priv```

II. Mamba version 2

1. Root CA certs
- Expire in: 15 years
- Signcert Location:
In rca pod: `/etc/hyperledger/fabric-ca/ca-cert.pem`
In efs pod:
    ```
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/${ROOTCA_NAME}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/tlsca/tlsca.${ORG_NAME}.${ORG_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/admin/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/user-${ORG_NAME}/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/ca/ca.${ORG_NAME}.${ORG_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/tlscacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/tlscacerts/ca.crt
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/tls/ca.crt
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/tls/tlscacerts/tls-ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/msp/tlscacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/admin/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/user-orderer/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/tls/ca.crt
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/tls/tlscacerts/tls-ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/msp/tlscacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/msp/tlscacerts/tlsca.${ORG_NAME}.${ORG_DOMAIN}-cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/cacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/tlscacerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/tlscacerts/tlsca.${ORG_NAME}.${ORG_DOMAIN}-cert.pem
    ```
- Key Location:
In rca pod: `/etc/hyperledger/fabric-ca/msp/keystore/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_sk`

2. Intermediate CA
- Expire in: 5 years
- Signcert Location:
In ica pod: `/etc/hyperledger/fabric-ca/ca-chain.pem`
In efs pod:
    ```
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/ica-${ORG_NAME}-ca-chain.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/admin/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/user-${ORG_NAME}/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/tlsintermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/tls/tlsintermediatecerts/tls-ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/msp/tlsintermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem


    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/ica-${ORDERER_NAME}-ca-chain.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/admin/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/user-orderer/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/tls/tlsintermediatecerts/tls-ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/msp/tlsintermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/tlsintermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/intermediatecerts/ica-${ORG_NAME}-${ORG_DOMAIN}-7054.pem
    ```
3. Org MSP
- Expire in: 1 year
- 
    ```
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/msp/signcerts/cert.pem
    ```
4. Peer MSP
- Expire in: 1 year
- Location:
In efs pod:
    ```
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/msp/signcerts/cert.pem
    ```
5. Peer TLS
- Expire in: 1 year
- Location:
In efs pod:
    ```
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/tls/server.crt
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/peers/peer0-${ORG_NAME}.${ORG_DOMAIN}/tls/signcerts/cert.pem

    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/tls/server.crt
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/orderers/orderer0-${ORG_NAME}.${ORG_DOMAIN}/tls/signcerts/cert.pem
    ```

6. User is used by peer/cli
- Expire in: 1 year
- Location:
    ```
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/admin/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/user-${ORG_NAME}/msp/signcerts/cert.pem

    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/admin/msp/signcerts/cert.pem
    /tmp/artifact/${CLUSTER_NAME}/akc-ca-data/crypto-config/${ORG_NAME}.${ORG_DOMAIN}/users/user-orderer/msp/signcerts/cert.pem
    ```

7. User is used by application
- 1 y
- 
    ```
    /tmp/artifact/${CLUSTER_NAME}/admin-v2/wallet/ica-${ORG_NAME}-admin.id
    /tmp/artifact/${CLUSTER_NAME}/admin-v2/wallet/${ORG_NAME}.id
    ```
