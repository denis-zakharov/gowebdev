function app() {
    fetch('http://localhost:9090/products')
        .then(r => r.json())
        .then(ps => ps.forEach(p =>
            document.getElementById('products').innerHTML += (`<li>${JSON.stringify(p)}</li>`)))
}

document.addEventListener("DOMContentLoaded", app);