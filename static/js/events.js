let initData = null;
let WebApp = null;

function waitForWebApp() {
    return new Promise((resolve) => {
        if (window.WebApp) {
            WebApp = window.WebApp;
            initData = window.WebApp?.initData;
            console.log('WebApp загружен:', WebApp);
            resolve();
            return;
        }

        let attempts = 0;
        const maxAttempts = 50;
        
        const check = () => {
            attempts++;
            if (window.WebApp) {
                WebApp = window.WebApp;
                initData = window.WebApp?.initData;
                console.log('WebApp загружен:', WebApp);
                console.log('InitData:', initData);
                resolve();
            } else if (attempts < maxAttempts) {
                setTimeout(check, 100);
            } else {
                console.warn('WebApp не загрузился, продолжаем без него');
                resolve();
            }
        };
        
        check();
    });
}

async function getCurrentUser() {
    try {
        await waitForWebApp();
        
        if (!initData) {
            console.error('No init data found');
            return null;
        }

        let decodedString;
        
        if (typeof initData === 'object') {
            const user = initData.user || initData;
            return user.id || null;
        }

        if (typeof initData === 'string') {
            decodedString = decodeURIComponent(initData);

            const params = new URLSearchParams(decodedString);
            const receivedHash = params.get('hash');
            
            if (!receivedHash) {
                console.error('Hash not found in init data');
                const userParam = params.get('user');
                if (userParam) {
                    try {
                        const userData = JSON.parse(userParam);
                        return userData.id || null;
                    } catch (e) {
                        console.error('Error parsing user data:', e);
                    }
                }
                return null;
            }

            const userParam = params.get('user');
            
            const dataPairs = [];
            for (const [key, value] of params) {
                if (key !== 'hash') {
                    dataPairs.push(`${key}=${value}`);
                }
            }
            dataPairs.sort();
            
            const dataCheckString = dataPairs.join('\n');

            const botToken = 'f9LHodD0cOLRQi29OdyXpiSqLM-SyPUJnePMbZQH3ceilC7cKmf12ib4C7Oeda975ZN_gzuX6fJmQVKE5j1e';
            
            const encoder = new TextEncoder();

            const secretKey = await crypto.subtle.importKey(
                'raw',
                encoder.encode('WebAppData'),
                { name: 'HMAC', hash: 'SHA-256' },
                false,
                ['sign']
            );

            const cryptoKey = await crypto.subtle.sign(
                'HMAC',
                secretKey,
                encoder.encode(botToken)
            );

            const hmacKey = await crypto.subtle.importKey(
                'raw',
                cryptoKey,
                { name: 'HMAC', hash: 'SHA-256' },
                false,
                ['sign']
            );

            const signature = await crypto.subtle.sign(
                'HMAC',
                hmacKey,
                encoder.encode(dataCheckString)
            );
            
            const calculatedHash = Array.from(new Uint8Array(signature))
                .map(b => b.toString(16).padStart(2, '0'))
                .join('');
            
            console.log('Calculated hash:', calculatedHash);
            console.log('Received hash:', receivedHash);

            if (calculatedHash === receivedHash) {
                console.log('Hash validation successful');
                
                if (userParam) {
                    try {
                        const userData = JSON.parse(userParam);
                        console.log('User data:', userData);
                        return userData.id || null;
                    } catch (parseError) {
                        console.error('Error parsing user data:', parseError);
                        return null;
                    }
                }
            } else {
                console.log('Hash validation failed');
                return null;
            }
        }
        
        return null;
    } catch (error) {
        console.error('Validation error:', error);
        return null;
    }
}

const API_BASE_URL = 'http://localhost:8080';

let currentEvents = [];
let currentFlames = [];
let selectedEventId = null;

async function initApp() {
    try {
        await waitForWebApp();
        console.log('Приложение инициализировано');
        
        await loadEvents();
        setupNavigation();
        
    } catch (error) {
        console.error('Ошибка инициализации приложения:', error);
        await loadEvents();
        setupNavigation();
    }
}

document.addEventListener('DOMContentLoaded', initApp);

async function loadEvents() {
    try {
        console.log('Загрузка мероприятий...');
        const response = await fetch(`${API_BASE_URL}/events`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            currentEvents = data.events || [];
            console.log(`Загружено мероприятий: ${currentEvents.length}`);
            displayEvents(currentEvents);
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка загрузки мероприятий:', error);
        showMessage('Не удалось загрузить мероприятия', 'error');
    }
}

function displayEvents(events) {
    const eventsList = document.getElementById('eventsList');
    
    if (!eventsList) {
        console.error('Элемент eventsList не найден');
        return;
    }
    
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
            <div class="event-url">${event.Url}</div>
        </div>
    `).join('');
}

async function openFlamesModal(eventId) {
    try {
        selectedEventId = eventId;
        
        const event = currentEvents.find(e => e.ID === eventId);
        if (event) {
            const modalEventTitle = document.getElementById('modalEventTitle');
            if (modalEventTitle) {
                modalEventTitle.textContent = event.Name;
            }
        }
        
        console.log('Загрузка лобби для мероприятия:', eventId);
        const response = await fetch(`${API_BASE_URL}/flames`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ event_id: eventId })
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            currentFlames = data.flames || [];
            console.log(`Загружено лобби: ${currentFlames.length}`);
            await displayFlames(currentFlames);
            
            const flamesModal = document.getElementById('flamesModal');
            if (flamesModal) {
                flamesModal.style.display = 'flex';
            }
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка загрузки лобби:', error);
        showMessage('Не удалось загрузить лобби', 'error');
    }
}

function closeFlamesModal() {
    const flamesModal = document.getElementById('flamesModal');
    if (flamesModal) {
        flamesModal.style.display = 'none';
    }
    selectedEventId = null;
    currentFlames = [];
}

async function displayFlames(flames) {
    const flamesList = document.getElementById('flamesList');
    if (!flamesList) {
        console.error('Элемент flamesList не найден');
        return;
    }
    
    const currentUserId = await getCurrentUser();
    console.log('Текущий пользователь ID:', currentUserId);
    
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
        const isOwnFlame = flame.user_id == currentUserId;
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

async function likeUser(userId, button) {
    try {
        console.log('Отправка лайка пользователю:', userId);
        const response = await fetch(`${API_BASE_URL}/like-user`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ light_id: userId })
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            if (button) {
                button.classList.add('liked');
                button.innerHTML = '❤️ Лайк отправлен!';
                button.disabled = true;
                
                setTimeout(() => {
                    button.style.opacity = '0.7';
                }, 1000);
            }
            
            console.log('Лайк успешно отправлен пользователю:', userId);
        } else {
            throw new Error(data.error || 'Неизвестная ошибка');
        }
    } catch (error) {
        console.error('Ошибка отправки лайка:', error);
        showMessage('Не удалось отправить лайк', 'error');
    }
}

function openCreateFlameModal() {
    const createFlameModal = document.getElementById('createFlameModal');
    const flameDescription = document.getElementById('flameDescription');
    
    if (createFlameModal) {
        createFlameModal.style.display = 'flex';
    }
    
    if (flameDescription) {
        flameDescription.value = '';
    }
}

function closeCreateFlameModal() {
    const createFlameModal = document.getElementById('createFlameModal');
    if (createFlameModal) {
        createFlameModal.style.display = 'none';
    }
}

async function createFlame() {
    const flameDescription = document.getElementById('flameDescription');
    if (!flameDescription) {
        showMessage('Элемент описания не найден', 'error');
        return;
    }
    
    const description = flameDescription.value.trim();
    
    if (!description) {
        showMessage('Введите описание лобби', 'error');
        return;
    }

    try {
        console.log('Создание лобби для мероприятия:', selectedEventId);
        const response = await fetch(`${API_BASE_URL}/flame`, {
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
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        
        if (data.status === 'ok') {
            closeCreateFlameModal();
            console.log('Лобби успешно создано для мероприятия:', selectedEventId);
            
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

function formatDate(dateString) {
    try {
        const date = new Date(dateString);
        return date.toLocaleDateString('ru-RU', {
            day: 'numeric',
            month: 'long',
            year: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
        });
    } catch (error) {
        console.error('Ошибка форматирования даты:', error);
        return dateString;
    }
}

function escapeHtml(text) {
    if (!text) return '';
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