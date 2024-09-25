list:
	@ \
    echo "quick-sorting" && \
	echo "bubble-sorting"

quick-sorting:
	go run cmd/cli/* -quick-sorting

bubble-sorting:
	go run cmd/cli/* -bubble-sorting