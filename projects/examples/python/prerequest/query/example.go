# encoding: utf-8
import requests

# GET http://xxxxx.com/?key2=value2&key1=value1
def PreRequest(request, ctx) :
    allparams = {'key1': 'value1', 'key2': 'value2'}

    request.params = allparams
