#include <iostream>
#include <stdlib.h>
#include <sstream>
#include <vector>


using namespace std;

vector<int> generate_fib_sequence(int);

void print_sequence(vector<int>);

int main(int argc, char *argv[]) {

    if (argc <= 1) {
        cout << "1";
    }
    else {
        int length = atoi(argv[1]);

        print_sequence(generate_fib_sequence(length));

    }

    return 0;
}

void print_sequence(vector<int> sequence) {

    for (vector<int>::size_type i = 0; i < sequence.size(); i++) {
        cout <<  sequence.at(i) << endl;
    }

}

vector<int> generate_fib_sequence(int length) {

    ostringstream ss;
    ss << length;

    cout << "include " + ss.str() + " digits" << endl;


    vector<int> vSequence;


    for (vector<int>::size_type i = 0; i <= length-1; ++i) {

        if (i <= 1) {
            vSequence.push_back(1);
        }
        else {
            vSequence.push_back( vSequence.at(i - 1) + vSequence.at(i - 2));
        }
    }

    return vSequence;

}