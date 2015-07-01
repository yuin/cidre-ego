===============================================================================
cidre-ego: an ego integration for the cidre web framework
===============================================================================

.. image:: https://godoc.org/github.com/yuin/cidre-ego?status.svg
    :target: http://godoc.org/github.com/yuin/cidre-ego

|

cidre-ego privides an easy way to integrate the `ego <https://github.com/benbjohnson/ego>`_ template engine into the `cidre <https://github.com/yuin/cidre/>`_ webframework.

----------------------------------------------------------------
Installation
----------------------------------------------------------------

.. code-block:: bash
   
   go get github.com/yuin/cidre-ego

----------------------------------------------------------------
Usage
----------------------------------------------------------------

Create ego templates and run the `ego` command.

    <%! func ShowItems(w io.Writer, items []string) error %>
    <%% import "strings" %%>
    <%% import "github.com/yuin/cidre-ego" %%>
        <ul>
          <% for _, item := range items { %>
            <li><%= item %></li>
          <% } %>
        </ul>
    <% ego.EgoLayout(w, MyLayout) %>

    <%! func MyLayout(w io.Writer, contents string) error %>
        <html><body>
        <%== contents %>
        </body></html>

.. code-block:: bash
    
    ego templates

Set EgoRenderer for a cidre app

.. code-block:: go

    app := cidre.NewApp(appConfig)
    app.Renderer = ego.NewEgoRenderer()
    items := app.MountPoint("/items/")
    items.Get("show_item", ".*", func(w http.ResponseWriter, r *http.Request) {
        app.Renderer.Html(w, ShowItems, []string{"a", "b", "c"})
    })

----------------------------------------------------------------
License
----------------------------------------------------------------
MIT

----------------------------------------------------------------
Author
----------------------------------------------------------------
Yusuke Inuzuka
