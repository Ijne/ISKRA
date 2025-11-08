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
    ]
};

let currentField = '';
const selectedItems = {
    career: [],
    personality: [],
    relationship: [],
    values: []
};

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
}

// Переключение капсулы
function toggleCapsule(category, text, capsule, tagsContainer) {
    const index = selectedItems[category].indexOf(text);
    const isSingleSelect = category !== 'values' && category !== 'career';
    
    if (index === -1) {
        if (isSingleSelect) {
            // Снимаем выделение со всех капсул в этой категории
            document.querySelectorAll(`#${category}Grid .capsule`).forEach(c => {
                c.classList.remove('selected');
            });
            selectedItems[category] = [text];
        } else {
            selectedItems[category].push(text);
        }
        capsule.classList.add('selected');
        addTag(category, text, tagsContainer);
    } else {
        selectedItems[category].splice(index, 1);
        capsule.classList.remove('selected');
        removeTag(category, text, tagsContainer);
    }
}

// Добавление тега
function addTag(category, text, container) {
    const tag = document.createElement('div');
    tag.className = 'selected-tag';
    tag.innerHTML = `
        ${text}
        <span class="remove-tag" onclick="removeTagByElement('${category}', this.parentElement)">×</span>
    `;
    container.appendChild(tag);
}

// Удаление тега по элементу
function removeTagByElement(category, tagElement) {
    const text = tagElement.textContent.replace('×', '').trim();
    selectedItems[category] = selectedItems[category].filter(item => item !== text);
    
    // Снимаем выделение с соответствующей капсулы
    const capsules = document.querySelectorAll(`#${category}Grid .capsule`);
    capsules.forEach(capsule => {
        if (capsule.textContent === text) {
            capsule.classList.remove('selected');
        }
    });
    
    tagElement.remove();
}

// Удаление тега
function removeTag(category, text, container) {
    const tags = container.querySelectorAll('.selected-tag');
    tags.forEach(tag => {
        if (tag.textContent.replace('×', '').trim() === text) {
            tag.remove();
        }
    });
}

// Модальное окно
function openModal(field) {
    currentField = field;
    const modal = document.getElementById('editModal');
    const title = document.getElementById('modalTitle');
    const input = document.getElementById('modalInput');
    
    const fieldTitles = {
        name: 'Имя',
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
    } else {
        valueElement.textContent = value;
    }
    
    closeModal();
}

function saveProfile() {
    const profileData = {
        name: document.getElementById('nameValue')?.textContent || '',
        age: document.getElementById('ageValue')?.textContent || '',
        city: document.getElementById('cityValue')?.textContent || '',
        ...selectedItems
    };
    
    console.log('Сохраненные данные:', profileData);
    alert('Профиль успешно сохранен! 🎉');
}

// Инициализация при загрузке DOM
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initCapsules);
} else {
    initCapsules();
}