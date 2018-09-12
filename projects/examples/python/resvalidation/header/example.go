# encoding: utf-8
import requests

def ResponseValidate(response, ctx) :
    print (response.headers['Date'])
    print (response.headers.get('Content-Type',"text/html"))

    if 'Connection' in response.headers:
        print ("Connection header is exist")
    else:
        print ("Connection header is not exist")

    for k, v in response.headers.items():
        print (k + " = " + v)

    return True
