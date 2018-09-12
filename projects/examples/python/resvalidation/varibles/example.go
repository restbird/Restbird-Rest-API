# encoding: utf-8
import requests

def ResponseValidate(response, ctx) :

    if 'counter' in vars:
        count = ctx.vars['counter']
        count = count +1
    else:
        count = 1

    ctx.vars['counter'] = count

    print (ctx.vars['counter'])

    ctx.setGlobal('globalcounter', count)
    print (ctx.getGlobal('globalcounter'))
    return True
