package onelogplus

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntry(t *testing.T) {
	t.Run("basic-info-entry", func(t *testing.T) {
		w := newWriter()
		logger := New(w, INFO)
		logger.InfoWith("hello").Int("test", 1).Write()
		json := `{"level":"info","message":"hello","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-info-entry-hook", func(t *testing.T) {
		w := newWriter()
		logger := New(w, INFO).Hook(func(e Entry) {
			e.String("hello", "world")
		})
		logger.InfoWith("hello").Int("test", 1).Write()
		json := `{"level":"info","message":"hello","hello":"world","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-info-entry-disabled", func(t *testing.T) {
		w := newWriter()
		logger := New(w, WARN)
		logger.InfoWith("hello").Int("test", 1).Write()
		json := ``
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-debug-entry", func(t *testing.T) {
		w := newWriter()
		logger := New(w, DEBUG)
		logger.DebugWith("hello").Int("test", 1).Write()
		json := `{"level":"debug","message":"hello","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-debug-entry-hook", func(t *testing.T) {
		w := newWriter()
		logger := New(w, DEBUG).Hook(func(e Entry) {
			e.String("hello", "world")
		})
		logger.DebugWith("hello").Int("test", 1).Write()
		json := `{"level":"debug","message":"hello","hello":"world","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-debug-entry-disabled", func(t *testing.T) {
		w := newWriter()
		logger := New(w, WARN)
		logger.DebugWith("hello").Int("test", 1).Write()
		json := ``
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-warn-entry", func(t *testing.T) {
		w := newWriter()
		logger := New(w, WARN)
		logger.WarnWith("hello").Int("test", 1).Write()
		json := `{"level":"warn","message":"hello","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-warn-entry-hook", func(t *testing.T) {
		w := newWriter()
		logger := New(w, WARN).Hook(func(e Entry) {
			e.String("hello", "world")
		})
		logger.WarnWith("hello").Int("test", 1).Write()
		json := `{"level":"warn","message":"hello","hello":"world","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-warn-entry-disabled", func(t *testing.T) {
		w := newWriter()
		logger := New(w, ERROR)
		logger.WarnWith("hello").Int("test", 1).Write()
		json := ``
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-error-entry", func(t *testing.T) {
		w := newWriter()
		logger := New(w, ERROR)
		logger.ErrorWith("hello").Int("test", 1).Write()
		json := `{"level":"error","message":"hello","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-error-entry-hook", func(t *testing.T) {
		w := newWriter()
		logger := New(w, ERROR).Hook(func(e Entry) {
			e.String("hello", "world")
		})
		logger.ErrorWith("hello").Int("test", 1).Write()
		json := `{"level":"error","message":"hello","hello":"world","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-error-entry-disabled", func(t *testing.T) {
		w := newWriter()
		logger := New(w, FATAL)
		logger.ErrorWith("hello").Int("test", 1).Write()
		json := ``
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-fatal-entry", func(t *testing.T) {
		w := newWriter()
		logger := New(w, FATAL)
		logger.FatalWith("hello").Int("test", 1).Write()
		json := `{"level":"fatal","message":"hello","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	t.Run("basic-fatal-entry-hook", func(t *testing.T) {
		w := newWriter()
		logger := New(w, FATAL).Hook(func(e Entry) {
			e.String("hello", "world")
		})
		logger.FatalWith("hello").Int("test", 1).Write()
		json := `{"level":"fatal","message":"hello","hello":"world","test":1}` + "\n"
		assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	})
	//t.Run("basic-fatal-entry-disabled", func(t *testing.T) {
	//	w := newWriter()
	//	logger := New(w, FATAL)
	//	logger.FatalWith("hello").Int("test", 1).Write()
	//	json := ``
	//	assert.Equal(t, json, string(w.b), "bytes written to the writer dont equal expected result")
	//})
}

func TestEntryFields(t *testing.T) {
	json := `{"level":"%s","message":"hello","testInt":1,"testInt64":2,"testFloat":1.15234,` +
		`"testString":"string","testBool":true,"testObj":{"testInt":100},` +
		`"testObj2":{"foo":"bar"},"testArr":[{"foo":"bar"},{"foo":"bar"}],` +
		`"testAnyString":"bar",` +
		`"testAnyFloat":10.1,` +
		`"testAnyInt":10,` +
		`"testErr":"my printer is on fire"}` + "\n"
	testCases := []struct {
		level       uint32
		disabled    uint32
		levelString string
		entryFunc   func(*Logger) ChainEntry
	}{
		{
			level:       INFO,
			disabled:    DEBUG,
			levelString: "info",
			entryFunc: func(l *Logger) ChainEntry {
				return l.InfoWith("hello")
			},
		},
		{
			level:       DEBUG,
			disabled:    INFO,
			levelString: "debug",
			entryFunc: func(l *Logger) ChainEntry {
				return l.DebugWith("hello")
			},
		},
		{
			level:       WARN,
			disabled:    ERROR,
			levelString: "warn",
			entryFunc: func(l *Logger) ChainEntry {
				return l.WarnWith("hello")
			},
		},
		{
			level:       ERROR,
			disabled:    WARN,
			levelString: "error",
			entryFunc: func(l *Logger) ChainEntry {
				return l.ErrorWith("hello")
			},
		},
		{
			level:       FATAL,
			disabled:    ERROR,
			levelString: "fatal",
			entryFunc: func(l *Logger) ChainEntry {
				return l.FatalWith("hello")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test-%s-entry-all-fields-enabled", testCase.levelString), func(t *testing.T) {
			w := newWriter()
			logger := New(w, testCase.level)
			testObj := &TestObj{"bar"}
			testArr := TestObjArr{testObj, testObj}
			testCase.entryFunc(logger).
				Int("testInt", 1).
				Int64("testInt64", 2).
				Float("testFloat", 1.15234).
				String("testString", "string").
				Bool("testBool", true).
				ObjectFunc("testObj", func(e Entry) {
					e.Int("testInt", 100)
				}).
				Object("testObj2", testObj).
				Array("testArr", testArr).
				Any("testAnyString", "bar").
				Any("testAnyFloat", 10.1).
				Any("testAnyInt", 10).
				Err("testErr", errors.New("my printer is on fire")).
				Write()
			assert.Equal(t, fmt.Sprintf(json, testCase.levelString), string(w.b), "bytes written to the writer dont equal expected result")
		})
		t.Run(fmt.Sprintf("test-%s-entry-all-fields-disabled", testCase.levelString), func(t *testing.T) {
			w := newWriter()
			logger := New(w, FATAL)
			testObj := &TestObj{"bar"}
			testArr := TestObjArr{testObj, testObj}
			testCase.entryFunc(logger).
				Int("testInt", 1).
				Int64("testInt64", 2).
				Float("testFloat", 1.15234).
				String("testString", "string").
				Bool("testBool", true).
				ObjectFunc("testObj", func(e Entry) {
					e.Int("testInt", 100)
				}).
				Object("testObj2", testObj).
				Array("testArr", testArr).
				Any("testAnyString", "bar").
				Any("testAnyFloat", 10.0).
				Any("testAnyInt", 10).
				Err("testErr", errors.New("my printer is on fire")).
				Write()
			if testCase.levelString == Levels[FATAL] {
				return
			}
			assert.Equal(t, ``, string(w.b), "bytes written to the writer dont equal expected result")

		})
	}
}
