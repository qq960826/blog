{% extends '_layout.swig' %}
{% import '_macro/post-collapse.swig' as post_template %}
{% import '_macro/sidebar.swig' as sidebar_template %}

{% block title %}{{ __('title.archive') }} | {{ title }}{% endblock %}

{% block page_class %}page-archive{% endblock %}

{% block content %}

  {#####################}
  {### ARCHIVE BLOCK ###}
  {#####################}
  <div class="post-block archive">
    <div id="posts" class="posts-collapse">

      {%- if theme.cheers %}
      <span class="archive-move-on"></span>
      <span class="archive-page-counter">
        {%- set cheers %}
        {%- set posts_length = site.posts.length %}
        {%- if posts_length > 210 %} {%- set cheers = 'excellent' %}
          {% elif posts_length > 130 %} {%- set cheers = 'great' %}
          {% elif posts_length > 80 %} {%- set cheers = 'good' %}
          {% elif posts_length > 50 %} {%- set cheers = 'nice' %}
          {% elif posts_length > 30 %} {%- set cheers = 'ok' %}
        {% else %}
          {%- set cheers = 'um' %}
        {%- endif %}
        {{ __('cheers.' + cheers) }}! {{ _p("counter.archive_posts", site.posts.length) }} {{ __('keep_on') }}
      </span>
      {%- endif %}

      {{ post_template.render(page.posts) }}

    </div>
  </div>
  {#########################}
  {### END ARCHIVE BLOCK ###}
  {#########################}

  {% include '_partials/pagination.swig' %}

{% endblock %}

{% block sidebar %}
  {{ sidebar_template.render(false) }}
{% endblock %}
