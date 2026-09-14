package test

import (
	"testing"

	"github.com/sudzekai-web-os/mediator"
)

type IntQuery struct {
	Value int
}

type StringQuery struct {
	Value string
}

type BoolQuery struct {
	Value bool
}

type EmptyQuery struct{}

type IntQueryHandler struct{}

func (IntQueryHandler) Handle(query IntQuery) int {
	return query.Value * 2
}

type StringQueryHandler struct{}

func (StringQueryHandler) Handle(query StringQuery) string {
	return "result: " + query.Value
}

type BoolQueryHandler struct{}

func (BoolQueryHandler) Handle(query BoolQuery) bool {
	return !query.Value
}

type EmptyQueryHandler struct{}

func (EmptyQueryHandler) Handle(EmptyQuery) int {
	return 42
}

func TestDispatch(t *testing.T) {
	mediator.RegisterHandler(IntQueryHandler{})

	query := IntQuery{Value: 10}

	var result int

	err := mediator.Dispatch(query, &result)
	if err != nil {
		t.Fatalf("Dispatch вернул ошибку: %v", err)
	}

	if result != 20 {
		t.Fatalf("ожидалось 20, получено %d", result)
	}
}

func TestDispatchStringResult(t *testing.T) {
	mediator.RegisterHandler(StringQueryHandler{})

	query := StringQuery{
		Value: "hello",
	}

	var result string

	err := mediator.Dispatch(query, &result)
	if err != nil {
		t.Fatalf("Dispatch вернул ошибку: %v", err)
	}

	expected := "result: hello"

	if result != expected {
		t.Fatalf("ожидалось %q, получено %q", expected, result)
	}
}

func TestDispatchBoolResult(t *testing.T) {
	mediator.RegisterHandler(BoolQueryHandler{})

	query := BoolQuery{
		Value: true,
	}

	var result bool

	err := mediator.Dispatch(query, &result)
	if err != nil {
		t.Fatalf("Dispatch вернул ошибку: %v", err)
	}

	if result != false {
		t.Fatalf("ожидалось false, получено %t", result)
	}
}

func TestDispatchEmptyQuery(t *testing.T) {
	mediator.RegisterHandler(EmptyQueryHandler{})

	query := EmptyQuery{}

	var result int

	err := mediator.Dispatch(query, &result)
	if err != nil {
		t.Fatalf("Dispatch вернул ошибку: %v", err)
	}

	if result != 42 {
		t.Fatalf("ожидалось 42, получено %d", result)
	}
}

func TestDispatchHandlerNotRegistered(t *testing.T) {
	query := struct {
		Value string
	}{
		Value: "test",
	}

	var result int

	err := mediator.Dispatch(query, &result)

	if err == nil {
		t.Fatal("ожидалась ошибка при отсутствии обработчика")
	}

	expected := "обработчик для запроса типа"

	if len(err.Error()) < len(expected) ||
		err.Error()[:len(expected)] != expected {
		t.Fatalf("неожиданная ошибка: %q", err.Error())
	}
}

func TestDispatchMultipleQueries(t *testing.T) {
	mediator.RegisterHandler(IntQueryHandler{})
	mediator.RegisterHandler(StringQueryHandler{})

	t.Run("int", func(t *testing.T) {
		var result int

		err := mediator.Dispatch(
			IntQuery{Value: 5},
			&result,
		)

		if err != nil {
			t.Fatalf("Dispatch вернул ошибку: %v", err)
		}

		if result != 10 {
			t.Fatalf("ожидалось 10, получено %d", result)
		}
	})

	t.Run("string", func(t *testing.T) {
		var result string

		err := mediator.Dispatch(
			StringQuery{Value: "test"},
			&result,
		)

		if err != nil {
			t.Fatalf("Dispatch вернул ошибку: %v", err)
		}

		if result != "result: test" {
			t.Fatalf("неожиданный результат: %q", result)
		}
	})
}

func TestDispatchOverwritesResult(t *testing.T) {
	mediator.RegisterHandler(IntQueryHandler{})

	result := 999

	err := mediator.Dispatch(
		IntQuery{Value: 7},
		&result,
	)

	if err != nil {
		t.Fatalf("Dispatch вернул ошибку: %v", err)
	}

	if result != 14 {
		t.Fatalf("ожидалось 14, получено %d", result)
	}
}
