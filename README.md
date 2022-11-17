# Certificates service
===

### Запуск сервиса в Docker:
Собрать образ коммандой:  
**"docker build -t cert ."**  
Запустить сервис коммандой:  
**"docker run -p 1234:1234 --rm cert"**  
По умолчанию входящий порт **1234**  
Автоматически монтируется **Docker VOLUME** на локальной машине/сервере, для сохранения шаблонов и сертификатов:   
**certificatesstorage/templates**  
**certificatesstorage/certificates**

### Запуск сервиса локально:
Заранее задать переменные окружения - **"TEMPLATES_DIR"**, **"CERTIFICATES_DIR"**, с путями хранения файлов шаблонов и сертификатов.  
В системе должна быть установлена **"wkhtmltopdf"**, программа конвертации **HTML** файлов в **PDF**: https://wkhtmltopdf.org/downloads.html  
Исполняемый файл **"wkhtmltopdf"** должен находиться в каталоге запуска сервиса, или быть доступным по путям из переменной окружения **"PATH"**.  
Сгенерировать из **".proto"** вспомогательные файлы **gRPC**: **"protoc @protoc_options_generate.txt"**  
Запуск сервиса: **"cmd/rpc/main.go"**

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
```
<html><body><h1 style="color:red;">Test html color<h1><p>{{.CourseName}}</p><p>{{.CourseType}}</p><p>{{.CourseHours}}</p><p>{{.CourseDate}}</p><p>{{.CourseMentors}}</p><p>{{.StudentFirstname}}</p><p>{{.StudentLastname}}</p><p><img src={{.QrCodeLink}} width="128" height="128"></p></body></html>
```
