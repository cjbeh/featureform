"use strict";(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[5259],{43826:function(a){function b(a){a.languages.groovy=a.languages.extend("clike",{string:[{pattern:/("""|''')(?:[^\\]|\\[\s\S])*?\1|\$\/(?:[^/$]|\$(?:[/$]|(?![/$]))|\/(?!\$))*\/\$/,greedy:!0},{pattern:/(["'/])(?:\\.|(?!\1)[^\\\r\n])*\1/,greedy:!0}],keyword:/\b(?:abstract|as|assert|boolean|break|byte|case|catch|char|class|const|continue|def|default|do|double|else|enum|extends|final|finally|float|for|goto|if|implements|import|in|instanceof|int|interface|long|native|new|package|private|protected|public|return|short|static|strictfp|super|switch|synchronized|this|throw|throws|trait|transient|try|void|volatile|while)\b/,number:/\b(?:0b[01_]+|0x[\da-f_]+(?:\.[\da-f_p\-]+)?|[\d_]+(?:\.[\d_]+)?(?:e[+-]?\d+)?)[glidf]?\b/i,operator:{pattern:/(^|[^.])(?:~|==?~?|\?[.:]?|\*(?:[.=]|\*=?)?|\.[@&]|\.\.<|\.\.(?!\.)|-[-=>]?|\+[+=]?|!=?|<(?:<=?|=>?)?|>(?:>>?=?|=)?|&[&=]?|\|[|=]?|\/=?|\^=?|%=?)/,lookbehind:!0},punctuation:/\.+|[{}[\];(),:$]/}),a.languages.insertBefore("groovy","string",{shebang:{pattern:/#!.+/,alias:"comment"}}),a.languages.insertBefore("groovy","punctuation",{"spock-block":/\b(?:and|cleanup|expect|given|setup|then|when|where):/}),a.languages.insertBefore("groovy","function",{annotation:{pattern:/(^|[^.])@\w+/,lookbehind:!0,alias:"punctuation"}}),a.hooks.add("wrap",function(b){if("groovy"===b.language&&"string"===b.type){var c=b.content.value[0];if("'"!=c){var d=/([^\\])(?:\$(?:\{.*?\}|[\w.]+))/;"$"===c&&(d=/([^\$])(?:\$(?:\{.*?\}|[\w.]+))/),b.content.value=b.content.value.replace(/&lt;/g,"<").replace(/&amp;/g,"&"),b.content=a.highlight(b.content.value,{expression:{pattern:d,lookbehind:!0,inside:a.languages.groovy}}),b.classes.push("/"===c?"regex":"gstring")}}})}a.exports=b,b.displayName="groovy",b.aliases=[]}}])