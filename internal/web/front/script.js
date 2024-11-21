document.addEventListener('DOMContentLoaded', () => {
    const titleSpans = document.querySelectorAll('.game-title span');
    titleSpans.forEach(span => {
        const initialRotation = Math.floor(Math.random() * 20 - 10) + 'deg';
        const rotationDirection = Math.random() < 0.5 ? -1 : 1;
        const randomDelay = Math.random() * 1.5 + 's';
        span.style.setProperty('--initial-rotate', initialRotation);
        span.style.setProperty('--rotation-direction', rotationDirection);
        span.style.setProperty('--random-delay', randomDelay);
    });
});