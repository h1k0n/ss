#include <riscv_vector.h>
#include <stddef.h>
#include <iostream>

void matmul_rvv(const float *A, const float *B, float *C, int n) {
    for(int i = 0; i < n; i++) {
        for(int j = 0; j < n; j++) {
            vfloat32m1_t acc = __riscv_vfmv_v_f_f32m1(0.0f, 1);
            for(int k = 0; k < n; ) {
                size_t vl = __riscv_vsetvl_e32m1(n - k);
                vfloat32m1_t a_vec = __riscv_vle32_v_f32m1(&A[i*n + k], vl);
                vfloat32m1_t b_vec = __riscv_vlse32_v_f32m1(&B[k*n + j], (ptrdiff_t)(n*sizeof(float)), vl);
                acc = __riscv_vfmacc_vv_f32m1(acc, a_vec, b_vec, vl);
                k += vl;
            }
            vfloat32m1_t sum = __riscv_vfredusum_vs_f32m1_f32m1(__riscv_vfmv_v_f_f32m1(0.0f, 1), acc, 1);
	    float sum1 = __riscv_vfmv_f_s_f32m1_f32(sum);
            C[i*n + j] = sum1;
        }
    }
}

int main() {
    const int n = 4;
    float A[n*n], B[n*n], C[n*n];

    // Initialize A and B
    for (int i = 0; i < n*n; i++) {
        A[i] = static_cast<float>(i);
        B[i] = static_cast<float>(i);
    }

    // Compute matrix multiplication
    matmul_rvv(A, B, C, n);

    // Print the result
    for(int i = 0; i < n; i++) {
        for(int j = 0; j < n; j++) {
            std::cout << C[i*n + j] << " ";
        }
        std::cout << std::endl;
    }

    return 0;
}
