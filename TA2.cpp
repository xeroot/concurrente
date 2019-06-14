#include <bits/stdc++.h>
#include <omp.h>
using namespace std;
// parametros de lanzamiento '-std=c++14 -fopenmp'

void imprimir(int** C, int N, int M) {
	for(int i=0; i<N; i++) {
		for(int j=0; j<M; j++) cout << C[i][j] << " ";
		cout << endl;
	}
	cout << endl;
}

int** create_matrix(int N,int M){
	int **matriz = new int*[N];
	for(int i=0;i<M;i++) matriz[i]= new int[N];
	for(int i=0;i<N*M;i++) cin >> matriz[i/M][i%M];
	imprimir(matriz, N, M);
	return matriz;
}

int** get_empty_matrix(int N,int M){
	int **matriz = new int*[N];
	for(int i=0;i<M;i++) matriz[i]= new int[N];
	for(int i=0;i<N*M;i++) matriz[i/M][i%M] = 0;
	return matriz;
}

void suma() {
	int N;
	cout << "N: ";
	cin >> N;
	cout << "Matriz A ("<<N*N<<"): ";
	int **A = create_matrix(N,N);
	cout << "Matriz B ("<<N*N<<"): ";
	int **B = create_matrix(N,N);
	int **C = get_empty_matrix(N,N);
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		C[i][j] = A[i][j]+B[i][j];
	}
	cout << "Resultado Suma:" << endl;
	imprimir(C, N, N);
}

void resta() {
	int N;
	cout << "N: ";
	cin >> N;
	cout << "Matriz A ("<<N*N<<"): ";
	int **A = create_matrix(N,N);
	cout << "Matriz B ("<<N*N<<"): ";
	int **B = create_matrix(N,N);
	int **C = get_empty_matrix(N,N);
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		C[i][j] = A[i][j]-B[i][j];
	}
	cout << "Resultado Resta:" << endl;
	imprimir(C, N, N);
}
void multi_escalar() {
	int N;
	cout << "N: ";
	cin >> N;
	cout << "Matriz A ("<<N*N<<"): ";
	int **A = create_matrix(N,N);
	cout << "Punto escalar: ";
	int E; cin >> E;
	int **C = get_empty_matrix(N,N);
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		C[i][j] = A[i][j]*E;
	}
	cout << "Resultado Mult. Escalar:" << endl;
	imprimir(C, N, N);
}
void producto_punto() {
	int N;
	cout << "N: ";
	cin >> N;
	cout << "Matriz A ("<<N*N<<"): ";
	int **A = create_matrix(N,N);
	cout << "Matriz B ("<<N*N<<"): ";
	int **B = create_matrix(N,N);
	int **C = get_empty_matrix(N,N);
	#pragma omp parallel num_threads(N*N)
	{
		int id = omp_get_thread_num();
		int i = id%N;
		int j = id/N;
		C[i][j] = A[i][j]*B[i][j];
	}
	cout << "Resultado Producto Punto:" << endl;
	imprimir(C, N, N);
}

void producto_cruz() {
	int N1, N2, N3;
	cout << "N1, N2 y N3: ";
	cin >> N1 >> N2 >> N3;
	cout << "Matriz A ("<<N1*N2<<"): ";
	int **A = create_matrix(N1,N2);
	cout << "Matriz B ("<<N2*N3<<"): ";
	int **B = create_matrix(N2,N3);
	int **C = get_empty_matrix(N1,N3);
	#pragma omp parallel num_threads(N1*N3)
	{
		int id = omp_get_thread_num();
		int i = id%N2;
		int j = id/N2;
		for(int p=0; p<N2; p++) C[i][j]+=A[i][p]*B[p][j];
	}
	
	cout << "Resultado Producto Cruz:" << endl;
	imprimir(C, N1, N3);
}

void menu(){
	cout << "Opciones:" << endl;
	cout << "1. Suma" << endl;
	cout << "2. Resta" << endl;
	cout << "3. Mult. por Escalar" << endl;
	cout << "4. Prod. Punto" << endl;
	cout << "5. Prod. Cruz" << endl;
}

int main() {
	menu();
	int opcion=0;
	do{
		cout << "Escoja opcion: ";
		cin >> opcion;
	}while(opcion < 1 || opcion > 5);
	switch(opcion){
		case 1:
			suma();
		break;
		case 2:
			resta();
		break;
		case 3:
			multi_escalar();
		break;
		case 4:
			producto_punto();
		break;
		case 5:
			producto_cruz();
		break;
	}
	return 0;
}
