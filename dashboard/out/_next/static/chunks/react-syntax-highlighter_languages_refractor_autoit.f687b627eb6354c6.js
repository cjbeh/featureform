"use strict";(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[8333],{53358:function(a){function b(a){a.languages.autoit={comment:[/;.*/,{pattern:/(^[\t ]*)#(?:comments-start|cs)[\s\S]*?^[ \t]*#(?:ce|comments-end)/m,lookbehind:!0}],url:{pattern:/(^[\t ]*#include\s+)(?:<[^\r\n>]+>|"[^\r\n"]+")/m,lookbehind:!0},string:{pattern:/(["'])(?:\1\1|(?!\1)[^\r\n])*\1/,greedy:!0,inside:{variable:/([%$@])\w+\1/}},directive:{pattern:/(^[\t ]*)#[\w-]+/m,lookbehind:!0,alias:"keyword"},function:/\b\w+(?=\()/,variable:/[$@]\w+/,keyword:/\b(?:Case|Const|Continue(?:Case|Loop)|Default|Dim|Do|Else(?:If)?|End(?:Func|If|Select|Switch|With)|Enum|Exit(?:Loop)?|For|Func|Global|If|In|Local|Next|Null|ReDim|Select|Static|Step|Switch|Then|To|Until|Volatile|WEnd|While|With)\b/i,number:/\b(?:0x[\da-f]+|\d+(?:\.\d+)?(?:e[+-]?\d+)?)\b/i,boolean:/\b(?:False|True)\b/i,operator:/<[=>]?|[-+*\/=&>]=?|[?^]|\b(?:And|Not|Or)\b/i,punctuation:/[\[\]().,:]/}}a.exports=b,b.displayName="autoit",b.aliases=[]}}])