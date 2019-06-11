#include <bits/stdc++.h>
#include <omp.h>
#define N 3
using namespace std;
// parametros de lanzamiento '-std=c++14 -fopenmp'
void imprimir(int C[N][N]) {
	for(int i=0; i<N; i++) {
		for(int j=0; j<N; j++) cout << C[i][j] << " ";
		cout << endl;
	}
	cout << endl;
}
void suma(int A[N][N], int B[N][N]) {
	int C[N][N] = {{0,0,0},{0,0,0},{0,0,0}};
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		C[i][j] = A[i][j]+B[i][j];
	}
	cout << "Resultado Suma:" << endl;
	imprimir(C);
}
void resta(int A[N][N], int B[N][N]) {
	int C[N][N] = {{0,0,0},{0,0,0},{0,0,0}};
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		C[i][j] = A[i][j]-B[i][j];
	}
	cout << "Resultado Resta:" << endl;
	imprimir(C);
}
void multi_escalar(int A[N][N], double E) {
	int C[N][N] = {{0,0,0},{0,0,0},{0,0,0}};
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		C[i][j] = A[i][j]*E;
	}
	cout << "Resultado Multi. Escalar:" << endl;
	imprimir(C);
}
void producto_punto(int A[N][N], int B[N][N]) {
	int C[N][N] = {{0,0,0},{0,0,0},{0,0,0}};
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		for(int p=0; p<N; p++) C[i][j]+=A[i][p]*B[p][j];
	}
	cout << "Resultado Producto Punto:" << endl;
	imprimir(C);
}
void producto_cruz(int A[N][N], int B[N][N]) {
	int C[N][N] = {{0,0,0},{0,0,0},{0,0,0}};
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		// ...
	}
	cout << "Resultado Producto Cruz:" << endl;
	imprimir(C);
}

int main() {
	int A[N][N] = {
		{1,2,3},
		{4,5,6},
		{7,8,9}
	};
	int B[N][N] = {
		{1,0,2},
		{0,1,2},
		{1,0,0}
	};
	double E = 4;
	suma(A,B);
	resta(A,B);
	multi_escalar(A,E);
	producto_punto(A,B);
	producto_cruz(A,B);
	return 0;
}
