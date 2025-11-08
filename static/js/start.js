document.addEventListener('DOMContentLoaded', function() {
    const sparkContainer = document.getElementById('sparkContainer');
    const authContainer = document.getElementById('authContainer');
    const mainContainer = document.getElementById('mainContainer');
    const fireCanvas = document.getElementById('fireCanvas');
    const loadingContainer = document.getElementById('loadingContainer');
    const loadingProgress = document.getElementById('loadingProgress');

    let ctx = fireCanvas.getContext('2d');
    let particles = [];
    let isAnimating = false;
    let animationId;

    // Получаем позицию искры относительно страницы
    function getSparkPosition() {
        const rect = sparkContainer.getBoundingClientRect();
        return {
            x: rect.left + rect.width / 2,
            y: rect.top + rect.height / 2
        };
    }

    // Пререндеренные текстуры для частиц
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

    // Создаем текстуры заранее
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
            
            // Плавное уменьшение
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
        // Создаем взрыв во все стороны из позиции искры
        for (let i = 0; i < 24; i++) {
            const angle = (i / 24) * Math.PI * 2;
            const speed = Math.random() * 2 + 1.5;
            particles.push(new ElegantParticle(x, y, 'core', angle, speed));
        }
        
        // Добавляем большие glow-частицы
        for (let i = 0; i < 12; i++) {
            const angle = (i / 12) * Math.PI * 2;
            const speed = Math.random() * 1 + 0.8;
            particles.push(new ElegantParticle(x, y, 'glow', angle, speed));
        }
    }

    function animateElegantFire() {
        // Плавное затемнение вместо резкого очищения
        ctx.fillStyle = 'rgba(10, 10, 10, 0.08)';
        ctx.fillRect(0, 0, fireCanvas.width, fireCanvas.height);

        // Обновляем и рисуем частицы
        for (let i = particles.length - 1; i >= 0; i--) {
            if (!particles[i].update()) {
                particles.splice(i, 1);
            } else {
                particles[i].draw();
            }
        }

        // Добавляем новые частицы из позиции искры
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

    function redirectToHomepage() {
        // Замените '/' на нужный URL главной страницы
        window.location.href = 'base.html';
        // Или используйте, если нужно:
        // window.location.replace('/');
    }

    sparkContainer.addEventListener('click', function() {
        if (isAnimating) return;
        isAnimating = true;

        initCanvas();
        fireCanvas.classList.add('active');
        loadingContainer.style.display = 'block';

        // Получаем позицию искры для взрыва
        const sparkPos = getSparkPosition();

        // Плавное исчезновение искры
        sparkContainer.style.opacity = '0';
        sparkContainer.style.transition = 'opacity 0.5s ease';

        // Создаем круговой взрыв из позиции искры
        setTimeout(() => {
            createRadialExplosion(sparkPos.x, sparkPos.y);
        }, 200);

        animateElegantFire();

        // Стильная загрузка
        let progress = 0;
        const loadingInterval = setInterval(() => {
            progress += Math.random() * 8 + 2;
            if (progress >= 100) {
                progress = 100;
                clearInterval(loadingInterval);
                
                setTimeout(() => {
                    authContainer.style.display = 'none';
                    mainContainer.style.display = 'block';
                    fireCanvas.classList.remove('active');
                    isAnimating = false;
                    
                    // Перенаправление на главную страницу после небольшой задержки
                    setTimeout(redirectToHomepage, 1000);
                }, 600);
            }
            loadingProgress.style.width = progress + '%';
        }, 120);
    });

    window.addEventListener('resize', initCanvas);
});