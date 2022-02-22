M1LIB = ../SenselLib/m1/LibSenselStatic.a
INC = ../SenselLib/include

morphtestm1 : morphtest.c $(M1LIB)
	gcc morphtest.c -I$(INC) $(M1LIB) -o morphtestm1

WINX64LIB = ../SenselLib/winx64/LibSensel.lib

winx64test.exe : morphtest.c 
	gcc -I ../SenselLib/include morphtest.c $(WINX64LIB) -o winx64test

m1test :
	make clean
	make morphtestm1
	./morphtestm1

winx64test :
	$(MAKE) clean
	$(MAKE) winx64test.exe
	.\winx64test.exe

clean : 
	rm -f morphtest.o morphtest
