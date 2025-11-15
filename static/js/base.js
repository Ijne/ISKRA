function getCurrentUser() {
    return 1;
}

function setupNavigation() {
    const profileButton = document.querySelector('.nav-button:nth-child(1)');
    const mainButton = document.querySelector('.main-button');
    const eventsButton = document.querySelector('.nav-button:nth-child(3)');
    
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
    
    if (eventsButton) {
        eventsButton.addEventListener('click', function() {
            window.location.href = 'events.html';
        });
    }
}

document.addEventListener('DOMContentLoaded', function() {
    setupNavigation();
});