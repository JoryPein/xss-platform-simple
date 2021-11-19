# readme



## 服务端
默认监听
    http 80
    https 443
xss.exe -mode server -ip 0.0.0.0 -port 9999 -pass test1234567 -sslcrt c:\\1.crt -sslkey c:\\1.key

## 客服端
默认web面板监听
    http 9090
    
xss.exe -mode client -ip 1.1.1.1 -port 9999 -pass test1234567