# encoding: utf-8
import requests

def PreRequest(request, ctx) :
    request.method = ctx.getGlobal('method')
    request.url = ctx.vars.get('url', "http://www.google.com/")

    ctx.setGlobal('hello', 'world')
