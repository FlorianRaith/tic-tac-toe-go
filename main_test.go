package main

import "testing"

func TestPlayerXWins(t *testing.T) {
	board := Board{
		{'X', 'O', '_'},
		{'_', 'X', '_'},
		{'_', '_', 'X'},
	}

	if !checkWin(0, board) {
		t.Error("Expected true, got false")
	}

	if checkWin(1, board) {
		t.Error("Expected false, got true")
	}

	board = Board{
		{'_', '_', 'X'},
		{'_', 'X', '_'},
		{'X', 'O', '_'},
	}

	if !checkWin(0, board) {
		t.Error("Expected true, got false")
	}

	if checkWin(1, board) {
		t.Error("Expected false, got true")
	}

	board = Board{
		{'X', 'X', 'X'},
		{'_', 'O', '_'},
		{'_', '_', 'O'},
	}

	if !checkWin(0, board) {
		t.Error("Expected true, got false")
	}

	if checkWin(1, board) {
		t.Error("Expected false, got true")
	}

	board = Board{
		{'X', 'O', '_'},
		{'X', 'O', '_'},
		{'X', '_', '_'},
	}

	if !checkWin(0, board) {
		t.Error("Expected true, got false")
	}

	if checkWin(1, board) {
		t.Error("Expected false, got true")
	}
}

func TestPlayerOWins(t *testing.T) {
	board := Board{
		{'O', 'X', '_'},
		{'_', 'O', '_'},
		{'_', '_', 'O'},
	}

	if checkWin(0, board) {
		t.Error("Expected false, got true")
	}

	if !checkWin(1, board) {
		t.Error("Expected true, got false")
	}
}

func TestNoWinner(t *testing.T) {
	board := Board{
		{'X', 'O', '_'},
		{'_', '_', '_'},
		{'X', 'O', 'X'},
	}

	if checkWin(0, board) {
		t.Error("Expected false, got true")
	}

	if checkWin(1, board) {
		t.Error("Expected false, got true")
	}
}

func TestCheckSequenceTrue(t *testing.T) {
	if !checkSequence(0, 'X', 'X', 'X') {
		t.Error("Expected true, got false")
	}

	if !checkSequence(1, 'O', 'O', 'O') {
		t.Error("Expected true, got false")
	}
}

func TestCheckSequenceFalse(t *testing.T) {
	if checkSequence(0, 'X', 'O', 'X') {
		t.Error("Expected false, got true")
	}

	if checkSequence(0, 'X', 'X', '_') {
		t.Error("Expected false, got true")
	}

	if checkSequence(0, 'O', 'O', 'O') {
		t.Error("Expected false, got true")
	}

	if checkSequence(1, 'X', 'X', 'X') {
		t.Error("Expected false, got true")
	}
}

func TestPlayerChar(t *testing.T) {
	if playerChar(0) != 'X' {
		t.Error("Expected 'X', got 'O'")
	}

	if playerChar(1) != 'O' {
		t.Error("Expected 'O', got 'X'")
	}
}

func TestValidateCoordinates(t *testing.T) {
	board := Board{
		{'X', 'O', '_'},
		{'_', '_', '_'},
		{'_', '_', '_'},
	}

	if validateCoordinates(board, 0, 0) {
		t.Error("Expected false, got true")
	}

	if !validateCoordinates(board, 1, 1) {
		t.Error("Expected true, got false")
	}

	if validateCoordinates(board, -1, 0) {
		t.Error("Expected false, got true")
	}

	if validateCoordinates(board, 0, -1) {
		t.Error("Expected false, got true")
	}

	if validateCoordinates(board, 3, 0) {
		t.Error("Expected false, got true")
	}

	if validateCoordinates(board, 0, 3) {
		t.Error("Expected false, got true")
	}
}
