# Phoebe Makefile

CC=g++
CFLAGS=-O2 -Wall
SOURCES=main.cpp ray.cpp
OBJECTS=$(SOURCES:.cpp=.o)
EXECUTABLE=phoebe

all: $(SOURCES) $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
    $(CC) $(LDFLAGS) $(OBJECTS) -o $@

.cpp.o:
    $(CC) $(CFLAGS) $< -o $@ 
