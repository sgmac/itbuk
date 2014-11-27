all: itb
itb:
	cd bin && go build $@.go
clean:
	@if [ -e bin/itb ];\
		then \
		echo "Removing bin/itb"; \
		rm bin/itb; \
	fi
