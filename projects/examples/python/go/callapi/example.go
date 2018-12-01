# encoding: utf-8

import requests
from restapi import RestApi

def PythonScripts(ctx) :
    api = RestApi("Test/testpy", "api0", ctx)

    result = api.runApi()
    if result == True:
	print ("call api success")
    else:
	print ("call api failed")

    return result
