package cmd

import (
	"context"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/outputparser"
	"github.com/tmc/langchaingo/prompts"
)

func Explain(util string) (string, error) {
	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		return "", err
	}

	const _conversationTemplate = `The following is a friendly conversation between a human and an AI. The AI is talkative and provides lots of specific details from its context. If the AI does not know the answer to a question, it truthfully says it does not know.
		{{.history}}
		{{.input}}
		`

	prompt := prompts.NewPromptTemplate(
		_conversationTemplate,
		[]string{"input", "history"},
	)

	chain := chains.LLMChain{
		Prompt:       prompt,
		LLM:          llm,
		Memory:       memory.NewConversationTokenBuffer(llm, 2000),
		OutputParser: outputparser.NewSimple(),
		OutputKey:    "text",
	}

	target := "What does the bash function " + util + " do?"
	output, err := chains.Call(ctx, chain, map[string]any{
		"input": target,
	})
	if err != nil {
		return "", err
	}
	return output["text"].(string), nil
}

func init() {
	utilMap = make(map[string]bool)
	for _, util := range utils {
		utilMap[util] = true
	}
}

func lookupUtil(util string) bool {
	val, ok := utilMap[util]
	return ok && val
}

var utilMap map[string]bool
var utils = []string{
	"chgrp",
	"chown",
	"chmod",
	"cp",
	"dd",
	"df",
	"dir",
	"dircolors",
	"install",
	"ln",
	"ls",
	"mkdir",
	"mkfifo",
	"mknod",
	"mktemp",
	"mv",
	"realpath",
	"rm",
	"rmdir",
	"shred",
	"sync",
	"touch",
	"truncate",
	"vdir",
	"b2sum",
	"base32",
	"base64",
	"cat",
	"cksum",
	"comm",
	"csplit",
	"cut",
	"expand",
	"fmt",
	"fold",
	"head",
	"join",
	"md5sum",
	"nl",
	"numfmt",
	"od",
	"paste",
	"ptx",
	"pr",
	"sha1sum ,",
	"sha224sum ,",
	"sha256sum ,",
	"sha384sum ,",
	"sha512sum",
	"shuf",
	"sort",
	"split",
	"sum",
	"tac",
	"tail",
	"tr",
	"tsort",
	"unexpand",
	"uniq",
	"wc",
	"arch",
	"basename",
	"chroot",
	"date",
	"dirname",
	"du",
	"echo",
	"env",
	"expr",
	"factor",
	"false",
	"groups",
	"hostid",
	"id",
	"link",
	"logname",
	"nice",
	"nohup",
	"nproc",
	"pathchk",
	"pinky",
	"printenv",
	"printf",
	"pwd",
	"readlink",
	"runcon",
	"seq",
	"sleep",
	"stat",
	"stdbuf",
	"stty",
	"tee",
	"test",
	"timeout",
	"true",
	"tty",
	"uname",
	"unlink",
	"uptime",
	"users",
	"who",
	"whoami",
	"yes",
}
