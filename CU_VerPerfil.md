1. El usuario hace click en el boton ver perfil
2. El sistema invoca los siguientes endpoints:
   GET api/v1/users/{userId}/info-profile retorna
   {
       "username", "description", "image_url"
   }
3. El sistema invoca endpoint para obtener publicaciones
    GET api/v1/posts/get-by-user/{userId} retorna
   {
       "id", "title", "image_url"
   }
4. El sistema invoca endpoint para obtener historias destacadas
   GET api/v1/history/get-by-user/{userId} retorna
   {
       "id", "title", "image_url"
   }
5. el sistema dibuja en la pantalla
