#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define SIZE 100

typedef struct _Data {
  int value;
  unsigned char payload[SIZE];
} Data;

int compare_by_reference(const Data *a, const Data *b) {
  if (a->value < b->value)
    return -1;
  if (a->value > b->value)
    return 1;
  return 0;
}

int compare_by_value(Data a, Data b) {
  if (a.value < b.value)
    return -1;
  if (a.value > b.value)
    return 1;
  return 0;
}

void bblsort_by_reference(Data *a[], size_t len) {
  for (int i = 0; i < len; i++) {
    for (int j = 1; j < len - i; j++) {
      if (compare_by_reference(a[j], a[j - 1]) < 0) {
        Data *tmp = a[j];
        a[j] = a[j - 1];
        a[j - 1] = tmp;
      }
    }
  }
}

void bblsort_by_value(Data a[], size_t len) {
  for (int i = 0; i < len; i++) {
    for (int j = 1; j < len - i; j++) {
      if (compare_by_value(a[j], a[j - 1]) < 0) {
        Data tmp = a[j];
        a[j] = a[j - 1];
        a[j - 1] = tmp;
      }
    }
  }
}

#define LEN 10000
int main(int argc, char **argv) {
  Data data1[LEN];
  Data *data2[LEN];
  srand(82749522);
  for (int i = 0; i < LEN; i++) {
    data2[i] = (Data *)malloc(sizeof(Data));
    data2[i]->value = rand();
    data1[i] = *data2[i];
  }

  clock_t t0 = clock();
  bblsort_by_value(data1, LEN);
  clock_t t1 = clock();
  printf("by value    : %f[sec]\n", (double)(t1 - t0) / CLOCKS_PER_SEC);

  clock_t t2 = clock();
  bblsort_by_reference(data2, LEN);
  clock_t t3 = clock();
  printf("by reference: %f[sec]\n", (double)(t3 - t2) / CLOCKS_PER_SEC);

  for (int i = 0; i < LEN; i++) {
    free(data2[i]);
  }
  return EXIT_SUCCESS;
}