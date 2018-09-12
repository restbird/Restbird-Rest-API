# encoding: utf-8

import requests
from restapi import RestApi

def PythonScripts(vars) :
    pass
    sess = requests.Session()

    api = RestApi("Test/testpy", "api0", sess, "")

    #api.runApi({})
    api.runApi(vars)
