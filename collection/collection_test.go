/*
 * @author: onno
 * @date: 2020/4/30
 */
package collection

import (
	"fmt"
	"testing"
)

func TestToSlice(t *testing.T) {
	t.Run("toSlice", func(t *testing.T) {
		result := ToSlice(1, 2, 3)
		fmt.Println(result)
		if len(result) != 3 {
			t.Error("len not 3")
		}
	})
}
