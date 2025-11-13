const DEFAULT_PHOTO = './images/example.jpg';

let initData = null;
let WebApp = null;
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
    city: '',
    gender: '',
    preferredGender: '',
    vkProfile: ''
};
let userPhoto = null;
let recommendedUsers = [];
let currentUserIndex = 0;
let startX = 0;
let currentX = 0;
let isDragging = false;

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
                reject(new Error('WebApp не загрузился после всех попыток'));
            }
        };
        
        check();
    });
}

function parseInitData(initData) {
    console.log('Parsing initData:', initData);
    
    if (!initData) {
        return null;
    }

    let userData = null;

    if (typeof initData === 'object') {
        console.log('InitData is object, using directly');
        userData = initData.user || initData;
    } else if (typeof initData === 'string') {
        const decodedString = decodeURIComponent(initData);
        console.log('Decoded initData:', decodedString);

        const params = new URLSearchParams(decodedString);
        const userParam = params.get('user');
        
        if (userParam) {
            try {
                userData = JSON.parse(userParam);
                console.log('Parsed user data from string:', userData);
            } catch (e) {
                console.error('Error parsing user data from string:', e);
            }
        }
    }

    if (userData) {
        console.log(
            userData.id, userData.username, userData.first_name, userData.last_name, userData.language_code, userData
        )
        return {
            id: userData.id || null,
            username: userData.username || '',
            firstName: userData.first_name || '',
            lastName: userData.last_name || '',
            languageCode: userData.language_code || 'ru'
        };
    }

    return null;
}

async function getCurrentUser() {
    try {
        await waitForWebApp();
        
        if (!initData) {
            console.error('No init data found');
            return { id: 0 };
        }

        console.log('Raw initData:', initData);

        const userData = parseInitData(initData);
        
        if (!userData || !userData.id) {
            console.error('No user data found in initData');
            return { id: 0 };
        }

        console.log('Extracted user data:', userData);
        return userData;
        
    } catch (error) {
        console.error('Error getting current user:', error);
        return { id: 0 };
    }
}

async function checkUserAuthorization() {
    const userData = await getCurrentUser();
    console.log('Проверка пользователя:', userData);
    
    try {
        const response = await fetch(`http://localhost:8080/profile?id=${userData.id}`);
        
        if (!response.ok) {
            throw new Error('Ошибка HTTP: ' + response.status);
        }
        
        const serverUserData = await response.json();
        console.log('Данные пользователя с сервера:', serverUserData);
        
        return { authorized: true, userData: serverUserData };
    } catch (error) {
        console.error('Ошибка при проверке авторизации:', error);
        return { authorized: false, userData: null };
    }
}

async function loadRecommendations() {
    try {
        const userData = await getCurrentUser();
        
        console.log('Загрузка рекомендаций для пользователя:', userData.id);
        
        const response = await fetch(`http://localhost:8080/recommendations?id=${userData.id}`);
        
        if (!response.ok) {
            throw new Error('Ошибка HTTP: ' + response.status);
        }
        
        const users = await response.json();
        if (!users) {
            return [];
        }
        console.log('Получены рекомендации:', users);
        
        return users;
    } catch (error) {
        console.error('Ошибка при загрузке рекомендаций:', error);
        return [];
    }
}

function loadOnboarding() {
    console.log('Загрузка анкеты...');
    
    Object.keys(selectedOnboardingItems).forEach(key => {
        selectedOnboardingItems[key] = [];
    });
    userBasicInfo = { age: '', city: '', gender: '', preferredGender: '', vkProfile: '' };
    userPhoto = null;
    
    const mainContent = document.getElementById('mainContent');
    const body = document.body;
    
    body.classList.add('onboarding-mode');
    
    mainContent.innerHTML = `
        <div class="onboarding-container">
            <canvas id="fireCanvas"></canvas>

            <div class="auth-container" id="authContainer">
                <div class="spark-container" id="sparkContainer">
                    <div class="pulse-ring"></div>
                    <div class="spark-elegant">
                        <div class="spark-dot"></div>
                        <div class="spark-orbit">
                            <div class="orbit-particle"></div>
                        </div>
                        <div class="spark-orbit">
                            <div class="orbit-particle"></div>
                        </div>
                        <div class="spark-orbit">
                            <div class="orbit-particle"></div>
                        </div>
                        <div class="spark-orbit">
                            <div class="orbit-particle"></div>
                        </div>
                    </div>
                    <div class="click-hint">Коснись меня</div>
                </div>
                
                <h1 class="auth-title">ИСКРА</h1>
                <p class="auth-subtitle">Прикоснись к энергии новых встреч</p>
                
                <div class="loading-container" id="loadingContainer">
                    <div class="loading-text">Зажигание</div>
                    <div class="loading-bar">
                        <div class="loading-progress" id="loadingProgress"></div>
                    </div>
                </div>
            </div>

            <div class="onboarding-progress">
                <div class="onboarding-progress-fill" id="onboardingProgressFill"></div>
            </div>

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

                        <div class="onboarding-input-compact select-input">
                            <div class="onboarding-input-label">Ваш пол</div>
                            <select class="onboarding-input-field" id="genderInput" onchange="updateBasicInfo('gender', this.value)">
                                <option value="0">Не выбран</option>
                                <option value="0">Мужской</option>
                                <option value="1">Женский</option>
                            </select>
                            <span class="onboarding-input-arrow">▼</span>
                        </div>

                        <div class="onboarding-input-compact select-input">
                            <div class="onboarding-input-label">Людей какого пола вы хотите найти</div>
                            <select class="onboarding-input-field" id="preferredGenderInput" onchange="updateBasicInfo('preferredGender', this.value)">
                                <option value="2">Не выбран</option>
                                <option value="2">Не важно</option>
                                <option value="1">Женский</option>
                                <option value="0">Мужской</option>
                            </select>
                            <span class="onboarding-input-arrow">▼</span>
                        </div>

                        <div class="onboarding-input-compact full-width">
                            <div class="onboarding-input-label">Ссылка на профиль ВК (необходима для дальнейшего знакомства с людьми)</div>
                            <input type="text" class="onboarding-input-field" id="vkProfileInput" 
                                placeholder="https://vk.com/username"
                                oninput="updateBasicInfo('vkProfile', this.value)">
                            <span class="onboarding-input-edit">✎</span>
                        </div>
                    </div>

                    <div class="photo-upload-section">
                        <div class="photo-upload-container">
                            <div class="photo-preview" id="photoPreview">
                                <div class="photo-placeholder">
                                    <div class="camera-icon">📷</div>
                                    <div class="photo-text">Добавить фото</div>
                                </div>
                            </div>
                            <input type="file" id="photoInput" accept="image/*" style="display: none;">
                            <button class="photo-upload-btn" onclick="document.getElementById('photoInput').click()">
                                Выбрать фото
                            </button>
                            <div class="photo-optional">Необязательно</div>
                        </div>
                    </div>
                    
                    <div class="selection-required" id="screen2Message">Заполните все обязательные поля</div>
                    
                    <button class="onboarding-btn" id="screen2Button" onclick="nextOnboardingScreen(3)">
                        Продолжить
                    </button>
                </div>
            </div>

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
    initSparkAnimation();
    initPhotoUpload();
}

function initPhotoUpload() {
    const photoInput = document.getElementById('photoInput');
    const photoPreview = document.getElementById('photoPreview');
    
    if (!photoInput || !photoPreview) return;
    
    photoInput.addEventListener('change', function(e) {
        const file = e.target.files[0];
        if (file) {
            if (!file.type.startsWith('image/')) {
                alert('Выберите файл изображения');
                return;
            }
            
            if (file.size > 5 * 1024 * 1024) {
                alert('Фото должно быть меньше 5MB');
                return;
            }
            
            const reader = new FileReader();
            reader.onload = function(e) {
                const img = new Image();
                img.onload = function() {
                    // Простая проверка - вертикальное фото
                    if (img.height <= img.width) {
                        alert('Выберите вертикальное фото');
                        return;
                    }
                    
                    // Всё ок - сохраняем
                    userPhoto = e.target.result;
                    
                    // Просто показываем фото
                    photoPreview.innerHTML = `
                        <img src="${userPhoto}" alt="Фото" style="width: 100%; height: 100%; object-fit: cover; border-radius: 18px;">
                    `;
                    photoPreview.classList.add('has-photo');
                };
                img.src = e.target.result;
            };
            reader.readAsDataURL(file);
        }
    });
    
    // Клик по превью открывает выбор файла
    photoPreview.addEventListener('click', function() {
        photoInput.click();
    });
}

function initSparkAnimation() {
    const sparkContainer = document.getElementById('sparkContainer');
    const authContainer = document.getElementById('authContainer');
    const fireCanvas = document.getElementById('fireCanvas');
    const loadingContainer = document.getElementById('loadingContainer');
    const loadingProgress = document.getElementById('loadingProgress');

    if (!sparkContainer) {
        console.error('sparkContainer не найден');
        return;
    }

    let ctx = fireCanvas.getContext('2d');
    let particles = [];
    let isAnimating = false;
    let animationId;

    function getSparkPosition() {
        const rect = sparkContainer.getBoundingClientRect();
        return {
            x: rect.left + rect.width / 2,
            y: rect.top + rect.height / 2
        };
    }

    function createParticleTexture(size, colorStops) {
        const canvas = document.createElement('canvas');
        canvas.width = size;
        canvas.height = size;
        const ctx = canvas.getContext('2d');
        
        const gradient = ctx.createRadialGradient(
            size/2, size/2, 0,
            size/2, size/2, size/2
        );
        
        colorStops.forEach(stop => {
            gradient.addColorStop(stop.offset, stop.color);
        });
        
        ctx.fillStyle = gradient;
        ctx.fillRect(0, 0, size, size);
        
        return canvas;
    }

    const textures = {
        core: createParticleTexture(64, [
            { offset: 0, color: 'rgba(255, 255, 255, 1)' },
            { offset: 0.2, color: 'rgba(255, 255, 200, 0.8)' },
            { offset: 0.4, color: 'rgba(255, 200, 100, 0.6)' },
            { offset: 1, color: 'rgba(255, 100, 0, 0)' }
        ]),
        glow: createParticleTexture(128, [
            { offset: 0, color: 'rgba(255, 200, 100, 0.4)' },
            { offset: 0.3, color: 'rgba(255, 150, 50, 0.2)' },
            { offset: 1, color: 'rgba(255, 100, 0, 0)' }
        ])
    };

    class ElegantParticle {
        constructor(x, y, type, angle, speed) {
            this.x = x;
            this.y = y;
            this.type = type;
            this.angle = angle;
            this.speed = speed;
            this.vx = Math.cos(angle) * speed;
            this.vy = Math.sin(angle) * speed;
            this.life = 1;
            this.decay = Math.random() * 0.008 + 0.005;
            this.size = type === 'core' ? 
                Math.random() * 12 + 8 : 
                Math.random() * 35 + 25;
            this.rotation = Math.random() * Math.PI * 2;
            this.rotationSpeed = (Math.random() - 0.5) * 0.02;
        }

        update() {
            this.x += this.vx;
            this.y += this.vy;
            this.life -= this.decay;
            this.rotation += this.rotationSpeed;
            
            this.size *= 0.995;
            
            return this.life > 0;
        }

        draw() {
            const texture = textures[this.type];
            const alpha = this.life;
            const size = this.size;
            
            ctx.save();
            ctx.globalAlpha = alpha;
            ctx.translate(this.x, this.y);
            ctx.rotate(this.rotation);
            
            ctx.drawImage(
                texture, 
                -size/2, -size/2, 
                size, size
            );
            
            ctx.restore();
        }
    }

    function initCanvas() {
        fireCanvas.width = window.innerWidth;
        fireCanvas.height = window.innerHeight;
        ctx = fireCanvas.getContext('2d');
    }

    function createRadialExplosion(x, y) {
        for (let i = 0; i < 24; i++) {
            const angle = (i / 24) * Math.PI * 2;
            const speed = Math.random() * 2 + 1.5;
            particles.push(new ElegantParticle(x, y, 'core', angle, speed));
        }
        
        for (let i = 0; i < 12; i++) {
            const angle = (i / 12) * Math.PI * 2;
            const speed = Math.random() * 1 + 0.8;
            particles.push(new ElegantParticle(x, y, 'glow', angle, speed));
        }
    }

    function animateElegantFire() {
        ctx.fillStyle = 'rgba(10, 10, 10, 0.08)';
        ctx.fillRect(0, 0, fireCanvas.width, fireCanvas.height);

        for (let i = particles.length - 1; i >= 0; i--) {
            if (!particles[i].update()) {
                particles.splice(i, 1);
            } else {
                particles[i].draw();
            }
        }

        if (isAnimating && particles.length < 60) {
            const sparkPos = getSparkPosition();
            
            if (Math.random() < 0.4) {
                const angle = Math.random() * Math.PI * 2;
                const speed = Math.random() * 1.2 + 0.5;
                particles.push(new ElegantParticle(sparkPos.x, sparkPos.y, 'core', angle, speed));
            }
            if (Math.random() < 0.15) {
                const angle = Math.random() * Math.PI * 2;
                const speed = Math.random() * 0.8 + 0.3;
                particles.push(new ElegantParticle(sparkPos.x, sparkPos.y, 'glow', angle, speed));
            }
        }

        animationId = requestAnimationFrame(animateElegantFire);
    }

    sparkContainer.addEventListener('click', function() {
        if (isAnimating) return;
        isAnimating = true;

        initCanvas();
        fireCanvas.classList.add('active');
        loadingContainer.style.display = 'block';

        const sparkPos = getSparkPosition();

        sparkContainer.style.opacity = '0';
        sparkContainer.style.transition = 'opacity 0.5s ease';

        setTimeout(() => {
            createRadialExplosion(sparkPos.x, sparkPos.y);
        }, 200);

        animateElegantFire();

        let progress = 0;
        const loadingInterval = setInterval(() => {
            progress += Math.random() * 8 + 2;
            if (progress >= 100) {
                progress = 100;
                clearInterval(loadingInterval);
                
                setTimeout(() => {
                    authContainer.style.display = 'none';
                    fireCanvas.classList.remove('active');
                    isAnimating = false;

                    setTimeout(() => nextOnboardingScreen(2), 1000);
                }, 600);
            }
            loadingProgress.style.width = progress + '%';
        }, 120);
    });

    window.addEventListener('resize', initCanvas);
}

function splitStringByCommas(str) {
    if (!str) return [];
    return str.split(',').map(item => item.trim()).filter(item => item !== '');
}

async function loadMainContent() {
    const mainContent = document.getElementById('mainContent');
    const body = document.body;
    
    body.classList.remove('onboarding-mode');
    currentUserIndex = 0;

    mainContent.innerHTML = `
        <div class="main-app">
            <div class="cards-container">
                <div class="loading-message">
                    <div class="loading-spinner"></div>
                    <p>Ищем подходящие анкеты...</p>
                </div>
            </div>
        </div>
    `;

    recommendedUsers = await loadRecommendations();

    mainContent.innerHTML = `
        <div class="main-app">
            <div class="cards-container">
                <div class="no-users-message" id="noUsersMessage" style="display: none;">
                    <div class="message-icon">💫</div>
                    <h3>Анкеты закончились</h3>
                    <p>Возвращайтесь позже, чтобы увидеть новые рекомендации</p>
                </div>
                
                <div class="user-card" id="userCard">
                    <div class="card-background"></div>
                    <div class="swipe-overlay swipe-like">👍</div>
                    <div class="swipe-overlay swipe-dislike">👎</div>
                    <div class="card-content">
                        <div class="card-main-info" id="cardMainInfo">
                            <h2 class="user-name" id="userName">Имя</h2>
                            <div class="user-age-city" id="userAgeCity">Возраст • Город</div>
                            <div class="user-events-tags" id="userEventsTags"></div>
                        </div>
                        
                        <button class="show-more-btn" onclick="toggleUserDetails()">
                            Показать больше
                            <span class="arrow">▼</span>
                        </button>
                        
                        <div class="user-details" id="userDetails">
                            <div class="details-section">
                                <h4>О себе</h4>
                                <div class="detail-item">
                                    <span class="detail-label">Карьера:</span>
                                    <span class="detail-value" id="detailCareer">-</span>
                                </div>
                                <div class="detail-item">
                                    <span class="detail-label">Характер:</span>
                                    <span class="detail-value" id="detailPersonality">-</span>
                                </div>
                                <div class="detail-item">
                                    <span class="detail-label">Цели отношений:</span>
                                    <span class="detail-value" id="detailRelationship">-</span>
                                </div>
                                <div class="detail-item">
                                    <span class="detail-label">Ценности:</span>
                                    <span class="detail-value" id="detailValues">-</span>
                                </div>
                            </div>
                            
                            <div class="details-section">
                                <h4>Интересы</h4>
                                <div class="detail-item">
                                    <span class="detail-label">Музыка:</span>
                                    <span class="detail-value tags-container" id="detailMusic"></span>
                                </div>
                                <div class="detail-item">
                                    <span class="detail-label">Фильмы:</span>
                                    <span class="detail-value tags-container" id="detailMovies"></span>
                                </div>
                                <div class="detail-item">
                                    <span class="detail-label">Хобби:</span>
                                    <span class="detail-value tags-container" id="detailHobbies"></span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    `;
    
    loadNextUser();
    initSwipeHandlers();
}

function loadNextUser() {
    if (currentUserIndex >= recommendedUsers.length) {
        document.getElementById('noUsersMessage').style.display = 'flex';
        document.getElementById('userCard').style.display = 'none';
        return;
    }
    
    const user = recommendedUsers[currentUserIndex];
    const userCard = document.getElementById('userCard');
    
    userCard.style.opacity = '0';
    userCard.style.transform = 'translateY(20px)';
    
    setTimeout(() => {
        // Обновляем фон карточки с фото пользователя или заглушкой
        const cardBackground = userCard.querySelector('.card-background');
        if (user.photo) {
            cardBackground.style.backgroundImage = `url(${user.photo})`;
        } else {
            // Используем заглушку по умолчанию
            cardBackground.style.backgroundImage = `url(${DEFAULT_PHOTO})`;
        }
        
        document.getElementById('userName').textContent = user.name || 'Не указано';
        document.getElementById('userAgeCity').textContent = `${user.age || '?'} • ${user.city || 'Не указан'}`;
        
        const eventsTagsContainer = document.getElementById('userEventsTags');
        eventsTagsContainer.innerHTML = '';
        const events = splitStringByCommas(user.event_preferences);
        if (events.length > 0) {
            events.forEach(event => {
                const tag = document.createElement('span');
                tag.className = 'event-tag';
                tag.textContent = event;
                eventsTagsContainer.appendChild(tag);
            });
        } else {
            eventsTagsContainer.innerHTML = '<span class="no-data">Не указаны</span>';
        }
        
        document.getElementById('detailCareer').textContent = user.career_type || 'Не указана';
        document.getElementById('detailPersonality').textContent = user.personality_type || 'Не указан';
        document.getElementById('detailRelationship').textContent = user.relationship_goal || 'Не указаны';
        document.getElementById('detailValues').textContent = user.important_values || 'Не указаны';
        
        updateTagsContainer('detailMusic', user.music);
        updateTagsContainer('detailMovies', user.films);
        updateTagsContainer('detailHobbies', user.hobbies);
        
        document.getElementById('userDetails').classList.remove('active');
        resetSwipeOverlay();
        
        userCard.style.opacity = '1';
        userCard.style.transform = 'translateY(0)';
    }, 200);
}

function updateTagsContainer(containerId, data) {
    const container = document.getElementById(containerId);
    container.innerHTML = '';
    
    const tags = splitStringByCommas(data);
    if (tags.length > 0) {
        tags.forEach(tag => {
            const tagElement = document.createElement('span');
            tagElement.className = 'interest-tag';
            tagElement.textContent = tag;
            container.appendChild(tagElement);
        });
    } else {
        container.innerHTML = '<span class="no-data">Не указаны</span>';
    }
}

function toggleUserDetails() {
    const details = document.getElementById('userDetails');
    const arrow = document.querySelector('.arrow');
    
    details.classList.toggle('active');
    arrow.style.transform = details.classList.contains('active') ? 'rotate(180deg)' : 'rotate(0)';
}

async function sendInteraction(targetUserId, isLike) {
    try {
        const currentUser = await getCurrentUser();
        
        const interactionType = isLike ? 'like' : 'dislike';
        
        console.log(`Отправка взаимодействия: ${interactionType} для пользователя ${targetUserId}`);
        
        const response = await fetch('http://localhost:8080/interaction', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                user_id: currentUser.id,
                target_user_id: targetUserId,
                interaction_type: interactionType
            })
        });
        
        if (response.ok) {
            console.log('Взаимодействие успешно отправлено');
        } else {
            console.error('Ошибка при отправке взаимодействия:', response.status);
        }
    } catch (error) {
        console.error('Ошибка при отправке взаимодействия:', error);
    }
}

function initSwipeHandlers() {
    const card = document.getElementById('userCard');
    if (!card) return;
    
    card.addEventListener('touchstart', handleTouchStart, { passive: false });
    card.addEventListener('touchmove', handleTouchMove, { passive: false });
    card.addEventListener('touchend', handleTouchEnd);
    
    card.addEventListener('mousedown', handleMouseDown);
    document.addEventListener('mousemove', handleMouseMove);
    document.addEventListener('mouseup', handleMouseUp);
}

function handleTouchStart(e) {
    if (e.touches.length > 1) return;
    
    const touch = e.touches[0];
    startX = touch.clientX;
    currentX = startX;
    isDragging = true;
    
    const card = document.getElementById('userCard');
    card.style.transition = 'none';
    resetSwipeOverlay();
}

function handleTouchMove(e) {
    if (!isDragging || e.touches.length > 1) return;
    
    e.preventDefault();
    const touch = e.touches[0];
    currentX = touch.clientX;
    updateCardPosition();
    updateSwipeOverlay();
}

function handleTouchEnd() {
    if (!isDragging) return;
    
    isDragging = false;
    handleSwipeEnd();
}

function handleMouseDown(e) {
    startX = e.clientX;
    currentX = startX;
    isDragging = true;
    
    const card = document.getElementById('userCard');
    card.style.transition = 'none';
    resetSwipeOverlay();
}

function handleMouseMove(e) {
    if (!isDragging) return;
    
    currentX = e.clientX;
    updateCardPosition();
    updateSwipeOverlay();
}

function handleMouseUp() {
    if (!isDragging) return;
    
    isDragging = false;
    handleSwipeEnd();
}

function updateCardPosition() {
    const card = document.getElementById('userCard');
    const deltaX = currentX - startX;
    const rotation = deltaX * 0.1;
    
    card.style.transform = `translateX(${deltaX}px) rotate(${rotation}deg)`;
}

function updateSwipeOverlay() {
    const deltaX = currentX - startX;
    const swipeThreshold = 50;
    
    const likeOverlay = document.querySelector('.swipe-like');
    const dislikeOverlay = document.querySelector('.swipe-dislike');
    
    if (!likeOverlay || !dislikeOverlay) return;
    
    likeOverlay.style.opacity = '0';
    dislikeOverlay.style.opacity = '0';
    
    if (deltaX > swipeThreshold) {
        likeOverlay.style.opacity = Math.min((deltaX - swipeThreshold) / 100, 0.3).toString();
    } else if (deltaX < -swipeThreshold) {
        dislikeOverlay.style.opacity = Math.min(Math.abs(deltaX + swipeThreshold) / 100, 0.3).toString();
    }
}

function resetSwipeOverlay() {
    const likeOverlay = document.querySelector('.swipe-like');
    const dislikeOverlay = document.querySelector('.swipe-dislike');
    
    if (!likeOverlay || !dislikeOverlay) return;
    
    likeOverlay.style.opacity = '0';
    dislikeOverlay.style.opacity = '0';
}

function handleSwipeEnd() {
    const card = document.getElementById('userCard');
    const deltaX = currentX - startX;
    const swipeThreshold = 100;
    
    card.style.transition = 'all 0.5s ease';
    
    if (Math.abs(deltaX) > swipeThreshold) {
        const direction = deltaX > 0 ? 1 : -1;
        const isLike = deltaX > 0;
        
        card.style.transform = `translateX(${direction * 500}px) rotate(${direction * 30}deg)`;
        card.style.opacity = '0';
        
        const currentUser = recommendedUsers[currentUserIndex];
        if (currentUser) {
            sendInteraction(currentUser.id, isLike);
        }
        
        setTimeout(() => {
            currentUserIndex++;
            loadNextUser();
            resetCardPosition();
        }, 300);
        
        console.log(isLike ? 'Лайк' : 'Дизлайк', recommendedUsers[currentUserIndex]?.name);
        
    } else {
        resetCardPosition();
    }
    
    resetSwipeOverlay();
}

function resetCardPosition() {
    const card = document.getElementById('userCard');
    card.style.transform = 'translateX(0) rotate(0)';
    card.style.opacity = '1';
}

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

function updateBasicInfo(field, value) {
    console.log(`Обновление ${field}:`, value);
    userBasicInfo[field] = value;
    
    checkScreen2Complete();
}

function checkScreen2Complete() {
    const isComplete = userBasicInfo.age && 
                        userBasicInfo.city && 
                        userBasicInfo.gender && 
                        userBasicInfo.preferredGender &&
                        userBasicInfo.vkProfile;
    
    const button = document.getElementById('screen2Button');
    const message = document.getElementById('screen2Message');
    
    console.log('Проверка экрана 2:', userBasicInfo);
    
    if (button) {
        if (isComplete) {
            button.classList.add('active');
            if (message) message.style.display = 'none';
        } else {
            button.classList.remove('active');
            if (message) {
                const missingFields = [];
                if (!userBasicInfo.age) missingFields.push('возраст');
                if (!userBasicInfo.city) missingFields.push('город');
                if (!userBasicInfo.gender) missingFields.push('пол');
                if (!userBasicInfo.preferredGender) missingFields.push('предпочтительный пол');
                if (!userBasicInfo.vkProfile) missingFields.push('профиль ВК');
                
                message.textContent = `Заполните: ${missingFields.join(', ')}`;
                message.style.display = 'block';
            }
        }
    }
    return isComplete;
}

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

function getCategoryByScreen(screenNumber) {
    const categories = ['career', 'personality', 'relationship', 'values', 'music', 'movies', 'hobbies', 'events'];
    return categories[screenNumber - 3] || 'career';
}

function getScreenByCategory(category) {
    const categories = ['career', 'personality', 'relationship', 'values', 'music', 'movies', 'hobbies', 'events'];
    return categories.indexOf(category) + 3;
}

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
    
    if (currentOnboardingScreen >= 7 && currentOnboardingScreen <= 10) {
        const currentCategory = getCategoryByScreen(currentOnboardingScreen);
        if (selectedOnboardingItems[currentCategory].length === 0) {
            console.log('Нельзя перейти - не выбрано ни одного варианта');
            const message = document.getElementById(`screen${currentOnboardingScreen}Message`);
            if (message) {
                message.style.display = 'block';
            }
            return;
        }
    }
    
    const currentMessage = document.getElementById(`screen${currentOnboardingScreen}Message`);
    if (currentMessage) {
        currentMessage.style.display = 'none';
    }
    
    const currentScreen = document.getElementById(`screen${currentOnboardingScreen}`);
    const nextScreen = document.getElementById(`screen${screenNumber}`);
    
    if (currentScreen) currentScreen.classList.remove('active');
    if (nextScreen) nextScreen.classList.add('active');
    
    currentOnboardingScreen = screenNumber;
    updateOnboardingProgress();
}

function updateOnboardingProgress() {
    const progressFill = document.getElementById('onboardingProgressFill');
    if (!progressFill) return;
    
    const progress = (currentOnboardingScreen - 1) / 9 * 100;
    progressFill.style.width = progress + '%';
    console.log(`Прогресс: ${progress}%`);
}

async function completeOnboarding() {
    console.log('Завершение онбординга...');
    
    if (selectedOnboardingItems.events.length === 0) {
        const message = document.getElementById('screen10Message');
        if (message) {
            message.textContent = 'Выберите до 3 мероприятий для продолжения';
            message.style.display = 'block';
        }
        console.log('Нельзя завершить - не выбраны мероприятия');
        return;
    }
    
    if (!userBasicInfo.age || !userBasicInfo.city || !userBasicInfo.gender || 
        !userBasicInfo.preferredGender || !userBasicInfo.vkProfile) {
        console.log('Не заполнены обязательные поля:', { ...userBasicInfo, hasPhoto: !!userPhoto });
        alert('Пожалуйста, заполните все обязательные поля');
        return;
    }
    
    console.log('Собранные данные:', {
        basic: userBasicInfo,
        selections: selectedOnboardingItems,
        hasPhoto: !!userPhoto
    });
    
    try {
        const userData = await getCurrentUser();
        
        const name = userData.firstName || '';
        const surname = userData.lastName || '';
        const fullName = [name, surname].filter(Boolean).join(' ') || userData.username || 'Пользователь';
        
        const profileData = {
            id: userData.id,
            username: userBasicInfo.vkProfile || '',
            name: fullName,
            surname: surname,
            age: parseInt(userBasicInfo.age) || 0,
            city: userBasicInfo.city || '',
            gender: parseInt(userBasicInfo.gender) || 0,
            preferred_gender: parseInt(userBasicInfo.preferredGender) || 0,
            career_type: selectedOnboardingItems.career[0] || '',
            personality_type: selectedOnboardingItems.personality[0] || '',
            relationship_goal: selectedOnboardingItems.relationship[0] || '',
            important_values: selectedOnboardingItems.values[0] || '',
            music: selectedOnboardingItems.music.join(', ') || '',
            films: selectedOnboardingItems.movies.join(', ') || '',
            hobbies: selectedOnboardingItems.hobbies.join(', ') || '',
            event_preferences: selectedOnboardingItems.events.join(', ') || '',
            photo: userPhoto || DEFAULT_PHOTO
        };

        console.log('Отправка данных на сервер:', { 
            ...profileData, 
            hasPhoto: !!userPhoto 
        });
        console.log(profileData.photo)
        
        const button = document.getElementById('screen10Button');
        const originalText = button.textContent;
        button.textContent = 'Сохранение...';
        button.disabled = true;
        
        const response = await fetch('http://localhost:8080/createuser', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(profileData)
        });

        if (response.ok) {
            console.log('Профиль успешно создан');
            
            setTimeout(() => {
                location.reload();
            }, 1000);
            
        } else {
            const errorText = await response.text();
            console.error('Ошибка при создании профиля:', response.status, errorText);
            button.textContent = originalText;
            button.disabled = false;
        }
    } catch (error) {
        console.error('Ошибка при сохранении профиля:', error);
        const button = document.getElementById('screen10Button');
        button.textContent = 'Завершить профиль';
        button.disabled = false;
    }
}

function editProfile() {
    console.log('Редактирование профиля...');
    loadOnboarding();
}

async function initApp() {
    console.log('Инициализация приложения...');
    
    try {
        const authStatus = await checkUserAuthorization();
        console.log('Статус авторизации:', authStatus);
        
        if (authStatus.authorized) {
            await waitForWebApp();
            await loadMainContent();
        } else {
            await waitForWebApp();
            loadOnboarding();
        }
    } catch (error) {
        const authStatus = await checkUserAuthorization();
        if (authStatus.authorized) {
            await loadMainContent();
        } else {
            loadOnboarding();
        }
    }
}

document.addEventListener('DOMContentLoaded', initApp);