/*
Go1.13beta1で提供されているエラーを確認します
*/
package main_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
基本
*/
func Testエラーを新しく作成する(t *testing.T) {
	err := errors.New("it is error")
	assert.EqualError(t, err, "it is error")
}

/*
基本
エラーをラッピングし、元のエラーに情報を付与して利用します
*/
func Testエラーをエラーで包み込む(t *testing.T) {
	err1 := errors.New("this is error 1")
	wrapError := fmt.Errorf("wrap error: %w", err1)
	assert.EqualError(t, wrapError, "wrap error: this is error 1")
}

/*
従来の基本的な比較方法
*/
func Testエラーを比較する_Equal(t *testing.T) {
	err1 := errors.New("this is error 1")
	var result bool
	if err1 == err1 {
		result = true
	}
	assert.True(t, result)
}

/*
新機能:チェック Is
*/
func Testエラーの値を同値かチェックする_生成された同じエラーを使う(t *testing.T) {
	err1 := errors.New("this is error 1")
	assert.True(t, errors.Is(err1, err1))
	assert.True(t, err1 == err1)
}

/*
新機能:チェック Is
*/
func Testエラーの値を同値かチェックする_エラーのテキストは同じ(t *testing.T) {
	err1 := errors.New("this is error 1")
	err2 := errors.New("this is error 1")
	assert.Equal(t, false, errors.Is(err1, err2))
}

/*
新機能:チェック Is
*/
func Testエラーの値を同値かチェックする_エラーのテキストが別(t *testing.T) {
	err1 := errors.New("this is error 1")
	err2 := errors.New("this is error 2")
	assert.Equal(t, false, errors.Is(err1, err2))
}

func Test包まれた最初のエラーであることを同じかを判定する(t *testing.T) {
	err := errors.New("this is error")
	wrappedError := fmt.Errorf("wrap error: %w", err)
	assert.True(t, errors.Is(wrappedError, err))
}

func Test包まれた最初のエラーであること同じかを判定する_何回も包む(t *testing.T) {
	err := errors.New("this is error")
	wrappedError1 := fmt.Errorf("one wrap: %w", err)
	wrappedError2 := fmt.Errorf("two wrap: %w", wrappedError1)
	assert.True(t, errors.Is(wrappedError2, err))
}

func Testエラーを包む時にフォーマットしない場合は同一性が担保できない(t *testing.T) {
	err := errors.New("error")
	wrappedError := fmt.Errorf("wrap error: %s", err)
	assert.Equal(t, false, errors.Is(wrappedError, err))
}

func Testエラーを包む時に間違ったWrap方法の場合は同一性が担保できない(t *testing.T) {
	err := errors.New("error")
	wrappedError := fmt.Errorf("wrap error %v", err)
	assert.Equal(t, false, errors.Is(wrappedError, err))
}

type UnknownError struct {
	Err string
}

func (s UnknownError) Error() string {
	return s.Err
}

func FindSomething() error {
	return &UnknownError{
		Err: "unknown error",
	}
}

type NotFoundError struct {
	Err   string
	Cause string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("%s: %v", n.Err, n.Cause)
}

func FindHoge() error {
	return &NotFoundError{
		Err:   "not found error",
		Cause: "i dont know",
	}
}

func Testエラーに対しAsを利用して指定したエラー型に変換する(t *testing.T) {
	if err := FindHoge(); err != nil {
		var notFoundErr *NotFoundError
		if errors.As(err, &notFoundErr) {
			assert.Equal(t, "i dont know", notFoundErr.Cause)
			return
		}
	}
	t.Fail()
}

func TestWrapされたエラーに対しAsを利用して指定したエラー型に変換する(t *testing.T) {
	err := FindHoge()
	wrapError := fmt.Errorf("wrapping error: %w", err)
	var notFoundErr *NotFoundError
	if errors.As(wrapError, &notFoundErr) {
		assert.NotNil(t, notFoundErr)
		assert.EqualError(t, notFoundErr, "not found error: i dont know")
		return
	}
	t.Fail()
}

func Testエラーに対しAsを利用して指定したエラー型に変換したけど変換できなかった(t *testing.T) {
	err := FindSomething()
	var notFoundErr *NotFoundError
	assert.Equal(t, false, errors.As(err, &notFoundErr))
	assert.Nil(t, notFoundErr)
}

func TestWrapされたエラーに対しAsを利用して指定したエラー型に変換したけど変換できなかった(t *testing.T) {
	err := FindSomething()
	wrapError := fmt.Errorf("wrap: %w", err)
	var notFoundErr *NotFoundError
	assert.Equal(t, false, errors.As(wrapError, &notFoundErr))
}
