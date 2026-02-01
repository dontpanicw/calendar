package port

// EventRepository интерфейс для работы с репозиторием событий
type EventRepository interface {
	// CreateEvent создает события и возвращает id
	CreateEvent() int64
	// UpdateEvent изменяет событие
	UpdateEvent() int64
	// DeleteEvent удаляет событие
	DeleteEvent() int64
	// GetEvents получить события
	GetEvents()
}
