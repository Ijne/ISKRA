function getCurrentUser() {
    return 96419039;
}

const API_BASE_URL = 'http://localhost:8080';

let currentEvents = [];
let currentFlames = [];
let selectedEventId = null;
let userCity = '';
let currentSkip = 0;
const limit = 5;
let isLoading = false;
let hasMoreEvents = true;

// Загрузка мероприятий при загрузке страницы
document.addEventListener('DOMContentLoaded', function() {
    loadUserCity();
    setupNavigation();
    setupInfiniteScroll();
});

// Загрузка города пользователя
async function loadUserCity() {
    try {
        const response = await fetch(`${API_BASE_URL}/profile?id=${getCurrentUser()}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (response.ok) {
            const userData = await response.json();
            userCity = userData.city || '';
            console.log('Город пользователя:', userCity);
        }
    } catch (error) {
        console.error('Ошибка загрузки города пользователя:', error);
    } finally {
        loadEvents();
    }
}

// Настройка бесконечной прокрутки
function setupInfiniteScroll() {
    window.addEventListener('scroll', () => {
        if (isLoading || !hasMoreEvents) return;

        const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
        
        if (scrollTop + clientHeight >= scrollHeight - 100) {
            loadMoreEvents();
        }
    });
}

// Загрузка мероприятий
async function loadEvents() {
    try {
        const response = await fetch(`${API_BASE_URL}/events?id=${getCurrentUser()}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error('Ошибка при загрузке мероприятий');
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            currentEvents = data.events || [];
            displayEvents(currentEvents);
            
            // Если сохраненных событий мало, сразу загружаем дополнительные
            if (currentEvents.length < 5 && userCity) {
                loadMoreEvents();
            }
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка загрузки мероприятий:', error);
        showMessage('Не удалось загрузить мероприятия', 'error');
    }
}

// Загрузка дополнительных мероприятий
async function loadMoreEvents() {
    if (isLoading || !userCity || !hasMoreEvents) return;
    
    isLoading = true;
    showLoadingIndicator();

    try {
        const response = await fetch(`${API_BASE_URL}/events?id=${getCurrentUser()}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                limit: limit,
                skip: currentSkip
            })
        });

        if (!response.ok) {
            throw new Error('Ошибка при загрузке дополнительных мероприятий');
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            const newEvents = data.events || [];
            
            if (newEvents.length > 0) {
                currentEvents = [...currentEvents, ...newEvents];
                displayEvents(currentEvents);
                currentSkip += limit;
                
                // Если пришло меньше событий, чем запрошено, значит больше нет
                if (newEvents.length < limit) {
                    hasMoreEvents = false;
                    hideLoadingIndicator();
                    showNoMoreEvents();
                }
            } else {
                hasMoreEvents = false;
                showNoMoreEvents();
            }
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка загрузки дополнительных мероприятий:', error);
        showMessage('Не удалось загрузить дополнительные мероприятия', 'error');
    } finally {
        isLoading = false;
        hideLoadingIndicator();
    }
}

// Отображение мероприятий
function displayEvents(events) {
    const eventsList = document.getElementById('eventsList');
    
    // Удаляем индикаторы загрузки если они есть
    const existingIndicators = eventsList.querySelectorAll('.loading-indicator, .no-more-events');
    existingIndicators.forEach(indicator => indicator.remove());
    
    if (!events || events.length === 0) {
        eventsList.innerHTML = `
            <div class="no-events">
                Пока нет доступных мероприятий
            </div>
        `;
        return;
    }

    eventsList.innerHTML = events.map(event => `
        <div class="event-card" onclick="openFlamesModal(${event.ID})">
            ${event.Photo ? 
                `<img src="${event.Photo}" alt="${event.Name}" class="event-photo" onerror="this.style.display='none'; this.nextElementSibling.style.display='flex';">` : 
                ''
            }
            <div class="event-photo-placeholder" ${event.Photo ? 'style="display: none;"' : ''}>
                🎭
            </div>
            <h3 class="event-name">${escapeHtml(event.Name)}</h3>
            <div class="event-date">${formatDate(event.StartsAt)}</div>
            <div class="event-url">
                <a href="${event.Url}" target="_blank" onclick="event.stopPropagation()">
                    ${event.Url}
                </a>
            </div>
        </div>
    `).join('');
    
    // Добавляем индикатор загрузки если есть еще события
    if (hasMoreEvents && userCity) {
        eventsList.innerHTML += `<div class="loading-indicator">Загрузка...</div>`;
    }
}

// Показать индикатор загрузки
function showLoadingIndicator() {
    const eventsList = document.getElementById('eventsList');
    const existingIndicator = eventsList.querySelector('.loading-indicator');
    if (!existingIndicator) {
        eventsList.innerHTML += `<div class="loading-indicator">Загрузка...</div>`;
    }
}

// Скрыть индикатор загрузки
function hideLoadingIndicator() {
    const indicator = document.querySelector('.loading-indicator');
    if (indicator) {
        indicator.remove();
    }
}

// Показать сообщение о том, что события закончились
function showNoMoreEvents() {
    const eventsList = document.getElementById('eventsList');
    const existingMessage = eventsList.querySelector('.no-more-events');
    if (!existingMessage) {
        eventsList.innerHTML += `
            <div class="no-more-events">
                Все мероприятия загружены
            </div>
        `;
    }
}

// Открытие модального окна с лобби
async function openFlamesModal(eventId) {
    selectedEventId = eventId;
    
    const event = currentEvents.find(e => e.ID === eventId);
    if (event) {
        document.getElementById('modalEventTitle').textContent = event.Name;
    }
    
    try {
        const response = await fetch(`${API_BASE_URL}/flames?id=${getCurrentUser()}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ event_id: eventId })
        });

        if (!response.ok) {
            throw new Error('Ошибка при загрузке лобби');
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            currentFlames = data.flames || [];
            displayFlames(currentFlames);
            document.getElementById('flamesModal').style.display = 'flex';
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка загрузки лобби:', error);
        showMessage('Не удалось загрузить лобби', 'error');
    }
}

// Закрытие модального окна лобби
function closeFlamesModal() {
    document.getElementById('flamesModal').style.display = 'none';
    selectedEventId = null;
    currentFlames = [];
}

// Отображение лобби
function displayFlames(flames) {
    const flamesList = document.getElementById('flamesList');
    const currentUserId = getCurrentUser();
    
    if (!flames || flames.length === 0) {
        flamesList.innerHTML = `
            <div class="no-flames">
                <div class="no-flames-icon">🔥</div>
                <div class="no-flames-text">Пока нет лобби для этого мероприятия</div>
                <button class="create-flame-btn-inline" onclick="openCreateFlameModal()">
                    Создать первое лобби
                </button>
            </div>
        `;
        return;
    }

    flamesList.innerHTML = flames.map(flame => {
        const isOwnFlame = flame.user_id === currentUserId;
        const userInitials = getInitials(flame.name, flame.surname);
        
        return `
            <div class="flame-card ${isOwnFlame ? 'own-flame' : ''}">
                <div class="flame-header">
                    <div class="flame-user">
                        <div class="flame-avatar">${userInitials}</div>
                        <div class="flame-user-info">
                            <div class="flame-username">${flame.username || 'Пользователь'}</div>
                            <div class="flame-user-details">
                                ${flame.age ? flame.age + ' лет' : ''} 
                                ${flame.gender !== undefined ? (flame.gender === 0 ? '♂' : '♀') : ''}
                            </div>
                        </div>
                    </div>
                    ${!isOwnFlame ? `
                        <button class="like-btn" onclick="likeUser(${flame.user_id}, this)">
                            ❤️ Лайк
                        </button>
                    ` : ''}
                </div>
                <div class="flame-description">
                    ${escapeHtml(flame.description || 'Без описания')}
                </div>
            </div>
        `;
    }).join('');
}

// Лайк пользователя
async function likeUser(userId, button) {
    try {
        const response = await fetch(`${API_BASE_URL}/like-user?id=${getCurrentUser()}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ light_id: userId })
        });

        if (!response.ok) {
            throw new Error('Ошибка при отправке лайка');
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            button.classList.add('liked');
            button.innerHTML = '❤️ Лайк отправлен!';
            button.disabled = true;
            
            setTimeout(() => {
                button.style.opacity = '0.7';
            }, 1000);
            
            console.log('Лайк успешно отправлен пользователю:', userId);
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка отправки лайка:', error);
        showMessage('Не удалось отправить лайк', 'error');
    }
}

// Открытие модального окна создания лобби
function openCreateFlameModal() {
    document.getElementById('createFlameModal').style.display = 'flex';
    document.getElementById('flameDescription').value = '';
}

// Закрытие модального окна создания лобби
function closeCreateFlameModal() {
    document.getElementById('createFlameModal').style.display = 'none';
}

// Создание лобби
async function createFlame() {
    const description = document.getElementById('flameDescription').value.trim();
    
    if (!description) {
        showMessage('Введите описание лобби', 'error');
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/flame?id=${getCurrentUser()}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                event_id: selectedEventId,
                description: description
            })
        });

        if (!response.ok) {
            throw new Error('Ошибка при создании лобби');
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            closeCreateFlameModal();
            console.log('Лобби успешно создано для мероприятия:', selectedEventId);
            
            // Обновляем список лобби
            setTimeout(() => {
                openFlamesModal(selectedEventId);
            }, 500);
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка создания лобби:', error);
        showMessage('Не удалось создать лобби', 'error');
    }
}

// Вспомогательные функции
function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('ru-RU', {
        day: 'numeric',
        month: 'long',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

function getInitials(name, surname) {
    const first = name ? name[0] : 'П';
    const second = surname ? surname[0] : 'У';
    return (first + second).toUpperCase();
}

function showMessage(message, type = 'info') {
    console.log(`${type.toUpperCase()}: ${message}`);
}

// Настройка навигации
function setupNavigation() {
    const profileButton = document.querySelector('.nav-button:nth-child(1)');
    const mainButton = document.querySelector('.main-button');
    
    if (profileButton) {
        profileButton.addEventListener('click', function() {
            window.location.href = 'profile.html';
        });
    }
    
    if (mainButton) {
        mainButton.addEventListener('click', function() {
            window.location.href = 'index.html';
        });
    }
}

// Закрытие модальных окон при клике вне их
document.addEventListener('click', function(event) {
    const flamesModal = document.getElementById('flamesModal');
    const createFlameModal = document.getElementById('createFlameModal');
    
    if (event.target === flamesModal) {
        closeFlamesModal();
    }
    
    if (event.target === createFlameModal) {
        closeCreateFlameModal();
    }
});