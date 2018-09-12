# encoding: utf-8
import requests
from restapi import RestApi
from restBirdCtx import RestBirdCtx

"""
write any your Python scripts here
"""
sess = requests.Session()
ctx  = RestBirdCtx("", sess, "")

api = RestApi("gggss", "api2", ctx)
api.runApi()

ctx.setGlobal("hello", "world")
v = ctx.getGlobal("hello")
print (v)

