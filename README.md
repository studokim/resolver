# A toy DNS resolver

Implemented as an educational task for the Computer Networks course.

Inspired by [Julia Evans' blog post](https://jvns.ca/blog/2022/02/01/a-dns-resolver-in-80-lines-of-go/).

## Problem

> Безопасный DNS резолвер, который может фильтровать (выдавать не их реальный, а ранее заданный IP) сайты, заданные из конфигурационного файла.
>
> DNS-резолвер должен уметь кешировать ответы.
>
> Проверка будет проходить путем запуска кода на Linux машине и настройкой хоста на использование этого резолвера. Если все сайты работают — значит задача принимается.
>
> Можно использовать библиотеки для парсинга протокола. Запрещено использовать библиотеки для резолвинга.

## Usage

```bash
git clone https://github.com/studokim/resolver.git
cd resolver
go build
./resolver <example.com>
```

You can also fill the `filter.yml` with pre-defined `domain: ip` pairs. Such a domain will always be resolved to the specified ip.
