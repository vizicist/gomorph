MORPHOBJM1 = sensel.o \
	sensel_comm_mutex.o \
	sensel_decompress.o \
	sensel_register.o \
	sensel_serial_linux.o \
	sensel_watchdog_linux.o

morphtestm1 : morphtest.o $(MORPHOBJM1)
	gcc morphtest.o $(MORPHOBJM1) -o morphtestm1

MORPHLIBX64 = ../SenselLib/x64/LibSensel.lib

morphtestx64.exe : morphtest.c 
	gcc -I ../SenselLib/include morphtest.c $(MORPHLIBX64) -o morphtestx64

morphtest.o : morphtest.c
	gcc -c -I . morphtest.c

runtestm1 :
	make clean
	make morphtestm1
	./morphtestm1

runtestx64 :
	make clean
	make morphtestx64
	./morphtestx64

clean : 
	rm -f morphtest.o morphtest
