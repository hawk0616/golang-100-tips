package tips21

import "strconv"

// 非効率な配列の初期化

func convert(foos []Foo) ([]Bar,error) {
	bars := make([]Bar, 0)

	for _, foo := range foos {
		bar, err := fooToBar(&foo)
		if err != nil {
			return nil, err
		}
		bars = append(bars, *bar)
	}

	return bars, nil
}
// 最初にBarの初期化
// append を使用してBarの要素を追加している
// 最初が空で次に追加しようとするとGoは大きさ１の規程配列を作成する
// 規程配列がいっぱいになるごとにGoはその容量を二倍にしようとする
// 故に大きなiteretorを使用する際に配列から別の配列へコピーを作成して前の基底配列をGCが削除しなくてはいけない

type Bar struct {
	xxx int
}

type Foo struct {
	yyy string
}

func fooToBar(foo *Foo) (*Bar, error) {
	xxx, err := strconv.Atoi(foo.yyy)
	if err != nil {
		return nil, err
	}
	return &Bar{
		xxx: xxx,
	}, nil
}
