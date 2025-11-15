WITH range(1, 1000) AS ids
UNWIND ids AS id

WITH id,
     ['IT и технологии', 'Дизайн и UX', 'Медицина', 'Образование', 'Бизнес', 'Финансы', 'Маркетинг', 'Искусство', 'Музыка', 'Кино', 'Фотография', 'Архитектура', 'Инженерия', 'Недвижимость', 'Юриспруденция', 'Психология'] AS careers,
     ['Экстраверт', 'Интроверт', 'Амбиверт', 'Аналитик', 'Творец', 'Прагматик', 'Романтик', 'Реалист', 'Оптимист', 'Философ', 'Новатор', 'Лидер', 'Целеустремленный', 'Гибкий', 'Настойчивый', 'Командный'] AS personalities,
     ['Серьезные отношения', 'Дружба', 'Несерьезные отношения', 'Создание семьи', 'Поиск партнера', 'Романтика', 'Деловое партнерство', 'Творчество', 'Путешествия', 'Совместные проекты', 'Духовность', 'Карьера'] AS relationships,
     ['Любовь и забота', 'Семья', 'Карьера', 'Финансы', 'Духовность', 'Здоровье', 'Образование', 'Творчество', 'Свобода', 'Приключения', 'Безопасность', 'Экология'] AS values_list

CREATE (p:Person {
    id: id,
    name: 'User_' + id + '_t',
    username: '@user' + id + '_t',
    age: 18 + (id % 50),
    gender: id % 2,
    preferred_gender: id % 3,
    career_type: careers[(id * 7) % size(careers)],
    personality_type: personalities[(id * 11) % size(personalities)],
    relationship_goal: relationships[(id * 13) % size(relationships)],
    important_values: values_list[(id * 17) % size(values_list)],
    currentLimit: 30,
    lastLimitUpdate: toString(datetime()),
    maxLimit: 30,
    refreshAmount: 5,
    refreshRate: "PT1H",
    music: '',
    films: '',
    hobbies: '',
    event_preferences: '',
    photo: ''
});