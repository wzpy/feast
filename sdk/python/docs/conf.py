# Configuration file for the Sphinx documentation builder

# ============================================================
# Project Information
# ============================================================

project   = "feast"
copyright = "2019, Feast Authors"
author    = "Feast Authors"

# The short X.Y version
version   = "0.3"
# The full version, including alpha/beta/rc tags
release   = "0.3.0"

# ============================================================
# General Configuration
# ============================================================

needs_sphinx = "1.6.3"

extensions = [
    "sphinx.ext.autodoc",
    "sphinx.ext.doctest",
    "sphinx.ext.intersphinx",
    "sphinx.ext.todo",
    "sphinx.ext.viewcode",
    "sphinx.ext.autosummary",
    "m2r",
]

templates_path   = ["_templates"]
source_suffix    = [".rst", ".md"]
master_doc       = "index"
exclude_patterns = []
pygments_style   = "sphinx"

# ============================================================
# HTML Output 
# ============================================================

html_theme         = "alabaster"
html_theme_options = {
    "description": "Feast Python SDK",
    "github_user": "gojek",
    "github_repo": "feast",
    "github_type": "star",
    "github_banner": True,
    "font_family": "'Roboto', Georgia, sans",
    "head_font_family": "'Roboto', Georgia, serif",
    "code_font_family": "'Roboto Mono', 'Consolas', monospace",
}

html_static_path = ["_static"]
html_sidebars    = {
    "**": [
        "about.html",
        "navigation.html",
        "relations.html",
        "searchbox.html",
    ]
}
html_favicon     = "_static/favicon.png"

# ============================================================
# Extensions 
# ============================================================

autodoc_default_flags = ["members"]
autosummary_generate = True
