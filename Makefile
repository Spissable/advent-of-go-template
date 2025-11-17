init:
	for folder in day*; do \
		touch $$folder/input.txt; \
	done

day-%:
	go test ./day$* -v
