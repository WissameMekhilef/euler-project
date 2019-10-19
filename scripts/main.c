#include <stdio.h>

/**
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
 * The sum of these multiples is 23.
 * 
 * Find the sum of all the multiples of 3 or 5 below 1000.
*/

int * multipleOfUnder(int a, int b){

   static int  r[10]; // Define the array
   int i;

   b = b/2;

   int lastAdded;
   i = 0;

   lastAdded = a;
   while ( lastAdded < b ) {

   }
   for ( i = 1; i <= b; ++i) {
      r[i] = a * i;
      printf( "r[%d] = %d\n", i, r[i]);
   }

   return r;
}

int main(int argc, char** argv) {
   printf("Hello world !");
   
   int *p;
   p = multipleOfUnder(3, 10);

   return 0;
}

