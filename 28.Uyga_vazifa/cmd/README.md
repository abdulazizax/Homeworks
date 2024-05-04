## Uy vazifa

1. #### [PostresSQL tutorials](https://www.postgresqltutorial.com/) web-sayidagi `DATABASE CONSTRAINTS` bo'limi ohirgacha ko'rib chiqing
2. #### Demo databasega ulaning:
3. #### `users` nomli jadvalini yarating.
4. #### `user_id`, `username`, `email`, `password` va `created_at` kabi ustunlarni yarating. Har bir foydalanuvchi noyob identifikatorga ega bo'lishini ta'minlash uchun user_id ustuniga asosiy kalit cheklovini qo'shing. Har bir foydalanuvchi o'ziga xos elektron pochta manziliga ega bo'lishini ta'minlash uchun elektron pochta ustuniga noyob cheklovni (unique)kiriting.
5. #### Foydalanuvchilar o'rtasidagi aloqalarni (ya'ni, do'stlik) ifodalash uchun `friendships` nomli jadval yarating.
6. #### O'zaro munosabatlar holatini ko'rsatish uchun `user_id`, `friend_id` va `status` kabi ustunlarni belgilang (masalan, kutilayotgan, qabul qilingan, bloklangan).
7. #### `user_id` va `friend_id` foydalanuvchilar jadvalidagi to'g'ri foydalanuvchilarga havola qilishini ta'minlash uchun `foreign key` cheklovlarini qo'shing.
8. #### Doʻstlik soʻrovlari (request), qabul qilish (accept) va foydalanuvchilarni bloklash/blokdan chiqarish uchun `Go` funksiyalarini amalga oshiring. 