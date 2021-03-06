// +build go1.7

package service

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/nytimes/gizmo/server"
)

// JSONMiddleware provides a JSONContextEndpoint hook wrapped around all requests.
// In this implementation, we're using it to provide application logging and to check errors
// and provide generic responses.
func (s *RPCService) JSONMiddleware(j server.JSONContextEndpoint) server.JSONContextEndpoint {
	return func(ctx context.Context, r *http.Request) (int, interface{}, error) {
		status, res, err := j(ctx, r)
		if err != nil {
			server.LogWithFields(r).WithFields(logrus.Fields{
				"error": err,
			}).Error("problems with serving request")
			return http.StatusServiceUnavailable, nil, &jsonErr{"sorry, this service is unavailable"}
		}

		server.LogWithFields(r).Info("success!")
		return status, res, nil
	}
}
