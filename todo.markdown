- [x] скопировать на github 
- [x] запустить на rpi
- [x] сохранять в БД например sqlite или в файл
- [x] побороть проблему с использованием ? и & в отправке сообщений
- [x] отправлять список TODO одним сообщением, для вывода сообщений использовать шаблоны text/template
- [x] добавить время добавления сообщения и отображать это в сообщении
- [x] нумеровать сообщение от единицы
- [x] добавить команду `/delete_all` которая удаляет все сообщения от данного пользователя
- [x] если сообщение от пользователя начинается с `/delete 2` удалять сообщение с определённым номером
- [x] если сообщение от пользователя начинается с `/restore 2` восстанавливать сообщение с определённым номером
  (не так просто, т.к. нужно взять как-то определить порядок и ссылаться на нужное сообщение, предлагаю сохранять message id и по нему уже удалять)
- [x] добавить команду /list которая не добавляет в todo лист, а просто выводит список текущих
- [ ] обрабатывать весь Result а не Result[0]
- [x] в telegrambot иметь возможность сменить secret например задавать через переменную окружения (`os.GetEnv`)
- [ ] преобразовать метод get на использование map (один аргумент, вместо двух keys, values)
- [x] /done должен принимать число, а не удалять первое
- [x] /undone N проставить, что todo'шка не сделана
- [ ] написать тесты на бизнес логику
- [ ] сделать подсказки для пользователя в виде кнопок https://core.telegram.org/bots/api#inlinekeyboardmarkup
- [х] /list_deleted показывает удалённые сообщения
- [ ] по всему коду возвращать ошибки, а не паниковать. Обрабатывать, а не забивать на них через `_`
- [ ] вынести из main'а все действия (handler*) в отдельный пакет и по файлам, например в директорию commands/ файлы add_todo.go. delete_todo.go и т.п.
- [ ] bitcask: написать многопоточный стресс тест (две горутины одновременно пишут в базу 100 раз каждый два разных ключа, дождаться пока они допишут через waitgroup и потом через get проверить, что всё правильно записалось)
- [ ] bitcask: вынести errors.New("record not found") в отдельную публичную глобальную константу, чтобы можно было ошибку с ней сравнивать
- [ ] bitcask: сейчас если выбрать директорию для базы которая не найдена, то будет ошибка, возможно стоит создать директорию.
- [ ] bitcask: добавить функцию Compact(dir), которая оставляет в файле БД только последние значения, т.к. если 10000 раз перезаписать один и тот же ключ, то старые значения всё равно будут в БД
- [ ] перенести этот TODO лист в SharedTodoBot
- [ ] Get если не смог найти User'а в базе создаёт его, мне кажется такое поведение неправильное, возможно нужно что-то вроде /start сделать для пользователя и уже создавать пустого.
