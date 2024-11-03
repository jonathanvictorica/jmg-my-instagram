
# Documentación de Endpoints

## 1. Subir Imagen
- **Endpoint**: `POST /api/v1/images`
- **Método**: POST
- **Propósito**: Permite al usuario subir una imagen para el post y devuelve un identificador único para cada imagen cargada.
- **Parámetros**:
  - Cuerpo de la petición: archivo de imagen.
- **Respuesta**:
  - **200 OK**: 
    ```json
    { "id": "unique_image_id" }
    ```
  - **Errores**:
    - `image_upload_failed`: Error en la carga de la imagen.

---

## 2. Buscar Ubicaciones
- **Endpoint**: `GET /api/v1/locations/search`
- **Método**: GET
- **Propósito**: Realizar una búsqueda de ubicaciones según el nombre ingresado por el usuario.
- **Parámetros de Query**:
  - `name`: Nombre o parte del nombre de la ubicación a buscar.
  - `top`: Número máximo de resultados a devolver.
- **Respuesta**:
  - **200 OK**:
    ```json
    [
      "location_1",
      "location_2",
      "location_3"
    ]
    ```
  - **Errores**:
    - `location_not_found`: No se encontraron ubicaciones que coincidan con la búsqueda.

---

## 3. Buscar Usuarios
- **Endpoint**: `GET /api/v1/users/search`
- **Método**: GET
- **Propósito**: Buscar usuarios en la base de datos mediante el nombre.
- **Parámetros de Query**:
  - `name`: Nombre o parte del nombre del usuario.
  - `top`: Número máximo de resultados a devolver.
- **Respuesta**:
  - **200 OK**:
    ```json
    [
      { "user_id": "123", "username": "user1" },
      { "user_id": "456", "username": "user2" }
    ]
    ```
  - **Errores**:
    - `user_not_found`: No se encontraron usuarios que coincidan con la búsqueda.

---

## 4. Publicar Post
- **Endpoint**: `POST /api/v1/posts`
- **Método**: POST
- **Propósito**: Publicar un nuevo post con las imágenes, título, descripción, ubicación, hashtags y usuarios etiquetados.
- **Parámetros en el Cuerpo de la Petición**:
  ```json
  {
    "title": "Post title",
    "description": "Post description",
    "user_id": "user id",
    "sections": [
      {
        "images": ["image_id1", "image_id2"],
        "hashtags": ["#example", "#tag"],
        "users": ["user_id1", "user_id2"],
        "location": "location_id"
      }
    ]
  }
- **Respuesta**:
  - **200 OK**:
    ```json
    { "post_id": "new_post_id" }
    ```
  - **Errores**:
    - `post_can_not_create`: No se pudo crear el post.
    - `image_not_found`: Alguna imagen referenciada no se encontró.
    - `user_not_found`: Algún usuario etiquetado no se encontró.
    - `hashtag_not_found`: Algún hashtag referenciado no se encontró.
---

## 5. Obtener Información de Perfil
- **Endpoint**: `GET /api/v1/users/{userId}/info-profile`
- **Método**: GET
- **Propósito**: Recuperar la información del perfil del usuario especificado.
- **Parámetros**:
  - `userId` (en la URL): ID del usuario cuyo perfil se desea ver.
- **Respuesta**:
  - **200 OK**:
    ```json
    {
      "username": "nombre_usuario",
      "description": "Descripción del perfil",
      "image_url": "url_de_imagen_de_perfil"
    }
    ```
  - **Errores**:
    - `user_not_found`: No se encontró al usuario con el `userId` especificado.

---

## 6. Obtener Publicaciones por Usuario
- **Endpoint**: `GET /api/v1/posts/get-by-user/{userId}`
- **Método**: GET
- **Propósito**: Obtener las publicaciones asociadas al usuario especificado.
- **Parámetros**:
  - `userId` (en la URL): ID del usuario cuyas publicaciones se desean ver.
- **Respuesta**:
  - **200 OK**:
    ```json
    [
      {
        "id": "post_id_1",
        "title": "Título de la publicación",
        "image_url": "url_de_imagen_principal"
      },
      {
        "id": "post_id_2",
        "title": "Otro título de publicación",
        "image_url": "url_de_otra_imagen"
      }
    ]
    ```
  - **Errores**:
    - `user_not_found`: No se encontró al usuario con el `userId` especificado.
    - `posts_not_found`: No se encontraron publicaciones para el usuario.

---

## 7. Obtener Historias Destacadas por Usuario
- **Endpoint**: `GET /api/v1/history/get-by-user/{userId}`
- **Método**: GET
- **Propósito**: Recuperar las historias destacadas asociadas al usuario especificado.
- **Parámetros**:
  - `userId` (en la URL): ID del usuario cuyas historias destacadas se desean ver.
- **Respuesta**:
  - **200 OK**:
    ```json
    [
      {
        "id": "history_id_1",
        "title": "Título de la historia destacada",
        "image_url": "url_de_imagen_de_historia"
      },
      {
        "id": "history_id_2",
        "title": "Otro título de historia destacada",
        "image_url": "url_de_otra_imagen_de_historia"
      }
    ]
    ```
  - **Errores**:
    - `user_not_found`: No se encontró al usuario con el `userId` especificado.
    - `history_not_found`: No se encontraron historias destacadas para el usuario.

---
