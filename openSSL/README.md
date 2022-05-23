# openSSL

### 1.生成私钥
    openssl genrsa -out private.key 1024 (rsa)
    openssl ecparam -name secp384r1 -genkey -out private-ecc.key (ecdsa)
    
### 2.生成公钥
    openssl rsa -in private.key -pubout -out public.key (rsa)
    openssl pkey -in private-ecc.key -pubout -out public-ecc.key (ecdsa)

### 3.生成证书请求
    openssl req -new -key private.key -out private.csr (rsa)
    openssl req -new -key private-ecc.key -out private-ecc.csr (ecdsa)
    openssl req -new -key server.key -out server.csr -config server.conf (带配置文件)

### 4.生成ca私钥
    openssl genrsa -out ca.key 1024

### 5.生成ca根证书
    openssl req -new -x509 -days 3650 -key ca.key -out ca.crt

### 6.使用根证书签发证书
    openssl x509 -req -in private.csr -out private.crt -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650
    openssl x509 -req -in private-ecc.csr -out private-ecc.crt -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650
    openssl x509 -req -in server.csr -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extensions req_ext -extfile server.conf -days 3650 (带配置文件)

### 7.查看私钥明细
    openssl rsa -in private.key -noout -text
    openssl ec -in private-ecc.key -noout -text

### 8.查看证书明细
    openssl x509 -in private.crt -noout -text
    openssl x509 -in private-ecc.crt -noout -text

### 9.查看证书公钥信息
    openssl x509 -in private.crt -pubkey
    openssl x509 -in private-ecc.crt -pubkey
    openssl pkcs12 -info -in private.pk12 -nodes
    openssl pkcs12 -info -in private-ecc.pk12 -nodes

### 10.验证证书是否是根证书签发的
    openssl verify -verbose -CAfile ca.crt private.crt
    openssl verify -verbose -CAfile ca.crt private-ecc.crt

### 11.生成证书安装包
    openssl pkcs12 -export -in private.crt -inkey private.key -out private.pk12
    openssl pkcs12 -export -in private-ecc.crt -inkey private-ecc.key -out private-ecc.pk12