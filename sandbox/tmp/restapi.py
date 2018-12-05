# encoding: utf-8

import re
import time
import io
import os
import logging

import requests
import urllib3
import json

from pathlib import Path

import pkgutil
from importlib import import_module
import importlib.machinery

from requests import Request, Response
from requests.exceptions import (InvalidSchema, InvalidURL, MissingSchema,
                                 RequestException)
from os.path import join

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

absolute_http_url_regexp = re.compile(r"^https?://", re.I)
RestProjectBasePath = join("projects","rest")
RestProjectStatePath = join("state", "rest")

logger = logging.Logger('catch_all')

class RestApi():
    """
    A api definition
    """

    def __init__(self, casepath, apiid, ctx):
        self.local = ctx.local
        self.casepath = casepath
        self.apiid = apiid
        self.ctx = ctx
        self.sess = ctx.sess
        self.hisid = ctx.hisid

        self.abspath = join(self.local , RestProjectBasePath , self.casepath , self.apiid)
        self.reqJsonFile = join(self.abspath , 'req.json')

        self.preRequestScriptFile = join(self.abspath, 'prerequest.py')
        self.resvalScriptFile = join(self.abspath, 'resval.py')
        self.reqDataFile = join(self.abspath, 'req.body')   
        
        self.PythonScriptFile = join(self.abspath, 'pyscript.py')   
        
        #to be loaded from file
        self.reqJson = {}

        #to be set after request
        self.summary = {}

        self.absStatePath = join(self.local, RestProjectStatePath , self.casepath , str(self.hisid), self.apiid)
        self.stateReqJsonFile = join(self.absStatePath, 'req.json')
        self.stateReqBodyFile = join(self.absStatePath,  'req.body')
        self.stateResBodyFile = join(self.absStatePath,  'res.body')
        self.stateVarsBodyFile = join(self.absStatePath,  'vars.json')

        my_file = Path(self.absStatePath)
        if not my_file.exists():
            os.makedirs(self.absStatePath,exist_ok=True)

        pass

    def dynamic_import(self, abs_module_path, module_name, class_name):
        #loadedmodule = importlib.machinery.SourceFileLoader(module_name, abs_module_path).load_module()

        #target_class = getattr(loadedmodule, class_name)
        #For the reason load_module() has been depreated since 3.4
        loader = importlib.machinery.SourceFileLoader(module_name, abs_module_path)
        spec = importlib.util.spec_from_loader(loader.name, loader)
        loadedmodule = importlib.util.module_from_spec(spec)
        loader.exec_module(loadedmodule)

        target_class = getattr(loadedmodule, class_name)
         
        return target_class

    def load_reqjson(self):
        """ load json file and check file content format
        """
        with io.open(self.reqJsonFile, encoding='utf-8') as data_file:
            try:
                json_content = json.load(data_file)
            except json.JSONDecodeError:
                json_content = {}

            if not isinstance(json_content, (list, dict)):
                json_content = {}

        self.reqJson = json_content

    def findandreplace(self, orgstr, ctx):
        items = re.findall('{{(.+?)}}', orgstr)

        for item in items :
            if item in ctx.vars:
                orgstr = re.sub('{{'+ item +'}}', ctx.vars[item], orgstr)
            else :
                v = ctx.getGlobal(str(item))
                if v != None:
                    orgstr = re.sub('{{'+ item +'}}', v, orgstr)
        return orgstr

    def bindVaribles(self, org, ctx):
        if isinstance(org, dict):
            neworg = {}
            for k, v in org.items():
            	v = self.findandreplace(v, ctx)
            	neworg[k] = v
            	pass				
            return neworg                
            pass
        elif isinstance(org, str):
            return self.findandreplace(org, ctx)
            pass
        else:
            return org

    def loadapi(self):
        self.load_reqjson()

        apitype = self.reqJson.get('type', 0)
        if apitype == 0 :
            self.PreRequest = self.dynamic_import(self.preRequestScriptFile, 'prerequest', 'PreRequest')
            self.ResponseValidate = self.dynamic_import(self.resvalScriptFile, 'resval','ResponseValidate') 
        elif apitype == 1:
            pass
        elif apitype == 2:
            self.PythonScripts = self.dynamic_import(self.PythonScriptFile, 'pyscript', 'PythonScripts')
            pass
        else:
            return
        
                    
        pass
    
    def saveRequestRecord(self, req, resp, envmap):
        tosave_reqjson = {}
        tosave_reqbody = ""
        tosave_resbody = ""
        tosave_vars = {}

        if self.hisid == "":
            return
			
        reqheaders = []
        for k, v in req.headers.items():
            reqheaders.append({'isdisabled': False,'key': k, 'value':v})

        respheaders = []
        for k, v in resp.headers.items():
            respheaders.append({'isdisabled': False,'key': k, 'value':v})

        tosave_reqjson['request'] = {
            'method' : req.method,
            'url': req.url,
            'auth': {
                'method': 'no',
                'user': '',
                'password':''
            },
            'header': reqheaders,
            'body' : {
                'type': 'raw',
                'data': '',
                'kvdata':[],
            },
        }

        tosave_reqjson['response'] = {
            'status' : resp.reason,
            'statuscode' : resp.status_code,
            'header': respheaders,
            'body' : {
                'type': 'raw',
                'data': '',
                'kvdata':[],
            },
        }

        tosave_reqjson['responseval'] = {
            'result' : self.summary["result"],
        }

        tosave_reqjson['time'] = {
            'duration' : self.summary["response_time_ms"],
        }

        tosave_reqjson['apiid'] = self.apiid
        tosave_reqjson['status'] = 1

        with open(self.stateReqJsonFile, 'w+') as outfile:
            json.dump(tosave_reqjson, outfile, indent=4, sort_keys=False)

        with open(self.stateReqBodyFile, 'w+') as outfile:
            if req.body:
                outfile.write(req.body)
            else :
                outfile.write("")
        
        with open(self.stateResBodyFile, 'w+') as outfile:
            outfile.write(resp.text)
            pass

        with open(self.stateVarsBodyFile, 'w+') as outfile:
            json.dump(envmap, outfile, indent=4, sort_keys=False)
            pass        

    def saveScriptRecord(self, envmap):
        tosave_reqjson = {}
        tosave_vars = {}
		
        if self.hisid == "":
            return

        tosave_reqjson['request'] = {
            'describ' : self.reqJson.get('request', {}).get('describ', "python script"),
        }
		
        tosave_reqjson['responseval'] = {
            'result' : self.summary["result"],
        }

        tosave_reqjson['time'] = {
            'duration' : self.summary["response_time_ms"],
        }

        tosave_reqjson['apiid'] = self.apiid
        tosave_reqjson['status'] = 1
        tosave_reqjson['type'] = 2
		
        with open(self.stateReqJsonFile, 'w+') as outfile:
            json.dump(tosave_reqjson, outfile, indent=4, sort_keys=False)

        with open(self.stateVarsBodyFile, 'w+') as outfile:
            json.dump(envmap, outfile, indent=4, sort_keys=False)
            pass        
			        
    def runApi(self, **kwargs):
        self.loadapi()
        result = False

        apitype = self.reqJson.get('type', 0)
        if apitype == 0 :
            result = self.doRequest(self.ctx, kwargs)
        elif apitype == 1:
            pass
        elif apitype == 2:
            result = self.doScript(self.ctx, kwargs)
            pass

        return result

    def doScript(self, ctx, kwargs):
        # record original request info
        self.summary["request_time"] = time.time()
        result = False
        ctx.local = self.local
        try:
            result = self.PythonScripts(ctx)
        except Exception as e:
            result = False
            logger.error(e, exc_info=True)

        self.summary['result'] = result													 
        self.summary["response_time_ms"] = round(time.time() - self.summary["request_time"], 6) 
		
        self.saveScriptRecord(ctx.vars)
        return result

    def doRequest(self, ctx, kwargs):

        # record original request info
        self.summary["request_time"] = time.time()

        reqheaders = {}
        for onhdr in self.reqJson.get('request').get("header"):
            if onhdr['disabled'] == False:
                reqheaders[onhdr['key']] = onhdr['value']
		
        myurl = self.bindVaribles(self.reqJson.get('request').get("url"), ctx)
        reqheaders = self.bindVaribles(reqheaders, ctx)
		
        bodytype = self.reqJson.get('request').get('body').get('type')
        if bodytype == 'raw' and os.path.exists(self.reqDataFile):
            data = ""
            mode = 'r' 
            with io.open(self.reqDataFile, mode) as data_file:
                try:
                    data = data_file.read()
                except json.JSONDecodeError:
                    data = ""            
                data = 	self.bindVaribles(data, ctx)
        else:
            data = {}
            for oneform in self.reqJson.get('request').get('body').get('kvdata'):
                if oneform['disabled'] == False:			
                    data[oneform['key']] = self.bindVaribles(oneform['value'], ctx)				
				
        req = Request(self.reqJson.get('request').get("method"), 
            myurl, 
            data=data, 
            headers=reqheaders)

        try:
            self.PreRequest(req, ctx)
        except Exception as e:
            logger.error(e, exc_info=True)
        prepped = req.prepare()
        
        resp = self.sess.send(prepped)

        self.summary["url"] = (resp.history and resp.history[0] or resp).request.url
        self.summary["request_headers"] = resp.request.headers
        self.summary["request_body"] = resp.request.body

        # record response info
        self.summary["status_code"] = resp.status_code
        self.summary["response_headers"] = resp.headers
        try:
            self.summary["response_body"] = resp.json()
        except ValueError:
            self.summary["response_body"] = resp.content

        if kwargs.get("stream", False):
            self.summary["content_size"] = int(self.summary["response_headers"].get("content-length") or 0)
        else:
            self.summary["content_size"] = len(resp.content or "")
        
        self.summary["response_time_ms"] = round(time.time() - self.summary["request_time"], 6)
        self.summary["elapsed_ms"] = resp.elapsed.microseconds / 1000.0

        try:
            result = self.ResponseValidate(resp, ctx)
            if result :
                self.summary["result"] = True
            else:
                self.summary["result"] = False            
        except Exception as e:
            logger.error(e, exc_info=True)
            self.summary["result"] = False            

        try:  
            self.saveRequestRecord(prepped, resp, ctx.vars)
        except Exception as e:
            logger.error(e, exc_info=True)
        return result
    
    def result(self):
        pass