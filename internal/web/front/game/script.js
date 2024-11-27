function updatePlaceholder() {
    const input = document.getElementById('full-word-guess');
    const button = document.querySelector('.guess-word button');
    if (window.matchMedia('(hover: none) and (pointer: coarse)').matches) {
        input.placeholder = 'Full word';
        button.textContent = 'Guess';
    }
}

window.addEventListener('load', updatePlaceholder);
window.addEventListener('resize', updatePlaceholder);