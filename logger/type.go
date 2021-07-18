package logger

import "github.com/sirupsen/logrus"

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// convertFields convert Fields to the logrus.Fields
func convertFields(fields Fields) logrus.Fields {
	logrusFields := make(logrus.Fields, len(fields))
	for k, v := range fields {
		logrusFields[k] = v
	}
	return logrusFields
}

type Hook interface {
	logrus.Hook
}

func convertHook(hook Hook) logrus.Hook {
	return hook.(logrus.Hook)
}

func convertHookReverse(hook logrus.Hook) Hook {
	return hook.(Hook)
}

func convertHooks(hooks []Hook) []logrus.Hook {
	logrusHooks := make([]logrus.Hook, 0, len(hooks))
	for _, hook := range hooks {
		logrusHooks = append(logrusHooks, convertHook(hook))
	}
	return logrusHooks
}

func convertHooksReverse(logrusHooks []logrus.Hook) []Hook {
	hooks := make([]Hook, 0, len(logrusHooks))
	for _, hook := range logrusHooks {
		hooks = append(hooks, convertHookReverse(hook))
	}
	return hooks
}

type LevelHooks = logrus.LevelHooks

func convertLevelHooks(levelHooks LevelHooks) logrus.LevelHooks {
	//hooks := make(logrus.LevelHooks, len(levelHooks))
	//for k, v := range levelHooks {
	//	hooks[convertLevel(k)] = convertHooks(v)
	//}
	return levelHooks
}

func convertLevelHooksReverse(levelHooks logrus.LevelHooks) LevelHooks {
	//hooks := make(LevelHooks, len(levelHooks))
	//for k, v := range levelHooks {
	//	hooks[convertLevelReverse(k)] = convertHooksReverse(v)
	//}
	return levelHooks
}
