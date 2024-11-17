package main

type Board struct{
	Turn string 
	GameRunning bool
    Winner string
	Board [3][3]string
}