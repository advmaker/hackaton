hackaton
========
[![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/advmaker/hackaton?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

#Architecture
User
  Has many projects
  Has and belongs to many achievements

  - id
  - email
  - created_at
  - updated_at
  - password
  - deleted_at

Achievement
  Belongs to project
  Has many progresses

  - id
  - project_id
  - image_url
  - title
  - description
  - progress_limit
  - created_at
  - updated_at
  - deleted_at

Project
  Has many achievements

  - id
  - secret
  - created_at
  - updated_at
  - deleted_at

Progress
  Belongs to achievement

  - id
  - achievement_id
  - progress
  - player_id
  - player_extra
  - unlocked_at
  - created_at
  - updated_at


#TODO
1. [ ] Создать базу.
    - [ ] Настроить gorm.
    - [ ] Настроить goose.
    - [ ] Инициализировать структуру (миграции).
    - [ ] Создать модели.
2. [ ] Тесты.
3. [ ] Роутинг
4. [ ] Контроллеры
5. [ ] Шаблоны
...
