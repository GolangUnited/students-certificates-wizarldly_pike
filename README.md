# Certificates service
===

### Запуск сервиса:
*"cmd/rpc/main.go"*

### Перед запуском необходимо:
Задать переменные окружения - **"TEMPLATES_DIR"**, **"CERTIFICATES_DIR"**, с путями хранения файлов шаблонов и сертификатов.  
В системе должна быть установлена **"wkhtmltopdf"**, программа конвертации **HTML** файлов в **PDF**: https://wkhtmltopdf.org/downloads.html  
Исполняемый файл **"wkhtmltopdf"** должен находиться в каталоге запуска сервиса, или быть доступным по путям из переменной окружения **"PATH"**.  
Сгенерировать из **".proto"** вспомогательные файлы **gRPC**: *"protoc @protoc_options_generate.txt"*

### Требования к шаблонам:
В шаблоне могут содержаться следующие теги замены:
```
    {{.CourseName}}
    {{.CourseType}}
    {{.CourseHours}}
    {{.CourseDate}}
    {{.CourseMentors}}
    {{.StudentFirstname}}
    {{.StudentLastname}}
    {{.QrCodeLink}}
```
Шаблоны с любыми другими тегами замены будут отклонены валидатором.  
Вместо тега замены `{{.QrCodeLink}}` будет вставлен **QR** код в формате **PNG**: ссылка на сертификат.  
Например: `"<img src={{.QrCodeLink}} width="128" height="128">"`

### Пример простого HTML шаблона:
`"<html><body><h1 style="color:red;">Test html color<h1><p>{{.CourseName}}</p><p>{{.CourseType}}</p><p>{{.CourseHours}}</p><p>{{.CourseDate}}</p><p>{{.CourseMentors}}</p><p>{{.StudentFirstname}}</p><p>{{.StudentLastname}}</p><p><img src={{.QrCodeLink}} width="128" height="128"></p></body></html>"`
