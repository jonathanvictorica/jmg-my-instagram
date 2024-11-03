# Modelo de Datos para MyInsta
Este documento describe el modelo de datos para una plataforma de redes sociales tipo Instagram. El modelo incluye tablas para almacenar usuarios, contenido multimedia, publicaciones, historias, ubicaciones, hashtags y relaciones entre ellas.

---

## Entidades y Esquema de Tablas

### 1. Hashtag

| Columna       | Tipo de Dato | PK | FK | Constraints        | Descripción                |
|---------------|--------------|----|----|--------------------|----------------------------|
| hashtag_id    | INT          | ✔  |    | `NOT NULL`        | Identificador único del hashtag |
| name          | VARCHAR(100) |    |    | `NOT NULL`, UNIQUE | Nombre del hashtag         |

---

### 2. Location

| Columna           | Tipo de Dato  | PK | FK | Constraints       | Descripción                          |
|-------------------|---------------|----|----|-------------------|--------------------------------------|
| location_id       | INT           | ✔  |    | `NOT NULL`       | Identificador único de la ubicación |
| country           | VARCHAR(100)  |    |    | `NOT NULL`       | País de la ubicación                |
| city              | VARCHAR(100)  |    |    | `NOT NULL`       | Ciudad de la ubicación              |
| info_aditional    | TEXT          |    |    |                   | Información adicional (opcional)    |

---

### 3. Multimedia Content

| Columna               | Tipo de Dato  | PK | FK | Constraints        | Descripción                          |
|-----------------------|---------------|----|----|--------------------|--------------------------------------|
| multimedia_content_id | INT           | ✔  |    | `NOT NULL`        | Identificador único del contenido multimedia |
| content_url           | TEXT          |    |    | `NOT NULL`        | URL del archivo multimedia           |
| content_type          | ENUM          |    |    | `NOT NULL`, CHECK (`content_type` IN ('IMAGE', 'VIDEO')) | Tipo de contenido (`IMAGE`, `VIDEO`) |

---

### 4. User_Insta

| Columna                   | Tipo de Dato     | PK | FK                          | Constraints      | Descripción                     |
|---------------------------|------------------|----|-----------------------------|------------------|---------------------------------|
| user_id                   | INT              | ✔  |                             | `NOT NULL`       | Identificador único del usuario |
| username                  | VARCHAR(50)      |    |                             | `NOT NULL`, UNIQUE | Nombre de usuario           |
| email                     | VARCHAR(100)     |    |                             | `NOT NULL`, UNIQUE | Correo electrónico          |
| profile_multimedia_content_id | INT          |    | `multimedia_content_id`      |                  | ID de contenido multimedia del perfil |
| description               | TEXT             |    |                             |                  | Descripción del usuario         |

---

### 5. Post

| Columna       | Tipo de Dato  | PK | FK               | Constraints            | Descripción                         |
|---------------|---------------|----|------------------|------------------------|-------------------------------------|
| post_id       | INT           | ✔  |                  | `NOT NULL`             | Identificador único del post        |
| user_id       | INT           |    | `user_id`       | `NOT NULL`             | ID del usuario que creó el post     |
| title         | VARCHAR(100)  |    |                  |                        | Título de la publicación            |
| created_at    | TIMESTAMP     |    |                  | `NOT NULL`             | Fecha de creación                   |
| post_type     | ENUM          |    |                  | `NOT NULL`, CHECK (`post_type` IN ('POST', 'REEL')) | Tipo de post (`POST`, `REEL`) |

---

### 6. Post_Content

| Columna               | Tipo de Dato  | PK | FK                          | Constraints         | Descripción                            |
|-----------------------|---------------|----|-----------------------------|---------------------|----------------------------------------|
| post_content_id       | INT           | ✔  |                             | `NOT NULL`          | Identificador único del contenido de post |
| post_id               | INT           |    | `post.post_id`             | `NOT NULL`          | ID del post                            |
| multimedia_content_id | INT           |    | `multimedia_content.multimedia_content_id` | `NOT NULL`          | ID del contenido multimedia asociado   |
| location_id           | INT           |    | `location.location_id`     |                     | ID de la ubicación asociada            |
| description           | TEXT          |    |                             |                     | Descripción del contenido de post      |
| orden                 | INT           |    |                             |                     | Orden del contenido en el post         |
| enabled               | BOOLEAN       |    |                             | `DEFAULT TRUE`      | Si el contenido está habilitado        |

---

### 7. Post_Content_Hashtag

| Columna               | Tipo de Dato  | PK | FK              | Constraints        | Descripción                   |
|-----------------------|---------------|----|-----------------|--------------------|-------------------------------|
| post_content_hashtag_id | INT         | ✔  |                 | `NOT NULL`        | Identificador único de la asociación de hashtag |
| post_content_id       | INT           |    | `post_content.post_content_id` | `NOT NULL`        | ID del contenido del post   |
| hashtag_id            | INT           |    | `hashtag.hashtag_id`     | `NOT NULL`        | ID del hashtag asociado      |

---

### 8. Post_Content_Users

| Columna               | Tipo de Dato  | PK | FK              | Constraints        | Descripción                   |
|-----------------------|---------------|----|-----------------|--------------------|-------------------------------|
| post_content_users_id | INT           | ✔  |                 | `NOT NULL`        | Identificador único de la asociación de usuario etiquetado |
| post_content_id       | INT           |    | `post_content.post_content_id` | `NOT NULL`        | ID del contenido del post   |
| user_id               | INT           |    | `user.user_id`         | `NOT NULL`        | ID del usuario etiquetado    |

---

### 9. History

| Columna               | Tipo de Dato  | PK | FK                | Constraints        | Descripción                    |
|-----------------------|---------------|----|-------------------|--------------------|--------------------------------|
| history_id            | INT           | ✔  |                   | `NOT NULL`        | Identificador único de la historia |
| title                 | VARCHAR(100)  |    |                   |                    | Título de la historia          |
| multimedia_content_id | INT           |    | `multimedia_content.multimedia_content_id` | `NOT NULL` | ID del contenido multimedia de la historia |
| orden                 | INT  |    |                   |   `NOT NULL`      | Orden de la historia          |
---

### 10. History_Content

| Columna               | Tipo de Dato  | PK | FK                          | Constraints         | Descripción                            |
|-----------------------|---------------|----|-----------------------------|---------------------|----------------------------------------|
| history_content_id    | INT           | ✔  |                             | `NOT NULL`          | Identificador único del contenido de historia |
| history_id            | INT           |    | `history.history_id`        | `NOT NULL`          | ID de la historia asociada             |
| multimedia_content_id | INT           |    | `multimedia_content.multimedia_content_id` | `NOT NULL`          | ID del contenido multimedia asociado   |
| location_id           | INT           |    | `location.location_id`      |                     | ID de la ubicación asociada            |
| description           | TEXT          |    |                             |                     | Descripción del contenido de historia  |
| orden                 | INT           |    |                             |                     | Orden del contenido en la historia     |
| enabled               | BOOLEAN       |    |                             | `DEFAULT TRUE`      | Si el contenido está habilitado        |

---

### 11. History_Content_Hashtag

| Columna                   | Tipo de Dato  | PK | FK                | Constraints        | Descripción                   |
|---------------------------|---------------|----|-------------------|--------------------|-------------------------------|
| history_content_hashtag_id | INT          | ✔  |                   | `NOT NULL`        | Identificador único de la asociación de hashtag |
| history_id                | INT           |    | `history.history_id`      | `NOT NULL`        | ID de la historia             |
| hashtag_id                | INT           |    | `hashtag.hashtag_id`      | `NOT NULL`        | ID del hashtag asociado       |

---

### 12. History_Content_Users

| Columna                   | Tipo de Dato  | PK | FK                | Constraints        | Descripción                   |
|---------------------------|---------------|----|-------------------|--------------------|-------------------------------|
| history_content_users_id  | INT           | ✔  |                   | `NOT NULL`        | Identificador único de la asociación de usuario etiquetado |
| history_id                | INT           |    | `history.history_id`      | `NOT NULL`        | ID de la historia             |
| user_id                   | INT           |    | `user.user_id`         | `NOT NULL`        | ID del usuario etiquetado     |

---------
```markdown
# Script SQL para Crear Tablas

Aquí tienes el script SQL para crear las tablas del modelo de datos en PostgreSQL:

```sql
-- Tabla Hashtag
CREATE TABLE hashtag (
    hashtag_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

-- Tabla Location
CREATE TABLE location (
    location_id SERIAL PRIMARY KEY,
    country VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    info_aditional TEXT
);

-- Tabla Multimedia Content
CREATE TABLE multimedia_content (
    multimedia_content_id SERIAL PRIMARY KEY,
    content_url TEXT NOT NULL,
    content_type VARCHAR(10) NOT NULL CHECK (content_type IN ('IMAGE', 'VIDEO'))
);

-- Tabla User
CREATE TABLE "User" (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    profile_multimedia_content_id INT,
    description TEXT,
    FOREIGN KEY (profile_multimedia_content_id) REFERENCES multimedia_content(multimedia_content_id)
);

-- Tabla Post
CREATE TABLE "Post" (
    post_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    post_type VARCHAR(10) NOT NULL CHECK (post_type IN ('POST', 'REEL')),
    FOREIGN KEY (user_id) REFERENCES "User"(user_id)
);

-- Tabla Post_Content
CREATE TABLE post_content (
    post_content_id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    multimedia_content_id INT NOT NULL,
    location_id INT,
    description TEXT,
    orden INT,
    enabled BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (post_id) REFERENCES "Post"(post_id),
    FOREIGN KEY (multimedia_content_id) REFERENCES multimedia_content(multimedia_content_id),
    FOREIGN KEY (location_id) REFERENCES location(location_id)
);

-- Tabla Post_Content_Hashtag
CREATE TABLE post_content_hashtag (
    post_content_hashtag_id SERIAL PRIMARY KEY,
    post_content_id INT NOT NULL,
    hashtag_id INT NOT NULL,
    FOREIGN KEY (post_content_id) REFERENCES post_content(post_content_id),
    FOREIGN KEY (hashtag_id) REFERENCES hashtag(hashtag_id)
);

-- Tabla Post_Content_Users
CREATE TABLE post_content_users (
    post_content_users_id SERIAL PRIMARY KEY,
    post_content_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (post_content_id) REFERENCES post_content(post_content_id),
    FOREIGN KEY (user_id) REFERENCES "User"(user_id)
);

-- Tabla History
CREATE TABLE history (
    history_id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    multimedia_content_id INT NOT NULL,
orden INT,
    FOREIGN KEY (multimedia_content_id) REFERENCES multimedia_content(multimedia_content_id)
);

-- Tabla History_Content
CREATE TABLE history_content (
    history_content_id SERIAL PRIMARY KEY,
    history_id INT NOT NULL,
    multimedia_content_id INT NOT NULL,
    location_id INT,
    description TEXT,
    orden INT,
    enabled BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (history_id) REFERENCES history(history_id),
    FOREIGN KEY (multimedia_content_id) REFERENCES multimedia_content(multimedia_content_id),
    FOREIGN KEY (location_id) REFERENCES location(location_id)
);

-- Tabla History_Content_Hashtag
CREATE TABLE history_content_hashtag (
    history_content_hashtag_id SERIAL PRIMARY KEY,
    history_id INT NOT NULL,
    hashtag_id INT NOT NULL,
    FOREIGN KEY (history_id) REFERENCES history(history_id),
    FOREIGN KEY (hashtag_id) REFERENCES hashtag(hashtag_id)
);

-- Tabla History_Content_Users
CREATE TABLE history_content_users (
    history_content_users_id SERIAL PRIMARY KEY,
    history_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (history_id) REFERENCES history(history_id),
    FOREIGN KEY (user_id) REFERENCES "User"(user_id)
);
