# Certificates service
===

### Запуск сервиса в Docker:
#### Создание образа и генерация **protobuf** файлов из **".proto"**:
`docker image build -f Dockerfile.proto -t cert-proto .`

#### Сохранение в локальном проекте **protobuf** файлов из Docker контейнера **cert-proto**:
`docker container run --rm -v $PWD/protobuf/transport/certificate:/result cert-proto` // Linux  
`docker container run --rm -v %CD%/protobuf/transport/certificate:/result cert-proto` // Windows

#### Создание образа для запуска проекта:
`docker image build -f Dockerfile.app -t cert-app .`

#### Запуск проекта:
`docker container run --rm --env-file env.list -p 1234:1234 -it -v certificate-storage:/storage cert-app`   
В файле **env.list** задаются переменные откружения необходимые для запуска сервиса.   
Монтируется Docker VOLUME: **certificate-storage** на локальной машине/сервере, для сохранения шаблонов и сертификатов.

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

### Пример простого HTML шаблона:
```
<html><body><h1 style="color:red;">Test html color<h1>
<p>{{.CourseName}}</p>
<p>{{.CourseType}}</p>
<p>{{.CourseHours}}</p>
<p>{{.CourseDate}}</p>
<p>{{.CourseMentors}}</p>
<p>{{.StudentFirstname}}</p>
<p>{{.StudentLastname}}</p>
<p><img src={{.QrCodeLink}} width="128" height="128"></p>
</body></html>
```
