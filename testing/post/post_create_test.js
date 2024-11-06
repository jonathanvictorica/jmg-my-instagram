import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomSeed, randomItem, randomIntBetween } from 'k6';

// Inicializamos las variables de prueba
const locations = ["Praga", "Berlin", "Dinamarca"];
const usuarios = ["jonathan", "martin", "guillermo", "roberto"];
const multimedia = [1, 2];

// Función para generar 10 hashtags aleatorios
function generar10Hashtag() {
    const hashtags = [];
    for (let i = 0; i < 10; i++) {
        hashtags.push(`#hashtag${i + 1}`);
    }
    return hashtags;
}

// Función para obtener el ID de usuario
function getIdUser(nombreUsuario) {
    const res = http.get(`http://localhost:8080/api/v1/users/search?name=${nombreUsuario}&top=10`);
    const users = res.json();
    var value= users.length > 0 ? users[0].user_id : null;
    //console.log(value)
    return value;
}

// Función para obtener el ID location
function getIdLocation(nombreLocation) {
    const res = http.get(`http://localhost:8080/api/v1/locations/search?name=${nombreLocation}&top=10`);
    const locations = res.json();
    var value= locations.length > 0 ? locations[0].location_id : null;
    //console.log(value)
    return value;
}


export const options = {
    stages: [
        { duration: '10s', target: 30 },   // En los primeros 20 segundos, aumentamos a 90 VUs
        { duration: '10s', target: 60 },  // Durante los siguientes 20 segundos, aumentamos a 120 VUs
        { duration: '10s', target: 90 },   // Finalmente, durante 20 segundos, reducimos a 50 VUs
    ]
};

// La prueba principal de carga
export default function () {
    // Obtenemos IDs de usuario
    const userIdJonathan = getIdUser("jonathan");
    const userIdMartin = getIdUser("martin");
    const userIdGuillermo = getIdUser("guillermo");
    const userIdRoberto = getIdUser("roberto");

    // Cuerpo de la solicitud POST
    const postData = JSON.stringify({
        title: "My First Post",
        description: "This is a description of my first post.",
        user_id: userIdJonathan,
        sections: [
            {
                multimedia_content_id: multimedia[0],
                description: "This is the first section description.",
                hashtags: generar10Hashtag(),
                users: [userIdJonathan, userIdMartin],
                location: getIdLocation("Praga")
            },
            {
                multimedia_content_id: multimedia[1],
                description: "This is the second section description.",
                hashtags: generar10Hashtag(),
                users: [userIdGuillermo, userIdRoberto],
                location: getIdLocation("Berlin")
            }
        ]
    });

    // Realiza la solicitud POST
    const res = http.post("http://localhost:8080/api/v1/posts", postData, {
        headers: { 'Content-Type': 'application/json' },
    });

    // Verifica que la respuesta tenga un código 201 y que el contenido esté presente
    check(res, {
        'status es 201': (r) => r.status === 201,
        'response contiene user_id': (r) => r.json('post_id') !== undefined,
    });

    // Simula un tiempo de espera entre iteraciones de usuario
    //sleep(randomIntBetween(1, 5));
}
