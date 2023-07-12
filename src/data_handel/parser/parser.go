package parser

type parser struct {
	SchemaToEntity  schemaToEntity
	EntityToDbModel entityToDbModel
	HttpPars        httpPars
}

type schemaToEntity struct{}

type entityToDbModel struct{}

type httpPars struct{}

func New() parser {
	return parser{
		SchemaToEntity:  schemaToEntity{},
		EntityToDbModel: entityToDbModel{},
		HttpPars:        httpPars{},
	}
}
