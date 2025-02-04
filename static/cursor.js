window.addEventListener("mousemove", (e) => {
    let cursor = document.getElementById("cursor");
    
    setTimeout(() => {
        cursor.style.top = `${e.clientY}px`;
        cursor.style.left = `${e.clientX}px`;
    }, 50);

    // List of interactive elements
    const interactiveElements = [
        'artist-card',          
        'content-button',
        'marquee',
        'image-container',
        'guitar-player',
        'image1',
        'image2',
        'cd',
        'menu-item',
        'link',
    ];

    // Check if cursor is over any interactive element
    if (interactiveElements.some(element => 
        e.target.id === element || 
        e.target.classList.contains(element) ||
        e.target.closest(`.${element}`))) {
        cursor.classList.add("active");
    } else {
        cursor.classList.remove("active");
    }
});