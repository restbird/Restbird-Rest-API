# encoding: utf-8

import requests
import urllib3
from restapi import RestApi

def PythonScripts(ctx) :
	#print ("I am in PythonScript Api")
	#call project Demo_Rest/Rest_Python api0
    api = RestApi("Demo_Rest/Rest_Python", "api0", ctx)

    result = api.runApi()
    if result == True:
	    print ("call api success")
    else:
    	print ("call api failed")

    return result