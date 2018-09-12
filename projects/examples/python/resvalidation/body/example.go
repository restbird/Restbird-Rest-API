# encoding: utf-8
import requests

def ResponseValidate(response, ctx) :
    substr = "hello world"
    if substr in response.text:
        return True
    else:
        return False
