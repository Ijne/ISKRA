let initData = null;
let WebApp = null;

function waitForWebApp() {
    return new Promise((resolve, reject) => {
        if (window.WebApp) {
            WebApp = window.WebApp;
            initData = window.WebApp?.initData;
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
            return 0;
        }

        let decodedString;
        
        if (typeof initData === 'object') {
            const user = initData.user || initData;
            return user.id || 0;
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
                        return userData.id || 0;
                    } catch (e) {
                        console.error('Error parsing user data:', e);
                    }
                }
                return 0;
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
            

            if (calculatedHash === receivedHash) {
                
                if (userParam) {
                    try {
                        const userData = JSON.parse(userParam);
                        return userData.id || 0;
                    } catch (parseError) {
                        console.error('Error parsing user data:', parseError);
                        return 0;
                    }
                }
            } else {
                return 0;
            }
        }
        
        return null;
    } catch (error) {
        console.error('Validation error:', error);
        return null;
    }
}

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

let userGender = 0;
let preferredGender = 1;

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

async function fetchUserProfile() {
    try {
        const userId = await getCurrentUser();

        const response = await fetch(`http://localhost:8080/profile?id=${userId}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
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
        const userId = await getCurrentUser();
        console.log('Updating profile for user:', userId);
        console.log('Profile data:', profileData);
        
        profileData.id = userId;

        const response = await fetch(`http://localhost:8080/updateuser?id=${userId}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(profileData)
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const result = await response.json();
        console.log('Profile updated successfully:', result);
        return result;
    } catch (error) {
        
    }
}

function parseServerData(userData) {
    if (!userData) return {
        name: '',
        surname: '',
        age: '',
        city: '',
        careerPlace: ''
    };

    const parseArray = (str) => {
        if (!str || str === 'null' || str === 'undefined') return [];
        return str.split(',').map(item => item.trim()).filter(item => item && item !== 'null' && item !== 'undefined');
    };
    
    selectedItems.career = parseArray(userData.career_type).slice(0, maxSelections.career);
    selectedItems.personality = parseArray(userData.personality_type).slice(0, maxSelections.personality);
    selectedItems.relationship = parseArray(userData.relationship_goal).slice(0, maxSelections.relationship);
    selectedItems.values = parseArray(userData.important_values).slice(0, maxSelections.values);
    selectedItems.music = parseArray(userData.music).slice(0, maxSelections.music);
    selectedItems.hobbies = parseArray(userData.hobbies).slice(0, maxSelections.hobbies);
    selectedItems.movies = parseArray(userData.films).slice(0, maxSelections.movies);
    selectedItems.events = parseArray(userData.event_preferences).slice(0, maxSelections.events);

    userGender = userData.gender !== undefined ? parseInt(userData.gender) : 0;
    preferredGender = userData.preferred_gender !== undefined ? parseInt(userData.preferred_gender) : 1;
    
    return {
        name: userData.name || '',
        surname: userData.surname || '',
        age: userData.age || '',
        city: userData.city || '',
        careerPlace: userData.career_place || ''
    };
}

function prepareDataForServer(profileData) {
    return {
        id: currentUserId,
        name: profileData.name || '',
        surname: profileData.surname || '',
        age: parseInt(profileData.age) || 0,
        city: profileData.city || '',
        gender: userGender,
        preferred_gender: preferredGender,
        career_place: profileData.careerPlace || '',
        career_type: (selectedItems.career || []).join(','),
        personality_type: (selectedItems.personality || []).join(','),
        relationship_goal: (selectedItems.relationship || []).join(','),
        important_values: (selectedItems.values || []).join(','),
        music: (selectedItems.music || []).join(','),
        hobbies: (selectedItems.hobbies || []).join(','),
        films: (selectedItems.movies || []).join(','),
        event_preferences: (selectedItems.events || []).join(',')
    };
}

async function loadProfileData() {
    try {
        const userData = await fetchUserProfile();
        const parsedData = parseServerData(userData);

        updateUIElement('nameValue', parsedData.name);
        updateUIElement('ageValue', parsedData.age ? parsedData.age + ' лет' : '');
        updateUIElement('cityValue', parsedData.city);

        updateGenderDisplay();
        updatePreferredGenderDisplay();
        updateAllSelectedCapsules();
    } catch (error) {
        console.error('Ошибка загрузки данных профиля:', error);
    }
}

function updateUIElement(elementId, value) {
    const element = document.getElementById(elementId);
    if (element) {
        element.textContent = value;
    }
}

function updateGenderDisplay() {
    const genderValue = document.getElementById('genderValue');
    if (genderValue) {
        genderValue.textContent = userGender === 0 ? 'Мужской' : 'Женский';
    }
}

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

function updateAllSelectedCapsules() {
    Object.keys(selectedItems).forEach(category => {
        updateSelectedCapsulesForCategory(category);
    });
}

function updateSelectedCapsulesForCategory(category) {
    const grid = document.getElementById(`${category}Grid`);
    const tagsContainer = document.getElementById(`${category}Tags`);
    
    if (!grid || !tagsContainer) return;

    tagsContainer.innerHTML = '';

    selectedItems[category].forEach(item => {
        addTag(category, item, tagsContainer);
    });

    const capsules = grid.querySelectorAll('.capsule');
    capsules.forEach(capsule => {
        const capsuleText = capsule.textContent;

        capsule.classList.remove('selected', 'disabled');

        if (selectedItems[category].includes(capsuleText)) {
            capsule.classList.add('selected');
        }

        if (selectedItems[category].length >= maxSelections[category] && 
            !selectedItems[category].includes(capsuleText)) {
            capsule.classList.add('disabled');
        }
    });
}

function toggleExpand(gridId, button) {
    const grid = document.getElementById(gridId);
    if (!grid) return;

    const isExpanded = grid.classList.contains('expanded');
    
    if (isExpanded) {
        grid.classList.remove('expanded');
        button.querySelector('span:first-child').textContent = 'Показать все';
    } else {
        grid.classList.add('expanded');
        button.querySelector('span:first-child').textContent = 'Скрыть';
    }

    const icon = button.querySelector('.expand-icon');
    if (icon) {
        icon.style.transform = isExpanded ? 'rotate(0deg)' : 'rotate(180deg)';
    }
}

async function initCapsules() {
    try {
        await waitForWebApp();
        
        Object.keys(capsuleData).forEach(category => {
            const grid = document.getElementById(`${category}Grid`);
            const tagsContainer = document.getElementById(`${category}Tags`);
            
            if (!grid) {
                console.warn(`Element with id ${category}Grid not found`);
                return;
            }

            const shuffledItems = [...capsuleData[category]].sort(() => Math.random() - 0.5);
            
            shuffledItems.forEach(item => {
                const capsule = document.createElement('div');
                capsule.className = 'capsule';
                capsule.textContent = item;
                
                capsule.addEventListener('click', () => toggleCapsule(category, item, capsule, tagsContainer));
                grid.appendChild(capsule);
            });
        });

        await loadProfileData();
        makeNameFieldReadonly();

    } catch (error) {
        console.error('Ошибка инициализации капсул:', error);
    }
}

function makeNameFieldReadonly() {
    const nameField = document.querySelector('.readonly-field');
    if (nameField) {
        nameField.style.pointerEvents = 'none';
        nameField.style.opacity = '0.7';
        nameField.style.cursor = 'default';
        nameField.title = 'Имя нельзя изменить';
    }
}

function toggleCapsule(category, text, capsule, tagsContainer) {
    const maxSelect = maxSelections[category];
    const currentSelected = selectedItems[category];

    if (currentSelected.includes(text)) {
        selectedItems[category] = currentSelected.filter(item => item !== text);
    } 
    else if (currentSelected.length < maxSelect) {
        if (maxSelect === 1) {
            selectedItems[category] = [text];
        } else {
            selectedItems[category] = [...currentSelected, text];
        }
    }
    else {
        return;
    }

    updateSelectedCapsulesForCategory(category);
}

function addTag(category, text, container) {
    const tag = document.createElement('div');
    tag.className = 'selected-tag';
    tag.innerHTML = `
        ${text}
        <span class="remove-tag" onclick="removeTag('${category}', '${text}')">×</span>
    `;
    container.appendChild(tag);
}

function removeTag(category, text) {
    selectedItems[category] = selectedItems[category].filter(item => item !== text);
    updateSelectedCapsulesForCategory(category);
}

function openGenderModal() {
    const modal = document.getElementById('genderModal');
    if (modal) {
        modal.style.display = 'flex';
        
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

function openModal(field) {
    if (field === 'name') {
        return;
    }
    
    currentField = field;
    const modal = document.getElementById('editModal');
    const title = document.getElementById('modalTitle');
    const input = document.getElementById('modalInput');
    
    const fieldTitles = {
        age: 'Возраст',
        city: 'Город'
    };
    
    if (!modal || !title || !input) {
        console.error('Modal elements not found');
        return;
    }
    
    title.textContent = `Редактирование ${fieldTitles[field]}`;
    
    const valueElement = document.getElementById(`${field}Value`);
    if (valueElement) {
        if (field === 'age') {
            input.value = valueElement.textContent.replace(' лет', '');
        } else {
            input.value = valueElement.textContent;
        }
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
    } else if (currentField === 'city') {
        valueElement.textContent = value;
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
            gender: userGender,
            preferredGender: preferredGender
        };
        
        const serverData = prepareDataForServer(profileData);
        
        await updateUserProfile(serverData);
        
        console.log('Профиль успешно сохранен!');
    } catch (error) {
        console.error('Ошибка при сохранении профиля:', error);
    }
}

if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initCapsules);
} else {
    initCapsules();
}