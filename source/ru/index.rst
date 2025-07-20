OpenAPI генератор для Go
========================

.. image:: https://img.shields.io/github/go-mod/go-version/goppuchino/oag
   :alt: GitHub go.mod Go version

.. image:: https://img.shields.io/github/commit-activity/y/goppuchino/oag
   :alt: GitHub commit activity

.. image:: https://img.shields.io/github/last-commit/goppuchino/oag
   :alt: GitHub last commit

.. image:: https://img.shields.io/github/contributors/goppuchino/oag
   :alt: GitHub contributors

.. image:: https://img.shields.io/github/license/goppuchino/oag
   :alt: GitHub License

.. image:: https://raw.githubusercontent.com/goppuchino/oag/master/assets/oag.png
   :align: right
   :width: 180px

Быстрый и легкий инструмент для **генерации чистых спецификаций OpenAPI 3, соответствующих стандартам** напрямую из
вашего Go кода. Идеально подходит для:

* Автоматического документирования REST API 🏗️
* Устранения ручной поддержки спецификаций ✨
* Обеспечения совместимости со Swagger UI, Postman и другими инструментами 🔌

| Возможности:
| ✅ Автоматическое сопоставление структур и схем
| ✅ Встроенная валидация и проверка
| ✅ Поддержка пользовательских шаблонов

.. code-block:: go
   // Just annotate & generate!
   // @method get
   // @path /users/{id}
   // @summary Get user
   // @description Get user information
   // @param id path isRequired User ID for get user information
   func GetUser(w http.ResponseWriter, r *http.Request) { ... }

**Начало работы:**

.. code-block:: shell
   go install github.com/goppuchino/oag
   cd <ваш проект>
   oag

*Потому что никто не любит писать YAML вручную.* 🐹


.. toctree::
   :maxdepth: 2
   :caption: Содержание: