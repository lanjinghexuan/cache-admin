package cache

import (
	"cache-admin/examples/cache"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name  string
		page  int32
		limit int32
	}{
		{"测试1", 1, 10},
		{"测试2", 2, 10},
		{"测试3", 3, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.GetUser(tt.page, tt.limit)
			if err != nil {
				t.Errorf("cache.GetUser() error = %v", err)
			}
			fmt.Println(got)

		})
	}
}
