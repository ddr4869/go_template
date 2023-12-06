# 기본 이미지 선택
FROM golang:1.20-alpine

# 작업 디렉토리 설정
WORKDIR /app

# 필요한 파일 복사
COPY . .

# Go 애플리케이션 빌드
RUN go build -o myapp .

# 애플리케이션 실행 명령어
CMD ["./myapp"]