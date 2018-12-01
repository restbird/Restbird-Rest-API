# encoding: utf-8
"""
write any your Python scripts here
"""
import sys
from os.path import os,join
sys.path.append(os.path.os.path.dirname(__file__))
import setting
sys.path.append(join(setting.local, "libraryPy"))

import requests
from restapi import RestApi
from restBirdCtx import RestBirdCtx

#Call Demo_Rest/Rest_Python api0 with env Demo_Rest_Env
sess = requests.Session()
ctx  = RestBirdCtx(setting.local, "Demo_Rest_Env", sess, "")

api = RestApi("Demo_Rest/Rest_Python", "api0", ctx)
api.runApi()

#Get/Set global env
ctx.setGlobal("hello", "world")
v = ctx.getGlobal("hello")
print (v)

#Get Env
print(ctx.loadEnvVaribles("Demo_Rest_Env"))
print(ctx.vars['counter'])
