# encoding: utf-8
import requests
import json

def PreRequest(request, ctx) :
    payload ={'key1':'value1', 'key2':'value2'}

    request.data = json.dumps(payload)
