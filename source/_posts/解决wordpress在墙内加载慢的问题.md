---
title: 解决wordpress在墙内加载慢的问题
url: 110.html
id: 110
comments: false
categories:
  - Linux
date: 2015-11-21 09:33:41
tags:
---

**替换掉谷歌字体**  
```bash
nano /wp-content/themes/twentyfifteen/function.php
```
插入
```PHP
function wpdx_replace_open_sans() {
  wp_deregister_style('open-sans');
  wp_register_style( 'open-sans', '//fonts.useso.com/css?family=Open+Sans:300italic,400italic,600italic,300,400,600' );
  if(is_admin()) wp_enqueue_style( 'open-sans');
}
add_action( 'init', 'wpdx_replace_open_sans' );
```

**替换gravatar头像**
插入
```PHP
function duoshuo_avatar($avatar) {
    $avatar = str_replace(array("www.gravatar.com","0.gravatar.com","1.gravatar.com","2.gravatar.com"),"gravatar.duoshuo.com",$avatar);
    return $avatar;
}
add_filter( 'get_avatar', 'duoshuo_avatar', 10, 3 );
```

然后保存