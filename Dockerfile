ARG NODE_VERSION=16

FROM golang:1.20

WORKDIR /app

RUN apt-get update -qq

# You need librariy files and headers of tesseract and leptonica.
# When you miss these or LD_LIBRARY_PATH is not set to them,
# you would face an error: "tesseract/baseapi.h: No such file or directory"
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev

# In case you face TESSDATA_PREFIX error, you minght need to set env vars
# to specify the directory where "tessdata" is located.
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/4.00/tessdata/

# Load languages.
# These {lang}.traineddata would b located under ${TESSDATA_PREFIX}/tessdata.
RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-deu \
  tesseract-ocr-jpn

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y && \
    curl -sL https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get install -y nodejs

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY *.png ./
COPY *.jpg ./
COPY . ./

# Let's have gosseract in your project and test it.
RUN go get -t github.com/otiai10/gosseract/v2

# RUN go build -o /docker-bin

RUN npm i -g nodemon

EXPOSE 8080

CMD ["nodemon", "--exec", "go", "run", ".", "--signal", "SIGTERM"]