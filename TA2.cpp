#include <bits/stdc++.h>
#include <omp.h>
#define N 3
#define N1 2
#define N2 2
#define N3 3
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
		C[i][j] = A[i][j]*B[i][j];
	}
	cout << "Resultado Producto Punto:" << endl;
	imprimir(C);
}
void producto_cruz(int arr[]) {
	int A[N1][N2] = {
		{2,1},
		{1,0}
	};
	int B[N2][N3] = {
		{1,2,3},
		{3,2,1}
	};
	int C[N1][N3]= {
		{0,0,0},
		{0,0,0}
	};
	#pragma omp parallel num_threads(N1*N3)
	{
		int id = omp_get_thread_num();
		int i = id%N2;
		int j = id/N2;
		for(int p=0; p<N2; p++) C[i][j]+=A[i][p]*B[p][j];
	}
	
	cout << "Resultado Producto Cruz:" << endl;
	for(int i=0; i<N1; i++) {
		for(int j=0; j<N3; j++) cout << C[i][j] << " ";
		cout << endl;
	}
	cout << endl;
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
	producto_cruz();	
	
	return 0;
}
