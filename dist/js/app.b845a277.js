(function(e){function t(t){for(var n,o,u=t[0],s=t[1],c=t[2],l=0,f=[];l<u.length;l++)o=u[l],a[o]&&f.push(a[o][0]),a[o]=0;for(n in s)Object.prototype.hasOwnProperty.call(s,n)&&(e[n]=s[n]);d&&d(t);while(f.length)f.shift()();return i.push.apply(i,c||[]),r()}function r(){for(var e,t=0;t<i.length;t++){for(var r=i[t],n=!0,o=1;o<r.length;o++){var u=r[o];0!==a[u]&&(n=!1)}n&&(i.splice(t--,1),e=s(s.s=r[0]))}return e}var n={},o={app:0},a={app:0},i=[];function u(e){return s.p+"js/"+({"FaucetGet-vue":"FaucetGet-vue"}[e]||e)+"."+{"FaucetGet-vue":"6eb7dc99"}[e]+".js"}function s(t){if(n[t])return n[t].exports;var r=n[t]={i:t,l:!1,exports:{}};return e[t].call(r.exports,r,r.exports,s),r.l=!0,r.exports}s.e=function(e){var t=[],r={"FaucetGet-vue":1};o[e]?t.push(o[e]):0!==o[e]&&r[e]&&t.push(o[e]=new Promise(function(t,r){for(var n="css/"+({"FaucetGet-vue":"FaucetGet-vue"}[e]||e)+"."+{"FaucetGet-vue":"200292df"}[e]+".css",a=s.p+n,i=document.getElementsByTagName("link"),u=0;u<i.length;u++){var c=i[u],l=c.getAttribute("data-href")||c.getAttribute("href");if("stylesheet"===c.rel&&(l===n||l===a))return t()}var f=document.getElementsByTagName("style");for(u=0;u<f.length;u++){c=f[u],l=c.getAttribute("data-href");if(l===n||l===a)return t()}var d=document.createElement("link");d.rel="stylesheet",d.type="text/css",d.onload=t,d.onerror=function(t){var n=t&&t.target&&t.target.src||a,i=new Error("Loading CSS chunk "+e+" failed.\n("+n+")");i.code="CSS_CHUNK_LOAD_FAILED",i.request=n,delete o[e],d.parentNode.removeChild(d),r(i)},d.href=a;var p=document.getElementsByTagName("head")[0];p.appendChild(d)}).then(function(){o[e]=0}));var n=a[e];if(0!==n)if(n)t.push(n[2]);else{var i=new Promise(function(t,r){n=a[e]=[t,r]});t.push(n[2]=i);var c,l=document.createElement("script");l.charset="utf-8",l.timeout=120,s.nc&&l.setAttribute("nonce",s.nc),l.src=u(e),c=function(t){l.onerror=l.onload=null,clearTimeout(f);var r=a[e];if(0!==r){if(r){var n=t&&("load"===t.type?"missing":t.type),o=t&&t.target&&t.target.src,i=new Error("Loading chunk "+e+" failed.\n("+n+": "+o+")");i.type=n,i.request=o,r[1](i)}a[e]=void 0}};var f=setTimeout(function(){c({type:"timeout",target:l})},12e4);l.onerror=l.onload=c,document.head.appendChild(l)}return Promise.all(t)},s.m=e,s.c=n,s.d=function(e,t,r){s.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r})},s.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},s.t=function(e,t){if(1&t&&(e=s(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(s.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)s.d(r,n,function(t){return e[t]}.bind(null,n));return r},s.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return s.d(t,"a",t),t},s.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},s.p="/",s.oe=function(e){throw console.error(e),e};var c=window["webpackJsonp"]=window["webpackJsonp"]||[],l=c.push.bind(c);c.push=t,c=c.slice();for(var f=0;f<c.length;f++)t(c[f]);var d=l;i.push([0,"chunk-vendors"]),r()})({0:function(e,t,r){e.exports=r("56d7")},"1a5d":function(e,t,r){var n={"./FaucetGet.vue":["8e1c","FaucetGet-vue"]};function o(e){var t=n[e];return t?r.e(t[1]).then(function(){var e=t[0];return r(e)}):Promise.resolve().then(function(){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t})}o.keys=function(){return Object.keys(n)},o.id="1a5d",e.exports=o},2703:function(e,t){function r(e){return Promise.resolve().then(function(){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t})}r.keys=function(){return[]},r.resolve=r,e.exports=r,r.id="2703"},"4db1":function(e,t,r){e.exports=r.p+"img/logo-proximax-sirius-faucet_new.17bb0a03.svg"},"56d7":function(e,t,r){"use strict";r.r(t);var n=r("2b0e"),o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("div",{staticClass:"header"},[n("div",{staticClass:"logo"},[n("img",{staticClass:" responsive",attrs:{src:r("4db1")}}),e._v(" "+e._s(e.version)+" ")])]),n("main",[n("router-view"),n("footer",[n("p",{staticClass:"copyright"},[e._v("© ProximaX 2019. All Rights Reserved. "),n("b",[e._v(" Sirius Faucet "+e._s(e.version))])])])],1),n("notifications",{attrs:{group:"foo"}})],1)},a=[],i={data(){return{version:"v0.0.5"}}},u=i,s=(r("5c0b"),r("2877")),c=Object(s["a"])(u,o,a,!1,null,null,null),l=c.exports,f=r("8c4f");function d(e){return()=>r("1a5d")(`./${e}.vue`)}n["default"].use(f["a"]);var p=new f["a"]({routes:[{path:"/faucet",component:d("FaucetGet")},{path:"*",redirect:"/faucet"},{path:"/",component:d("FaucetGet")}]}),v=r("2f62");n["default"].use(v["a"]);var h=new v["a"].Store({state:{},mutations:{},actions:{}}),g=r("9483");Object(g["a"])("/service-worker.js",{ready(){console.log("App is being served from cache by a service worker.\nFor more details, visit https://goo.gl/AFskqB")},registered(){console.log("Service worker has been registered.")},cached(){console.log("Content has been cached for offline use.")},updatefound(){console.log("New content is downloading.")},updated(){console.log("New content is available; please refresh."),setTimeout(()=>{window.location.reload(!0)},1e3)},offline(){console.log("No internet connection found. App is running in offline mode.")},error(e){console.error("Error during service worker registration:",e)}});var m=r("ee98"),b=r.n(m);const y="",w="api";var _={API_BASE_URL:y,API_URL:w},O=r("bc3a"),F=r.n(O);class k{constructor(e,t,r){this.api=e.API_URL,this.axios=t,this.headers=r}get(e=null){return this.axios.get(`${this.api}/${e}`)}post(){console.log("post")}put(){console.log("put")}delete(){console.log("delete")}}const x=new k(_,F.a,{key:"Content-type",value:"application/json"});n["default"].config.productionTip=!1,n["default"].prototype.$apiService=x,n["default"].use(b.a),new n["default"]({router:p,store:h,render:function(e){return e(l)}}).$mount("#app")},"5c0b":function(e,t,r){"use strict";var n=r("5e27"),o=r.n(n);o.a},"5e27":function(e,t,r){}});
//# sourceMappingURL=app.b845a277.js.map