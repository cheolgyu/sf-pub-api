# stock-read-http-api

메인 : https://github.com/bxcodec/go-clean-arch   
참고   
    https://github.com/jfeng45/servicetmpl   

---
## 실행
gin -i --appPort 5001  --port 5000  run  main.go 
## 빌드
```
golang 환경변수 설정 (powershell)

$env:GOOS = 'linux'
$env:GOARCH = 'amd64'

go build -o bin/release main.go    

ssh -i "highserpot_stock.pem" ec2-user@ec2-3-35-30-100.ap-northeast-2.compute.amazonaws.com

ec2 업로드 전 기존 프로세스 kill -9 p_id 시키기.

scp -i "highserpot_stock.pem" stock-read-http-api/.env.prod  ec2-user@3.35.30.100:~/stock-read-http-api/.env.prod
scp -i "highserpot_stock.pem" stock-read-http-api/bin/release  ec2-user@3.35.30.100:~/stock-read-http-api/release
chmod +x ./stock-read-http-api/release
nohup ./stock-read-http-api/release -prod    > nohup.out &
```

```
할것:
project = 단타+월피그 인데 
    월피크 따로 뺴고, 
    월피크랑 단타랑 분리시키기 
    변수이름 다시 짓기.

```