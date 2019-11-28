// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tm_bot

import (
	"github.com/gardener/test-infra/pkg/testrunner"
	"github.com/gardener/test-infra/pkg/tm-bot/github"
	"github.com/gardener/test-infra/pkg/tm-bot/hook"
	"github.com/gardener/test-infra/pkg/tm-bot/tests"
	"github.com/gardener/test-infra/pkg/tm-bot/ui"
	"github.com/gardener/test-infra/pkg/tm-bot/ui/auth"
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func setup(log logr.Logger, stopCh chan struct{}) (*mux.Router, error) {
	w, err := testrunner.StartWatchController(log, kubeconfigPath, stopCh)
	if err != nil {
		return nil, err
	}
	runs := tests.NewRuns(w)

	ghClient, err := github.NewManager(log.WithName("github"), ghManagerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize github client")
	}
	hooks, err := hook.New(log.WithName("hooks"), ghClient, webhookSecretToken, w, runs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize webhooks handler")
	}

	r := mux.NewRouter()
	r.Use(loggingMiddleware(log.WithName("trace")))
	r.HandleFunc("/healthz", healthz(log.WithName("health"))).Methods(http.MethodGet)
	r.HandleFunc("/events", hooks.HandleWebhook).Methods(http.MethodPost)

	a := auth.NewAuth(log.WithName("authentication"), authOrg, oauthClientID, oauthClientSecret, cookieSecret)
	ui.Serve(log, runs, uiBasePath, a, r)
	return r, nil
}

func loggingMiddleware(log logr.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.V(10).Info(r.RequestURI, "method", r.Method)
			next.ServeHTTP(w, r)
		})
	}
}
