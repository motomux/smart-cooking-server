package resource

import (
	"fmt"
	"reflect"
	"strings"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"

	tarantool "github.com/tarantool/go-tarantool"
)

// RecipesRscInterface is an interface to test RecipesRsc
type RecipesRscInterface interface {
	GetOne(ID int) (*Recipe, error)
}

// RecipesRsc provides api to manipulate resouce on tarantool
type RecipesRsc struct {
	client    *tarantool.Connection
	spaceName string
}

// Recipe represents a document on Recipess collection in taratool
type Recipe struct {
	ID    uint     `json:"id"`
	Title string   `json:"title"`
	Photo string   `json:"photo"`
	Howto []string `json:"howto"`
	Video string   `json:"video"`
}

func init() {
	msgpack.Register(reflect.TypeOf(Recipe{}), encodeRecipe, decodeRecipe)
}

// NewRecipesRsc initiates RecipesRsc
func NewRecipesRsc(client *tarantool.Connection) *RecipesRsc {
	return &RecipesRsc{
		client:    client,
		spaceName: "recipes",
	}
}

// GetOne finds one document on MongoDB with RecipeID
func (rsc *RecipesRsc) GetOne(ID int) (*Recipe, error) {
	var recipes []Recipe
	err := rsc.client.SelectTyped(rsc.spaceName, "primary", 0, 1, tarantool.IterEq, []interface{}{ID}, &recipes)
	if err != nil {
		return nil, err
	}

	return &recipes[0], nil
}

func encodeRecipe(e *msgpack.Encoder, v reflect.Value) error {
	m := v.Interface().(Recipe)
	if err := e.EncodeSliceLen(5); err != nil {
		return err
	}
	if err := e.EncodeUint(m.ID); err != nil {
		return err
	}
	if err := e.EncodeString(m.Title); err != nil {
		return err
	}
	if err := e.EncodeString(m.Photo); err != nil {
		return err
	}
	if err := e.EncodeString(strings.Join(m.Howto[:], ",")); err != nil {
		return err
	}
	if err := e.EncodeString(m.Video); err != nil {
		return err
	}
	return nil
}

func decodeRecipe(d *msgpack.Decoder, v reflect.Value) error {
	var err error
	var l int
	m := v.Addr().Interface().(*Recipe)
	if l, err = d.DecodeSliceLen(); err != nil {
		return err
	}
	if l != 5 {
		return fmt.Errorf("array len doesn't match: %d", l)
	}
	if m.ID, err = d.DecodeUint(); err != nil {
		return err
	}
	if m.Title, err = d.DecodeString(); err != nil {
		return err
	}
	if m.Photo, err = d.DecodeString(); err != nil {
		return err
	}
	if howtos, err := d.DecodeString(); err == nil {
		m.Howto = strings.Split(howtos, ",")
	} else {
		return err
	}
	if m.Video, err = d.DecodeString(); err != nil {
		return err
	}
	return nil
}
