"use strict";(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[980],{63128:function(a){function b(a){a.languages["firestore-security-rules"]=a.languages.extend("clike",{comment:/\/\/.*/,keyword:/\b(?:allow|function|if|match|null|return|rules_version|service)\b/,operator:/&&|\|\||[<>!=]=?|[-+*/%]|\b(?:in|is)\b/}),delete a.languages["firestore-security-rules"]["class-name"],a.languages.insertBefore("firestore-security-rules","keyword",{path:{pattern:/(^|[\s(),])(?:\/(?:[\w\xA0-\uFFFF]+|\{[\w\xA0-\uFFFF]+(?:=\*\*)?\}|\$\([\w\xA0-\uFFFF.]+\)))+/,lookbehind:!0,greedy:!0,inside:{variable:{pattern:/\{[\w\xA0-\uFFFF]+(?:=\*\*)?\}|\$\([\w\xA0-\uFFFF.]+\)/,inside:{operator:/=/,keyword:/\*\*/,punctuation:/[.$(){}]/}},punctuation:/\//}},method:{pattern:/(\ballow\s+)[a-z]+(?:\s*,\s*[a-z]+)*(?=\s*[:;])/,lookbehind:!0,alias:"builtin",inside:{punctuation:/,/}}})}a.exports=b,b.displayName="firestoreSecurityRules",b.aliases=[]}}])