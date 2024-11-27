"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[2123],{1332:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>c,contentTitle:()=>s,default:()=>u,frontMatter:()=>a,metadata:()=>n,toc:()=>l});const n=JSON.parse('{"id":"how-to/create_project","title":"Creating a New Project","description":"When you log in to Minder without a project, Minder will automatically create a","source":"@site/docs/how-to/create_project.md","sourceDirName":"how-to","slug":"/how-to/create_project","permalink":"/how-to/create_project","draft":false,"unlisted":false,"tags":[],"version":"current","sidebarPosition":90,"frontMatter":{"title":"Creating a New Project","sidebar_position":90},"sidebar":"minder","previous":{"title":"Setting up a profile for GitHub Security Advisories","permalink":"/how-to/setup-alerts"},"next":{"title":"Writing custom rule types","permalink":"/how-to/custom-rules"}}');var i=r(74848),o=r(28453);const a={title:"Creating a New Project",sidebar_position:90},s=void 0,c={},l=[{value:"Prerequisites",id:"prerequisites",level:2},{value:"Creating a New Project",id:"creating-a-new-project",level:2}];function d(e){const t={a:"a",code:"code",em:"em",h2:"h2",li:"li",p:"p",ul:"ul",...(0,o.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsxs)(t.p,{children:["When you log in to Minder without a project, Minder will automatically create a\nnew project to manage your entities (repositories, artifacts, etc). It is also\npossible to create additional projects after you have created your Minder\nprofile, for example, to create different projects for different organizations\nor teams you are a part of. Note that\n",(0,i.jsx)(t.a,{href:"/understand/projects",children:"Minder Projects"})," can collect resources from several\nupstream resource providers such as different GitHub organizations, so you can\nregister several ",(0,i.jsx)(t.a,{href:"/understand/providers",children:"entity providers"})," within a\nproject."]}),"\n",(0,i.jsx)(t.h2,{id:"prerequisites",children:"Prerequisites"}),"\n",(0,i.jsxs)(t.ul,{children:["\n",(0,i.jsx)(t.li,{children:"A Minder account"}),"\n",(0,i.jsx)(t.li,{children:"A GitHub organization you are an administrator for which does not have the\nMinder app installed on."}),"\n"]}),"\n",(0,i.jsx)(t.h2,{id:"creating-a-new-project",children:"Creating a New Project"}),"\n",(0,i.jsxs)(t.p,{children:["To create a new project, enable a minder ",(0,i.jsx)(t.a,{href:"/run_minder_server/config_provider",children:"GitHub application"}),"\non a GitHub organization or user account. If the GitHub App is installed on a GitHub organization\nwhich is not already registered in Minder, Minder will create a new project to\nmanage those resources. Using ",(0,i.jsx)(t.a,{href:"/ref/cli/minder_provider_enroll",children:(0,i.jsx)(t.code,{children:"minder provider enroll"})})," within a\nproject to add a new GitHub provider will ",(0,i.jsx)(t.em,{children:"not"})," create a new project and will\nadd the selected organization to an existing project."]})]})}function u(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},28453:(e,t,r)=>{r.d(t,{R:()=>a,x:()=>s});var n=r(96540);const i={},o=n.createContext(i);function a(e){const t=n.useContext(o);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function s(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:a(e.components),n.createElement(o.Provider,{value:t},e.children)}}}]);