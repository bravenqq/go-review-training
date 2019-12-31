#include<stdio.h>

struct B
{
    char a;     //1
    double b;   //8
    int c;      //4
};

struct A{
    char a;
    int c;
    double d;
};

struct C{
    char a;
    char b;
    int c;
    int d;
    char e;
};

int main(int argc, char *argv[])
{
    struct A x; 
    struct B y;
    struct C z;
    printf("A %d,B %d C %d\n",sizeof(x),sizeof(y),sizeof(z));
    return 0;
}


