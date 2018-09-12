# encoding: utf-8
import requests

def ResponseValidate(response, ctx) :
    jsoncontent = response.json()

    print (jsoncontent['message'])
    return True
