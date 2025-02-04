window.addEventListener("mousemove", (e) => {
    let cursor = document.getElementById("cursor");
    
    setTimeout(() => {
        cursor.style.top = `${e.clientY}px`;
        cursor.style.left = `${e.clientX}px`;
    }, 50);

    // Get computed style of the element
    const computedStyle = window.getComputedStyle(e.target);
    
    // Check if cursor is pointer
    if (computedStyle.cursor === 'pointer') {
        cursor.classList.add("active");
    } else {
        cursor.classList.remove("active");
    }
});