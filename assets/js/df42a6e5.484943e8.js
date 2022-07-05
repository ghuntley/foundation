"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[445],{3905:function(e,n,t){t.d(n,{Zo:function(){return c},kt:function(){return g}});var r=t(7294);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function s(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function i(e,n){if(null==e)return{};var t,r,a=function(e,n){if(null==e)return{};var t,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var l=r.createContext({}),p=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):s(s({},n),e)),t},c=function(e){var n=p(e.components);return r.createElement(l.Provider,{value:n},e.children)},u={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},d=r.forwardRef((function(e,n){var t=e.components,a=e.mdxType,o=e.originalType,l=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),d=p(t),g=a,m=d["".concat(l,".").concat(g)]||d[g]||u[g]||o;return t?r.createElement(m,s(s({ref:n},c),{},{components:t})):r.createElement(m,s({ref:n},c))}));function g(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var o=t.length,s=new Array(o);s[0]=d;var i={};for(var l in n)hasOwnProperty.call(n,l)&&(i[l]=n[l]);i.originalType=e,i.mdxType="string"==typeof e?e:a,s[1]=i;for(var p=2;p<o;p++)s[p]=t[p];return r.createElement.apply(null,s)}return r.createElement.apply(null,t)}d.displayName="MDXCreateElement"},4088:function(e,n,t){t.r(n),t.d(n,{assets:function(){return c},contentTitle:function(){return l},default:function(){return g},frontMatter:function(){return i},metadata:function(){return p},toc:function(){return u}});var r=t(3117),a=t(102),o=(t(7294),t(3905)),s=["components"],i={},l="Using PostgreSQL",p={unversionedId:"guides/persistency/postgresql",id:"guides/persistency/postgresql",title:"Using PostgreSQL",description:"After completing these steps, you'll get:",source:"@site/docs/guides/persistency/postgresql.mdx",sourceDirName:"guides/persistency",slug:"/guides/persistency/postgresql",permalink:"/guides/persistency/postgresql",draft:!1,editUrl:"https://github.com/namespacelabs/docs/tree/main/docs/guides/persistency/postgresql.mdx",tags:[],version:"current",frontMatter:{},sidebar:"sidebar",previous:{title:"Data persistency",permalink:"/guides/persistency/introduction"},next:{title:"Using S3",permalink:"/guides/persistency/s3"}},c={},u=[],d={toc:u};function g(e){var n=e.components,t=(0,a.Z)(e,s);return(0,o.kt)("wrapper",(0,r.Z)({},d,t,{components:n,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"using-postgresql"},"Using PostgreSQL"),(0,o.kt)("p",null,"After completing these steps, you'll get:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"A configured PostgreSQL client ready to use in your codebase"),": regardless of\nenvironment: dev, prod, etc."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"A PostgreSQL server running in your stack"),", and managed for you: no need to\nmanage docker compose, or figure out how to run PostgreSQL in production,\nNamespace manages it for you, including the authentication."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"A database schema update system hooked up into the deployment workflow"),": no\nmore manual schema updates and migrations."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"A set of PostgreSQL-specific tools available at the tip of your fingers"),": also\nregardless of environment, jump into a ",(0,o.kt)("inlineCode",{parentName:"li"},"psql")," shell in a single command.")),(0,o.kt)("p",null,"In the ",(0,o.kt)("inlineCode",{parentName:"p"},"service.cue")," of the service where you want to use PostgreSQL, first add an import:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'import (\n    // ...\n    // highlight-next-line\n    postgresql "namespacelabs.dev/foundation/universe/db/postgres/incluster"\n)\n')),(0,o.kt)("p",null,"And then instantiate a database:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'service: fn.#Service & {\n    // ...\n\n    instantiate: {\n        // highlight-start\n        db: postgresql.#Exports.Database & {\n            name:       "foobar"\n            schemaFile: inputs.#FromFile & {\n                path: "foobar_schema.sql"\n            }\n        }\n        // highlight-end\n    }\n}\n')),(0,o.kt)("p",null,"And finally create a schema file along side your service implementation\n(",(0,o.kt)("inlineCode",{parentName:"p"},"foobar_schema.sql")," in the example above), with the database definition:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-sql"},"CREATE TABLE IF NOT EXISTS foobar (\n    Id INT GENERATED ALWAYS AS IDENTITY,\n    Item varchar(255) NOT NULL,\n    PRIMARY KEY(Id)\n);\n")),(0,o.kt)("p",null,"After you run ",(0,o.kt)("inlineCode",{parentName:"p"},"ns dev"),", you'll find a configured instance in your codebase:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func WireService(ctx context.Context, srv server.Registrar, deps ServiceDeps) {\n    svc := &Service{\n        // highlight-start\n        // deps.db is a postgresql client, attached to the right server and database.\n        db: deps.db,\n        // highlight-end\n    }\n    // ...\n}\n")),(0,o.kt)("p",null,"As you deploy to production later on, a database will also be instantiated. And\nas other developers run your service, they'll also automatically get a managed\ndatabase instance. No more managing dependencies manually."),(0,o.kt)("p",null,"You can explore more of the functionality by running ",(0,o.kt)("inlineCode",{parentName:"p"},"ns help use"),". For example,\nrun ",(0,o.kt)("inlineCode",{parentName:"p"},"ns use psql")," to jump into a ",(0,o.kt)("inlineCode",{parentName:"p"},"psql")," shell of your database."))}g.isMDXComponent=!0}}]);