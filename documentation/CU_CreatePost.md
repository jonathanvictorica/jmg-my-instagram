# Caso de Uso: Crear Post

## Descripción
Este caso de uso describe el proceso que un usuario sigue para crear un post en la aplicación, incluyendo la carga de imágenes, la adición de título, descripción, ubicación opcional, hashtags y usuarios etiquetados.

## Flujo de Proceso

1. **Iniciar Creación de Post**
   - El usuario selecciona la opción "Crear Post" en la aplicación.

2. **Subir Imágenes**
   - El usuario puede subir una o más imágenes al post.
   - Por cada imagen, el sistema invoca el endpoint:

     ```
     POST /api/v1/images
     ```

     - **Propósito**: Enviar la imagen mediante FTP y recibir un identificador único de la imagen.
     - **Respuesta Esperada**: 
       ```json
       { "id": "unique_image_id" }
       ```
   
3. **Agregar Título y Descripción**
   - El usuario agrega un título al post.
   - El usuario agrega una descripción opcional al post.

4. **Seleccionar Ubicación (Opcional)**
   - El usuario puede seleccionar una ubicación para el post utilizando un campo de búsqueda.
   - El sistema invoca el siguiente endpoint para filtrar ubicaciones disponibles:

     ```
     GET /api/v1/locations/search?name={search_term}&top={number_of_results}
     ```

     - **Parámetros**:
       - `name`: Nombre de la ubicación que el usuario desea buscar.
       - `top`: Número máximo de resultados que deben devolverse.
     - **Respuesta Esperada**: Una lista de ubicaciones en formato de cadena de texto.

5. **Agregar Hashtags (Opcional)**
   - El usuario puede añadir hasta 10 hashtags al post.
   - Los hashtags deben comenzar con `#` y escribirse sin espacios.

6. **Etiquetar Usuarios (Opcional)**
   - El usuario puede buscar y etiquetar otros usuarios en el post mediante un campo de búsqueda.
   - El sistema invoca el siguiente endpoint para buscar usuarios:

     ```
     GET /api/v1/users/search?name={search_term}&top={number_of_results}
     ```

     - **Parámetros**:
       - `name`: Nombre del usuario a buscar.
       - `top`: Número máximo de resultados que deben devolverse.
     - **Respuesta Esperada**: Lista de usuarios coincidentes en formato de texto.

7. **Publicar Post**
   - El usuario selecciona "Publicar" para crear el post final.
   - El sistema invoca el siguiente endpoint:

     ```
     POST /api/v1/posts
     ```

     - **Parámetros en el Request**:
       - `title`: Título del post.
       - `description`: Descripción del post.
       - `sections`: Lista de secciones que contienen:
         - `images`: Lista de identificadores de imágenes cargadas.
         - `hashtags`: Lista de hashtags seleccionados.
         - `users`: Lista de IDs de usuarios etiquetados.
         - `location`: ID de la ubicación seleccionada.

     - **Respuesta Esperada**:
       - En caso de éxito, retorna un identificador único del post.
       - En caso de error, retorna un mensaje de error.

## Posibles Errores

- **post_can_not_create**: No se pudo crear el post por un error inesperado.
- **image_not_found**: Alguna de las imágenes referenciadas no se encontró.
- **user_not_found**: Alguno de los usuarios etiquetados no se encontró.
- **hashtag_not_found**: Algún hashtag referenciado no se encontró.
