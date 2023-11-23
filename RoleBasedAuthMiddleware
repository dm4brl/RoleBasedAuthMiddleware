package main

import (
	"net/http"
)

// RoleBasedAuthMiddleware ограничивает доступ к ресурсам в зависимости от роли пользователя.
func RoleBasedAuthMiddleware(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Header.Get("X-User-Role")

		// Проверка, есть ли роль пользователя в списке разрешенных ролей.
		roleAllowed := false
		for _, role := range allowedRoles {
			if role == userRole {
				roleAllowed = true
				break
			}
		}

		// Если роль пользователя разрешена, передаем управление следующему обработчику.
		if roleAllowed {
			next.ServeHTTP(w, r)
		} else {
			// В противном случае, возвращаем http.StatusForbidden.
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// AdminHandler обработчик для ресурса "/admin".
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	userRole := r.Header.Get("X-User-Role")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Admin Resource\nUser Role: " + userRole))
}

// UserHandler обработчик для ресурса "/user".
func UserHandler(w http.ResponseWriter, r *http.Request) {
	userRole := r.Header.Get("X-User-Role")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Resource\nUser Role: " + userRole))
}

func main() {
	allowedAdminRoles := []string{"admin", "superadmin"}
	allowedUserRoles := []string{"user"}

	// Создание маршрута и применение Middleware для пути "/admin".
	adminHandler := RoleBasedAuthMiddleware(allowedAdminRoles, http.HandlerFunc(AdminHandler))
	http.Handle("/admin", adminHandler)

	// Создание маршрута и применение Middleware для пути "/user".
	userHandler := RoleBasedAuthMiddleware(allowedUserRoles, http.HandlerFunc(UserHandler))
	http.Handle("/user", userHandler)

	// Запуск веб-сервера на порту 8080.
	http.ListenAndServe(":8080", nil)
}
