# encoding: utf-8

import re
import time
import io
import os
import logging

import json

from pathlib import Path
from os.path import join

logger = logging.Logger('catch_all')


RestProjectGlobalPath = join("state","global")

class RestBirdCtx():
    """
    RestBird ctx definition
    """

    def __init__(self, local, env, sess, hisid):
        self.env = env
        self.sess = sess
        self.hisid = hisid
        self.local = local
        self.globalPath = join(local, RestProjectGlobalPath)

        my_file = Path(self.globalPath)
        if not my_file.exists():
            os.makedirs(self.globalPath,exist_ok=True)

        self.vars = self.loadEnvVaribles(env)
        pass

    def getGlobal(self, key):
        """ get a global string """
        globalpath = join(self.globalPath, key)

        my_file = Path(globalpath)
        if not my_file.exists():
            return None

        with io.open(globalpath, "r") as data_file:
            try:
                #data = data_file.read()
                data = json.load(data_file)                
            except Exception as e:
                logger.error(e, exc_info=True)	
                return None
        return str(data)

    def setGlobal(self, key, value):
        """ set a global string """
        globalpath = join(self.globalPath, key)

        with io.open(globalpath, "w") as data_file:
            try:
                #data_file.write(value)
                json.dump(value, data_file)				
            except Exception as e:
                logger.error(e, exc_info=True)			
                return False
        return True

    def loadEnvVaribles(self, env):
	    envmap = {}

	    if env == "" or env == "EMPTY" or env == "No Environment":
		    return envmap

	    envpath = join(self.local, 'projects','env', env ,'env.json')
	    with io.open(envpath, encoding='utf-8') as data_file:
		    try:
			    envjson = json.load(data_file)
			    if not isinstance(envjson, (list, dict)):
				    return envmap
			    else:
				    for ite in envjson['envs']:
					    if ite['disabled'] == False:
						    envmap[ite['key']] = ite['value']
		    except json.JSONDecodeError:
			    envmap = {}
	    return envmap	