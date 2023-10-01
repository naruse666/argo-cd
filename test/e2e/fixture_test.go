package e2e

import (
	"testing"
	"time"

	. "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	. "github.com/argoproj/argo-cd/v2/test/e2e/fixture"
	. "github.com/argoproj/argo-cd/v2/test/e2e/fixture/app"
	"github.com/stretchr/testify/assert"
)

func TestSetEnableManifestGeneration(t *testing.T) {
	Given(t).
		SetAppNamespace(AppNamespace()).
		SetTrackingMethod("annotation").
		Path("guestbook").
		When().
		CreateApp().
		Refresh(RefreshTypeHard).
		Then().
		And(func(app *Application) {
			assert.Equal(t, app.Status.SourceType, ApplicationSourceTypeKustomize)
		}).
		When().
		And(func() {
			SetEnableManifestGeneration(map[ApplicationSourceType]bool{
				ApplicationSourceTypeKustomize: false,
			})
		}).
		Refresh(RefreshTypeHard).
		Then().
		And(func(app *Application) {
			time.Sleep(1 * time.Second)
		}).
		And(func(app *Application) {
			assert.Equal(t, app.Status.SourceType, ApplicationSourceTypeDirectory)
		})
}
