OpenAPI Generator for Go
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

A fast and lightweight tool to **generate clean, standards-compliant OpenAPI 3 specifications** directly from your Go code. Perfect for:

* Auto-documenting REST APIs 🏗️
* Eliminating manual spec maintenance ✨
* Ensuring compatibility with Swagger UI, Postman & more 🔌

| Features:
| ✅ Struct-to-Schema auto-mapping
| ✅ Built-in validation & linting
| ✅ Custom template support

.. code-block:: go
   :linenos:

   // Just annotate & generate!
   // @method get
   // @path /users/{id}
   // @summary Get user
   // @description Get user information
   // @param id path isRequired User ID for get user information
   func GetUser(w http.ResponseWriter, r *http.Request) { ... }

**Get started:**

.. code-block:: shell

   $ go install github.com/goppuchino/oag
   $ cd <Your Project>
   $ oag

*Because nobody loves writing YAML by hand.* 🐹
