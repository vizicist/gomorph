M1LIB = ../SenselLib/m1/LibSenselStatic.a
INC = ../SenselLib/include

morphtestm1 : morphtest.c $(M1LIB)
	gcc morphtest.c -I$(INC) $(M1LIB) -o morphtestm1

WINX64LIB = ../SenselLib/winx64/LibSensel.lib

WINCC = cl
WINCC = gcc
winx64test.exe : morphtest.c 
	$(WINCC) -I ../SenselLib/include morphtest.c $(WINX64LIB) -o winx64test

m1test :
	rm -f morphtestm1
	make morphtestm1
	./morphtestm1

winx64test :
	del winx64test.exe
	$(MAKE) winx64test.exe
	.\winx64test.exe
