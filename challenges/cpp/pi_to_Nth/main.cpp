#include <iostream>
#include <sstream>
#include <stdlib.h>

using namespace std;

int parse_args(char *);

string calculate_pi_to_the(int);

const string pi_expo = "14159265358979323846264338327950288419";

int main(int argc, char *argv[]) {
    if (argc <= 1) {
        cout << "\n3";
    }
    else {
        cout << calculate_pi_to_the(parse_args(argv[1]));
    }
}

int parse_args(char *arg) {
    return atoi(arg);
}

string calculate_pi_to_the(int nth) {

    std::ostringstream ss;
    ss << nth;

    cout << "include " + ss.str() + " digits";

    if (nth != 0) {
        return "\n3." + pi_expo.substr(0, nth);
    }
    else {
        return "\n3";
    }
}