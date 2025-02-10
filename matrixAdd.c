#include <riscv_vector.h>
#include <stdio.h>

// Matrix dimensions
#define ROWS 4
#define COLS 4

void matrix_add(double *a, double *b, double *c, size_t rows, size_t cols) {
    for (size_t i = 0; i < rows; i++) {
        size_t vl;
        for (size_t j = 0; j < cols; j += vl) {
            // Calculate vector length for this iteration
            vl = __riscv_vsetvl_e64m1(cols - j);
            
            // Load vectors from matrices
            vfloat64m1_t va = __riscv_vle64_v_f64m1(&a[i * cols + j], vl);
            vfloat64m1_t vb = __riscv_vle64_v_f64m1(&b[i * cols + j], vl);
            
            // Perform vector addition
            vfloat64m1_t vc = __riscv_vfadd_vv_f64m1(va, vb, vl);
            
            // Store result
            __riscv_vse64_v_f64m1(&c[i * cols + j], vc, vl);
        }
    }
}

int main() {
    double matrix_a[ROWS][COLS] = {
        {1.0, 2.0, 3.0, 4.0},
        {5.0, 6.0, 7.0, 8.0},
        {9.0, 10.0, 11.0, 12.0},
        {13.0, 14.0, 15.0, 16.0}
    };
    
    double matrix_b[ROWS][COLS] = {
        {1.0, 1.0, 1.0, 1.0},
        {2.0, 2.0, 2.0, 2.0},
        {3.0, 3.0, 3.0, 3.0},
        {4.0, 4.0, 4.0, 4.0}
    };
    
    double result[ROWS][COLS];
    
    // Perform matrix addition
    matrix_add((double *)matrix_a, (double *)matrix_b, (double *)result, ROWS, COLS);
    
    // Print result
    for (int i = 0; i < ROWS; i++) {
        for (int j = 0; j < COLS; j++) {
            printf("%.1f ", result[i][j]);
        }
        printf("\n");
    }
    
    return 0;
}
