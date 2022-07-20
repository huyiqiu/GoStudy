#include<iostream>
using namespace std;

int main() {
	int b = 100;
	int *a = &b;
	int &c = b;
	cout << a << endl;
	cout << b << endl;
	cout << c << endl;
	cout << *a << endl;
	cout << (*a == c) << endl;
	return 0;
}