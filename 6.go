package main
import "fmt"
const NMAX int = 1024
type arrBilangan[NMAX]int
func main(){
	var bilBul arrBilangan
	var nBilangan, bilangan int
	nBilangan = 0
	fmt.Scan(&bilangan)
	for bilangan != 0 {
		bilBul[nBilangan] = bilangan
		nBilangan++
		fmt.Scan(&bilangan)
	}
	fmt.Print(cariNilaiMax(bilBul, nBilangan), " ")
	fmt.Print(cariNilaiMin(bilBul, nBilangan))
}
func cariNilaiMax(arrayBilBul arrBilangan, nBilangan int)int{
	var i, max int
	max = arrayBilBul[0]
	for i = 0;i < nBilangan; i++{
		if max < arrayBilBul[i]{
			max = arrayBilBul[i]
		}
	}
	return max
}
func cariNilaiMin(arrayBilBul arrBilangan, nBilangan int)int{
	var i, min int
	min = arrayBilBul[0]
	for i = 0;i < nBilangan; i++{
		if min > arrayBilBul[i]{
			min = arrayBilBul[i]
		}
	}
	return min
}