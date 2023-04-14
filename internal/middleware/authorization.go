package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"template/internal/appctx"
	"template/internal/entities"
	h "template/internal/handler"
	"template/internal/repository"
	"template/utils/jwt"
	p "template/utils/password"

	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
)

func Authorization(db *gorm.DB) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logrus.Info("Authorization middleware is executed")
			startTime := time.Now()
			hd := &h.Handler{}

			// Enforce Route
			e_route, err := casbin.NewEnforcer("./internal/config/casbin/route_model.conf", "./internal/config/casbin/route_policy.csv")
			if err != nil {
				log.Fatal(err)
			}

			var user entities.User
			var roles []entities.Role
			userRepo := repository.NewUserRepository(db)

			if res, _ := e_route.Enforce(r.URL.Path, r.Method); res {
				authHeader := r.Header.Get("Authorization")
				if strings.Contains(authHeader, "Bearer") {
					logrus.Info("JWT authorization")

					// Parse token
					token := strings.ReplaceAll(authHeader, "Bearer ", "")
					claims, err := jwt.ParseToken(token)
					if err != nil {
						resp := appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusUnauthorized)
						hd.Response(w, *resp, startTime, time.Now())
						return
					}

					// Get User
					user, err := userRepo.Get(user, claims.UserID)
					if err != nil {
						resp := appctx.NewResponse().WithErrors("user not found").WithCode(http.StatusUnauthorized)
						hd.Response(w, *resp, startTime, time.Now())
						return
					}
					roles = user.Roles

					r.Header.Add("user", strconv.Itoa(user.ID))
				} else if strings.Contains(authHeader, "Basic") {
					logrus.Info("Basic Auth authorization")
					username, password, ok := r.BasicAuth()
					if !ok {
						resp := appctx.NewResponse().WithErrors("Wrong basic auth header").WithCode(http.StatusBadRequest)
						hd.Response(w, *resp, startTime, time.Now())
						return
					}

					// Get User
					user, err := userRepo.GetByEmail(username)
					if err != nil {
						resp := appctx.NewResponse().WithErrors("user not found").WithCode(http.StatusUnauthorized)
						hd.Response(w, *resp, startTime, time.Now())
						return
					}
					roles = user.Roles

					if match := p.CheckPasswordHash(password, user.Password); !match {
						resp := appctx.NewResponse().WithErrors("wrong password").WithCode(http.StatusUnauthorized)
						hd.Response(w, *resp, startTime, time.Now())
						return
					}

					r.Header.Add("user", strconv.Itoa(user.ID))

				} else {
					logrus.Error("Wrong authorizatio header")
					resp := appctx.NewResponse().WithErrors("Wrong authorization header").WithCode(http.StatusBadRequest)
					hd.Response(w, *resp, startTime, time.Now())
					return
				}

				e, err := casbin.NewEnforcer("./internal/config/casbin/auth_model.conf", "./internal/config/casbin/policy.csv")
				if err != nil {
					log.Fatal(err)
				}

				isAuthorized := false

				for _, role := range roles {
					if res, _ := e.Enforce(role.Name, r.URL.Path, r.Method); res {
						isAuthorized = true
						break
					}
				}

				if !isAuthorized {
					resp := appctx.NewResponse().WithErrors("Unauthorized role").WithCode(http.StatusForbidden)
					hd.Response(w, *resp, startTime, time.Now())
					return
				}
			}

			handler.ServeHTTP(w, r)
		})
	}
}
