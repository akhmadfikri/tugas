package main
import ("fmt")
type student struct {
	var name = [] string {"budi", "ani", "badu"}
	var score = [] int {40, 50, 70}
}
func (s student) Avarage () float64{
	jumlahTotal= 0
	for index :=0; index< len(s.score); index++{
		jumlahTotal = jumlahTotal + score[index]
	}
	rata := float64(jumlahTotal) / float64(index)
}
func main(){

	var data student
	data.score = [40, 50, 70]
	data.name = ["budi", "ani", "badu"]
	fmt.Println("Average score", rata)
}