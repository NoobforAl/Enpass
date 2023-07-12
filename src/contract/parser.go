package contract

type Parser interface {
	ParserPassword
	ParserService
	ParserUser
	httpReq
}
