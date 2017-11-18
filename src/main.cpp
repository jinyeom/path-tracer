#include <iostream>

#define WINDOW_WIDTH 800
#define WINDOW_HEIGHT 600

int main(int argc, char** argv)
{
    std::cout << "Phoebe: A Simple Ray Tracer" << std::endl
              << "Copyright (c) 2017 by Jin Yeom" << std::endl;

    // Trace for each pixel of the screen.
    for (int i = 0; i < WINDOW_WIDTH; ++i) {
        for (int j = 0; i < WIDNOW_HEIGHT; ++j) {
            // TODO: get pixel value
        }
    } 

    return EXIT_SUCCESS;
}
