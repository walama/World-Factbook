FROM golang:latest
WORKDIR /scratch_api
ADD . /scratch_api

FROM python:latest
WORKDIR /scratch_api
ADD . /scratch_api
RUN pip install FLask Requests
