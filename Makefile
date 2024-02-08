BINARY_NAME=mycalculator
setvar:
	export ORDER_SVC_URL="http://localhost:8081/orders"	
#	http://order-service.<namespace>.svc.cluster.local:<port>/endpoint.
build:
	go build ./user-service/cmd
	mv cmd ./user-service/bin
	go build ./order-service/cmd
	mv cmd ./order-service/bin
	go build ./payment-service/cmd
	mv cmd ./payment-service/bin

runusc: setvar
	./user-service/bin/cmd
runosc: 	
	./order-service/bin/cmd
runpsc: 	
	./payment-service/bin/cmd

clean:
	//go clean
	//rm -f $(BINARY_NAME)