// Pakcage testbeIMPserver
package testbeIMPserver

//go:generate swagger generate server --exclude-main -A test-be-IMP-server -t gen -f ./api/swagger.yml --principal models.Principal
//go:generate swagger generate client -A client-collection -t gen -f ./api/collection.yml
