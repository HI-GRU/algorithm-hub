# [Gold II] 늑대 사냥꾼 - 2917 

[문제 링크](https://www.acmicpc.net/problem/2917) 

### 성능 요약

메모리: 15472 KB, 시간: 40 ms

### 분류

그래프 이론, 그래프 탐색, 너비 우선 탐색, 최단 경로, 데이크스트라, 격자 그래프

### 제출 일자

2025년 7월 27일 15:50:50

### 문제 설명

<p>늑대 현우는 피에 굶주린 사냥꾼들에게 벗어나려고 도망치고 있다. 사냥꾼은 매우 똑똑해서 나무 뒤에 숨어있다. 현우도 이 사실을 알고있다. 하지만, 어떤 나무 뒤에 사냥꾼들이 숨어있는지 알지 못한다. 현우는 사냥꾼에게 잡히지 않기 위해서 숲의 한 오두막으로 대피하려고 한다. 현우는 나무와 거리가 최대한 떨어지는 경로로 대피하려고 한다.</p>

<p>숲은 N×M 크기의 그리드로 나타낼 수 있다. 각 칸이 빈 목초지라면 '.', 나무가 있다면 '+', 현재 현우의 위치는 'V', 오두막의 위치는 'J'로 나타나있다. 현우는 현재 있는 위치에서 동서남북으로 인접한 칸으로 이동할 수 있고, 나무가 있는 칸으로도 이동할 수 있다.</p>

<p>만약 현우가 지금 R행 C열에 있고, 나무가 A행 B열에 있다면, 현우와 나무의 거리는 |R-A| + |C-B|이다.</p>

<p>현우는 오두막으로 도망치는 가장 안전한 길로 이동하려고 한다. 가장 안전한 길이란 현우가 이동하는 모든 칸에서 나무와 거리의 최솟값이 가장 큰 경로이다. 또, 오두막이 있는 칸도 경로의 일부이기 때문에 나무와 거리를 계산해야 한다.</p>

<p>숲의 지도가 주어졌을 때, 가장 안전한 길을 찾는 프로그램을 작성하시오.</p>

### 입력 

 <p>첫째 줄에 N과 M (1 ≤ N, M ≤ 500)이 주어진다. 둘째 줄부터 N개 줄에는 숲의 지도가 주어진다. 지도에 'V'와 'J'는 딱 하나만 있고, 적어도 하나의 '+'가 있다.</p>

### 출력 

 <p>첫째 줄에 가장 안전한 경로에서 나무와 현우와 거리의 최솟값을 출력한다.</p>

