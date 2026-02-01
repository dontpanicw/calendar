package port

// EventRepository интерфейс для работы с репозиторием событий
type EventUsecases interface {
	// CreateEvent создает события и возвращает id
	CreateEvent() int64
	// UpdateEvent изменяет событие
	UpdateEvent() int64
	// DeleteEvent удаляет событие
	DeleteEvent() int64
	// GetEventsForDay получить события за день
	GetEventsForDay()
	// GetEventsForWeek получить события за неделю
	GetEventsForWeek()
	// GetEventsForMonth получить события за месяц
	GetEventsForMonth()
}
