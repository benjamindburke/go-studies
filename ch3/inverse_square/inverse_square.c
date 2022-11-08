#include <math.h>
#include <stdio.h>

float q_rsqrt(float number, unsigned int magic_number) {
  long i;
  float x2, y;
  const float threehalfs = 1.5F;

  x2 = number * 0.5F;
  y = number;

  i = * ( long * ) &y;
  i = magic_number - (i >> 1);
  y = * ( float * ) &i;

  // newtonian calculus approximation function
  y = y * ( threehalfs - ( x2 * y * y) );
  y = y * ( threehalfs - ( x2 * y * y) ); // approximate a second time if margin of error is not low enough

  return y;
}

float error_margin(float known, float calculated) {
  return fabs((calculated / known) - 1.0F) * 100.0F;
}

void compare(float known_inverse_sqrt, float num, unsigned int magic_number) {
  float inverse_sqrt = q_rsqrt(num, magic_number);
  float error = error_margin(known_inverse_sqrt, inverse_sqrt);
  printf("num=%.2f, inverse_square=%.15f\n", num, inverse_sqrt);
  printf("known_inverse_square=%.2f, error_margin=%.15f%%\n", known_inverse_sqrt, error);
}

int main(void) {
  float num = 4.0F;
  float known_inverse_sqrt = 0.5F;

  compare(known_inverse_sqrt, num, 0x5f3759df);
  compare(known_inverse_sqrt, num, 0x5f375a86);

  return 0;
}