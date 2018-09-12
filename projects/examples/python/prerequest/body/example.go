# encoding: utf-8
import requests

def PreRequest(request, ctx) :
    request.data = "hello world"
