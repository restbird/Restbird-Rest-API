# encoding: utf-8
import requests
import urllib3

def ResponseValidate(response, ctx) :
    print ("I am in ResponseValidate")
    jsoncontent = response.json()
    
    print (jsoncontent)
    
    if 'counter' in ctx.vars:
        count = ctx.vars['counter']
        try:
            count = int(count) + 1
        except TypeError:
            count = 1
    else:
        count = 1

    ctx.vars['counter'] = count

    print (ctx.vars['counter'])

    ctx.setGlobal('globalcounter', count)
    print (ctx.getGlobal('globalcounter'))
    return True
