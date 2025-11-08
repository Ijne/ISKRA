function getCurrentUser() {
    return 2;
}

// Глобальные переменные
let currentOnboardingScreen = 1;
const selectedOnboardingItems = {
    career: [],
    personality: [],
    relationship: [],
    values: [],
    music: [],
    movies: [],
    hobbies: [],
    events: []
};
let userBasicInfo = {
    age: '',
    city: ''
};

// Проверка авторизации
async function checkUserAuthorization() {
    const userId = getCurrentUser();
    console.log('Проверка пользователя:', userId);
    
    try {
        const response = await fetch(`http://localhost:8080/profile?id=${userId}`);
        
        if (!response.ok) {
            throw new Error('Ошибка HTTP: ' + response.status);
        }
        
        const userData = await response.json();
        console.log('Данные пользователя с сервера:', userData);
        
        if (userData.id) {
            return { authorized: true, userData };
        } else {
            return { authorized: false, userData };
        }
    } catch (error) {
        console.error('Ошибка при проверке авторизации:', error);
        return { authorized: false, userData: null };
    }
}

function isProfileComplete(userData) {
    const requiredFields = [
        'name', 'age', 'city', 
        'career_type', 'personality_type', 
        'relationship_goal', 'important_values'
    ];
    
    const isComplete = requiredFields.every(field => 
        userData[field] && userData[field].toString().trim() !== ''
    );
    
    console.log('Проверка заполненности профиля:', isComplete, userData);
    return isComplete;
}

// Загрузка анкеты
function loadOnboarding() {
    console.log('Загрузка анкеты...');
    
    // Сбрасываем данные
    Object.keys(selectedOnboardingItems).forEach(key => {
        selectedOnboardingItems[key] = [];
    });
    userBasicInfo = { age: '', city: '' };
    
    const mainContent = document.getElementById('mainContent');
    const body = document.body;
    
    body.classList.add('onboarding-mode');
    
    mainContent.innerHTML = `
        <div class="onboarding-container">
            <div class="onboarding-progress">
                <div class="onboarding-progress-fill" id="onboardingProgressFill"></div>
            </div>

            <!-- Экран 1: Приветствие -->
            <div class="onboarding-screen active" id="screen1">
                <div class="onboarding-header">
                    <h1 class="onboarding-title">ISKRA</h1>
                    <p class="onboarding-subtitle">Создадим твой уникальный профиль</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="avatar-section">
                        <div class="onboarding-avatar">
                            <span>IS</span>
                        </div>
                    </div>
                    
                    <p class="onboarding-subtitle">Расскажи о себе, и мы найдем тебе идеальную пару</p>
                    
                    <button class="onboarding-btn active" onclick="nextOnboardingScreen(2)">
                        Начать заполнение
                    </button>
                </div>
            </div>

            <!-- Экран 2: Основная информация -->
            <div class="onboarding-screen" id="screen2">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Основная информация</h2>
                    <p class="onboarding-subtitle">Расскажи немного о себе</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-fields-grid">
                        <div class="onboarding-input-compact">
                            <div class="onboarding-input-label">Возраст</div>
                            <input type="number" class="onboarding-input-field" id="ageInput" 
                                   placeholder="Укажите возраст" min="18" max="100"
                                   oninput="updateBasicInfo('age', this.value)">
                            <span class="onboarding-input-edit">✎</span>
                        </div>
                        
                        <div class="onboarding-input-compact">
                            <div class="onboarding-input-label">Город</div>
                            <input type="text" class="onboarding-input-field" id="cityInput" 
                                   placeholder="Укажите город"
                                   oninput="updateBasicInfo('city', this.value)">
                            <span class="onboarding-input-edit">✎</span>
                        </div>
                    </div>
                    
                    <div class="selection-required" id="screen2Message">Заполните возраст и город</div>
                    
                    <button class="onboarding-btn" id="screen2Button" onclick="nextOnboardingScreen(3)">
                        Продолжить
                    </button>
                </div>
            </div>

            <!-- Экран 3: Карьера -->
            <div class="onboarding-screen" id="screen3">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Карьера</h2>
                    <p class="onboarding-subtitle">Чем ты занимаешься?</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="careerTags"></div>
                    <div class="selection-required" id="screen3Message">Выберите вариант для продолжения</div>
                    <div class="onboarding-capsules-grid" id="careerGrid"></div>
                    
                    <button class="onboarding-btn" id="screen3Button" onclick="nextOnboardingScreen(4)">
                        Продолжить
                    </button>
                </div>
            </div>

            <!-- Экран 4: Характер -->
            <div class="onboarding-screen" id="screen4">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Характер</h2>
                    <p class="onboarding-subtitle">Какой ты человек?</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="personalityTags"></div>
                    <div class="selection-required" id="screen4Message">Выберите вариант для продолжения</div>
                    <div class="onboarding-capsules-grid" id="personalityGrid"></div>
                    
                    <button class="onboarding-btn" id="screen4Button" onclick="nextOnboardingScreen(5)">
                        Далее
                    </button>
                </div>
            </div>

            <!-- Экран 5: Цели отношений -->
            <div class="onboarding-screen" id="screen5">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Цели отношений</h2>
                    <p class="onboarding-subtitle">Что ты ищешь?</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="relationshipTags"></div>
                    <div class="selection-required" id="screen5Message">Выберите вариант для продолжения</div>
                    <div class="onboarding-capsules-grid" id="relationshipGrid"></div>
                    
                    <button class="onboarding-btn" id="screen5Button" onclick="nextOnboardingScreen(6)">
                        Далее
                    </button>
                </div>
            </div>

            <!-- Экран 6: Ценности -->
            <div class="onboarding-screen" id="screen6">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Ценности</h2>
                    <p class="onboarding-subtitle">Что для тебя важно?</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="valuesTags"></div>
                    <div class="selection-required" id="screen6Message">Выберите вариант для продолжения</div>
                    <div class="onboarding-capsules-grid" id="valuesGrid"></div>
                    
                    <button class="onboarding-btn" id="screen6Button" onclick="nextOnboardingScreen(7)">
                        Далее
                    </button>
                </div>
            </div>

            <!-- Экран 7: Музыка -->
            <div class="onboarding-screen" id="screen7">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Любимая музыка</h2>
                    <p class="onboarding-subtitle">Выбери до 3 любимых жанров</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="musicTags"></div>
                    <div class="selection-counter" id="musicCounter">Выбрано: 0/3</div>
                    <div class="selection-required" id="screen7Message">Выберите до 3 жанров</div>
                    <div class="onboarding-capsules-grid" id="musicGrid"></div>
                    
                    <button class="onboarding-btn" id="screen7Button" onclick="nextOnboardingScreen(8)">
                        Продолжить
                    </button>
                </div>
            </div>

            <!-- Экран 8: Фильмы -->
            <div class="onboarding-screen" id="screen8">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Любимые фильмы</h2>
                    <p class="onboarding-subtitle">Выбери до 3 любимых жанров</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="moviesTags"></div>
                    <div class="selection-counter" id="moviesCounter">Выбрано: 0/3</div>
                    <div class="selection-required" id="screen8Message">Выберите до 3 жанров</div>
                    <div class="onboarding-capsules-grid" id="moviesGrid"></div>
                    
                    <button class="onboarding-btn" id="screen8Button" onclick="nextOnboardingScreen(9)">
                        Продолжить
                    </button>
                </div>
            </div>

            <!-- Экран 9: Хобби -->
            <div class="onboarding-screen" id="screen9">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Хобби и увлечения</h2>
                    <p class="onboarding-subtitle">Выбери до 3 своих увлечений</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="hobbiesTags"></div>
                    <div class="selection-counter" id="hobbiesCounter">Выбрано: 0/3</div>
                    <div class="selection-required" id="screen9Message">Выберите до 3 увлечений</div>
                    <div class="onboarding-capsules-grid" id="hobbiesGrid"></div>
                    
                    <button class="onboarding-btn" id="screen9Button" onclick="nextOnboardingScreen(10)">
                        Продолжить
                    </button>
                </div>
            </div>

            <!-- Экран 10: Мероприятия -->
            <div class="onboarding-screen" id="screen10">
                <div class="onboarding-header">
                    <h2 class="profile-section-title">Мероприятия</h2>
                    <p class="onboarding-subtitle">Куда бы хотел сходить с кем-то?</p>
                </div>
                
                <div class="onboarding-board">
                    <div class="onboarding-selected-tags" id="eventsTags"></div>
                    <div class="selection-counter" id="eventsCounter">Выбрано: 0/3</div>
                    <div class="selection-required" id="screen10Message">Выберите до 3 мероприятий</div>
                    <div class="onboarding-capsules-grid" id="eventsGrid"></div>
                    
                    <button class="onboarding-btn" id="screen10Button" onclick="completeOnboarding()">
                        Завершить профиль
                    </button>
                </div>
            </div>
        </div>
    `;
    
    initOnboarding();
}

// Загрузка основного контента
function loadMainContent(userData) {
    console.log('Загрузка основного контента:', userData);
    
    const mainContent = document.getElementById('mainContent');
    const body = document.body;
    
    body.classList.remove('onboarding-mode');

    mainContent.innerHTML = `

    `;
}

// Инициализация анкеты
function initOnboarding() {
    console.log('Инициализация анкеты...');
    
    const capsuleData = {
        career: ['IT и технологии', 'Дизайн и UX', 'Медицина', 'Образование', 'Бизнес', 'Финансы', 'Маркетинг', 'Искусство', 'Музыка', 'Кино', 'Фотография', 'Архитектура', 'Инженерия', 'Недвижимость', 'Юриспруденция', 'Психология'],
        personality: ['Экстраверт', 'Интроверт', 'Амбиверт', 'Аналитик', 'Творец', 'Прагматик', 'Романтик', 'Реалист', 'Оптимист', 'Философ', 'Новатор', 'Лидер', 'Целеустремленный', 'Гибкий', 'Настойчивый', 'Командный'],
        relationship: ['Серьезные отношения', 'Дружба', 'Несерьезные отношения', 'Создание семьи', 'Поиск партнера', 'Романтика', 'Деловое партнерство', 'Творчество', 'Путешествия', 'Совместные проекты', 'Духовность', 'Карьера'],
        values: ['Любовь и забота', 'Семья', 'Карьера', 'Финансы', 'Духовность', 'Здоровье', 'Образование', 'Творчество', 'Свобода', 'Приключения', 'Безопасность', 'Экология'],
        music: ['Поп', 'Рок', 'Хип-хоп', 'Электроника', 'Джаз', 'Классика', 'R&B', 'Метал', 'Инди', 'Фолк', 'Кантри', 'Регги', 'Блюз', 'Соул', 'Диско', 'Альтернатива', 'Рэп'],
        movies: ['Комедия', 'Драма', 'Боевик', 'Триллер', 'Ужасы', 'Фантастика', 'Фэнтези', 'Мелодрама', 'Детектив', 'Приключения', 'Аниме', 'Документальный', 'Артхаус', 'Исторический', 'Криминал', 'Мюзикл'],
        hobbies: ['Спорт', 'Путешествия', 'Кулинария', 'Фотография', 'Рисование', 'Танцы', 'Йога', 'Велоспорт', 'Гейминг', 'Чтение', 'Садоводство', 'Рукоделие', 'Музыка', 'Театр', 'Кино', 'Настолки', 'Рыбалка', 'Охота', 'Авто', 'Технологии'],
        events: ['Концерты', 'Кино', 'Выставки', 'Театры', 'Фестивали', 'Спортивные события', 'Вечеринки', 'Клубы', 'Рестораны', 'Кафе', 'Пикники', 'Походы', 'Мастер-классы', 'Лекции', 'Йога-сессии', 'Танцы', 'Настольные игры', 'Караоке', 'Боулинг', 'Картинг']
    };

    Object.keys(capsuleData).forEach(category => {
        const grid = document.getElementById(`${category}Grid`);
        if (!grid) {
            console.error('Не найден контейнер для:', category);
            return;
        }

        grid.innerHTML = '';
        capsuleData[category].forEach(item => {
            const capsule = document.createElement('div');
            capsule.className = 'onboarding-capsule';
            
            if (['music', 'movies', 'hobbies', 'events'].includes(category)) {
                capsule.classList.add('multiple');
            }
            
            capsule.textContent = item;
            capsule.addEventListener('click', () => toggleOnboardingCapsule(category, item, capsule));
            grid.appendChild(capsule);
        });
        
        console.log(`Загружено ${capsuleData[category].length} элементов для ${category}`);
    });

    updateOnboardingProgress();
}

// Обновление базовой информации
function updateBasicInfo(field, value) {
    console.log(`Обновление ${field}:`, value);
    userBasicInfo[field] = value;
    checkScreen2Complete();
}

// Проверка заполненности второго экрана
function checkScreen2Complete() {
    const isComplete = userBasicInfo.age && userBasicInfo.city;
    const button = document.getElementById('screen2Button');
    const message = document.getElementById('screen2Message');
    
    console.log('Проверка экрана 2:', { isComplete, age: userBasicInfo.age, city: userBasicInfo.city });
    
    if (button) {
        if (isComplete) {
            button.classList.add('active');
            if (message) message.textContent = '';
        } else {
            button.classList.remove('active');
            if (message) message.textContent = 'Заполните возраст и город';
        }
    }
    return isComplete;
}

// Переключение капсулы
function toggleOnboardingCapsule(category, text, capsule) {
    console.log(`Клик по капсуле: ${category} - ${text}`);
    
    const index = selectedOnboardingItems[category].indexOf(text);
    const isMultiple = ['music', 'movies', 'hobbies', 'events'].includes(category);
    const maxSelection = 3;
    
    if (index === -1) {
        if (isMultiple) {
            if (selectedOnboardingItems[category].length >= maxSelection) {
                console.log('Достигнут лимит выбора для', category);
                return;
            }
            selectedOnboardingItems[category].push(text);
            capsule.classList.add('selected');
        } else {
            document.querySelectorAll(`#${category}Grid .onboarding-capsule`).forEach(c => {
                c.classList.remove('selected');
            });
            selectedOnboardingItems[category] = [text];
            capsule.classList.add('selected');
        }
    } else {
        selectedOnboardingItems[category].splice(index, 1);
        capsule.classList.remove('selected');
    }
    
    console.log(`Текущий выбор для ${category}:`, selectedOnboardingItems[category]);
    
    updateOnboardingTags(category);
    updateSelectionCounter(category);
    
    if (isMultiple) {
        updateMultipleSelectionButtonState(category);
    } else {
        updateCapsulesButtonState(category);
    }
}

// Обновление тегов
function updateOnboardingTags(category) {
    const tagsContainer = document.getElementById(`${category}Tags`);
    if (!tagsContainer) return;
    
    tagsContainer.innerHTML = '';
    
    selectedOnboardingItems[category].forEach(item => {
        const tag = document.createElement('div');
        tag.className = 'onboarding-selected-tag';
        tag.innerHTML = `${item} <span class="remove-tag" onclick="removeSelectedItem('${category}', '${item}')">×</span>`;
        tagsContainer.appendChild(tag);
    });
}

// Удаление выбранного элемента
function removeSelectedItem(category, item) {
    console.log(`Удаление: ${category} - ${item}`);
    
    const index = selectedOnboardingItems[category].indexOf(item);
    if (index !== -1) {
        selectedOnboardingItems[category].splice(index, 1);
        
        const grid = document.getElementById(`${category}Grid`);
        if (grid) {
            const capsules = grid.querySelectorAll('.onboarding-capsule');
            capsules.forEach(capsule => {
                if (capsule.textContent === item) {
                    capsule.classList.remove('selected');
                }
            });
        }
        
        updateOnboardingTags(category);
        updateSelectionCounter(category);
        
        if (['music', 'movies', 'hobbies', 'events'].includes(category)) {
            updateMultipleSelectionButtonState(category);
        } else {
            updateCapsulesButtonState(category);
        }
    }
}

// Обновление счетчика выбора
function updateSelectionCounter(category) {
    const counter = document.getElementById(`${category}Counter`);
    if (!counter) return;
    
    const count = selectedOnboardingItems[category].length;
    const maxSelection = 3;
    counter.textContent = `Выбрано: ${count}/${maxSelection}`;
    
    if (count >= maxSelection) {
        counter.style.color = '#ffaa00';
    } else {
        counter.style.color = 'rgba(255, 170, 0, 0.7)';
    }
}

// Обновление состояния кнопок для одиночного выбора
function updateCapsulesButtonState(category) {
    const screenNumber = getScreenByCategory(category);
    const button = document.getElementById(`screen${screenNumber}Button`);
    const message = document.getElementById(`screen${screenNumber}Message`);
    
    if (button) {
        const hasSelection = selectedOnboardingItems[category].length > 0;
        console.log(`Обновление кнопки экрана ${screenNumber}:`, hasSelection);
        
        if (hasSelection) {
            button.classList.add('active');
            if (message) message.textContent = '';
        } else {
            button.classList.remove('active');
            if (message) message.textContent = 'Выберите вариант для продолжения';
        }
    }
}

// Обновление состояния кнопки для множественного выбора
function updateMultipleSelectionButtonState(category) {
    const screenNumber = getScreenByCategory(category);
    const button = document.getElementById(`screen${screenNumber}Button`);
    const message = document.getElementById(`screen${screenNumber}Message`);
    
    if (button) {
        const count = selectedOnboardingItems[category].length;
        console.log(`Обновление кнопки множественного выбора ${screenNumber}:`, count);
        
        if (count > 0) {
            button.classList.add('active');
            if (message) message.textContent = '';
        } else {
            button.classList.remove('active');
            if (message) message.textContent = 'Выберите до 3 вариантов';
        }
    }
}

// Вспомогательные функции
function getCategoryByScreen(screenNumber) {
    const categories = ['career', 'personality', 'relationship', 'values', 'music', 'movies', 'hobbies', 'events'];
    return categories[screenNumber - 3] || 'career';
}

function getScreenByCategory(category) {
    const categories = ['career', 'personality', 'relationship', 'values', 'music', 'movies', 'hobbies', 'events'];
    return categories.indexOf(category) + 3;
}

// Переход между экранами
function nextOnboardingScreen(screenNumber) {
    console.log(`Переход с экрана ${currentOnboardingScreen} на ${screenNumber}`);
    
    if (currentOnboardingScreen === 2 && !checkScreen2Complete()) {
        console.log('Нельзя перейти - не заполнены основные поля');
        return;
    }
    
    if (currentOnboardingScreen >= 3 && currentOnboardingScreen <= 6) {
        const currentCategory = getCategoryByScreen(currentOnboardingScreen);
        if (selectedOnboardingItems[currentCategory].length === 0) {
            console.log('Нельзя перейти - не выбран вариант');
            return;
        }
    }
    
    if (currentOnboardingScreen >= 7) {
        const currentCategory = getCategoryByScreen(currentOnboardingScreen);
        if (selectedOnboardingItems[currentCategory].length === 0) {
            console.log('Нельзя перейти - не выбрано ни одного варианта');
            return;
        }
    }
    
    const currentScreen = document.getElementById(`screen${currentOnboardingScreen}`);
    const nextScreen = document.getElementById(`screen${screenNumber}`);
    
    if (currentScreen) currentScreen.classList.remove('active');
    if (nextScreen) nextScreen.classList.add('active');
    
    currentOnboardingScreen = screenNumber;
    updateOnboardingProgress();
}

// Обновление прогресса
function updateOnboardingProgress() {
    const progressFill = document.getElementById('onboardingProgressFill');
    if (!progressFill) return;
    
    const progress = (currentOnboardingScreen - 1) / 9 * 100;
    progressFill.style.width = progress + '%';
    console.log(`Прогресс: ${progress}%`);
}

// Завершение онбординга
async function completeOnboarding() {
    console.log('Завершение онбординга...');
    console.log('Собранные данные:', {
        basic: userBasicInfo,
        selections: selectedOnboardingItems
    });
    
    try {
        const profileData = {
            id: getCurrentUser(),
            age: parseInt(userBasicInfo.age),
            city: userBasicInfo.city,
            career_type: selectedOnboardingItems.career[0] || '',
            personality_type: selectedOnboardingItems.personality[0] || '',
            relationship_goal: selectedOnboardingItems.relationship[0] || '',
            important_values: selectedOnboardingItems.values[0] || '',
            music: selectedOnboardingItems.music.join(',') || '',
            films: selectedOnboardingItems.movies.join(',') || '',
            hobbies: selectedOnboardingItems.hobbies.join(',') || '',
            event_preferences: selectedOnboardingItems.events.join(',') || ''
        };

        console.log('Отправка данных на сервер:', profileData);
        
        const response = await fetch('http://localhost:8080/createuser', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(profileData)
        });
        
        if (response.ok) {
            console.log(response)
            const result = await response.json();
            console.log('Успешный ответ сервера:', result);
            //alert('Профиль успешно сохранен! 🎉');
            //setTimeout(() => {
            //    location.reload();
            //}, 1500);
        } else {
            //const errorText = await response.text();
            //console.error('Ошибка сервера:', response.status, errorText);
            //throw new Error(`Ошибка сервера: ${response.status} - ${errorText}`);
            location.reload()
        }
    } catch (error) {
        //console.error('Ошибка при сохранении профиля:', error);
        //alert('Ошибка при сохранении профиля. Попробуйте еще раз.');
        location.reload()
    }
}

function editProfile() {
    console.log('Редактирование профиля...');
    loadOnboarding();
}

// Основная функция
async function initApp() {
    console.log('Инициализация приложения...');
    const authStatus = await checkUserAuthorization();
    console.log('Статус авторизации:', authStatus);
    
    if (authStatus.authorized) {
        loadMainContent(authStatus.userData);
    } else {
        loadOnboarding();
    }
}

// Запуск
document.addEventListener('DOMContentLoaded', initApp);