module github.com/MCDevKit/jsonte

go 1.26.4

require (
	github.com/Bedrock-OSS/go-burrito v1.0.4
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20220911224424-aa1f1f12a846
	github.com/faiface/beep v1.1.0
	github.com/fatih/color v1.16.0
	github.com/gammazero/deque v0.2.1
	github.com/gobwas/glob v0.2.3
	github.com/james-barrow/golang-ipc v1.2.4
	github.com/paul-mannino/go-fuzzywuzzy v0.0.0-20200127021948-54652b135d0e
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d
	github.com/stirante/molang-go v0.0.0-20260917122549-6abea1d70a0d
	github.com/zeebo/xxh3 v1.0.2
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f
	golang.org/x/net v0.24.0
	golang.org/x/text v0.14.0
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	github.com/jfreymuth/oggvorbis v1.0.5 // indirect
	github.com/jfreymuth/vorbis v1.0.2 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.19.0 // indirect
)

// Since Go 1.24 rand.Seed is NOP, we need to bring it back for --seed flag
// TODO: switch to own rand implementation, that replicated that of Go
godebug randseednop=0
