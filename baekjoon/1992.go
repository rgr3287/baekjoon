package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	image := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&image[i])
	}

	result := compressQuadTree(image)
	fmt.Println(result)
}

func compressQuadTree(image []string) string {
	n := len(image)

	// 이미지가 모두 0으로 구성되어 있는 경우 "0" 반환
	if allZeroes(image) {
		return "0"
	}

	// 이미지가 모두 1로 구성되어 있는 경우 "1" 반환
	if allOnes(image) {
		return "1"
	}

	// 이미지를 분할하여 각 부분을 압축하고 괄호로 묶음
	half := n / 2
	topLeft := make([]string, half)
	topRight := make([]string, half)
	bottomLeft := make([]string, half)
	bottomRight := make([]string, half)

	for i := 0; i < half; i++ {
		topLeft[i] = image[i][:half]
		topRight[i] = image[i][half:]
		bottomLeft[i] = image[i+half][:half]
		bottomRight[i] = image[i+half][half:]
	}

	return "(" + compressQuadTree(topLeft) + compressQuadTree(topRight) + compressQuadTree(bottomLeft) + compressQuadTree(bottomRight) + ")"
}

func allZeroes(image []string) bool {
	for _, row := range image {
		if strings.Contains(row, "1") {
			return false
		}
	}
	return true
}

func allOnes(image []string) bool {
	for _, row := range image {
		if strings.Contains(row, "0") {
			return false
		}
	}
	return true
}
