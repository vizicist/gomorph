MORPHOBJ = sensel.o \
	sensel_comm_mutex.o \
	sensel_decompress.o \
	sensel_register.o \
	sensel_serial_linux.o \
	sensel_watchdog_linux.o

morphtest : morphtest.o $(MORPHOBJ)
	gcc morphtest.o $(MORPHOBJ) -o morphtest

morphtest.o : morphtest.c
	gcc -c -I . morphtest.c

runtest :
	make clean
	make morphtest
	./morphtest

clean : 
	rm -f morphtest.o morphtest
