# barts

Brother code art project


## Pages


 - `/` logged out hello world
 - `/auth/login` login / auth w/ phone number
 - `/auth/logout`
 - `/posts/\w+` actual post page
   - example `/posts/abcdef`
   - Post types
     - Images
     - Audio
     - HTML
 - `/nav` page with tree and navigation

## Database


 - `users` table
   - `id`
   - `phonenumber`
 - `posts` table
   - `id` integers
   - `previous` integer
   - `slug` piece in url
   - `type`
   - `content_link` - http url to thing

