# encoding: utf-8
import requests

def ResponseValidate(response, ctx) :
    print (response.status_code) #status code
    print (response.reason)      #status line

    if response.status_code ==  200:
        return True
    else:
        return False
