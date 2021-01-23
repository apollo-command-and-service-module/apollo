package repo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)



func TestNewClone(t *testing.T) {
	structValue := Repo{Url: "http://github.com/", Branch: "main", ConfigFile: "config.yaml", Since: time.Now()}
	fields := reflect.TypeOf(structValue)
	values := reflect.ValueOf(structValue)

	clone := NewClone(structValue.Url, structValue.Branch, structValue.ConfigFile, structValue.Since)

	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fieldValue := value.Interface()

		switch v := fieldValue.(type) {
		case string:
			fmt.Print( field.Name, "=", v, "\n")
			if clone.Branch != v {
				t.Errorf("Branch was incorrect, got: %s, want: %s.", clone.Branch, v)
			}
		case int:
			fmt.Print( field.Name, "=", v, "\n")
		case int32:
			fmt.Print(field.Name, "=", v, "\n")
		case int64:
			fmt.Print(field.Name, "=", v, "\n")
		case time.Time:
			fmt.Print(field.Name, "=", v, "\n")
			if clone.Since != v {
				t.Errorf("Branch was incorrect, got: %s, want: %s.", clone.Branch, v)
			}
		default:
			assert.Fail(t, "Not support type of struct")
		}
	}

}

func TestRepo_ReadIntoMemory(t *testing.T) {

}

