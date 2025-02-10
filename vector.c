#include <stdio.h>
#include <riscv_vector.h>

#define N 16

// 向量加法函数
void vector_add(int *a, int *b, int *c, int n) {
    // 每次循环处理的元素数量由VL决定
    size_t vl;
    for (; n > 0; n -= vl, a += vl, b += vl, c += vl) {
        // 设置向量长度（自动适配硬件支持的最大长度）
        vl = __riscv_vsetvl_e32m1(n);
        
        // 加载向量a和b
        vint32m1_t va = __riscv_vle32_v_i32m1(a, vl);
        vint32m1_t vb = __riscv_vle32_v_i32m1(b, vl);
        
        // 向量相加
        vint32m1_t vc = __riscv_vadd_vv_i32m1(va, vb, vl);
        
        // 存储结果到c
        __riscv_vse32_v_i32m1(c, vc, vl);
    }
}

int main() {
    int a[N] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15};
    int b[N] = {15,14,13,12,11,10,9,8,7,6,5,4,3,2,1,0};
    int c[N] = {0};

    vector_add(a, b, c, N);

    printf("Vector addition result:\n");
    for (int i = 0; i < N; i++) {
        printf("%d + %d = %d\n", a[i], b[i], c[i]);
    }

    return 0;
}
