# Phoebe Makefile

CC=g++
CFLAGS=-c -O2 -Wall

all: phoebe

phoebe: main.o ray.o
	$(CC) main.o ray.o -o phoebe

main.o: src/main.cpp
	$(CC) $(CFLAGS) src/main.cpp

ray.o: src/ray.cpp
	$(CC) $(CFLAGS) src/ray.cpp

clean:
	rm *.o phoebe
