function getCurrentUser() {
    return 2;
}

// Богатые наборы данных для капсул
const capsuleData = {
    career: [
        'IT и технологии', 'Дизайн и UX', 'Медицина и здоровье', 'Образование и наука',
        'Бизнес и предпринимательство', 'Финансы и инвестиции', 'Маркетинг и реклама',
        'Искусство и творчество', 'Музыка и звук', 'Кино и видео', 'Фотография',
        'Архитектура и дизайн', 'Инженерия и техника', 'Строительство и недвижимость',
        'Юриспруденция и право', 'Психология и коучинг', 'Консалтинг и аналитика',
        'HR и рекрутинг', 'Продажи и переговоры', 'Логистика и снабжение',
        'Туризм и гостиничный бизнес', 'Ресторанный бизнес', 'Фитнес и спорт',
        'Мода и стиль', 'Красота и уход', 'Журналистика и медиа', 'Писательство',
        'Экология и устойчивое развитие', 'Некоммерческий сектор', 'Государственная служба'
    ],
    personality: [
        'Экстраверт', 'Интроверт', 'Амбиверт', 'Аналитический склад ума',
        'Творческая личность', 'Прагматик', 'Романтик', 'Реалист',
        'Оптимист', 'Пессимист', 'Мечтатель', 'Лидер',
        'Исполнитель', 'Новатор', 'Традиционалист', 'Спонтанный',
        'Планировщик', 'Эмпат', 'Логик', 'Энергичный', 'Спокойный',
        'Целеустремленный', 'Гибкий', 'Настойчивый', 'Командный игрок'
    ],
    relationship: [
        'Серьезные отношения', 'Дружба и общение', 'Несерьезные отношения',
        'Создание семьи', 'Поиск партнера для жизни', 'Романтические отношения',
        'Деловое партнерство', 'Творческое сотрудничество', 'Путешествия вместе',
        'Совместные проекты', 'Духовное развитие', 'Карьерный рост вместе',
        'Спортивные увлечения', 'Интеллектуальное общение', 'Культурный обмен',
        'Взаимное развитие', 'Эмоциональная поддержка', 'Финансовое партнерство'
    ],
    values: [
        'Любовь и забота', 'Семейные ценности', 'Карьерный рост', 'Финансовая стабильность',
        'Духовное развитие', 'Здоровый образ жизни', 'Образование и знания', 'Творчество',
        'Свобода и независимость', 'Приключения и путешествия', 'Стабильность и безопасность',
        'Социальная ответственность', 'Экологичность', 'Технологический прогресс',
        'Традиции и культура', 'Инновации и изменения', 'Гармония и баланс',
        'Успех и достижения', 'Слава и признание', 'Скромность и простота',
        'Щедрость и благотворительность', 'Справедливость и равенство', 'Честность и прозрачность'
    ],
    music: [
        'Рок', 'Поп', 'Хип-хоп', 'Электронная', 'Джаз', 'Классическая',
        'Альтернативная', 'Инди', 'R&B', 'Фолк', 'Кантри', 'Регги',
        'Метал', 'Панк', 'Блюз', 'Соул', 'Фанк', 'Диско',
        'Хаус', 'Техно', 'Транс', 'Драм-н-бейс', 'Лоу-фай', 'Эмбиент'
    ],
    hobbies: [
        'Путешествия', 'Фотография', 'Кулинария', 'Спорт', 'Йога', 'Танцы',
        'Рисование', 'Пение', 'Игра на гитаре', 'Программирование', 'Чтение',
        'Писательство', 'Садоводство', 'Рыбалка', 'Охота', 'Велоспорт',
        'Скалолазание', 'Серфинг', 'Горные лыжи', 'Шахматы', 'Настольные игры',
        'Коллекционирование', 'Астрология', 'Медитация', 'Бег', 'Плавание'
    ],
    movies: [
        'Комедии', 'Драмы', 'Триллеры', 'Ужасы', 'Фантастика', 'Фэнтези',
        'Боевики', 'Приключения', 'Мелодрамы', 'Детективы', 'Криминал',
        'Исторические', 'Биографии', 'Документальные', 'Аниме', 'Мультфильмы',
        'Арт-хаус', 'Нуар', 'Вестерны', 'Мюзиклы', 'Романтические',
        'Мистика', 'Катастрофы', 'Военные', 'Спортивные'
    ],
    events: [
        'Концерты', 'Фестивали', 'Выставки', 'Театры', 'Кино', 'Вечеринки',
        'Спортивные события', 'Конференции', 'Воркшопы', 'Мастер-классы',
        'Гастрономические туры', 'Винные дегустации', 'Квесты', 'Караоке',
        'Настольные игры', 'Пикники', 'Походы', 'Кемпинг', 'Пляжные вечеринки',
        'Боулинг', 'Бильярд', 'Картинг', 'Каток', 'Йога на природе',
        'Благотворительные мероприятия', 'Косплей-фесты', 'Книжные клубы'
    ]
};

let currentField = '';
let currentUserId = null;
const selectedItems = {
    career: [],
    personality: [],
    relationship: [],
    values: [],
    music: [],
    hobbies: [],
    movies: [],
    events: []
};

// Данные о поле и предпочтительном поле
let userGender = 0; // 0 - мужской, 1 - женский
let preferredGender = 1; // 0 - мужчин, 1 - женщин, 2 - всех

// Максимальное количество выборов для каждой категории
const maxSelections = {
    career: 1,
    personality: 1,
    relationship: 1,
    values: 1,
    music: 3,
    hobbies: 3,
    movies: 3,
    events: 3
};

// Базовый URL API
const API_BASE_URL = 'http://localhost:8080/profile?id=';

// Функции для работы с API
async function fetchUserProfile() {
    try {
        const response = await fetch(API_BASE_URL + getCurrentUser(), {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        
        if (!response.ok) {
            throw new Error('Ошибка при загрузке профиля');
        }
        
        const userData = await response.json();
        currentUserId = userData.id;
        return userData;
    } catch (error) {
        console.error('Ошибка загрузки профиля:', error);
        return null;
    }
}

async function updateUserProfile(profileData) {
    try {
        const response = await fetch(`http://localhost:8080/updateuser?id=`+getCurrentUser(), {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(profileData)
        });
        
        if (!response.ok) {
            throw new Error('Ошибка при обновлении профиля');
        }
        
        return await response.json();
    } catch (error) {
        console.error('Ошибка обновления профиля:', error);
        throw error;
    }
}

// Преобразование данных из формата сервера в формат клиента
function parseServerData(userData) {
    // Преобразуем строки с разделителями в массивы
    const parseArray = (str) => str ? str.split(',').map(item => item.trim()).filter(item => item) : [];
    
    selectedItems.career = parseArray(userData.career_type).slice(0, maxSelections.career);
    selectedItems.personality = parseArray(userData.personality_type).slice(0, maxSelections.personality);
    selectedItems.relationship = parseArray(userData.relationship_goal).slice(0, maxSelections.relationship);
    selectedItems.values = parseArray(userData.important_values).slice(0, maxSelections.values);
    selectedItems.music = parseArray(userData.music).slice(0, maxSelections.music);
    selectedItems.hobbies = parseArray(userData.hobbies).slice(0, maxSelections.hobbies);
    selectedItems.movies = parseArray(userData.films).slice(0, maxSelections.movies);
    selectedItems.events = parseArray(userData.event_preferences).slice(0, maxSelections.events);
    
    // Обрабатываем пол и предпочтительный пол
    userGender = userData.gender || 0;
    preferredGender = userData.preferred_gender !== undefined ? userData.preferred_gender : 1;
    
    return {
        name: userData.name || '',
        surname: userData.surname || '',
        age: userData.age || '',
        city: userData.city || '',
        careerPlace: userData.career_place || ''
    };
}

// Преобразование данных из формата клиента в формат сервера
function prepareDataForServer(profileData) {
    return {
        id: currentUserId,
        name: profileData.name,
        surname: profileData.surname || '',
        age: parseInt(profileData.age) || 0,
        city: profileData.city,
        gender: userGender,
        preferred_gender: preferredGender,
        career_place: profileData.careerPlace || '',
        career_type: profileData.career.join(', '),
        personality_type: profileData.personality.join(', '),
        relationship_goal: profileData.relationship.join(', '),
        important_values: profileData.values.join(', '),
        music: profileData.music.join(', '),
        hobbies: profileData.hobbies.join(', '),
        films: profileData.movies.join(', '),
        event_preferences: profileData.events.join(', ')
    };
}

// Загрузка данных профиля с сервера
async function loadProfileData() {
    const userData = await fetchUserProfile();
    if (userData) {
        const parsedData = parseServerData(userData);
        
        // Обновляем интерфейс
        document.getElementById('nameValue').textContent = parsedData.name;
        document.getElementById('ageValue').textContent = parsedData.age ? parsedData.age + ' лет' : '';
        document.getElementById('cityValue').textContent = parsedData.city;
        
        // Обновляем отображение пола
        updateGenderDisplay();
        updatePreferredGenderDisplay();
        
        // Обновляем выбранные капсулы для всех категорий
        updateAllSelectedCapsules();
    }
}

// Обновление отображения пола
function updateGenderDisplay() {
    const genderValue = document.getElementById('genderValue');
    if (genderValue) {
        genderValue.textContent = userGender === 0 ? 'Мужской' : 'Женский';
    }
}

// Обновление отображения предпочтительного пола
function updatePreferredGenderDisplay() {
    const preferredGenderValue = document.getElementById('preferredGenderValue');
    if (preferredGenderValue) {
        switch(preferredGender) {
            case 0:
                preferredGenderValue.textContent = 'Мужчин';
                break;
            case 1:
                preferredGenderValue.textContent = 'Женщин';
                break;
            case 2:
                preferredGenderValue.textContent = 'Всех';
                break;
            default:
                preferredGenderValue.textContent = 'Женщин';
        }
    }
}

// Обновление ВСЕХ выбранных капсул в интерфейсе (только при загрузке)
function updateAllSelectedCapsules() {
    Object.keys(selectedItems).forEach(category => {
        updateSelectedCapsulesForCategory(category);
    });
}

// Обновление выбранных капсул для КОНКРЕТНОЙ категории
function updateSelectedCapsulesForCategory(category) {
    const grid = document.getElementById(`${category}Grid`);
    const tagsContainer = document.getElementById(`${category}Tags`);
    
    if (!grid || !tagsContainer) return;
    
    // Очищаем контейнер тегов
    tagsContainer.innerHTML = '';
    
    // Добавляем теги для выбранных элементов
    selectedItems[category].forEach(item => {
        addTag(category, item, tagsContainer);
    });
    
    // Обновляем состояние капсул
    const capsules = grid.querySelectorAll('.capsule');
    capsules.forEach(capsule => {
        const capsuleText = capsule.textContent;
        
        // Сбрасываем класс selected
        capsule.classList.remove('selected');
        
        // Добавляем класс selected только если элемент выбран
        if (selectedItems[category].includes(capsuleText)) {
            capsule.classList.add('selected');
        }
        
        // Для категорий с ограничением показываем, если достигнут лимит
        if (selectedItems[category].length >= maxSelections[category] && 
            !selectedItems[category].includes(capsuleText)) {
            capsule.classList.add('disabled');
        } else {
            capsule.classList.remove('disabled');
        }
    });
}

// Переключение раскрытия/скрытия
function toggleExpand(gridId, button) {
    const grid = document.getElementById(gridId);
    const isExpanded = grid.classList.contains('expanded');
    
    if (isExpanded) {
        grid.classList.remove('expanded');
        button.querySelector('span:first-child').textContent = 'Показать все';
    } else {
        grid.classList.add('expanded');
        button.querySelector('span:first-child').textContent = 'Скрыть';
    }
    
    // Анимируем иконку
    const icon = button.querySelector('.expand-icon');
    icon.style.transform = isExpanded ? 'rotate(0deg)' : 'rotate(180deg)';
}

// Инициализация капсул
function initCapsules() {
    Object.keys(capsuleData).forEach(category => {
        const grid = document.getElementById(`${category}Grid`);
        const tagsContainer = document.getElementById(`${category}Tags`);
        
        if (!grid) {
            console.error(`Element with id ${category}Grid not found`);
            return;
        }
        
        // Перемешиваем массив для случайного порядка
        const shuffledItems = [...capsuleData[category]].sort(() => Math.random() - 0.5);
        
        shuffledItems.forEach(item => {
            const capsule = document.createElement('div');
            capsule.className = 'capsule';
            capsule.textContent = item;
            
            capsule.addEventListener('click', () => toggleCapsule(category, item, capsule, tagsContainer));
            grid.appendChild(capsule);
        });
    });
    
    // Загружаем данные профиля после инициализации капсул
    loadProfileData();
    
    // Делаем имя и город неизменяемыми
    makeFieldsReadonly();
}

// Делаем поля имя и город неизменяемыми
function makeFieldsReadonly() {
    const nameField = document.querySelector('[onclick="openModal(\'name\')"]');
    const cityField = document.querySelector('[onclick="openModal(\'city\')"]');
    
    if (nameField) {
        nameField.style.pointerEvents = 'none';
        nameField.style.opacity = '0.7';
        nameField.style.cursor = 'default';
        nameField.title = 'Имя нельзя изменить';
    }
    
    if (cityField) {
        cityField.style.pointerEvents = 'none';
        cityField.style.opacity = '0.7';
        cityField.style.cursor = 'default';
        cityField.title = 'Город нельзя изменить';
    }
}

// Переключение капсулы - с учетом ограничений
function toggleCapsule(category, text, capsule, tagsContainer) {
    const maxSelect = maxSelections[category];
    const currentSelected = selectedItems[category];
    
    // Если уже выбран этот элемент - убираем его
    if (currentSelected.includes(text)) {
        selectedItems[category] = currentSelected.filter(item => item !== text);
    } 
    // Если можно добавить еще элементы
    else if (currentSelected.length < maxSelect) {
        // Для категорий с одним выбором заменяем, для нескольких добавляем
        if (maxSelect === 1) {
            selectedItems[category] = [text];
        } else {
            selectedItems[category] = [...currentSelected, text];
        }
    }
    // Если достигнут лимит - ничего не делаем
    else {
        return;
    }
    
    // Обновляем отображение ТОЛЬКО для этой категории
    updateSelectedCapsulesForCategory(category);
}

// Добавление тега с возможностью удаления
function addTag(category, text, container) {
    const tag = document.createElement('div');
    tag.className = 'selected-tag';
    tag.innerHTML = `
        ${text}
        <span class="remove-tag" onclick="removeTag('${category}', '${text}')">×</span>
    `;
    container.appendChild(tag);
}

// Удаление тега
function removeTag(category, text) {
    selectedItems[category] = selectedItems[category].filter(item => item !== text);
    updateSelectedCapsulesForCategory(category);
}

// Функции для работы с модальными окнами пола
function openGenderModal() {
    const modal = document.getElementById('genderModal');
    if (modal) {
        modal.style.display = 'flex';
        
        // Подсвечиваем текущий выбор
        const options = modal.querySelectorAll('.gender-option');
        options.forEach(option => option.classList.remove('selected'));
        
        if (userGender === 0) {
            options[0].classList.add('selected');
        } else {
            options[1].classList.add('selected');
        }
    }
}

function closeGenderModal() {
    const modal = document.getElementById('genderModal');
    if (modal) {
        modal.style.display = 'none';
    }
}

function selectGender(gender, label) {
    userGender = gender;
    updateGenderDisplay();
    closeGenderModal();
}

function openPreferredGenderModal() {
    const modal = document.getElementById('preferredGenderModal');
    if (modal) {
        modal.style.display = 'flex';
        
        // Подсвечиваем текущий выбор
        const options = modal.querySelectorAll('.gender-option');
        options.forEach(option => option.classList.remove('selected'));
        options[preferredGender].classList.add('selected');
    }
}

function closePreferredGenderModal() {
    const modal = document.getElementById('preferredGenderModal');
    if (modal) {
        modal.style.display = 'none';
    }
}

function selectPreferredGender(gender, label) {
    preferredGender = gender;
    updatePreferredGenderDisplay();
    closePreferredGenderModal();
}

// Модальное окно - только для возраста
function openModal(field) {
    // Блокируем открытие модального окна для имени и города
    if (field === 'name' || field === 'city') {
        return;
    }
    
    currentField = field;
    const modal = document.getElementById('editModal');
    const title = document.getElementById('modalTitle');
    const input = document.getElementById('modalInput');
    
    const fieldTitles = {
        age: 'Возраст'
    };
    
    if (!modal || !title || !input) {
        console.error('Modal elements not found');
        return;
    }
    
    title.textContent = `Редактирование ${fieldTitles[field]}`;
    
    const valueElement = document.getElementById(`${field}Value`);
    if (valueElement) {
        input.value = valueElement.textContent.replace(' лет', '');
    }
    
    input.placeholder = `Введите ${fieldTitles[field].toLowerCase()}`;
    
    modal.style.display = 'flex';
    setTimeout(() => input.focus(), 100);
}

function closeModal() {
    const modal = document.getElementById('editModal');
    if (modal) {
        modal.style.display = 'none';
    }
}

function saveField() {
    const value = document.getElementById('modalInput').value;
    const valueElement = document.getElementById(`${currentField}Value`);
    
    if (!valueElement) {
        console.error(`Element with id ${currentField}Value not found`);
        return;
    }
    
    if (currentField === 'age') {
        valueElement.textContent = value + ' лет';
    }
    
    closeModal();
}

async function saveProfile() {
    try {
        const profileData = {
            name: document.getElementById('nameValue')?.textContent || '',
            age: document.getElementById('ageValue')?.textContent.replace(' лет', '') || '',
            city: document.getElementById('cityValue')?.textContent || '',
            careerPlace: '',
            career: selectedItems.career,
            personality: selectedItems.personality,
            relationship: selectedItems.relationship,
            values: selectedItems.values,
            music: selectedItems.music,
            hobbies: selectedItems.hobbies,
            movies: selectedItems.movies,
            events: selectedItems.events
        };
        
        const serverData = prepareDataForServer(profileData);
        
        await updateUserProfile(serverData);
        
        //alert('Профиль успешно сохранен!');
    } catch (error) {
        //console.error('Ошибка при сохранении профиля:', error);
        //alert('Ошибка при сохранении профиля. Попробуйте еще раз.');
    }
}

// Инициализация при загрузке DOM
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initCapsules);
} else {
    initCapsules();
}