package js_test

import (
	"context"
	"fmt"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v3"

	"go.k6.io/k6/js"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/testutils"
	"go.k6.io/k6/loader"
	"go.k6.io/k6/metrics"
)

type CheckModule struct {
	t             testing.TB
	initCtxCalled int
	vuCtxCalled   int
}

func (cm *CheckModule) InitCtx(ctx context.Context) {
	cm.initCtxCalled++
}

func (cm *CheckModule) VuCtx(ctx context.Context) {
	cm.vuCtxCalled++
}

func TestNewJSRunnerWithCustomModule(t *testing.T) {
	t.Parallel()

	var uniqueModuleNumber int64
	checkModule := &CheckModule{t: t}
	moduleName := fmt.Sprintf("k6/x/check-%d", atomic.AddInt64(&uniqueModuleNumber, 1))
	modules.Register(moduleName, checkModule)

	script := fmt.Sprintf(`
		var check = require("%s");
		check.initCtx();

		module.exports.options = { vus: 1, iterations: 1 };
		module.exports.default = function() {
			check.vuCtx();
		};
	`, moduleName)

	logger := testutils.NewLogger(t)
	rtOptions := lib.RuntimeOptions{CompatibilityMode: null.StringFrom("base")}
	registry := metrics.NewRegistry()
	builtinMetrics := metrics.RegisterBuiltinMetrics(registry)
	runner, err := js.New(
		&lib.TestPreInitState{
			Logger:         logger,
			BuiltinMetrics: builtinMetrics,
			Registry:       registry,
			RuntimeOptions: rtOptions,
		},
		&loader.SourceData{
			URL:  &url.URL{Path: "blah", Scheme: "file"},
			Data: []byte(script),
		},
		map[string]afero.Fs{"file": afero.NewMemMapFs(), "https": afero.NewMemMapFs()}, redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	)
	require.NoError(t, err)
	assert.Equal(t, checkModule.initCtxCalled, 1)
	assert.Equal(t, checkModule.vuCtxCalled, 0)

	vu, err := runner.NewVU(1, 1, make(chan metrics.SampleContainer, 100), redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}))
	require.NoError(t, err)
	assert.Equal(t, checkModule.initCtxCalled, 2)
	assert.Equal(t, checkModule.vuCtxCalled, 0)

	vuCtx, vuCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer vuCancel()

	activeVU := vu.Activate(&lib.VUActivationParams{RunContext: vuCtx})
	require.NoError(t, activeVU.RunOnce())
	assert.Equal(t, checkModule.initCtxCalled, 2)
	assert.Equal(t, checkModule.vuCtxCalled, 1)
	require.NoError(t, activeVU.RunOnce())
	assert.Equal(t, checkModule.initCtxCalled, 2)
	assert.Equal(t, checkModule.vuCtxCalled, 2)

	arc := runner.MakeArchive()
	assert.Equal(t, checkModule.initCtxCalled, 2) // shouldn't change, we're not executing the init context again
	assert.Equal(t, checkModule.vuCtxCalled, 2)

	runnerFromArc, err := js.NewFromArchive(
		&lib.TestPreInitState{
			Logger:         logger,
			BuiltinMetrics: builtinMetrics,
			Registry:       registry,
			RuntimeOptions: rtOptions,
		}, arc, redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}))
	require.NoError(t, err)
	assert.Equal(t, checkModule.initCtxCalled, 3) // changes because we need to get the exported functions
	assert.Equal(t, checkModule.vuCtxCalled, 2)
	vuFromArc, err := runnerFromArc.NewVU(2, 2, make(chan metrics.SampleContainer, 100), redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}))
	require.NoError(t, err)
	assert.Equal(t, checkModule.initCtxCalled, 4)
	assert.Equal(t, checkModule.vuCtxCalled, 2)
	activeVUFromArc := vuFromArc.Activate(&lib.VUActivationParams{RunContext: vuCtx})
	require.NoError(t, activeVUFromArc.RunOnce())
	assert.Equal(t, checkModule.initCtxCalled, 4)
	assert.Equal(t, checkModule.vuCtxCalled, 3)
	require.NoError(t, activeVUFromArc.RunOnce())
	assert.Equal(t, checkModule.initCtxCalled, 4)
	assert.Equal(t, checkModule.vuCtxCalled, 4)
}
