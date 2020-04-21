#go package 使用
import 方式：

1 別名: 把原本package name換成 另一個名稱

import bar ${package repo}

2 _ : 什麼都不引用 只載入init function

import _ ${package repo}

3 .: 把package function 直接載入global

import . ${pacakge repo}

4 載入原本的package name

import ${package repo}