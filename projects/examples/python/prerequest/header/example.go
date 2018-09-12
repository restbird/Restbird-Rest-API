# encoding: utf-8
import requests

def PreRequest(request, ctx) :
    request.headers['aa'] = 'bb'
    request.headers['counter'] = '100'
