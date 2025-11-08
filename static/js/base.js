function getCurrentUser() {
    return 2;
}

// Функция для редиректа на нужные страницы
function setupNavigation() {
    // Получаем все кнопки навигации
    const profileButton = document.querySelector('.nav-button:nth-child(1)'); // Профиль
    const mainButton = document.querySelector('.main-button'); // Главная
    const eventsButton = document.querySelector('.nav-button:nth-child(3)'); // Мероприятия
    
    // Добавляем обработчики событий
    if (profileButton) {
        profileButton.addEventListener('click', function() {
            window.location.href = 'profile.html'; // Редирект на страницу профиля
        });
    }
    
    if (mainButton) {
        mainButton.addEventListener('click', function() {
            window.location.href = 'index.html'; // Редирект на главную страницу
        });
    }
    
    if (eventsButton) {
        eventsButton.addEventListener('click', function() {
            window.location.href = 'events.html'; // Редирект на страницу мероприятий
        });
    }
}

// Запускаем настройку навигации когда DOM загружен
document.addEventListener('DOMContentLoaded', function() {
    setupNavigation();
});