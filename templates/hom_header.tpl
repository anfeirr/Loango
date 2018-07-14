{{ define "hom_header"}}

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
    <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta name="apple-mobile-web-app-title" content="聚合信用">
    <meta name="apple-mobile-web-app-capable" content="yes" />
    <meta name="apple-mobile-web-app-status-bar-style" content="blue" />
    <meta name="format-detection" content="telephone=no, email=no" />
    <meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
    <meta http-equiv="Pragma" content="no-cache" />
    <meta http-equiv="Expires" content="0" />
    <meta name="keywords" content="{{.title}}" />
    <meta name="description" content="{{.title}}" />

    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon">
    <script type="text/javascript" src="statics/js/home/htmlrem.js"></script>
    <script type="text/javascript" src="statics/js/home/jquery-3.2.1.min.js"></script>
    <script type="text/javascript" src="statics/js/home/global.js"></script>
    <script type="text/javascript" src="statics/js/home/iscroll-probe.min.js"></script>
    <script type="text/javascript" src="statics/js/home/swiper.js"></script>
    <script type="text/javascript" src="statics/js/home/t-layer.js"></script>

    <link rel="stylesheet" href="statics/css/home/swiper.min.css" />
    <link rel="stylesheet" href="statics/css/home/reset.css" />
    <link rel="stylesheet" href="statics/css/home/common.css?20170814" />

</head>

{{end}}