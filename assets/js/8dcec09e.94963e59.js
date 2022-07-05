"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[834],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return f}});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var c=r.createContext({}),u=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},p=function(e){var t=u(e.components);return r.createElement(c.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,c=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),d=u(n),f=o,m=d["".concat(c,".").concat(f)]||d[f]||l[f]||i;return n?r.createElement(m,a(a({ref:t},p),{},{components:n})):r.createElement(m,a({ref:t},p))}));function f(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,a=new Array(i);a[0]=d;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:o,a[1]=s;for(var u=2;u<i;u++)a[u]=n[u];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},2748:function(e,t,n){n.r(t),n.d(t,{assets:function(){return p},contentTitle:function(){return c},default:function(){return f},frontMatter:function(){return s},metadata:function(){return u},toc:function(){return l}});var r=n(3117),o=n(102),i=(n(7294),n(3905)),a=["components"],s={},c="Data persistency",u={unversionedId:"guides/persistency/introduction",id:"guides/persistency/introduction",title:"Data persistency",description:"Namespace helps you manage your dependencies across multiple environments, and",source:"@site/docs/guides/persistency/introduction.mdx",sourceDirName:"guides/persistency",slug:"/guides/persistency/introduction",permalink:"/guides/persistency/introduction",draft:!1,editUrl:"https://github.com/namespacelabs/docs/tree/main/docs/guides/persistency/introduction.mdx",tags:[],version:"current",frontMatter:{},sidebar:"sidebar",previous:{title:"Getting started",permalink:"/getting-started"},next:{title:"Using PostgreSQL",permalink:"/guides/persistency/postgresql"}},p={},l=[{value:"How it works",id:"how-it-works",level:2}],d={toc:l};function f(e){var t=e.components,n=(0,o.Z)(e,a);return(0,i.kt)("wrapper",(0,r.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"data-persistency"},"Data persistency"),(0,i.kt)("p",null,"Namespace helps you manage your dependencies across multiple environments, and\nthat includes databases and object stores. Currently, we support the following\ndatabases and object stores out of the box, but are adding more as we go:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"../postgresql"},"PostgreSQL")," (in-cluster, or AWS RDS)"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"../s3"},"AWS S3")),(0,i.kt)("li",{parentName:"ul"},"Redis (in-cluster only)"),(0,i.kt)("li",{parentName:"ul"},"MariaDB (in-cluster only)")),(0,i.kt)("h2",{id:"how-it-works"},"How it works"),(0,i.kt)("p",null,"To dramatically simplify using resources, Namespace manages both the\nprovisioning (i.e. creating and maintaining) of the resource (e.g. databases,\nservers, hosted instances), authentication, and even how your application is\nconfigured to use them."),(0,i.kt)("p",null,"We hook into the build process, generating code which bring in the required SDKs\nto support a dependency, and into the deployment workflow to ensure that\ndependencies that an application ",(0,i.kt)("em",{parentName:"p"},"instantiates")," are created and updated before\nyour application code actually runs."),(0,i.kt)("p",null,"You can explore more of how Namespace works in our ",(0,i.kt)("a",{parentName:"p",href:"/reference/architecture"},"Architecture")," reference."))}f.isMDXComponent=!0}}]);