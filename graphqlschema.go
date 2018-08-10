package terse

const GRAPHQL_SCHEMA = `

schema {
	query: Query
	mutation: Mutation
}

type Query {
	link(short: String, long: String): [Link!]!
}

type Mutation {
	createLink(link: LinkInput!): Link
}

type Link{
	id: ID!
	long: String!
	short: String!
}

input LinkInput{
	long: String!
	short: String!
}

`
